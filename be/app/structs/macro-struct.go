package structs

type MacroRequest struct {
	Macro string `json:"macro"`
}

type Step struct {
	Type      string `json:"type"`
	Key       string `json:"key"`
	Code      string `json:"code"`
	Location  int    `json:"location"`
	Direction string `json:"direction"`
	Value     int    `json:"value"`
}

type NewMacro struct {
	Name  string `json:"name"`
	Steps []Step `json:"steps"`
}

type MacroInfo struct {
	Name      string `json:"name"`
	Macroname string `json:"macroname"`
}
