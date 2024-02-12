package dbutils

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/uptrace/bun"
)

func UrlToDbPagination(dbQuery *bun.SelectQuery, pagination model.Pagination) {
	dbQuery.Limit(pagination.PageSize)
	dbQuery.Offset(pagination.PageSize * pagination.Page)
}
