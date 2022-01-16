package data

type Display struct {
	Entity      Entity 			   `json:"entity,omitempty" bson:"entity"`
	Badge       DisplayBadge       `json:"badge,omitempty" bson:"badge"`
	Description DisplayDescription `json:"description,omitempty" bson:"description"`
}

type DisplayDescription struct {
	Content []DisplayDescriptionContent `json:"content" bson:"content"`
	Full    string                      `json:"full" bson:"full"`
	Links   []string                    `json:"links" bson:"links"`
	Lite    string                      `json:"lite" bson:"lite"`
}

type DisplayDescriptionContent struct {
	Alt  string `json:"alt" bson:"alt"`
	ID   string `json:"id" bson:"id"`
	Type string `json:"type" bson:"type"`
	URL  string `json:"url" bson:"url"`
}

type DisplayBadge struct {
	Group string          `json:"group" bson:"group"`
	Tag   DisplayBadgeTag `json:"tag" bson:"tag"`
	Title string          `json:"title" bson:"title"`
}

type DisplayBadgeTag struct {
	Error string `json:"error" bson:"error"`
	Fail  string `json:"fail" bson:"fail"`
	Pass  string `json:"pass" bson:"pass"`
}
