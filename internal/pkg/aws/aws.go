package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/savioxavier/termlink"
)

type Aws struct {
	Profile string
	Region  string
	Filters []string
}

func (a *Aws) GetConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(a.Region),
		config.WithSharedConfigProfile(a.Profile),
	)

	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load AWS configuration for profile '%s' and region '%s': %w", a.Profile, a.Region, err)
	}

	return cfg, nil
}

func (a *Aws) GetEc2Instances() ([]Instance, error) {
	cfg, err := a.GetConfig()
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

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
		return nil, fmt.Errorf("failed to describe instances: %w", err)
	}

	if len(resp.Reservations) == 0 {
		return nil, nil
	}

	return a.getInstancesFromEc2Response(resp), nil
}

// Refactored getInstancesFromEc2Response to improve readability and reduce nesting
func (a *Aws) getInstancesFromEc2Response(resp *ec2.DescribeInstancesOutput) []Instance {
	var instances []Instance

	for _, reservation := range resp.Reservations {
		for _, ec2Instance := range reservation.Instances {
			instance := Instance{
				InstanceId: *ec2Instance.InstanceId,
				State:      string(ec2Instance.State.Name),
				LaunchTime: ec2Instance.LaunchTime.Format("2006-01-02 15:04:05"),
			}

			if ec2Instance.PublicIpAddress != nil {
				instance.Ip = *ec2Instance.PublicIpAddress
			}

			for _, tag := range ec2Instance.Tags {
				if *tag.Key == "Name" {
					instance.Name = *tag.Value
					break
				}
			}

			instance.Link = a.generateInstanceLink(*ec2Instance.InstanceId)
			instance.SshCommand = a.generateSshCommand(*ec2Instance.InstanceId)
			instance.TerminateCommand = a.generateTerminateCommand(*ec2Instance.InstanceId)

			instances = append(instances, instance)
		}
	}

	return instances
}

func (a *Aws) generateInstanceLink(instanceId string) string {
	link := fmt.Sprintf(
		"https://%s.console.aws.amazon.com/ec2/v2/home?region=%s#ConnectToInstance:instanceId=%s",
		a.Region,
		a.Region,
		instanceId,
	)

	if termlink.SupportsHyperlinks() {
		return termlink.ColorLink(fmt.Sprintf("Link to: %s", instanceId), link, "italic green")
	}

	return link
}

func (a *Aws) generateSshCommand(instanceId string) string {
	return fmt.Sprintf(
		"aws ssm start-session --target %s --region %s --profile %s",
		instanceId,
		a.Region,
		a.Profile,
	)
}

func (a *Aws) generateTerminateCommand(instanceId string) string {
	return fmt.Sprintf(
		"aws ec2 terminate-instances --instance-ids %s --region %s --profile %s",
		instanceId,
		a.Region,
		a.Profile,
	)
}
