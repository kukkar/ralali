package hello_test

import (
	"fmt"
	"testing"

	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"

	service "github.com/ralali/apitest/src/common"
	reflorestservice "github.com/BoutiqaatREPO/getsetgo/src/core/service"
	webserver "github.com/BoutiqaatREPO/getsetgo/src/testtools/fakers/webserver"
)

var testHTTPServer *webserver.TestWebserver

func TestHelloAPI(t *testing.T) {
	gm.RegisterFailHandler(gk.Fail)
	gk.RunSpecs(t, "Hello API TEST Suite")
}

var _ = gk.BeforeSuite(func() {
	//Set the APPlication to run in test mode
	reflorestservice.SetAppMode(reflorestservice.MODE_TEST)
	fmt.Println("Starting webserver")
	service.Register()

	//@todo: need to set init manager in reflorest so that its not needed to be set here explicitely.
	initMgr := new(reflorestservice.InitManager)
	initMgr.Execute()
	testHTTPServer = new(webserver.TestWebserver)
})

var _ = gk.AfterSuite(func() {
	//No Need to revert the App Mode.
	fmt.Println("Crashing webserver")
	testHTTPServer = nil
})
