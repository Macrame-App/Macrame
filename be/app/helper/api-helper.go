package helper

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strings"

	"be/app/structs"
	. "be/app/structs"
)

func EndpointAccess(w http.ResponseWriter, r *http.Request) (bool, string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)

	if err != nil {
		return false, "", errors.New(r.URL.Path + ": SplitHostPort error: " + err.Error())
	}

	if (isLocal(ip) && isEndpointAllowed("Local", r.URL.Path)) ||
		(isLanRemote(ip) && isEndpointAllowed("Remote", r.URL.Path)) {
		return true, "", nil
	} else if isLanRemote(ip) && isEndpointAllowed("Auth", r.URL.Path) {
		data, err := decryptAuth(r)

		if err != nil {
			return false, "", err
		}

		return data != "", data, nil
	}

	return false, "", errors.New(r.URL.Path + ": not authorized or accessible")
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
		return false
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

func decryptAuth(r *http.Request) (string, error) {
	var req structs.Authcall

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Uuid == "" || req.Data == "" {
		return "", errors.New("DecryptAuth Error: " + err.Error() + "; UUID: " + req.Uuid + "; Data: " + req.Data)
	}

	deviceKey, errKey := GetKeyByUuid(req.Uuid)
	decryptData, errDec := DecryptAES(deviceKey, req.Data)

	if errKey != nil && errDec != nil || decryptData == "" {
		return "", errors.New("DecryptAuth Error: " + errKey.Error() + "; " + errDec.Error() + "; UUID: " + req.Uuid + "; Data: " + req.Data)
	}

	return decryptData, nil
}

func ParseRequest(req interface{}, data string, r *http.Request) (d interface{}, err error) {
	if data != "" {
		dataBytes := []byte(data)
		return req, json.Unmarshal(dataBytes, &req)
	} else {
		return req, json.NewDecoder(r.Body).Decode(&req)
	}
}
