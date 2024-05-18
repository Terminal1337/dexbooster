package main

import (
	"dex/internal/boost"
	"dex/internal/utils"
	"dex/pkg/logging"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	Config        utils.Configuration
	wrappedConfig boost.WrappedConfiguration
	total_boosts  = 0
	mu            sync.Mutex
)

func init() {
	var err error
	Config, err = utils.ReadConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	wrappedConfig = boost.WrappedConfiguration{Configuration: &Config}

}
func checker() {
	for {
		if total_boosts >= wrappedConfig.Configuration.DexSettings.MaxBoostCount {
			fmt.Println("\n Max Boost Count Reached.... Force Quiting!")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}
}
func handler() {
	for {
		success, err := wrappedConfig.Post()
		if err != nil {
			logging.Logger.Error().
				Err(err).
				Msg("Boost Request Fail >> ")
		} else if success {
			mu.Lock()
			total_boosts += 1
			mu.Unlock()
			logging.Logger.Info().
				Str("coin", wrappedConfig.DexSettings.Coin).
				Str("type", wrappedConfig.DexSettings.BoostType).
				Str("id", wrappedConfig.DexSettings.ID).
				Int("boosts", total_boosts).
				Msg("Boosted")
		}
	}
}
func main() {

	if wrappedConfig.DexSettings.MaxBoost {
		fmt.Println("started")
		go checker()
	}
	for i := 0; i < wrappedConfig.Program.Threads; i++ {
		go handler()
	}

	time.Sleep(10000 * time.Hour)

}
