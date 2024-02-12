package domainutils_test

import (
	"testing"

	domainutils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/utils"
	"github.com/stretchr/testify/suite"
)

type DomainGidGeneratorTestSuite struct {
	suite.Suite
	domainGidGenerator *domainutils.DomainGidGenerator
}

type testGenerateIfEmpty struct {
	name  string
	input string
}

func (suite *DomainGidGeneratorTestSuite) SetupSuite() {
	const PREFIX = "9999"
	suite.domainGidGenerator = domainutils.NewDomainGidGenerator(PREFIX)
}

func (suite *DomainGidGeneratorTestSuite) TestGenerateIfEmpty() {
	testCases := []testGenerateIfEmpty{
		{
			name:  "valid gid",
			input: "59eaffa9-cff9-4e5c-90e6-99ff6815ee6b",
		},
		{
			name:  "empty gid",
			input: "",
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.name, func() {
			gid := suite.domainGidGenerator.GenerateIfEmpty(&testCase.input)

			suite.Equal(len(*gid), 41)
		})
	}

}

func TestDomainGidGeneratorSuite(t *testing.T) {
	suite.Run(t, new(ServiceValidatorTestSuite))
}
