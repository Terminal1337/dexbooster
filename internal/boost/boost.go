package boost

import (
	"dex/internal/constants"
	"fmt"
	"io/ioutil"
	"log"

	http "github.com/bogdanfinn/fhttp"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func (wc *WrappedConfiguration) Post(proxy string) (bool, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_120),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithServerNameOverwrite(fmt.Sprintf("io.dexscreener.com")),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return false, err
	}
	Payload := fmt.Sprintf(`{"emoji":"%s"}`, wc.Configuration.DexSettings.BoostType)
	fmt.Println(Payload)
	req, err := http.NewRequest(http.MethodGet, "https://dexscreener.com/hype/reactions/dexPair/bsc:0xEb555C4ca3Ba9C0Ae7f9dfBcd9383CaEBEA6C8b9", nil)
	if err != nil {
		log.Println(err)
		return false, err
	}
	req.Header = constants.HeaderC
	req.Header.Set("Origin", "https://dexscreener.com")
	for {
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(resp.StatusCode)
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	}
	return false, nil

}
