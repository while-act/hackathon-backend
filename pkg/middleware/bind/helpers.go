package bind

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var (
	NameRegexp     = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,18}([a-zA-Z0-9])$`)
	InnRegexp      = regexp.MustCompile(`^(\s+)?(\d{10})$`)
	LinkRegexp     = regexp.MustCompile(`^https?://(?:www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b[-a-zA-Z0-9()@:%_+.~#?&/=]*$`)
	TitleRegexp    = regexp.MustCompile(`^\p{L}+(?:([ \-']|(\. ))\p{L}+)*$`)
	EmailRegexp    = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	PasswordRegexp = regexp.MustCompile(`^\w{4,20}$`)
	UUID4          = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`)
)

func validateName(fl validator.FieldLevel) bool {
	return NameRegexp.MatchString(fl.Field().String())
}

func validateInn(fl validator.FieldLevel) bool {
	return InnRegexp.MatchString(fl.Field().String())
}

func validateLink(fl validator.FieldLevel) bool {
	return LinkRegexp.MatchString(fl.Field().String())
}

func validateTitle(fl validator.FieldLevel) bool {
	return TitleRegexp.MatchString(fl.Field().String())
}

func validateEmail(fl validator.FieldLevel) bool {
	return EmailRegexp.MatchString(fl.Field().String())
}

func validateUUID4(fl validator.FieldLevel) bool {
	return UUID4.MatchString(fl.Field().String())
}

func validateEnum(fl validator.FieldLevel) bool {
	enums := strings.Split(fl.Param(), "*")
	field := fl.Field().String()

	for _, v := range enums {
		if v == field {
			return true
		}
	}

	return false
}

func validatePassword(fl validator.FieldLevel) bool {
	return PasswordRegexp.MatchString(fl.Field().String())
}
