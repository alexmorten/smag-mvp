package db

import (
	"fmt"

	"database/sql"

	"github.com/alexmorten/smag-mvp/config"
	"github.com/jinzhu/gorm"

	// necessary for "database/sql"
	_ "github.com/lib/pq"

	// necessary for gorm :pointup:
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateOrUpdate checks whether a specific (gorm) database entry already exists using a model filter,
// creates it in case no record is found and updates the first in case of existing record(s)
func CreateOrUpdate(db *gorm.DB, out interface{}, where interface{}, update interface{}) error {
	var err error

	tx := db.Begin()
	if tx.Where(where).First(out).RecordNotFound() {
		err = tx.Create(update).Scan(out).Error
	} else {
		err = tx.Model(out).Update(update).Scan(out).Error
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

//Create just create a new entry in the database
func Create(db *gorm.DB, out interface{}, update interface{}) error {
	return db.Create(update).Scan(out).Error
}

// NewConn establishes a db connection
func NewConn(conf config.PostgresConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=postgres dbname=instascraper sslmode=disable", conf.Host)
	if conf.Password != "" {
		connectionString += " " + "password=" + conf.Password
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
