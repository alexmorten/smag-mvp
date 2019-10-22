package models

import (
	"time"

	"github.com/codeuniversity/smag-mvp/utils"
)

// TwitterUserList is a custom type of TwitterUser to be used for easier handling
// of relation users in twitter inserters
type TwitterUserList []*TwitterUser

// TwitterUserRaw holds the follow graph info, only relating userNames
type TwitterUserRaw struct {
	// Meta
	ID   string `json:"id"`
	URL  string `json:"url"`
	Type string `json:"type"`

	// User info
	Name            string `json:"name"`
	Username        string `json:"username"`
	Bio             string `json:"bio"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`

	// Profile stats
	Location   string `json:"location"`
	JoinDate   string `json:"join_date"`
	JoinTime   string `json:"join_time"`
	IsPrivate  int    `json:"is_private"`
	IsVerified int    `json:"is_verified"`

	// Follows
	Following     int      `json:"following"`
	FollowingList []string `json:"following_list"`
	Followers     int      `json:"followers"`
	FollowersList []string `json:"followers_list"`

	// Usage stats
	Tweets     int `json:"tweets"`
	Likes      int `json:"likes"`
	MediaCount int `json:"media_count"`
}

// TwitterUser holds the follow graph info, only relating userNames
type TwitterUser struct {
	GormModelWithoutID

	// Meta
	UserIdentifier string
	URL            string
	Type           string

	// User info
	Name            string
	Username        string `gorm:"primary_key"`
	Bio             string
	Avatar          string
	BackgroundImage string

	// Profile stats
	Location   string
	JoinDate   time.Time
	IsPrivate  bool
	IsVerified bool

	// Follows
	Following     int
	FollowingList []*TwitterUser `gorm:"many2many:twitter_followings;association_jointable_foreignkey:following_user_id"`
	Followers     int
	FollowersList []*TwitterUser `gorm:"many2many:twitter_followers;association_jointable_foreignkey:followers_user_id"`

	// Usage stats
	Tweets     int
	Likes      int
	MediaCount int
}

// ConvertTwitterUser converts the raw TwitterUser structure
// from kafka into the database model
func ConvertTwitterUser(raw *TwitterUserRaw) *TwitterUser {

	followingList := make([]*TwitterUser, len(raw.FollowingList))
	followersList := make([]*TwitterUser, len(raw.FollowersList))

	for index, item := range raw.FollowingList {
		followingList[index] = &TwitterUser{
			Username: item,
		}
	}

	for index, item := range raw.FollowersList {
		followersList[index] = &TwitterUser{
			Username: item,
		}
	}

	joinDate, _ := utils.ConvertDateStrToTime(raw.JoinDate)

	isPrivate := utils.ConvertIntToBool(raw.IsPrivate)
	isVerified := utils.ConvertIntToBool(raw.IsVerified)

	return &TwitterUser{
		UserIdentifier: raw.ID,
		URL:            raw.URL,
		Type:           raw.Type,

		Name:            raw.Name,
		Username:        raw.Username,
		Bio:             raw.Bio,
		Avatar:          raw.Avatar,
		BackgroundImage: raw.BackgroundImage,

		Location:   raw.Location,
		JoinDate:   joinDate,
		IsPrivate:  isPrivate,
		IsVerified: isVerified,

		Following:     raw.Following,
		FollowingList: followingList,
		Followers:     raw.Followers,
		FollowersList: followersList,

		Tweets:     raw.Tweets,
		Likes:      raw.Likes,
		MediaCount: raw.MediaCount,
	}
}

// NewTwitterUserList converts multiple TwitterUser slices into a single custom
// TwitterUserList structure
func NewTwitterUserList(slices ...[]*TwitterUser) *TwitterUserList {
	list := &TwitterUserList{}

	for _, slice := range slices {
		*list = append(*list, slice...)
	}

	return list
}

// RemoveDuplicates removes duplicated users from TwitterUserList
func (list *TwitterUserList) RemoveDuplicates() {
	uniqueSlice := make([]*TwitterUser, 0, len(*list))
	set := make(map[string]bool)

	for i, user := range *list {
		if !set[user.Username] {
			set[user.Username] = true
			uniqueSlice = append(uniqueSlice, (*list)[i])
		}
	}

	*list = uniqueSlice
}
