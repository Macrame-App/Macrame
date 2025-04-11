package structs

type PanelJSON struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	AspectRatio string            `json:"aspectRatio"`
	Macros      map[string]string `json:"macros"`
}

type PanelSaveJSON struct {
	// Name        string            `json:"name"`
	// Description string            `json:"description"`
	// AspectRatio string            `json:"aspectRatio"`
	// Macros      map[string]string `json:"macros"`
	Dir         string      `json:"dir"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	AspectRatio string      `json:"aspectRatio"`
	Macros      interface{} `json:"macros"`
}

type PanelInfo struct {
	Dir         string            `json:"dir"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	AspectRatio string            `json:"aspectRatio"`
	Macros      map[string]string `json:"macros"`
	Thumb       string            `json:"thumb"`
}

type PanelRequest struct {
	Dir string `json:"dir"`
}

type PanelResponse struct {
	Dir         string            `json:"dir"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	AspectRatio string            `json:"aspectRatio"`
	Macros      map[string]string `json:"macros"`
	Thumb       string            `json:"thumb"`
	HTML        string            `json:"html"`
	Style       string            `json:"style"`
}
