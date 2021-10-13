package model

import (
	"encoding/json"
	"errors"
	"github.com/pomment/backend-next/server/dao"
	"sort"
)

type Thread struct {
	Url          string `json:"url"`
	UUID         string `json:"uuid"`
	Title        string `json:"title"`
	FirstPostAt  int64  `json:"firstPostAt"`
	LatestPostAt int64  `json:"latestPostAt"`
	Amount       int64  `json:"amount"`
}

type ThreadList []Thread

func (p ThreadList) Len() int {
	return len(p)
}
func (p ThreadList) Less(i, j int) bool {
	return p[i].LatestPostAt < p[j].LatestPostAt
}
func (p ThreadList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var Threads = map[string]Thread{}
var ThreadReady = false

func InitThreadsForce() (err error) {
	rawData, err := dao.Read("index.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(rawData), &Threads)
	ThreadReady = true
	return err
}

func InitThreads() (err error) {
	err = InitThreadsForce()
	return err
}

func SaveThreads() (err error) {
	rawData, err := json.Marshal(Threads)
	if err != nil {
		return err
	}
	err = dao.Write("index.json", string(rawData))
	return err
}

func GetThread(url string) (data *Thread, err error) {
	err = InitThreads()
	if err != nil {
		return nil, err
	}
	out := Threads[url]
	return &out, nil
}

func GetThreads() (data *ThreadList, err error) {
	err = InitThreads()
	if err != nil {
		return nil, err
	}
	list := make(ThreadList, len(Threads))
	i := 0
	for _, v := range Threads {
		list[i] = v
		i++
	}
	sort.Sort(sort.Reverse(list))
	return &list, nil
}

func SetThread(url string, title string) (err error) {
	err = InitThreads()
	if err != nil {
		return err
	}

	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if data, ok := Threads[url]; ok {
		data.Title = title
		Threads[url] = data
		err = SaveThreads()
		return err
	}

	return errors.New("thread not found")
}

func UpdateThreadStatus(url string, latestPostAt int64, amount int64) (err error) {
	err = InitThreads()
	if err != nil {
		return err
	}

	if data, ok := Threads[url]; ok {
		data.LatestPostAt = latestPostAt
		data.Amount = amount
		Threads[url] = data
		err = SaveThreads()
		return err
	}

	return errors.New("thread not found")
}
