package dbutils

import "math"

func GetPagesTotal(recordCount int, pageSize int) int {
	return int(math.Ceil(float64(recordCount) / float64(pageSize)))
}
