package utils

type DexSettings struct {
	Coin          string `json:"coin"`
	ID            string `json:"id"`
	BoostType     string `json:"boost_type"`
	MaxBoost      bool   `json:"max_boost"`
	MaxBoostCount int    `json:"max_boost_count"`
}

type Program struct {
	Threads   int    `json:"threads"`
	Proxies   string `json:"proxies"`
	ProxyType string `json:"proxy_type"`
}

type Configuration struct {
	DexSettings DexSettings `json:"dex_settings"`
	Program     Program     `json:"program"`
}
