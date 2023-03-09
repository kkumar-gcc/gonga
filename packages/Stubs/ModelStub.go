package Stubs

import (
	"fmt"
	"strings"
)

func GetModelStub(modelName string) string {
	modelName = strings.Title(modelName)

	return fmt.Sprintf(`package Models

import (
    "gorm.io/gorm"
)

type %s struct {
    gorm.Model
    Name    string   `+"`gorm:\"not null\"`"+`
    Email   string   `+"`gorm:\"unique;not null\"`"+`
}

func (%s) TableName() string {
    return "%ss"
}`, modelName, modelName, strings.ToLower(modelName))


}
