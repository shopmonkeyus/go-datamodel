package datatypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type StringArray []string

func (arr *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}
	err := json.Unmarshal(bytes, &arr)
	return err
}

var emptyStringArray = "[]"

func (arr StringArray) Value() (driver.Value, error) {
	if len(arr) == 0 {
		return emptyStringArray, nil
	}
	return json.Marshal(arr)
}

type NullableStringArray []string

func (arr *NullableStringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}
	err := json.Unmarshal(bytes, &arr)
	return err
}

func (arr NullableStringArray) Value() (driver.Value, error) {
	if len(arr) == 0 {
		return nil, nil
	}
	return json.Marshal(arr)
}

// GormDBDataType gorm db data type
func (StringArray) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return getJSONDataType(db)
}

// GormDBDataType gorm db data type
func (NullableStringArray) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return getJSONDataType(db)
}
