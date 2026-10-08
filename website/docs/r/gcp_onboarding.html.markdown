---
subcategory: "Cloud WAF - Site Management"
layout: "incapsula"
page_title: "IGC Site Onboarding with Terraform"
description: |-
  End-to-end guide for onboarding IGC (Imperva for Google Cloud) sites via Terraform.
---

# IGC Site Onboarding with Terraform

This guide walks through onboarding IGC (Imperva for Google Cloud) sites using the
Incapsula Terraform provider. It covers a **load-balancer site** (the site IGC routes traffic
through) and a **domain site**, plus reading the site's `authority_header`, discovering regions
and resolving the GCP PSC target.

## 1. Create the IGC load-balancer site

`is_load_balancer_site = true` is valid only for `PUBLIC_CLOUD` + `GCP`.
It is set on creation only. Changing it on an existing site fails at apply time, and the site is not replaced.
To recreate the site intentionally, run `terraform apply -replace=<resource address>`.

```hcl
resource "incapsula_site_v3" "gcp_lb" {
  name                  = "my-gcp-lb-site"
  type                  = "PUBLIC_CLOUD"
  cloud_type            = "GCP"
  is_load_balancer_site = true
}

output "authority_header" {
  value = incapsula_site_v3.gcp_lb.authority_header
}
```

The `authority_header` (format `<key>-<accountId>.<domain>`) is populated on read and is used to route requests through the connector.

## 2. Discover supported regions

```hcl
data "incapsula_gcp_regions" "all" {}
```

The list is sorted alphabetically but its contents can change.

## 3. Resolve the PSC target for a chosen region

The PSC target lookup requires the account to already have a load-balancer site.
Use `depends_on` so the lookup runs after the site from step 1 is created.

```hcl
data "incapsula_gcp_psc_target" "target" {
  region     = "us-east1"
  depends_on = [incapsula_site_v3.gcp_lb]
}
```

Use `data.incapsula_gcp_psc_target.target.address` as the service attachment when creating the
GCP Private Service Connect endpoint that links your VPC to Imperva.

## 4. Create a domain site

A domain site is a regular site (`is_load_balancer_site = false`, or the field omitted); it has no `authority_header`.
Add its protected domain with the `incapsula_domain` resource.

```hcl
resource "incapsula_site_v3" "gcp_domain" {
  name                  = "my-domain-site"
  type                  = "PUBLIC_CLOUD"
  cloud_type            = "GCP"
  is_load_balancer_site = false
}

resource "incapsula_domain" "gcp_domain" {
  site_id = incapsula_site_v3.gcp_domain.id
  domain  = "www.example.com"
}
```

## 5. Inspect the site's domains

```hcl
data "incapsula_gcp_domains" "lb_domains" {
  site_id = incapsula_site_v3.gcp_lb.id
}
```

Lists all domains of the site, including auto-discovered ones. A domain site returns its protected domain.

## 6. Look up a site created elsewhere by name

```hcl
data "incapsula_site_v3" "existing" {
  name = "my-gcp-lb-site"
}

output "existing_authority_header" {
  value = data.incapsula_site_v3.existing.authority_header
}
```

## 7. Attach security policies

Security-policy resources apply to IGC sites as-is. See the
[IGC security guide](gcp_security.html) for details.
