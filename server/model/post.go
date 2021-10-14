package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pomment/backend-next/server/dao"
	"github.com/pomment/backend-next/server/utils"
	"path"
	"time"
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

func GetPost(url string, uuid string) (data *Post, posts *[]Post, err error) {
	posts, err = GetPosts(url)
	if err != nil {
		return nil, nil, err
	}
	for _, e := range *posts {
		if e.UUID == uuid {
			data = &e
			break
		}
	}
	if data == nil {
		return nil, nil, errors.New("post not found")
	}
	return data, posts, nil
}

type SetPostParam struct {
	Url               string  `json:"url"`
	UUID              string  `json:"uuid"`
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	Website           string  `json:"website"`
	Content           string  `json:"content"`
	Hidden            bool    `json:"hidden"`
	ByAdmin           bool    `json:"byAdmin"`
	ReceiveEmail      bool    `json:"receiveEmail"`
	Avatar            string  `json:"avatar"`
	ResetEditKey      bool    `json:"resetEditKey"`
	PreserveUpdatedAt bool    `json:"preserveUpdatedAt"`
}

func SetPost(param SetPostParam) (data *Post, err error) {
	data, posts, err := GetPost(param.Url, param.UUID)
	if err != nil {
		return nil, err
	}
	data.Name = param.Name
	data.Email = param.Email
	data.Website = param.Website
	data.Content = param.Content
	data.Hidden = param.Hidden
	data.ByAdmin = param.ByAdmin
	data.ReceiveEmail = param.ReceiveEmail
	data.Avatar = param.Avatar
	if !param.PreserveUpdatedAt {
		now := time.Now()
		data.UpdatedAt = now.Unix() * 1000
	}
	if param.ResetEditKey {
		data.EditKey = utils.GetEditKey()
	}
	postList := *posts
	target := -1
	for i, _ := range postList {
		if postList[i].UUID == param.UUID {
			target = i
			break
		}
	}
	if target < 0 {
		return nil, errors.New("post not found when saving")
	}
	postList[target] = *data
	err = SavePosts(param.Url, postList)
	return data, err
}

func SavePosts(url string, data []Post) (err error) {
	rawData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	postPath := path.Join("threads", utils.EncodeURIComponent(url))
	err = dao.Write(fmt.Sprintf("%s.json", postPath), string(rawData))
	return err
}
