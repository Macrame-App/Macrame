package structs

type Settings struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type RemoteWebhook struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type Check struct {
	Uuid string `json:"uuid"`
}

type Request struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Handshake struct {
	Uuid  string `json:"uuid"`
	Shake string `json:"shake"`
}

type Authcall struct {
	Uuid string `json:"uuid"`
	Data string `json:"d"`
}
