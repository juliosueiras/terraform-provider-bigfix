# Terraform BigFix Provider

This is the repository for the Terraform bigfix Provider, which one can use
with Terraform to work with [bigfix][1].

Coverage is currently limited to a one resource only ie. Multiple action group and two data sources called Computer and Fixlet  

[1]: https://www.hcltechsw.com/products/bigfix

For general information about Terraform, visit the [official website][3] and the
[GitHub project page][4].

[3]: https://terraform.io/
[4]: https://github.com/hashicorp/terraform

# Using the Provider

The current version of this provider requires Terraform v0.12 or higher to
run.

Note that you need to run `terraform init` to fetch the provider before
deploying. Read about the provider split and other changes to TF v0.12 in the
official document found [here][5].

[5]: https://www.terraform.io/docs/extend/terraform-0.12-compatibility.html

## Full Provider Documentation

The provider is useful for creating [multiple action group][12] in bigfix for particular target computer.

Additionally you can mention custom attributes for multiple action group using xml file along with terraform configuration file.

[12]: https://help.hcltechsw.com/bigfix/9.5/platform/Platform/Console/c_take_multiple_actions.html

### Example

Checkout complete example [here](/Examples)

```hcl
# Configure the Bigfix Provider
provider "bigfix" {
  port = var.bigfix_port
  username = var.bigfix_username
  password = var.bigfix_password
  server = var.bigfix_server
}

# data source bigfix fixlet
data "bigfix_fixlet" "myfixlet"{
  name = var.fixlet_name
}

###########################################
# example: patch a linux machine
###########################################

# data source bigfix computer
data  "bigfix_computer" "linux_vm"{
  name = var.linux-vm-name
}

# resource bigfix multiple_action_group
resource "bigfix_multiple_action_group" "test" {
  # path of xml file containing custom attributes
  input_file_name = "MAG.xml"
  target_computer_id = data.bigfix_computer.linux_vm.id
  site_name = var.linux-sites
}

output "MAG-linux-vm" {
  value = bigfix_multiple_action_group.test.title
}

output "action-state-linux" {
  value = bigfix_multiple_action_group.test.state
}
```

# Building The Provider

**NOTE:** Unless you are [developing][6] or require a pre-release bugfix or feature,
you will want to use the officially released version of the provider (see [the
section above][7]).

[6]: #developing-the-provider
[7]: #using-the-provider

## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/terraform-providers/terraform-provider-bigfix`:

```sh
mkdir -p $GOPATH/src/github.com/terraform-providers
cd $GOPATH/src/github.com/terraform-providers
git clone git@github.com:terraform-providers/terraform-provider-bigfix
```

## Running the Build

After the clone has been completed, you can enter the provider directory and
build the provider.

```sh
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-bigfix
make build
```

## Installing the Local Plugin

After the build is complete, copy the `terraform-provider-bigfix` binary into
the same path as your `terraform` binary, and re-run `terraform init`.

After this, your project-local `.terraform/plugins/ARCH/lock.json` (where `ARCH`
matches the architecture of your machine) file should contain a SHA256 sum that
matches the local plugin. Run `shasum -a 256` on the binary to verify the values
match.

# Developing the Provider

If you wish to work on the provider, you'll first need [Go][8] installed on your
machine (version 1.13.1+ is **required**). You'll also need to correctly setup a
[GOPATH][9], as well as adding `$GOPATH/bin` to your `$PATH`.

[8]: https://golang.org/
[9]: http://golang.org/doc/code.html#GOPATH

See [Building the Provider][10] for details on building the provider.

[10]: #building-the-provider


## Checking the Logs
To persist logged output you can set TF_LOG_PATH in order to force the log to always be appended to a specific file when logging is enabled. Note that even when TF_LOG_PATH is set, TF_LOG must be set in order for any logging to be enabled.

To check logs use the following commands :
```sh
# Specify Log Level
export TF_LOG=DEBUG
# Specify Log File Path
export TF_LOG_PATH='. . .'
```

## Configuring Environment Variables

Most of the tests in this provider require a comprehensive list of environment
variables to run. See the individual `*_test.go` files in the
[`bigfix/`](bigfix/) directory for more details. The next section also
describes how you can manage a configuration file of the test environment
variables.

## Running the Acceptance Tests
In order to perform acceptance tests of bigfix, first set in your environment variables required for the connection (`BFX_SERVER`, `BFX_PORT`, `BFX_USERNAME`, `BFX_PASSWORD`, `BFX_COMPUTER`, `BFX_INPUT_FILE_NAME`, `BFX_FIXLET`).

For `BFX_INPUT_FILE_NAME`, give complete path of file.

After this is done, you can run the acceptance tests by running:

```sh
$ make testacc
```


# Building The Provider

**NOTE:** Unless you are [developing][7] or require a pre-release bugfix or feature,
you will want to use the officially released version of the provider (see [the
section above][8]).

[7]: #developing-the-provider
[8]: #using-the-provider

