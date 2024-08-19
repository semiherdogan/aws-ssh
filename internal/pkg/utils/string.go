package utils

import (
	"fmt"
	"strings"
)

func CapitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}

	return fmt.Sprintf("%s%s", strings.ToUpper(input[0:1]), input[1:])
}
