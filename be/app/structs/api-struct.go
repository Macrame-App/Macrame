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
	},
	Remote: []string{
		"/macro/list",
		"/device/access",
		"/device/auth",
	},
	Auth: []string{
		"/macro/play",
	},
}
