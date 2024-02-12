package utils_test

import (
	"testing"

	utils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/utils"
)

func TestContains(t *testing.T) {
	t.Run("should return true if element is in array", func(t *testing.T) {
		array := []string{"foo", "bar", "baz"}
		element := "bar"
		result := utils.ArrayContains(array, element)
		if result != true {
			t.Errorf("expected %v, but got %v", true, result)
		}
	})
}

func TestNoContains(t *testing.T) {
	t.Run("should return false if element is not in array", func(t *testing.T) {
		array := []string{"foo", "bar", "baz"}
		element := "qux"
		result := utils.ArrayContains(array, element)
		if result != false {
			t.Errorf("expected %v, but got %v", false, result)
		}
	})
}
