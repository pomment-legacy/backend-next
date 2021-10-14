package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pomment/backend-next/server/dao"
	"github.com/pomment/backend-next/server/utils"
	"path"
)

type Post struct {
	UUID         string  `json:"uuid"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Website      string  `json:"website,omitempty"`
	Parent       string  `json:"parent,omitempty"`
	ParentName   string  `json:"parentName,omitempty"`
	Child        []Post  `json:"child,omitempty"`
	Content      string  `json:"content"`
	Hidden       bool    `json:"hidden"`
	ByAdmin      bool    `json:"byAdmin"`
	ReceiveEmail bool    `json:"receiveEmail"`
	EditKey      string  `json:"editKey,omitempty"`
	CreatedAt    int64   `json:"createdAt"`
	UpdatedAt    int64   `json:"updatedAt"`
	OrigContent  string  `json:"origContent"`
	Avatar       string  `json:"avatar,omitempty"`
	Rating       float64 `json:"rating,omitempty"`
}

func GetPostsRaw(url string) (data string, err error) {
	postPath := path.Join("threads", utils.EncodeURIComponent(url))
	rawData, err := dao.Read(fmt.Sprintf("%s.json", postPath))
	return rawData, err
}

func GetPosts(url string) (data *[]Post, err error) {
	rawData, err := GetPostsRaw(url)
	if err != nil {
		return nil, err
	}
	var posts []Post
	err = json.Unmarshal([]byte(rawData), &posts)
	return &posts, err
}

func GetPost(url string, uuid string) (data *Post, err error) {
	posts, err := GetPosts(url)
	if err != nil {
		return nil, err
	}
	for _, e := range *posts {
		if e.UUID == uuid {
			data = &e
			break
		}
	}
	if data == nil {
		return nil, errors.New("post not found")
	}
	return data, nil
}
