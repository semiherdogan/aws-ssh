package aws

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
	"log"
)

// GetLocalAwsProfiles Taken from: https://github.com/aws/aws-sdk-go/issues/3656#issuecomment-2017510038
func GetLocalAwsProfiles() (profiles []string) {
	f, err := ini.Load(config.DefaultSharedCredentialsFilename()) // Load ini file

	if err != nil {
		log.Fatal("Error while loading config file", err)
	}

	sections := f.SectionStrings()

	if len(sections) < 1 {
		log.Fatal("No profiles section found")
	}

	for _, s := range sections[1:] {
		profiles = append(profiles, s)
	}

	return
}
