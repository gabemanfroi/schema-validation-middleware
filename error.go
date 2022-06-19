package schema_validation_middleware

type ValidationError struct {
	Field string
	Tag   string
	Value string
}
