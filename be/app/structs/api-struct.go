package structs

type Allowed struct {
	Local  []string
	Remote []string
	Auth   []string
}

var Endpoints = Allowed{
	Local: []string{
		"/macro/record",
		"/macro/list",
		"/macro/delete",
		"/macro/play",
		"/device/list",
		"/device/access/check",
		"/device/access/request",
		"/device/link/ping",
		"/device/link/start",
		"/device/link/poll",
		"/device/link/remove",
		"/device/handshake",
	},
	Remote: []string{
		"/macro/list",
		"/device/access/check",
		"/device/access/request",
		"/device/link/ping",
		"/device/link/end",
		"/device/handshake",
		"/device/auth",
	},
	Auth: []string{
		"/macro/play",
		"/device/link/remove",
	},
}
