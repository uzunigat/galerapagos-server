package domainutils_test

import (
	"testing"

	domainutils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/utils"
	testhelpers "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/test-helpers"
	"github.com/stretchr/testify/suite"
)

type ServiceValidatorTestSuite struct {
	suite.Suite
	validator *domainutils.ServiceValidator
}

type testValidate struct {
	name     string
	input    interface{}
	expected error
}

func (suite *ServiceValidatorTestSuite) SetupSuite() {
	suite.validator = domainutils.NewServiceValidator()
}

func (suite *ServiceValidatorTestSuite) TestValidate() {
	validPlayer := testhelpers.ExamplePlayer()

	testCases := []testValidate{
		{
			name:     "valid player",
			input:    validPlayer,
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.name, func() {
			err := suite.validator.ValidateStruct(testCase.input)

			suite.Equal(testCase.expected, err)
		})
	}
}

func TestServiceValidatorSuite(t *testing.T) {
	suite.Run(t, new(ServiceValidatorTestSuite))
}
