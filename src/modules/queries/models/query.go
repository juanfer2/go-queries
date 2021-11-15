package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"gorm.io/gorm"
)

type Query struct {
	ID          uint   `json:"id" db:"id" gorm:"primary_key"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`

	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" db:"deleted_at" gorm:"index"`
}

func (query Query) Type() {
	fmt.Println(reflect.TypeOf(query))
}

func (query Query) TableName() string {
	return "queries"
}

func (query *Query) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

type Queries []Query

func (queriesList Queries) String() string {
	q, _ := json.Marshal(queriesList)
	return string(q)
}
