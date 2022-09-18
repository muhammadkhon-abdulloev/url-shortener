package models

type GetURLResponse struct {
	Comment     string `json:"comment"`
	Redirect    bool   `json:"redirect"`
	RedirectURI string `json:"redirectURI"`
	ShortURI    string `json:"shortURI"`
}

type NewURLParams struct {
	Shortened string
	URL       string `json:"url" validate:"required"`
}

type NewURLResponse struct {
	Message   string `json:"message"`
	Shortened string `db:"shortened" json:"shortened"`
}
