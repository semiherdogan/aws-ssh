package ui

import (
	"github.com/manifoldco/promptui"
	"os"
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
