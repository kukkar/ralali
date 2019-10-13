package appconfig

import (
	"errors"
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
)

type ApplicationConfig struct {
	TestConfig TestConfig
}

type TestConfig struct {
	Alpha int
	Beta  int
}

func GetApplicationConfig() (*ApplicationConfig, error) {
	c := config.GlobalAppConfig.ApplicationConfig
	appConfig, ok := c.(*ApplicationConfig)
	if !ok {
		msg := fmt.Sprintf("ApplicationConfig Not correct %+v", c)
		logger.Error(msg)
		return nil, errors.New(msg)
	}
	return appConfig, nil
}

func (this *ApplicationConfig) DisplayOnCli() {

	fmt.Println("Following is the Application Config:")

	fmt.Printf(
		"TestConfig      := %s\n",
		this.TestConfig,
	)

}
