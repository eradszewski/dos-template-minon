package main

import (
	_ "github.com/dimiro1/banner/autoload"
	"github.com/eradszewski/dos-template-minon/cmd/client"
	"github.com/eradszewski/dos-template-minon/cmd/domain"
	loggerI "github.com/eradszewski/dos-template-minon/cmd/logger"
	"github.com/eradszewski/dos-template-minon/cmd/server"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
)

var logger = loggerI.NewLogger()

func readFile(cfg *domain.Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		logger.Error(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		logger.Error(err)
	}
}

func readEnv(cfg *domain.Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		logger.Error(err)
	}
}

func main() {

	var config domain.Config
	readFile(&config)
	readEnv(&config)

	//isEnabled := true
	//isColorEnabled := true
	//banner.Init(os.Stdout, isEnabled, isColorEnabled, bytes.NewBufferString("My Custom Banner"))

	if config.Client.Enabled == true {
		client.Run(&config)

	} else {
		err := server.Run(&config)
		if err != nil {
			logger.Error(err)
		}
	}

}
