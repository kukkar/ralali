package config

import (
	"strings"
)

var allowedHosts []string

func getAllowedHosts() []string {
	if GlobalAppConfig.AllowedHosts == "" {
		return nil
	}
	return strings.Split(GlobalAppConfig.AllowedHosts, ",")
}

func RegisterAllowedHosts() {
	allowedHosts = getAllowedHosts()
}

func GetAllowedHosts() []string {
	return allowedHosts
}
