package model

import (
	"time"
)

// omit is the bool type for omitting a field of struct.
type omit bool

// CommentList is list that contains comments and meta.
type CommentList struct {
	Comments    []Comment `json:"comments"`
	HasPrev     bool      `json:"hasPrev"`
	HasNext     bool      `json:"hasNext"`
	Count       int       `json:"count"`
	CurrentPage int       `json:"currentPage"`
}

// Comment is a comment model.
type Comment struct {
	Id          uint      `json:"id"`
	Content     string    `json:"content"`
	UserId      uint      `json:"userId"`
	LikingCount int       `json:"likingCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	User        User      `json:"user"`
}

// LikingList is list that contains likings and meta.
type LikingList struct {
	Likings     []User `json:"likings"`
	HasPrev     bool   `json:"hasPrev"`
	HasNext     bool   `json:"hasNext"`
	Count       int    `json:"count"`
	CurrentPage int    `json:"currentPage"`
	IsLiked     bool   `json:"isLiked"`
}

// LikedList is list that contains liked and meta.
type LikedList struct {
	Liked       []User `json:"liked"`
	HasPrev     bool   `json:"hasPrev"`
	HasNext     bool   `json:"hasNext"`
	Count       int    `json:"count"`
	CurrentPage int    `json:"currentPage"`
	
}

// Image is a image model.
type Image struct {
	Id        uint      `json:"id"`
	Kind      int       `json:"kind"`
	Large     string    `json:"large"`
	Medium    string    `json:"medium"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"createdAt"`
}

// Tag is a tag model.
type Tag struct {
	Id   uint   `json:"id"`
	Name string `json:"name",sql:"size:255"`
}

// Link is a link model.
type Link struct {
	Id        uint      `json:"id"`
	Kind      int       `json:"kind"`
	Name      string    `json:"title",sql:"size:255"`
	Url       string    `json:"url",sql:"size:512"`
	CreatedAt time.Time `json:"createdAt"`
	Icon      string    `json:"icon",sql:"size:255"`
}
