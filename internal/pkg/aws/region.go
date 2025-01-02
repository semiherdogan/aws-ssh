package aws

type Region struct {
	Name  string
	Value string
}

func GetRegions() []Region {
	// Taken from: https://docs.aws.amazon.com/general/latest/gr/rande.html
	// Find:    ^(.+)\t(.+)$
	// Replace: {\n\tName:  "$1",\n\tValue: "$2",\n},

	return []Region{
		{
			Name:  "Europe (Frankfurt)",
			Value: "eu-central-1",
		},
		{
			Name:  "US East (N. Virginia)",
			Value: "us-east-1",
		},
	}
}
