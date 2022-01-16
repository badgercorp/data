package data

type Signature struct {
	UUID      string  `json:"uuid" bson:"uuid"`
	Key       string  `json:"key" bson:"key"`
	Signature string  `json:"signature" bson:"signature"`
	Status    string  `json:"status" bson:"status"`
	Errors    []error `json:"errors" bson:"errors"`
}
