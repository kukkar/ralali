package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BoutiqaatREPO/getsetgo/src/core/middleware"
	"github.com/sanksons/gowraps/filesystem"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/BoutiqaatREPO/getsetgo/src/common/depchecker"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	figure "github.com/common-nighthawk/go-figure"
)

type Webserver struct {
}

func (ws Webserver) PreStart(a func(), b func()) {
	//BootStrap the Application
	a()
	initMgr := new(InitManager)
	initMgr.Execute()
	logger.Info(fmt.Sprintln("Web server Initialization done"))
	b()
}

//
// Access version info from version file.
//
func (ws Webserver) getVersionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//read version info.
		dataBytes, err := filesystem.GetFile("version.txt")
		if err != nil {
			msg := fmt.Sprintf("Could not read version info, %s", err.Error())
			w.Write([]byte(msg))
			return
		}
		w.Write(dataBytes)
	}
}

//
// Check for all dependencies.
//
func (ws Webserver) getDependencyCheckerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//read version info.
		result := depchecker.CheckDependencies()
		dataBytes, _ := json.Marshal(result)
		w.Write(dataBytes)
	}
}

//
// Define Handler chain.
//
func (ws Webserver) getHandler() http.HandlerFunc {
	// chain middlewares.
	// below funcs are executed in the order top 2 down sequentially.
	// credits: this technique is taken from: github.com/justinas/alice
	handlerFuncList := []func(http.HandlerFunc) http.HandlerFunc{
		middleware.EntryHandler,
		middleware.ProfilerHandler,
		middleware.RequestHeaderHandler,
		middleware.ExecuteHandler,
		middleware.ResponseHeaderHandler,
		middleware.GzipHandler,
		middleware.LoggerHandler,
	}

	var wrapper http.HandlerFunc = middleware.Flush
	noOfHandlers := len(handlerFuncList)
	for key := range handlerFuncList {
		wrapper = handlerFuncList[noOfHandlers-(1+key)](wrapper)
	}

	return wrapper
}

func (ws Webserver) Start() {

	http.HandleFunc("/", ws.getHandler())
	http.HandleFunc("/version.txt", ws.getVersionHandler())
	http.HandleFunc("/dependency-check", ws.getDependencyCheckerHandler())

	//Start the web server
	url := ":" + config.GlobalAppConfig.ServerPort

	logger.Info(fmt.Sprintln("Web server Starting......"))

	ws.displayConfigOnCli()
	ws.displayLogoOnCli()
	fmt.Printf("\nWeb server Starting......on port %s\n", config.GlobalAppConfig.ServerPort)

	serr := http.ListenAndServe(url, nil)
	if serr != nil {
		log.Printf("Could not start web server %s\n", serr.Error())
		logger.Error(fmt.Sprintln("Could not start web server ", serr))
	}
	if serr == nil {
		log.Printf("Web server Started on port %v\n", config.GlobalAppConfig.ServerPort)
		logger.Info(fmt.Sprintln("Web server Started on port : ", config.GlobalAppConfig.ServerPort))
	}

}

// swagger handler
func (ws Webserver) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}

func (ws Webserver) displayLogoOnCli() {
	file, _ := os.Open("conf/standard.flf")

	myFigure := figure.NewFigureWithFont(config.GlobalAppConfig.AppName, file, false)
	myFigure.Print()
}

func (ws Webserver) displayConfigOnCli() {
	fmt.Printf("\n\nUsing Configuration: \n")
	fmt.Printf("%s\n", config.GlobalAppConfig.ShowConfig())
}
