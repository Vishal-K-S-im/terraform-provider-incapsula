---
subcategory: "Cloud WAF - Site Management"
layout: "incapsula"
page_title: "incapsula_site_lookup"
description: |-
  Retrieves an existing site by its name.
---

# incapsula_site_lookup (Data Source)

Retrieves an existing V3 site by name and returns its onboarding parameters, including the `authority_header` for Imperva for Google Cloud load balancer websites. Use this data source to reference a site that was created outside Terraform or in a different Terraform state. It applies to all site types.

## Example Usage

```hcl
data "incapsula_site_lookup" "lb" {
  name = "my-gcp-lb-site"
}

output "authority_header" {
  value = data.incapsula_site_lookup.lb.authority_header
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The exact, case-sensitive name of the site. Site names are not unique. The lookup fails if no site or more than one site has this name.
* `account_id` - (Optional) The account ID (`caid`) used to scope the request. When omitted, the account of the API credentials is used.

## Attributes Reference

The following attributes are exported:

* `id` / `site_id` - The numeric identifier of the site.
* `cname` - The CNAME provided by Imperva for the site.
* `type` - The website type of the site, such as `CLOUD_WAF` or `PUBLIC_CLOUD`.
* `cloud_type` - The cloud provider type of the site, `AWS` or `GCP`. Empty for sites that are not `PUBLIC_CLOUD`.
* `is_load_balancer_site` - Indicates whether the site is an Imperva for Google Cloud load balancer website. `false` for domain websites, legacy Imperva for Google Cloud default sites, and sites that are not part of Imperva for Google Cloud.
* `authority_header` - The authority header assigned by Imperva for an Imperva for Google Cloud load balancer website (format `<key>-<accountId>.<domain>`). Available for Imperva for Google Cloud load balancer websites and legacy Imperva for Google Cloud default sites. Empty for domain websites and sites that are not part of Imperva for Google Cloud.
