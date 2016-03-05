package models

import (
	"labix.org/v2/mgo/bson"
)

type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `from:"title" json:"title"`
	Description string        `from:"description" json:"description"`
	Price       float64       `from:"price" json:"price"`
}
