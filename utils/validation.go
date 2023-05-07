package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validation = map[string]interface{}{
	"accepted" : "The :attribute must be accepted.",
    "accepted_if" : "The :attribute must be accepted when :other is :value.",
    "active_url" : "The :attribute is not a valid URL.",
    "after" : "The :attribute must be a date after :date.",
    "after_or_equal" : "The :attribute must be a date after or equal to :date.",
    "alpha" : "The :attribute must only contain letters.",
    // "alpha_dash" : "The :attribute must only contain letters, numbers, dashes and underscores.",
    "alphanum" : "The :attribute must only contain letters and numbers.",
    "array" : "The :attribute must be an array.",
    "before" : "The :attribute must be a date before :date.",
    "before_or_equal" : "The :attribute must be a date before or equal to :date.",
    "between" : map[string]string{
        "array" : "The :attribute must have between :min and :max items.",
        "file" : "The :attribute must be between :min and :max kilobytes.",
        "numeric" : "The :attribute must be between :min and :max.",
        "string" : "The :attribute must be between :min and :max characters.",
	},
    "boolean" : "The :attribute field must be true or false.",
    "confirmed" : "The :attribute confirmation does not match.",
    "current_password" : "The password is incorrect.",
    "date" : "The :attribute is not a valid date.",
    "date_equals" : "The :attribute must be a date equal to :date.",
    "datetime" : "The :attribute does not match the format :format.",
    "declined" : "The :attribute must be declined.",
    "declined_if" : "The :attribute must be declined when :other is :value.",
    "different" : "The :attribute and :other must be different.",
    "digits" : "The :attribute must be :digits digits.",
    "digits_between" : "The :attribute must be between :min and :max digits.",
    "dimensions" : "The :attribute has invalid image dimensions.",
    "distinct" : "The :attribute field has a duplicate value.",
    "endsnotwith" : "The :attribute may not end with one of the following: :values.",
    "startsnotwith" : "The :attribute may not start with one of the following: :values.",
    "email" : "The :attribute must be a valid email address.",
    "endswith" : "The :attribute must end with one of the following: :values.",
    "enum" : "The selected :attribute is invalid.",
    "exists" : "The selected :attribute is invalid.",
    "file" : "The :attribute must be a file.",
    "filled" : "The :attribute field must have a value.",
    "gtfield" : map[string]string{
        "array" : "The :attribute must have more than :value items.",
        "file" : "The :attribute must be greater than :value kilobytes.",
        "numeric" : "The :attribute must be greater than :value.",
        "string" : "The :attribute must be greater than :value characters.",
	},
    "gtefield" : map[string]string{
        "array" : "The :attribute must have :value items or more.",
        "file" : "The :attribute must be greater than or equal to :value kilobytes.",
        "numeric" : "The :attribute must be greater than or equal to :value.",
        "string" : "The :attribute must be greater than or equal to :value characters.",
	},
    "image" : "The :attribute must be an image.",
    "in" : "The selected :attribute is invalid.",
    "in_array" : "The :attribute field does not exist in :other.",
    "integer" : "The :attribute must be an integer.",
    "ip" : "The :attribute must be a valid IP address.",
    "ipv4" : "The :attribute must be a valid IPv4 address.",
    "ipv6" : "The :attribute must be a valid IPv6 address.",
    "json" : "The :attribute must be a valid JSON string.",
    "ltfield" : map[string]string{
        "array" : "The :attribute must have less than :value items.",
        "file" : "The :attribute must be less than :value kilobytes.",
        "numeric" : "The :attribute must be less than :value.",
        "string" : "The :attribute must be less than :value characters.",
	},
    "ltefield" : map[string]string{
        "array" : "The :attribute must not have more than :value items.",
        "file" : "The :attribute must be less than or equal to :value kilobytes.",
        "numeric" : "The :attribute must be less than or equal to :value.",
        "string" : "The :attribute must be less than or equal to :value characters.",
	},
    "mac" : "The :attribute must be a valid MAC address.",
    "max" : map[string]string{
        "array" : "The :attribute must not have more than :max items.",
        "file" : "The :attribute must not be greater than :max kilobytes.",
        "numeric" : "The :attribute must not be greater than :max.",
        "string" : "The :attribute must not be greater than :max characters.",
	},
    "max_digits" : "The :attribute must not have more than :max digits.",
    "mimes" : "The :attribute must be a file of type: :values.",
    "mimetypes" : "The :attribute must be a file of type: :values.",
    "min" : map[string]string{
        "array" : "The :attribute must have at least :min items.",
        "file" : "The :attribute must be at least :min kilobytes.",
        "numeric" : "The :attribute must be at least :min.",
        "string" : "The :attribute must be at least :min characters.",
	},
    "min_digits" : "The :attribute must have at least :min digits.",
    "multiple_of" : "The :attribute must be a multiple of :value.",
    "not_in" : "The selected :attribute is invalid.",
    "not_regex" : "The :attribute format is invalid.",
    "numeric" : "The :attribute must be a number.",
    "password" : map[string]string{
        "letters" : "The :attribute must contain at least one letter.",
        "mixed" : "The :attribute must contain at least one uppercase and one lowercase letter.",
        "numbers" : "The :attribute must contain at least one number.",
        "symbols" : "The :attribute must contain at least one symbol.",
        "uncompromised" : "The given :attribute has appeared in a data leak. Please choose a different :attribute.",
	},
    "present" : "The :attribute field must be present.",
    // "prohibited" : "The :attribute field is prohibited.",
    // "prohibited_if" : "The :attribute field is prohibited when :other is :value.",
    // "prohibited_unless" : "The :attribute field is prohibited unless :other is in :values.",
    // "prohibits" : "The :attribute field prohibits :other from being present.",
    // "regex" : "The :attribute format is invalid.",
    "required" : "The :attribute field is required.",
    // "required_array_keys" : "The :attribute field must contain entries for: :values.",
    "required_if" : "The :attribute field is required when :other is :value.",
    "required_unless" : "The :attribute field is required unless :other is in :values.",
    "required_with" : "The :attribute field is required when :values is present.",
    "required_with_all" : "The :attribute field is required when :values are present.",
    "required_without" : "The :attribute field is required when :values is not present.",
    "required_without_all" : "The :attribute field is required when none of :values are present.",
    "eqcsfield" : "The :attribute and :other must match.",
    "size" : map[string]string{
        "array" : "The :attribute must contain :size items.",
        "file" : "The :attribute must be :size kilobytes.",
        "numeric" : "The :attribute must be :size.",
        "string" : "The :attribute must be :size characters.",
	},
    "startswith" : "The :attribute must start with one of the following: :values.",
    "string" : "The :attribute must be a string.",
    "timezone" : "The :attribute must be a valid timezone.",
    "unique" : "The :attribute has already been taken.",
    // "uploaded" : "The :attribute failed to upload.",
    "url" : "The :attribute must be a valid URL.",
    "uuid" : "The :attribute must be a valid UUID.",
	"uuid3" : "The :attribute must be a valid UUID V3.",
	"uuid4" : "The :attribute must be a valid UUID V4.",
	"uuid3_rfc4122" : "The :attribute must be a valid UUID v3 RFC4122.",
}

// ValidationMessage returns a formatted error message for the given validator.FieldError.
// It replaces placeholders in the error message with values from the validation error.
//
// Parameters:
//   - e (validator.FieldError): The validation error to format.
//
// Returns:
//  - message (string): The formatted error message for the given validation error.
func ValidationMessage(e validator.FieldError) string {
	messages := Validation
	// replace placeholders with values
    attribute := e.Field()
    value := e.Param()

    message, ok := messages[e.Tag()].(string)
    if !ok {
        nested, ok := messages[e.Tag()].(map[string]string)
        if !ok {
            return ""
        }
        message, ok = nested[e.Type().Name()]
        if !ok {
            return ""
        }
    }

    message = strings.Replace(message, ":attribute", attribute, -1)
    message = strings.Replace(message, ":other", e.Param(), -1)
    message = strings.Replace(message, ":value", value, -1)
    message = strings.Replace(message, ":min", e.Param(), -1)
    message = strings.Replace(message, ":max", e.Param(), -1)
    message = strings.Replace(message, ":date", e.Param(), -1)

	return message
}