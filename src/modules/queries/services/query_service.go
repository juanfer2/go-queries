package services

import (
	"juanfer2/go-queries/src/modules/queries/respositories"
	"juanfer2/go-queries/src/types"
)

func CreateQuery(inputQuery types.PayloadQuery) interface{} {
	data := respositories.CreateRepository(inputQuery)
	return data
}

func GetAllQueries() interface{} {
	qr := respositories.QueryRepository{}
	return qr.All()
}
