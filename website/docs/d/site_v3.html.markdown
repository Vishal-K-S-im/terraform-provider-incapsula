---
subcategory: "Cloud WAF - Site Management"
layout: "incapsula"
page_title: "incapsula_site_v3"
description: |-
  Looks up an Incapsula V3 site by name.
---

# incapsula_site_v3 (Data Source)

Looks up an existing V3 site by its name and exposes its onboarding parameters, including the IGC `authority_header` for load-balancer sites. Use this when the site was created outside of Terraform.

## Example Usage

```hcl
data "incapsula_site_v3" "lb" {
  name = "my-gcp-lb-site"
}

output "authority_header" {
  value = data.incapsula_site_v3.lb.authority_header
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The exact (case-sensitive) name of the site to look up. Site names are not unique: the lookup fails if no site or more than one site has this name.
* `account_id` - (Optional) The account ID (`caid`) to scope the lookup to.

## Attributes Reference

The following attributes are exported:

* `id` / `site_id` - The numeric identifier of the matched site.
* `is_load_balancer_site` - Whether the matched site is an IGC load-balancer site.
* `authority_header` - The authority header of the site. Populated for IGC load-balancer sites and for a legacy IGC default site; empty for domain sites and non-IGC sites.
* `type` - The website type of the matched site.
* `cloud_type` - The cloud provider type of the matched site.
