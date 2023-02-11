package model

type Approve struct {
	ID   string `json:"id" bson:"_id"`
	IP   string `json:"ip" bson:"ip"`
	Name string `json:"name" bson:"name"`
}
