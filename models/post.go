package models

type Post struct {
	Id              int16  `json:"id"`
	Caption         string `json:"caption"`
	ImageURL        string `json:"image_url"`
	PostedTimeStamp int32  `json:"posted_timestamp"`
}
