package schema_validation_middleware

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func ValidateSchema(c *fiber.Ctx, schema interface{}) error {
	var errors []*ValidationError
	body := reflect.New(reflect.TypeOf(schema)).Interface()

	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e ValidationError
			e.Field = err.Field()
			e.Tag = err.Tag()
			e.Value = err.Param()
			errors = append(errors, &e)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
