package helper

import (
	"log"
	"net"
	"net/http"
	"strings"

	. "be/app/structs"
)

func EndpointAccess(w http.ResponseWriter, r *http.Request) bool {
	log.Println("endpoint access")

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	if (isLocal(ip) && isEndpointAllowed("Local", r.URL.Path)) ||
		(isLanRemote(ip) && isEndpointAllowed("Remote", r.URL.Path)) {
		log.Println("accessible")
		return true
	} else if isLanRemote(ip) && isEndpointAllowed("auth", r.URL.Path) && isDeviceAuthorized() {
		log.Println("authorized")
	}

	log.Println(r.URL.Path, "not authorized or accessible")

	return false
}

func isLocal(ip string) bool {
	return ip == "127.0.0.1" || ip == "::1"
}

func isLanRemote(ip string) bool {
	return strings.HasPrefix(ip, "192.168.")
}

func isEndpointAllowed(source string, endpoint string) bool {
	var endpoints, err = getAllowedEndpoints(source)
	if err != "" {
		log.Println(err)
	}

	if (endpoints != nil) && (len(endpoints) > 0) {
		for _, e := range endpoints {
			if e == endpoint {
				return true
			}
		}
	}

	return false
}

func getAllowedEndpoints(source string) (endpoints []string, err string) {
	if source == "Local" {
		return Endpoints.Local, ""
	}
	if source == "Remote" {
		return Endpoints.Remote, ""
	}
	if source == "Auth" {
		return Endpoints.Auth, ""
	}

	return []string{}, "No allowed endpoints"
}

func isDeviceAuthorized() bool {
	return false
}
