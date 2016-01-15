package config

import (
	"github.com/memerelics/asana/utils"

	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"os"
)

type Conf struct {
	Personal_access_token string
	Workspace int
}

func Load() Conf {
	var dat []byte
	var err error
	dat, err = ioutil.ReadFile(utils.Home() + "/.asana.yml")
	if err != nil {
		fmt.Println("Config file isn't set.\n  ==> $ asana config")
		os.Exit(1)
	}
	conf := Conf{}
	err = yaml.Unmarshal(dat, &conf)
	utils.Check(err)
	return conf
}
