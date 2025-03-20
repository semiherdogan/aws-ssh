package aws

import (
	"log"
)

func GetLocalAwsProfiles() []string {
	sections := GetIniConfig().SectionStrings()

	if len(sections) < 1 {
		log.Fatal("No profiles found in config file")
	}

	return sections[1:]
}
