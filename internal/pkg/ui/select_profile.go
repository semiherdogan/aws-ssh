package ui

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func SelectProfile(profiles []string) string {
	prompt := promptui.Select{
		Label:    "Profile",
		Items:    profiles,
		Size:     5,
		HideHelp: true,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Println("Cancelled.")
		os.Exit(1)
	}

	return profiles[i]
}
