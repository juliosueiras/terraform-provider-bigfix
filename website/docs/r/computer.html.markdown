---
layout: "bigfix"
page_title: "Bigfix: bigfix_computer"
sidebar_current: "docs-bigfix-resource-inventory-folder"
description: |-
  Get details of existing computers in BES environment.
---

# bigfix\_computer

Get details of existing computers in BES environment.

## Example Usage

```hcl
# data source bigfix computer
data  "bigfix_computer" "linux_vm"{
    name = var.linux-vm-name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The distinguished name of the computer .
* `id` - (Computed) The distinguished id of the computer.
* `ip_address` - (Computed) ip_address of the computer.
* `os` - (Computed) os of the computer.
* `cpu` - (Computed) cpu of the computer.
* `last_report_time` - (Computed) last report time of the computer.
* `dns_name` - (Computed) The dns_name of the computer.
* `ram` - (Computed) ram of the computer.
* `subnet_address` - (Computed) subnet address of the computer.
* `computer_type` - (Computed) Computer type.
* `relay` - (Computed) relay.
