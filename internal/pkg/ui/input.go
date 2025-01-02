package ui

import (
	"os"

	"github.com/manifoldco/promptui"
)

func GetInput(s string) string {
	prompt := promptui.Prompt{
		Label:   s,
		Default: "",
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
	}

	return result
}
