package model

type GidGenerator interface {
	Generate() string
	GenerateIfEmpty(gid *string) *string
}
