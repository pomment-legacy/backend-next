package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	APIHost   string `yaml:"apiHost"`
	APIPort   int    `yaml:"apiPort"`
	APIURL    string `yaml:"apiURL"`
	SiteAdmin struct {
		Name     string `yaml:"name"`
		Email    string `yaml:"email"`
		Password string `yaml:"password"`
	} `yaml:"siteAdmin"`
	ReCAPTCHA struct {
		Enabled      bool    `yaml:"enabled"`
		SecretKey    string  `yaml:"secretKey"`
		MinimumScore float64 `yaml:"minimumScore"`
	} `yaml:"reCAPTCHA"`
	GuestNotify struct {
		Mode          string `yaml:"mode"`
		Title         string `yaml:"title"`
		SMTPSender    string `yaml:"smtpSender"`
		SMTPHost      string `yaml:"smtpHost"`
		SMTPPort      int    `yaml:"smtpPort"`
		SMTPUsername  string `yaml:"smtpUsername"`
		SMTPPassword  string `yaml:"smtpPassword"`
		SMTPSecure    bool   `yaml:"smtpSecure"`
		MailgunAPIKey string `yaml:"mailgunAPIKey"`
		MailgunDomain string `yaml:"mailgunDomain"`
	} `yaml:"guestNotify"`
	Webhook struct {
		Enabled bool     `yaml:"enabled"`
		Targets []string `yaml:"targets"`
	} `yaml:"webhook"`
}

var Content = Config{}

func Read(basePath string) (err error) {
	absPath := filepath.Join(basePath, "config.yaml")
	jsonFile, err := os.Open(absPath)
	defer jsonFile.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Content)
	return err
}
