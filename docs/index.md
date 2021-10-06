---
page_title: "Provider: adman dns"
subcategory: "dns"
description: |-
  Terraform provider for interacting with dns records hosted on adman.com.
---

# Adman DNS Provider

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "adman" {
  mdpass = "mdpass"
  login = "login"
}
```

## Schema

### Required

- **login** (String, Optional) Username to authenticate to adman API
- **mdpass** (String, Optional) Password to authenticate to adman API
