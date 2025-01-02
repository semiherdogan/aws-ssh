package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/semiherdogan/aws-ssh/internal/pkg/aws"
)

func SelectRegion(regions []aws.Region) aws.Region {
	template := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .Name | cyan }} ({{ .Value | cyan }})",
		Inactive: "  {{ .Name | faint }} ({{ .Value | faint }})",
		Selected: "--> {{ .Name | cyan }} ({{ .Value | cyan }})",
	}

	searcher := func(input string, index int) bool {
		region := regions[index]
		name := strings.Replace(strings.ToLower(region.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(
			strings.ToLower(region.Value),
			input,
		)
	}

	prompt := promptui.Select{
		Label:             "Region:",
		Items:             regions,
		Templates:         template,
		Size:              8,
		Searcher:          searcher,
		HideHelp:          true,
		StartInSearchMode: len(regions) > 5,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Println("Cancelled.")
		os.Exit(1)
	}

	return regions[i]
}
