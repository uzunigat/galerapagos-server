package utils_test

import (
	utils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/utils"
	"github.com/stretchr/testify/suite"
)

// create suite
type IsProdTestSuite struct {
	suite.Suite
}

type testIsProd struct {
	name  string
	input string
}

func (suite *IsProdTestSuite) TestIsProd() {
	testCases := []testIsProd{
		{
			name:  "Is production",
			input: "production",
		},
		{
			name:  "Is not production",
			input: "dev",
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.name, func() {
			result := utils.IsProd(testCase.input)
			suite.Equal(result, testCase.input == "production")
		})
	}
}
