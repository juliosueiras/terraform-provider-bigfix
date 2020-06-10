---
layout: "Bigfix"
page_title: "Provider: Bigfix"
sidebar_current: "docs-bigfix-index"
description: |-
  The Bigfix provider is used to interact with the resources supported by
  HCL Bigfix. The provider needs to be configured with the proper credentials
  before it can be used.
---

# Bigfix Provider

The Bigfix provider is used to interact with the resources supported by HCL Bigfix.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

~> **NOTE:** The Bigfix Provider currently represents _initial support_
and therefore may undergo significant changes as the community improves it. This
provider at this time only supports adding Multiple Action Group for a targeted computer.

## Example Usage

```hcl
# data source bigfix fixlet
data "bigfix_fixlet" "myfixlet"{
  name = var.fixlet_name
}

# data source bigfix computer
data  "bigfix_computer" "linux_vm"{
  name = var.linux-vm-name
}

# resource bigfix multiple_action_group
resource "bigfix_multiple_action_group" "test" {
  input_file_name = "MAG.xml" 
  target_computer_id = data.bigfix_computer.linux_vm.id
  site_name = var.linux-sites
}

# provider config
provider "bigfix" {
  port = var.bigfix_port
  username = var.bigfix_username
  password = var.bigfix_password
  server = var.bigfix_server
}

```

## Argument Reference

The following arguments are used to configure the Bigfix Provider:

* `username` - (Required) This is username of BES Console operator.
* `password` - (Required) This is password of BES Console operator.
* `server` - (Required) This is BES Server address.
* `port` - (Required) This is port of BES Server.

## Acceptance Tests

The Bigfix provider's acceptance tests require the following environment variables, and must be set to valid values for your Bigfix environment:

* BFX\_USERNAME
* BFX\_PASSWORD
* BFX\_SERVER
* BFX\_PORT
* BFX\_INPUT\_FILE\_NAME
* BFX\_FIXLET
* BFX\_COMPUTER
 
Once all these variables are in place, the tests can be run like this:

```
make testacc
```
