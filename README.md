# AWS SSH

Cli tool for connection to EC2 through session manager.

It's a wrapper UI around AWS CLI tool, designed to simplify the process of connecting to your EC2 instances. This tool leverages the AWS CLI and SSM plugin to provide a user-friendly interface for selecting and connecting to EC2 instances.

### Before installing this tool:
> This tool requires you to install AWS CLI tool and SSM plugin.
For installing, you may follow the link below:

[Installing the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)

[Installing the AWS SSM Plugin](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html)

### Installation

To install AWS SSH, you can use the `go get` command:

```bash
go get github.com/semiherdogan/aws-ssh
```

### Usage

```bash
aws-ssh [OPTIONS] searchparam1 searchparam2 ...
```

This will launch a UI that lists all your EC2 instances. You can search and select the instance you wish to connect to. The tool will then initiate a connection using the AWS Session Manager.

### Features
* Easy to use UI for selecting EC2 instances
* Search functionality to quickly find your instances
* Connects to EC2 through AWS Session Manager for secure access

### Contributing
Contributions are welcome! Please feel free to submit a pull request.

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
