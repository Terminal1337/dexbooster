package boost

import (
	"fmt"
	"strings"

	http "github.com/bogdanfinn/fhttp"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func (wc *WrappedConfiguration) Post() (bool, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_124),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithProxyUrl(fmt.Sprintf("%s://%s", wc.Program.ProxyType, wc.Program.Proxies)),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://io.dexscreener.com/hype/reactions/dexPair/%s:%s", wc.DexSettings.Coin, wc.DexSettings.ID), strings.NewReader(fmt.Sprintf(`{"emoji":"%s"}`, wc.Configuration.DexSettings.BoostType)))
	if err != nil {
		return false, err
	}

	req.Header = http.Header{
		"Accept":             {"*/*"},
		"Accept-Language":    {"en-GB,en;q=0.8"},
		"Content-Type":       {"application/json"},
		"Origin":             {"https://dexscreener.com"},
		"Priority":           {"u=1, i"},
		"Referer":            {"https://dexscreener.com/"},
		"Sec-Ch-Ua":          {`"Brave";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`},
		"Sec-Ch-Ua-Mobile":   {"?0"},
		"Sec-Ch-Ua-Platform": {`"macOS"`},
		"Sec-Fetch-Dest":     {"empty"},
		"Sec-Fetch-Mode":     {"cors"},
		"Sec-Fetch-Site":     {"same-site"},
		"Sec-Gpc":            {"1"},
		"User-Agent":         {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	}

	return false, nil

}
