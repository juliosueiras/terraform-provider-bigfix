---
layout: "bigfix"
page_title: "Bigfix: bigfix_fixlet"
sidebar_current: "docs-bigfix-resource-inventory-folder"
description: |-
  Get details of existing fixlet in BES environment.
---

# bigfix\_computer

Get details of existing fixlet in BES environment.

## Example Usage

```hcl
# data source bigfix fixlet
data "bigfix_fixlet" "myfixlet"{
    name = var.fixlet_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The distinguished name of the fixlet.
* `id` - (Computed) The distinguished id of the fixlet.
* `default_action` - (Computed) default action mention in fixlet.
* `site_name` - (Computed) name of site to which fixlet belongs to.
