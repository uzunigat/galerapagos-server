package domainutils

import (
	"fmt"

	"github.com/google/uuid"
)

type DomainGidGenerator struct {
	prefix string
}

func NewDomainGidGenerator(prefix string) *DomainGidGenerator {
	return &DomainGidGenerator{
		prefix: prefix,
	}
}

func (generator DomainGidGenerator) Generate() string {
	return fmt.Sprintf("%s.%s", generator.prefix, uuid.New().String())
}

func (generator DomainGidGenerator) GenerateIfEmpty(gid *string) *string {
	if gid == nil || *gid == "" {
		newGid := generator.Generate()
		return &newGid
	}
	return gid
}
