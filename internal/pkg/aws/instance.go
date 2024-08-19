package aws

type Instance struct {
	Name             string
	InstanceId       string
	State            string
	LaunchTime       string
	Ip               string
	Link             string
	SshCommand       string
	TerminateCommand string
}
