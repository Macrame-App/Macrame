package structs

type Settings struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type RemoteWebhook struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
