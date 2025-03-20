package aws

import (
	"strings"
)

type Region struct {
	Name  string
	Value string
}

func GetRegions(profile string) []Region {
	regions := strings.Split(GetIniConfig().Section(profile).Key("regions").String(), ",")

	var result []Region
	allRegions := GetAllRegions()

	for _, region := range regions {
		trimmedRegion := strings.TrimSpace(region)
		for _, ar := range allRegions {
			if ar.Value == trimmedRegion {
				result = append(result, ar)
				break
			}
		}
	}

	return result
}

func GetAllRegions() []Region {
	// Taken from: https://docs.aws.amazon.com/general/latest/gr/rande.html
	// Find:    ^(.+)\t(.+)$
	// Replace: {\n\tName:  "$1",\n\tValue: "$2",\n},

	return []Region{
		{
			Name:  "US East (Ohio) ",
			Value: "us-east-2",
		},
		{
			Name:  "US East (N. Virginia) ",
			Value: "us-east-1",
		},
		{
			Name:  "US West (N. California) ",
			Value: "us-west-1",
		},
		{
			Name:  "US West (Oregon) ",
			Value: "us-west-2",
		},
		{
			Name:  "Africa (Cape Town) ",
			Value: "af-south-1",
		},
		{
			Name:  "Asia Pacific (Hong Kong) ",
			Value: "ap-east-1",
		},
		{
			Name:  "Asia Pacific (Hyderabad) ",
			Value: "ap-south-2",
		},
		{
			Name:  "Asia Pacific (Jakarta) ",
			Value: "ap-southeast-3",
		},
		{
			Name:  "Asia Pacific (Malaysia) ",
			Value: "ap-southeast-5",
		},
		{
			Name:  "Asia Pacific (Melbourne) ",
			Value: "ap-southeast-4",
		},
		{
			Name:  "Asia Pacific (Mumbai) ",
			Value: "ap-south-1",
		},
		{
			Name:  "Asia Pacific (Osaka) ",
			Value: "ap-northeast-3",
		},
		{
			Name:  "Asia Pacific (Seoul) ",
			Value: "ap-northeast-2",
		},
		{
			Name:  "Asia Pacific (Singapore) ",
			Value: "ap-southeast-1",
		},
		{
			Name:  "Asia Pacific (Sydney) ",
			Value: "ap-southeast-2",
		},
		{
			Name:  "Asia Pacific (Thailand) ",
			Value: "ap-southeast-7",
		},
		{
			Name:  "Asia Pacific (Tokyo) ",
			Value: "ap-northeast-1",
		},
		{
			Name:  "Canada (Central) ",
			Value: "ca-central-1",
		},
		{
			Name:  "Canada West (Calgary) ",
			Value: "ca-west-1",
		},
		{
			Name:  "Europe (Frankfurt) ",
			Value: "eu-central-1",
		},
		{
			Name:  "Europe (Ireland) ",
			Value: "eu-west-1",
		},
		{
			Name:  "Europe (London) ",
			Value: "eu-west-2",
		},
		{
			Name:  "Europe (Milan) ",
			Value: "eu-south-1",
		},
		{
			Name:  "Europe (Paris) ",
			Value: "eu-west-3",
		},
		{
			Name:  "Europe (Spain) ",
			Value: "eu-south-2",
		},
		{
			Name:  "Europe (Stockholm) ",
			Value: "eu-north-1",
		},
		{
			Name:  "Europe (Zurich) ",
			Value: "eu-central-2",
		},
		{
			Name:  "Israel (Tel Aviv) ",
			Value: "il-central-1",
		},
		{
			Name:  "Mexico (Central) ",
			Value: "mx-central-1",
		},
		{
			Name:  "Middle East (Bahrain) ",
			Value: "me-south-1",
		},
		{
			Name:  "Middle East (UAE) ",
			Value: "me-central-1",
		},
		{
			Name:  "South America (São Paulo) ",
			Value: "sa-east-1",
		},
		{
			Name:  "AWS GovCloud (US-East) ",
			Value: "us-gov-east-1",
		},
		{
			Name:  "AWS GovCloud (US-West) ",
			Value: "us-gov-west-1",
		},
	}

}
