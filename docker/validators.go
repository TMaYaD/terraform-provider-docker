package docker

import (
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func validateIntegerInRange(min, max int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if value < min {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than %d: %d", k, min, value))
		}
		if value > max {
			errors = append(errors, fmt.Errorf(
				"%q cannot be higher than %d: %d", k, max, value))
		}
		return
	}
}

func validateIntegerGeqThan0() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if value < 0 {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than 0", k))
		}
		return
	}
}

func validateFloatRatio() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(float64)
		if value < 0.0 || value > 1.0 {
			errors = append(errors, fmt.Errorf(
				"%q has to be between 0.0 and 1.0", k))
		}
		return
	}
}

func validateDurationGeq0() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		dur, err := time.ParseDuration(value)
		if err != nil {
			errors = append(errors, fmt.Errorf(
				"%q is not a valid duration", k))
		}
		if dur < 0 {
			errors = append(errors, fmt.Errorf(
				"duration must not be negative"))
		}
		return
	}
}

func validateStringMatchesPattern(pattern string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		compiledRegex, err := regexp.Compile(pattern)
		if err != nil {
			errors = append(errors, fmt.Errorf(
				"%q regex does not compile", pattern))
			return
		}

		value := v.(string)
		if !compiledRegex.MatchString(value) {
			errors = append(errors, fmt.Errorf(
				"%q doesn't match the pattern (%q): %q",
				k, pattern, value))
		}

		return
	}
}