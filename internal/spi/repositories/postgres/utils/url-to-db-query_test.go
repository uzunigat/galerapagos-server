//go:build unit

package dbutils_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	dbutils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/repositories/postgres/utils"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

/*
The query interceptor attaches before and after query hooks in order to intercept the generated DB query before it's executed.
This allows us to test the expected vs. the actually generated SQL string.
*/
type QueryInterceptor struct {
	Query string
}

func (interceptor *QueryInterceptor) BeforeQuery(c context.Context, queryEvent *bun.QueryEvent) context.Context {
	interceptor.Query = queryEvent.Query
	return c
}

func (interceptor *QueryInterceptor) AfterQuery(c context.Context, queryEvent *bun.QueryEvent) {
}

type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusInactive        = "INACTIVE"
)

func TestUrlToDomainQuery(t *testing.T) {
	type urlQueryStruct struct {
		CountryCode string `db:"country_code"`
		Language    string `db:"language"`
		Status      Status `db:"status"`
	}
	type testData struct {
		description string
		urlQuery    urlQueryStruct
		expected    string
	}

	testCases := []testData{
		{
			"Generates correct db query when all string and non-string url parameters are set",
			urlQueryStruct{
				CountryCode: "US",
				Language:    "en-US",
				Status:      StatusActive,
			},
			"SELECT * FROM test WHERE (country_code = 'US') AND (language = 'en-US') AND (status = 'ACTIVE')",
		},
		{
			"Generates correct db query when only some url parameters are set",
			urlQueryStruct{
				CountryCode: "US",
			},
			"SELECT * FROM test WHERE (country_code = 'US')",
		},
		{
			"Generates correct db query for url parameters with multiple values",
			urlQueryStruct{
				CountryCode: "US,CA,CH",
			},
			"SELECT * FROM test WHERE (country_code IN ('US', 'CA', 'CH'))",
		},
	}

	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	queryInterceptor := QueryInterceptor{}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(&queryInterceptor)
	defer db.Close()

	for _, testCase := range testCases {
		fmt.Print(testCase.description)
		dbQuery := db.NewSelect().ColumnExpr("*").TableExpr("test")
		dbutils.UrlToDbQuery(dbQuery, testCase.urlQuery)
		dbQuery.Exec(context.TODO())
		actual := queryInterceptor.Query
		passed := assert.Equal(t, testCase.expected, actual)
		if passed {
			fmt.Println(" -- passed")
		} else {
			fmt.Println(" -- failed")
		}
	}
}
