package aws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
)

func GetIniConfig() (f *ini.File) {
	f, err := ini.Load(config.DefaultSharedCredentialsFilename())

	if err != nil {
		log.Fatal("Error while loading config file", err)
	}

	return
}
