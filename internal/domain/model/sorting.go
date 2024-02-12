package model

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var SortDirections = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (SortDirection) Enum() []interface{} {
	enums := []interface{}{}
	for _, element := range SortDirections {
		enums = append(enums, element)
	}
	return enums
}

type Sorting struct {
	SortBy *string        `json:"sortBy" query:"sortBy"`
	Sort   *SortDirection `json:"sort" query:"sort" validate:"omitempty,oneof=ASC DESC" `
}
