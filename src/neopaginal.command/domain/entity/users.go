package entity

type Users struct {
	Id                       string `json:"Id" bson:"_id"`
	Name                   string             `json:"Name,omitempty"`
}


