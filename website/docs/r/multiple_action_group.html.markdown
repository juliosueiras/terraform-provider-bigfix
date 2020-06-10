---
layout: "bigfix"
page_title: "Bigfix: bigfix_multiple_action_group"
sidebar_current: "docs-bigfix-resource-inventory-folder"
description: |-
  Initiates Multiple action group for targeted computer.
---

# bigfix\_multiple\_action\_group

Initiates Multiple action group for targeted computer.

**Note:** As a requirement, each Fixlet or task involved in the group must have associated a default action, ie. it will load only fixlets with default actions.

## Example Usage

```hcl

###########################################
# Example: Patch a linux machine
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

```

## Argument Reference

The following arguments are supported:

* `target_computer_id` - (Required) The distinguished id of computer we want to patch.
* `input_file_name` - (Required) The input file which contains custom attributes of multiple action group.
* `site_name` - (Required) List of sites from which relevant fixlets will be fetched. Note that it will load only those fixlets which has default action.
* `title` - (Optional) Name of multiple action group.
* `relevance` - (Optional) Relevance for multiple action group.
* `state` - (Computed) State of multiple action group after creation.
