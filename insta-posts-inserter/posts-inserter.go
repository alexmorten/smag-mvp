package insta_posts_inserter

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/codeuniversity/smag-mvp/models"
	"github.com/codeuniversity/smag-mvp/service"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"time"
)

type InstaPostInserter struct {
	postQReader *kafka.Reader
	errQWriter  *kafka.Writer
	*service.Executor
	db *sql.DB
}

func New(postgresHost, postgresPassword string, postQReader *kafka.Reader, errQWriter *kafka.Writer) *InstaPostInserter {
	p := &InstaPostInserter{}
	p.postQReader = postQReader
	p.errQWriter = errQWriter
	p.Executor = service.New()

	connectionString := fmt.Sprintf("host=%s user=postgres dbname=instascraper sslmode=disable", postgresHost)
	if postgresPassword != "" {
		connectionString += " " + "password=" + postgresPassword
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	p.db = db
	return p
}

func (i *InstaPostInserter) findOrCreateUser(username string) (userID int, err error) {
	err = i.db.QueryRow("Select id from users where user_name = $1", username).Scan(&userID)

	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}

		var insertedUserID int
		err := i.db.QueryRow(`INSERT INTO users(user_name) VALUES($1) RETURNING id`, username).Scan(&insertedUserID)
		if err != nil {
			return 0, err
		}

		userID = int(insertedUserID)
	}

	return userID, nil
}

func (i *InstaPostInserter) Run() {
	defer func() {
		i.MarkAsStopped()
	}()

	fmt.Println("starting inserter")
	for i.IsRunning() {

		message, err := i.postQReader.FetchMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}

		var post models.InstagramPost
		err = json.Unmarshal(message.Value, &post)
		if err != nil {
			panic(err)
		}

		err = i.insertPost(post)

		if err != nil {
			panic(fmt.Errorf("comments inserter failed %s ", err))
		}
		fmt.Println("Insert Post: ", post.ShortCode)
		i.postQReader.CommitMessages(context.Background(), message)
	}
}

func (i *InstaPostInserter) insertPost(post models.InstagramPost) error {

	userID, err := i.findOrCreateUser(post.UserId)

	if err != nil {
		return err
	}

	_, err = i.db.Exec(`INSERT INTO posts(user_id, post_id, short_code, picture_url) VALUES($1,$2,$3,$4) ON CONFLICT(post_id) DO UPDATE SET short_code=$2, picture_url=$4`, userID, post.PostId, post.ShortCode, post.PictureUrl)

	if err != nil {
		return err
	}

	return nil
}

func (i *InstaPostInserter) Close() {
	i.Stop()
	i.WaitUntilStopped(time.Second * 3)

	i.errQWriter.Close()
	i.postQReader.Close()
	i.MarkAsClosed()
}
