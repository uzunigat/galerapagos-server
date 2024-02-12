package dbutils

import "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"

func GetDefaultPagination(pagination model.Pagination) model.Pagination {
	defaultPagination := model.Pagination{
		Page:     0,
		PageSize: 10,
	}
	if pagination.Page > 0 {
		defaultPagination.Page = pagination.Page
	}
	if pagination.PageSize > 10 {
		defaultPagination.PageSize = pagination.PageSize
	}
	return defaultPagination
}
