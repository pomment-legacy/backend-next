package model

import (
	"encoding/json"
	"github.com/pomment/backend-next/server/dao"
	"sort"
)

type Thread struct {
	Url          string `json:"url"`
	UUID         string `json:"uuid"`
	Title        string `json:"title"`
	FirstPostAt  int64  `json:"firstPostAt"`
	LatestPostAt int64  `json:"latestPostAt"`
	Amount       int    `json:"amount"`
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
