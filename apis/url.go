package apis

import "net/url"

// HostWithPort returns the URL host, stripping the port if it is the
// default for the scheme (so detection-friendly hostnames are used for
// canonical URLs) while preserving non-default ports for self-hosted
// instances where the port is load-bearing for API calls.
func HostWithPort(parsedURL *url.URL) string {
	port := parsedURL.Port()
	if port == "" || isDefaultPort(parsedURL.Scheme, port) {
		return parsedURL.Hostname()
	}
	return parsedURL.Host
}

func isDefaultPort(scheme, port string) bool {
	switch scheme {
	case "https", "wss":
		return port == "443"
	case "http", "ws":
		return port == "80"
	case "ssh", "git+ssh":
		return port == "22"
	}
	return false
}
