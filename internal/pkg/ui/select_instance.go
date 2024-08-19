package ui

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/semiherdogan/aws-ssh/internal/pkg/aws"
	"os"
	"strings"
)

func SelectInstance(instances []aws.Instance) aws.Instance {
	searcher := func(input string, index int) bool {
		instance := instances[index]
		name := strings.Replace(strings.ToLower(instance.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:             fmt.Sprintf("Select Instance (total: %d)", len(instances)),
		Items:             instances,
		Size:              13,
		HideHelp:          true,
		Searcher:          searcher,
		StartInSearchMode: len(instances) > 8,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "-> {{ .Name | cyan }} ({{ .State | cyan }}) ({{ .LaunchTime | cyan }})",
			Inactive: "  {{ .Name | faint }} ({{ .State | faint }}) ({{ .LaunchTime | faint }})",
			Selected: "--> {{ .Name | cyan }} ({{ .InstanceId | cyan }}) ({{ .LaunchTime | cyan }})",
			Details: `
{{ "--------- Instance ----------" | cyan }}
{{ "Name:" | faint }}	{{ .Name }}
{{ "State:" | faint }}	{{ .State }}
{{ "LaunchTime:" | faint }}	{{ .LaunchTime }}
{{ "Ip:" | faint }}	{{ .Ip }}
{{ "Link:" | faint }}	{{ .Link }}
{{ "$" | red }} {{ .SshCommand }}
`,
		},
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Println("Cancelled.")
		os.Exit(1)
	}

	return instances[i]
}
