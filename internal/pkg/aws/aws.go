package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/savioxavier/termlink"
	"log"
)

type Aws struct {
	Profile string
	Region  string
	Filters []string
}

func (a *Aws) GetConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(a.Region),
		config.WithSharedConfigProfile(a.Profile),
	)

	if err != nil {
		log.Fatalf("Failed to load aws configuration, %v", err)
	}

	return cfg
}

func (a *Aws) GetEc2Instances() (result []Instance) {
	client := ec2.NewFromConfig(a.GetConfig())

	resp, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []string{"running", "pending"},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: a.Filters,
			},
		},
	})

	if err != nil {
		log.Fatalf("Failed to describe instances, %v", err)
	}

	if len(resp.Reservations) == 0 {
		return
	}

	if len(resp.Reservations[0].Instances) == 0 {
		return
	}

	result = a.getInstancesFromEc2Response(resp)

	return
}

func (a *Aws) getInstancesFromEc2Response(resp *ec2.DescribeInstancesOutput) []Instance {
	var instances []Instance

	for _, r := range resp.Reservations {
		for _, i := range r.Instances {
			var instance Instance

			if i.PublicIpAddress != nil {
				instance.Ip = *i.PublicIpAddress
			}

			instance.InstanceId = *i.InstanceId

			for _, t := range i.Tags {
				if *t.Key == "Name" {
					instance.Name = *t.Value
				}
			}

			instance.State = string(i.State.Name)

			t := *i.LaunchTime
			instance.LaunchTime = t.Format("2006-01-02 15:04:05")

			link := fmt.Sprintf(
				"https://%s.console.aws.amazon.com/ec2/v2/home?region=%s#ConnectToInstance:instanceId=%s",
				a.Region,
				a.Region,
				*i.InstanceId,
			)

			if termlink.SupportsHyperlinks() {
				link = termlink.ColorLink(fmt.Sprintf("Link to: %s", *i.InstanceId), link, "italic green")
			}

			instance.Link = link

			instance.SshCommand = fmt.Sprintf(
				"aws ssm start-session --target %s --region %s --profile %s",
				*i.InstanceId,
				a.Region,
				a.Profile,
			)

			instance.TerminateCommand = fmt.Sprintf(
				"aws ec2 terminate-instances --instance-ids %s --region %s --profile %s",
				*i.InstanceId,
				a.Region,
				a.Profile,
			)

			instances = append(instances, instance)
		}
	}

	return instances
}
