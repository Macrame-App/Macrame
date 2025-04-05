package structs

type PanelJSON struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	AspectRatio string            `json:"aspectRatio"`
	Macros      map[string]string `json:"macros"`
}

type PanelInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	AspectRatio string            `json:"aspectRatio"`
	Macros      map[string]string `json:"macros"`
	Thumb       string            `json:"thumb"`
}
