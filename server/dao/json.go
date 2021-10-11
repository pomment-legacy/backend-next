package dao

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	BasePath = ""
)

func Init(p string) {
	BasePath = p
}

func Read(p string)(content string, err error) {
	absPath := filepath.Join(BasePath, p)
	jsonFile, err := os.Open(absPath)
	defer jsonFile.Close()
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Write(p string, content string)(err error) {
	absPath := filepath.Join(BasePath, p)
	jsonFile, err := os.Create(absPath)
	defer jsonFile.Close()
	if err != nil {
		return err
	}

	n, _ := jsonFile.Seek(0, io.SeekEnd)
	_, err = jsonFile.WriteAt([]byte(content), n)
	return err
}
