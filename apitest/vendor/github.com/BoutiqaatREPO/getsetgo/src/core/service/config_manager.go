package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/env"
)

const DefaultConfFile = "conf/conf.json"
const TestConfFile = "../conf/conf.json"

type ConfigManager struct {
	ConfFile string
}

func (cm *ConfigManager) GetConfFile() string {
	fmt.Println("APP mode is " + AppMode)
	if AppMode == MODE_PROD {
		return DefaultConfFile
	}
	return TestConfFile

}

func (cm *ConfigManager) InitializeGlobalConfig(confFile string) {

	cm.Initialize(confFile, config.GlobalAppConfig)
	//log.Printf("Global Config is listed below: \n%+v", config.GlobalAppConfig)
	//log.Printf("Application Config is listed below: \n%+v", config.GlobalAppConfig.ApplicationConfig)
}

func (cm *ConfigManager) Initialize(filePath string, conf interface{}) {
	//log.Println("Initializing Application Config")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error loading App Config file %s \n %s", filePath, err))
	}
	err = json.Unmarshal(file, conf)
	if err != nil {
		panic(fmt.Sprintf("Incorrect Json in %s \n %s", filePath, err))
	}
	//log.Println("Application Config Initialized")
}

// UpdateConfigFromEnv updates provided config from environment variables
func (cm *ConfigManager) UpdateConfigFromEnv(conf interface{}, ty string) {
	if conf == nil {
		return
	}
	localConfigMap := make(map[string]string)
	if ty == "global" {
		if globalEnvUpdateMap == nil {
			log.Println("Global Environment variables Not Found")
			return
		}
		localConfigMap = globalEnvUpdateMap
	} else {
		if configEnvUpdateMap == nil {
			log.Println("Application Environment variables Not Found")
			return
		}
		localConfigMap = configEnvUpdateMap
	}

	configEnvUpdateValuesMap := make(map[string]string)
	for k, v := range localConfigMap {
		updatedVal, envValfound := env.GetOsEnviron().Get(v)

		if !envValfound {
			fmt.Printf("\n>> Environment variable %s not found\n", v)
			continue
		}

		if strings.Trim(updatedVal, " ") == "" {
			fmt.Printf("\n>> Environment variable %s is empty\n", v)
			continue
		}

		configEnvUpdateValuesMap[k] = updatedVal
	}

	byt, _ := json.Marshal(conf)

	newbyt, juerr := cm.updateJSONPath(configEnvUpdateValuesMap, byt, ".")
	if juerr != nil {
		panic(juerr)
	}

	if uerr := json.Unmarshal(newbyt, &conf); uerr != nil {
		panic(uerr)
	}
	if ty == "global" {
		//log.Printf("Updated config from environment variables: %+v\n", config.GlobalAppConfig)
	}
}

func (cm *ConfigManager) updateJSONPath(queries map[string]string, byt []byte, pathSep string) (newByt []byte, err error) {
	unMarshallObj := make(map[string]interface{})
	jerr := json.Unmarshal(byt, &unMarshallObj)
	if jerr != nil {
		return byt, jerr
	}

	for query, newNodeVal := range queries {
		path := strings.Split(query, pathSep)
		var v map[string]interface{}

		jsPath := unMarshallObj
		for _, node := range path {

			nextJsPath, found := jsPath[node]
			if !found {
				return byt, fmt.Errorf("Not found node %s in path", node)
			}
			v = jsPath
			jsPath, _ = nextJsPath.(map[string]interface{})

		}

		leafNode := path[len(path)-1]

		var newNodeValConv interface{}
		var convErr error

		switch v[leafNode].(type) {
		case float64:
			newNodeValConv, convErr = strconv.ParseFloat(newNodeVal, 64)
		case int64:
			newNodeValConv, convErr = strconv.ParseInt(newNodeVal, 10, 64)
		case uint64:
			newNodeValConv, convErr = strconv.ParseUint(newNodeVal, 10, 64)
		case string:
			newNodeValConv, convErr = newNodeVal, nil
		case bool:
			newNodeValConv, convErr = strconv.ParseBool(newNodeVal)
		default:
			newNodeValConv, convErr = nil, errors.New("Unsupported json value for json xpath update")
		}
		if convErr != nil {
			return byt, convErr
		}
		v[leafNode] = newNodeValConv
	}
	return json.Marshal(unMarshallObj)
}
