package model

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/slices"
)

type UserType string

const (
	UserTypeSystem UserType = "SYSTEM"
	UserTypePerson UserType = "PERSON"
)

var UserTypes = []UserType{
	UserTypeSystem,
	UserTypePerson,
}

func (UserType) Enum() []interface{} {
	enums := []interface{}{}
	for _, element := range UserTypes {
		enums = append(enums, element)
	}
	return enums
}

func (userType *UserType) UnmarshalJSON(byteArray []byte) error {
	str := string(byteArray)
	if str == "null" {
		*userType = ""
		return nil
	}

	type _UserType UserType
	var stringValue *_UserType = (*_UserType)(userType)
	err := json.Unmarshal(byteArray, &stringValue)

	if err != nil {
		return err
	}

	if slices.Contains(UserTypes, *userType) {
		return nil
	}

	return fmt.Errorf("invalid user type: %s", *stringValue)
}
