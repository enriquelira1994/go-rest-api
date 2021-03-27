package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Data struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Interruptor   string             `json:"interruptor,omitempty" bson:"interruptor,omitempty"`
	Dato  string             `json:"dato" bson:"dato,omitempty"`
}