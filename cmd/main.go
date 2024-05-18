package main

import (
	"dex/internal/boost"
	"dex/internal/utils"
	"log"
	"os"
)

var (
	Config utils.Configuration
)

func init() {
	var err error
	Config, err = utils.ReadConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	wrappedConfig := boost.WrappedConfiguration{Configuration: &Config}
	_, err := wrappedConfig.Post("terminaluwubjuiebwbz4z9-session-sakc23gq-package-residential:cwz0eoj4sv5v@resi.proxies.fo:1337")
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}

}
