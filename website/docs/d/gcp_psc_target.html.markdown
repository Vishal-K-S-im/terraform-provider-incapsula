---
subcategory: "Cloud WAF - Site Management"
layout: "incapsula"
page_title: "incapsula_gcp_psc_target"
description: |-
  Provides the GCP PSC target address for a region.
---

# incapsula_gcp_psc_target (Data Source)

Provides the GCP PSC target address for a given region. Use the returned `address` when configuring the GCP PSC endpoint that connects your VPC to Imperva. Pick the `region` from the [`incapsula_gcp_regions`](gcp_regions.html) data source.

~> **Note:** The account must already have an IGC load-balancer site (`incapsula_site_v3` with `is_load_balancer_site = true`). A legacy IGC default site (created automatically before multi-site support) also satisfies this check.

## Example Usage

```hcl
data "incapsula_gcp_psc_target" "target" {
  region     = "us-east1"
  depends_on = [incapsula_site_v3.gcp-lb-site]
}

output "psc_target" {
  value = data.incapsula_gcp_psc_target.target.address
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Required) The region for which to retrieve the GCP PSC target.
* `account_id` - (Optional) The account ID (`caid`) to scope the request to.

## Attributes Reference

The following attributes are exported:

* `id` - Set to `<account_id>/<region>` when `account_id` is set, otherwise to the requested region.
* `address` - The GCP PSC target address for the specified region.
