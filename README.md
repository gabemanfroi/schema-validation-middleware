# GO Schema Validation Middleware

## Overview

---

Suppose that we have a fiber POST route that handles the creation of a user entity on our application:

```go
func UserRoutes(router fiber.Router){

router.Post("/myRoute", func create(){ // logic here})

}

```

And that we want to validate the request body based on the following Schema:

```go
type CreateUserValidator struct {
	FirstName string `validate:"required,min=3"`
	LastName  string `validate:"required,min=3"`
	Email     string `validate:"required,email"`
	Age       uint8  `validate:"omitempty,gte=0,lte=130"`
	City      string `validate:"omitempty,min=3"`
	Country   string `validate:"omitempty,min=3"`
	Pronouns  string `validate:"omitempty,min=3"`
	Password  string `validate:"required"`
}
```

Using the Validator middleware we can simply call the ValidateSchema function on our route like this:
```go

func UserRoutes(router fiber.Router){

router.Post("/myRoute", 
	func(c *fiber.Ctx) error {return ValidateSchema(c, CreateUserValidator{})} ,
	func create(){ // logic here})}

```

