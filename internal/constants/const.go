package constants

import (
	http "github.com/bogdanfinn/fhttp"
)

var (
	HeaderC = http.Header{
		"Cache-Control":             {"max-age=0"},
		"Sec-Ch-Ua":                 {`"Chromium";v="8", "Not(A:Brand";v="24", "Google Chrome";v="120"`},
		"Sec-Ch-Ua-Mobile":          {"?0"},
		"Sec-Ch-Ua-Platform":        {`\"Windows\"`},
		"Upgrade-Insecure-Requests": {"1"},
		"User-Agent":                {`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36`},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"Sec-Fetch-Site":            {"same-origin"},
		"Sec-Fetch-Mode":            {"navigate"},
		"Sec-Fetch-User":            {"?1"},
		"Sec-Fetch-Dest":            {"document"},
		"Accept-Encoding":           {"gzip, deflate, br"},
		"Accept-Language":           {"en-GB,en-US;q=0.9,en;q=0.8"},
		"Priority":                  {"u=0, i"},
		"X-Requested-With":          {"XMLHttpRequest"},
		http.HeaderOrderKey:         {"cache-control", "sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "upgrade-insecure-requests", "user-agent", "accept", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-user", "sec-fetch-dest", "accept-encoding", "accept-language", "priority", "x-requested-with"},
	}
)
