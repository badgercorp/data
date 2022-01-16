package data

type Display struct {
	Badge       DisplayBadge       `json:"badge"`
	Description DisplayDescription `json:"description"`
}

type DisplayDescription struct {
	Content []DisplayDescriptionContent `json:"content"`
	Full    string                      `json:"full"`
	Links   []string                    `json:"links"`
	Lite    string                      `json:"lite"`
}

type DisplayDescriptionContent struct {
	Alt  string `json:"alt"`
	ID   string `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type DisplayBadge struct {
	Group string          `json:"group"`
	Tag   DisplayBadgeTag `json:"tag"`
	Title string          `json:"title"`
}

type DisplayBadgeTag struct {
	Error string `json:"error"`
	Fail  string `json:"fail"`
	Pass  string `json:"pass"`
}
