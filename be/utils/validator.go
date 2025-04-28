package utils

import "github.com/go-playground/validator/v10"

func FormatValidationError(err error) map[string]string {
	errs := make(map[string]string)
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := fe.Field()
			tag := fe.Tag()
			switch field {
			case "Username":
				if tag == "min" {
					errs["username"] = "username must be at least 3 characters"
				} else if tag == "required" {
					errs["username"] = "username is required"
				} else if tag == "max" {
					errs["username"] = "username must not exceed 16 characters"
				}

			case "Email":
				if tag == "required" {
					errs["email"] = "email is required"
				} else if tag == "email" {
					errs["email"] = "email must be a valid email address"
				}

			case "Password":
				if tag == "required" {
					errs["password"] = "password is required"
				} else if tag == "min" {
					errs["password"] = "password must be at least 6 characters"
				}

			case "AboutMessage":
				if tag == "max" {
					errs["about_message"] = "about message must not exceed 255 characters"
				}
			}
		}
	}

	return errs
}
