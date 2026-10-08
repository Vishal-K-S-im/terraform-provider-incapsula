---
subcategory: "Cloud WAF - Site Management"
layout: "incapsula"
page_title: "incapsula_site_v3"
description: |- 
  Provides a Incapsula Site resource.
---

# incapsula_site_v3

Provides an Incapsula V3 site resource. A V3 site resource is the core resource that is required by all other resources. incapsula_site_v3 is a newer version of incapsula_site. Site can be managed by incapsula_site_v3 or incapsula_site, but not both simultaneously.
Full site onboarding example with TF can be found [here](https://docs.imperva.com/bundle/cloud-application-security/page/website-certificate-terraform.htm).

## Example Usage

### Cloud WAF Site (default)

```hcl
resource "incapsula_site_v3" "example-site-v3" {
  name = "example.com"
}
```

### Imperva for AWS (PUBLIC_CLOUD) Site

```hcl
resource "incapsula_site_v3" "aws-site" {
  name       = "my-aws-site"
  type       = "PUBLIC_CLOUD"
  cloud_type = "AWS"
  ref_id     = "cf-dist-E1234567890"
}
```

### IGC (Imperva for Google Cloud) Load-Balancer Site

```hcl
resource "incapsula_site_v3" "igc-lb-site" {
  name                  = "my-igc-lb-site"
  type                  = "PUBLIC_CLOUD"
  cloud_type            = "GCP"
  is_load_balancer_site = true
}

output "authority_header" {
  value = incapsula_site_v3.igc-lb-site.authority_header
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) The account to operate on. If not specified, operation will be performed on the account identified by the authentication parameters.
* `name` - (Required) The site name.
* `type` - (Optional) The website type. Indicates which kind of website is created. Supported values: `CLOUD_WAF` (default) for a website onboarded to Imperva Cloud WAF, `PUBLIC_CLOUD` for a website onboarded to Imperva for a public cloud provider (e.g., Imperva for AWS).
* `cloud_type` - (Optional) The cloud provider type. Required when `type` is `PUBLIC_CLOUD`. Supported values: `AWS`, `GCP`. This field cannot be changed after creation.
* `ref_id` - (Optional) Sets the Reference ID. A free-text field that enables you to add a unique identifier to correlate a website in our service with an object on the customer side.
* `active` - (Optional) Whether the site is active or bypassing the Imperva network.
* `is_load_balancer_site` - (Optional) Whether this site is an IGC load-balancer site. Can be set **only on creation** and **only** when `type` is `PUBLIC_CLOUD` and `cloud_type` is `GCP`. A load-balancer site (`true`) is the site through which IGC routes traffic and for which an `authority_header` is issued; a domain site (`false`, or the field omitted on a `PUBLIC_CLOUD` + `GCP` site) is a regular site with no IGC onboarding parameters. See the [IGC onboarding guide](gcp_onboarding.html) and the [IGC security guide](gcp_security.html).

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the site.
* `creation_time` - Creation time of the site.
* `cname` - The CNAME provided by Imperva that is used for pointing your website traffic to the Imperva network.
* `is_load_balancer_site` - Whether the site is an IGC load-balancer site. A legacy IGC default site (created automatically before multi-site support) reports `false`; keep the field omitted or set to `false` when importing it.
* `authority_header` - The authority header assigned by Imperva for an IGC load-balancer site (format `<key>-<accountId>.<domain>`). Populated for load-balancer sites and for a legacy IGC default site; empty for domain sites and non-IGC sites.

## Import

Site can be imported using the `account Id`/`id`, e.g.:

```
$ terraform import incapsula_site_v3.example 543/1234
```
