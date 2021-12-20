package data

import (
	"github.com/google/uuid"
)

type Note struct {
	Context   Context     `json:"context" bson:"context"`
	Metadata  Metadata    `json:"metadata" bson:"metadata"`
	Entity    Entity      `json:"entity" bson:"entity"`
	Detail    interface{} `json:"detail" bson:"detail"`
	Policy    interface{} `json:"policy,omitempty" bson:"policy,omitempty"`
	Procedure interface{} `json:"procedure,omitempty" bson:"procedure,omitempty"`
}

type Context struct {
	Domain string `json:"domain" bson:"domain"`
	Group  string `json:"group" bson:"group"`
}

type Metadata struct {
	Type      string        `json:"type" bson:"type"`
	Path      string        `json:"path" bson:"path"`
	Platform  string        `json:"platform" bson:"platform"`
	URI       URI           `json:"uri" bson:"uri"`
	Timestamp string        `json:"timestamp" bson:"timestamp"`
	UUID      string        `json:"uuid" bson:"uuid"`
	Status    string        `json:"status" bson:"status"`
	Errors    []error       `json:"errors" bson:"errors"`
	Parents   string        `json:"parent" bson:"parent"`
}

type URI struct {
	API string `json:"api" bson:"api"`
	Web string `json:"web" bson:"web"`
}

type Entity struct {
	ID           int    `json:"id" bson:"id"`
	Component    string `json:"component" bson:"component"`
	Organization string `json:"organization" bson:"organization"`
	Version      string `json:"version" bson:"version"`
}

func (n *Note) WithContext(domain string, group string) *Note {

	n.Context = Context{
		Domain: domain,
		Group:  group,
	}

	return n
}

func (n *Note) WithMetadata(notetype string, path string, platform string, uri URI, timestamp string, status string, errors []error, parent string) *Note {

	n.Metadata = Metadata{
		Type:     notetype,
		Path:     path,
		Platform: platform,
		URI: URI{
			API: uri.API,
			Web: uri.Web,
		},
		Timestamp: timestamp,
		UUID:      uuid.New().String(),
		Status:    status,
		Errors:    errors,
		Parents:   parent,
	}

	return n
}

func (n *Note) WithEntity(id int, component string, organization string, version string) *Note {

	n.Entity = Entity{
		ID:           id,
		Component:    component,
		Organization: organization,
		Version:      version,
	}

	return n
}

func (n *Note) WithDetail(detail interface{}) *Note {

	n.Detail = detail

	return n
}

func (n *Note) WithPolicy(policy interface{}) *Note {

	n.Policy = policy

	return n
}

func (n *Note) WithProcedure(procedure interface{}) *Note {

	n.Procedure = procedure

	return n
}
