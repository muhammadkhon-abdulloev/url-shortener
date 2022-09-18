package repository

const (
	queryGetLongURL   = `select url from shortener.urls where shortened = $1;`
	queryIsExist      = `select shortened from shortener.urls where url = $1;`
	queryNewShortened = `insert into shortener.urls (shortened, url) values ($1, $2) returning shortened;`
)
