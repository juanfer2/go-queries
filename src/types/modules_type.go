package types

import (
	"juanfer2/go-queries/src/modules/queries/models"
)

type IRepository interface {
	Create(input models.Query) interface{}
	New() interface{}
	All() interface{}
}

type PayloadQuery struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
