package cli

type Snapshot struct {
	Timestamp   string                 `json:"timestamp" bson:"timestamp"`
	Commit      Commit                 `json:"commit" bson:"commit"`
	Version     string                 `json:"version" bson:"version"`
	Repository  Repository             `json:"repository" bson:"repository"`
	Badgerfile  Badgerfile             `json:"badgerfile" bson:"badgerfile"`
	Environment map[string]interface{} `json:"environment" bson:"environment"`
}

type Author struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type Commit struct {
	ID        string   `json:"id" bson:"id"`
	Parent    string   `json:"parent" bson:"parent"`
	Tags      []string `json:"tags" bson:"tags"`
	Timestamp string   `json:"timestamp" bson:"timestamp"`
	Author    Author   `json:"author" bson:"author"`
}

type Repository struct {
	ID           string `json:"id"`
	Organization string `json:"organization" bson:"organization"`
	Name         string `json:"name" bson:"name"`
	Slug         string `json:"slug" bson:"slug"`
	APIURL       string `json:"apiurl" bson:"apiurl"`
	WebURL       string `json:"weburl" bson:"weburl"`
}

type Version struct {
	Manager string `json:"manager" bson:"manager"`
	Auto    bool   `json:"auto" bson:"auto"`
	Env     string `json:"env" bson:"env"`
}

type Properties struct {
	Group     string `json:"group" bson:"group"`
	Lob       string `json:"lob" bson:"lob"`
	Portfolio string `json:"portfolio" bson:"portfolio"`
	Product   string `json:"product" bson:"product"`
}

type Badgerfile struct {
	APIVersion   string     `json:"apiVersion" bson:"apiVersion"`
	Organization string     `json:"organization" bson:"organization"`
	Component    string     `json:"component" bson:"component"`
	Version      Version    `json:"version" bson:"version"`
	Properties   Properties `json:"properties" bson:"properties"`
}
