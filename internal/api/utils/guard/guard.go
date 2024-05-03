package guard

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"

	"github.com/quinntas/go-rest-template/internal/api/utils"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

func AgainstBadEmail(key string, value string) error {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return web.NewHttpError(
			http.StatusUnprocessableEntity,
			fmt.Sprintf("%s is not a valid email", key),
			utils.Map[interface{}]{
				"key": key,
			},
		)
	}
	return nil
}

func AgainstBadRegex(key string, value string, regex string) error {
	compiledRegex := regexp.MustCompile(regex)
	if result := compiledRegex.MatchString(value); !result {
		return web.NewHttpError(
			http.StatusUnprocessableEntity,
			fmt.Sprintf("%s is invalid", key),
			utils.Map[interface{}]{
				"key": key,
			},
		)
	}
	return nil
}

func AgainstBetween(key string, value interface{}, min int, max int) error {
	err := web.NewHttpError(
		http.StatusUnprocessableEntity,
		fmt.Sprintf("%s is not between %d and %d", key, min, max),
		utils.Map[interface{}]{
			"key": key,
		},
	)

	switch v := value.(type) {
	case string:
		if min > len(v) || len(v) > max {
			return err
		}
	case float32:
	case float64:
	case int:
		if min > v || v > max {
			return err
		}
	default:
		return web.NewHttpError(
			http.StatusUnprocessableEntity,
			fmt.Sprintf("%s is not string nor a number", key),
			utils.Map[interface{}]{
				"key": key,
			},
		)
	}
	return nil
}
