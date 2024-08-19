package main

import (
	"fmt"
	"github.com/semiherdogan/aws-ssh/internal/pkg/aws"
	"github.com/semiherdogan/aws-ssh/internal/pkg/ui"
	"github.com/semiherdogan/aws-ssh/internal/pkg/utils"
	"os"
	"strings"
)

func main() {
	cli := NewCli()

	err := cli.Run()
	if err != nil {
		return
	}

	profile := "default"
	cli.AddCommand("--profile", "-p", func() {
		profile = ui.SelectProfile(aws.GetLocalAwsProfiles())
	})

	var userSelectedRegion aws.Region
	cli.AddCommand("--region", "-r", func() {
		userSelectedRegion = ui.SelectRegion(aws.GetRegions())
	})

	var userFilter string
	if len(cli.Args) > 0 {
		userFilter = strings.Join(cli.Args, " ")
	} else {
		userFilter = ui.GetInput("Filter")
	}

	var userFilters = strings.Split(
		strings.TrimSpace(userFilter),
		" ",
	)

	pkgAws := aws.Aws{}
	var awsInstances []aws.Instance
	var currentRegion aws.Region

	regions := aws.GetRegions()
	if userSelectedRegion.Name != "" {
		regions = []aws.Region{userSelectedRegion}
	}

	for _, r := range regions {
		currentRegion = r
		pkgAws = aws.Aws{
			Region:  r.Value,
			Profile: profile,
			Filters: []string{
				fmt.Sprintf("*%s*", userFilters[0]),
				fmt.Sprintf("*%s*", utils.CapitalizeFirstLetter(userFilters[0])),
				fmt.Sprintf("*%s*", strings.ToLower(userFilters[0])),
				fmt.Sprintf("*%s*", strings.ToUpper(userFilters[0])),
			},
		}

		awsInstances = pkgAws.GetEc2Instances()

		if len(awsInstances) > 0 {
			break
		}
	}

	instances := utils.Filter(awsInstances, func(i aws.Instance) bool {
		instanceName := strings.ToLower(i.Name)

		for _, f := range userFilters[1:] {
			if !strings.Contains(instanceName, strings.ToLower(f)) {
				return false
			}
		}

		return true
	})

	if len(instances) == 0 {
		fmt.Println("No instances found.")
		os.Exit(0)
	}

	if len(regions) > 1 {
		fmt.Printf("===== %s (%s) =====\n", currentRegion.Name, currentRegion.Value)
	}

	selectedInstance := ui.SelectInstance(instances)

	utils.RunShellCommand(selectedInstance.SshCommand)
}
