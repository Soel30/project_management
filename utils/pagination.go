package utils

import (
	models "pm/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GenerateUserPagination(c *gin.Context) models.PaginationUser {

	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}
	return models.PaginationUser{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
