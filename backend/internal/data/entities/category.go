package entities

import "time"

type Category struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
