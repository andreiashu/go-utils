package url

import (
	"net/url"
	"strings"
)

// UrlIsInternal will return true for URLs
func UrlIsInternal(u string, host string) bool {
	// exit early if the path is relative
	if strings.HasPrefix(u, "/") {
		return true
	}

	parsedUrl, err := url.Parse(u)
	if err != nil {
		return false
	}
	if strings.IndexByte(parsedUrl.Host, ":"[0]) == -1 {
		parsedUrl.Host = parsedUrl.Host + ":80"
	}

	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}
	parsedHost, err := url.Parse(host)
	if err != nil {
		return false
	}
	if strings.IndexByte(parsedHost.Host, ":"[0]) == -1 {
		parsedHost.Host = parsedHost.Host + ":80"
	}

	if !strings.HasPrefix(parsedUrl.Host, parsedHost.Host) ||
		!strings.HasPrefix(parsedHost.Host, parsedUrl.Host) {
		return false
	}

	return true
}
