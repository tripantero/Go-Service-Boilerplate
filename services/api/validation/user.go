package validation

import (
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/softbrewery/gojoi/pkg/joi"
)

var schema = joi.Struct().Keys(joi.StructKeys{
	"Name": joi.String().Max(64).Min(4),
	"Age":  joi.Int().Min(4).Max(120),
})

func ValidateUser(user *models.User) error {
	return joi.Validate(*user, schema)
}
