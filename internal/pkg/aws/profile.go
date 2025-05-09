package aws

import (
	"log"
)

// GetLocalAwsProfiles retrieves the list of AWS profiles from the shared credentials file.
func GetLocalAwsProfiles() []string {
	sections := GetIniConfig().SectionStrings()

	if len(sections) < 1 {
		log.Fatal("No profiles found in config file")
	}

	return sections[1:]
}
