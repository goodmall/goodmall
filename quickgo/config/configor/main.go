package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func main() {
	mainWithFlag()
}

func basicMain() {
	configor.Load(&Config, "./conf/config.yml")
	fmt.Printf("config: %#v", Config)
}

func mainWithFlag() {
	config := flag.String("file", "./conf/config.yml", "configuration file")
	flag.StringVar(&Config.APPName, "name", "", "app name")
	flag.StringVar(&Config.DB.Name, "db-name", "", "database name")
	flag.StringVar(&Config.DB.User, "db-user", "root", "database user")
	flag.Parse()

	os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	// Earlier configurations have higher priority
	configor.Load(&Config, *config)
	// configor.Load(&Config) // only load configurations from shell env & flag
	fmt.Printf("config: %#v", Config)
}
