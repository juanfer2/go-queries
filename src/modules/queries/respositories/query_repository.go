package respositories

import (
	databases "juanfer2/go-queries/src/database"
	"juanfer2/go-queries/src/modules/queries/models"
	"juanfer2/go-queries/src/types"
)

type QueryRepository struct {
	models.Query
}

func NewQuery(title string, description string) types.IRepository {
	return &QueryRepository{
		Query: models.Query{
			Title:       title,
			Description: description,
		},
	}
}

func CreateRepository(input types.PayloadQuery) interface{} {
	db := databases.Conn()

	newQuery := models.Query{
		Title:       input.Title,
		Description: input.Description,
	}

	db.Create(&newQuery)

	return newQuery
}

func (queryRepository QueryRepository) Create(input models.Query) interface{} {
	db := databases.Conn()
	newQuery := input
	db.Create(&newQuery)

	return newQuery
}

func (queryRepository *QueryRepository) New() interface{} {
	return true
}

func (queryRepository QueryRepository) All() interface{} {
	db := databases.Conn()
	var list []models.Query
	db.Find(&list)

	if list == nil {
		return []string{}
	}

	return list
}
