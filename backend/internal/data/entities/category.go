package entities

import "time"

type Category struct {
	ID        string    `bson:"id"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
