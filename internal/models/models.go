package models

type GetURLResponse struct {
	Comment     string `json:"comment"`
	Redirect    bool   `json:"redirect"`
	RedirectURI string `json:"redirectURI"`
	ShortURI    string `json:"shortURI"`
}

type NewURLParams struct {
	URL string `json:"url"`
}

type NewURLResponse struct {
	Message string `json:"message"`
	Shorted string `json:"shorted"`
}
