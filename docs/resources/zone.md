---
page_title: "zone Resource - terraform-provider-adman"
subcategory: "dns"
description: |-
  dns zone resource.
---

## Example Usage

```terraform
resource "adman_dns_zone" "test_com" {
  domain  = "test.com"
}
```

## Argument Reference

- `domain` - (Required) Domain name.
