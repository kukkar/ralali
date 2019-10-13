package common

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/core/service"
	"github.com/ralali/apitest/src/common/appconfig"
	"github.com/ralali/apitest/src/common/appconstant"
	get "github.com/ralali/apitest/src/get"
	"github.com/ralali/apitest/src/getstats"
	"github.com/ralali/apitest/src/hello"
	post "github.com/ralali/apitest/src/post"
)

//main is the entry point of the florest web service

func StartServer() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)
	Register()
	webserver.PreStart(func() {}, func() {})
	webserver.Start()
}

func Register() {
	registerConfig()
	registerErrors()
	registerAllApis()
}

func registerAllApis() {
	service.RegisterAPI(new(hello.HelloAPI))
	service.RegisterAPI(new(post.ShotenPostAPI))
	service.RegisterAPI(new(get.ShortenAPI))
	service.RegisterAPI(new(getstats.ShortenStatsAPI))
}

func registerConfig() {
	service.RegisterConfig(new(appconfig.ApplicationConfig))
}

func registerErrors() {
	service.RegisterHTTPErrors(appconstant.APPErrorCodeToHTTPCodeMap)
}
