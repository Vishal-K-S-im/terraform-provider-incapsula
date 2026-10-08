---
subcategory: "Cloud WAF - Site Management"
layout: "incapsula"
page_title: "incapsula_gcp_regions"
description: |-
  Provides the list of supported GCP regions.
---

# incapsula_gcp_regions (Data Source)

Provides the list of supported GCP regions. Pick a region from this list before requesting a PSC target via [`incapsula_gcp_psc_target`](gcp_psc_target.html).

## Example Usage

```hcl
data "incapsula_gcp_regions" "all" {}

data "incapsula_gcp_psc_target" "target" {
  region = "us-east1"
}

output "region_supported" {
  value = contains(data.incapsula_gcp_regions.all.regions, "us-east1")
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) The account ID (`caid`) to scope the request to.

## Attributes Reference

The following attributes are exported:

* `id` - A stable identifier for the data source.
* `regions` - The list of supported GCP region identifiers (e.g., `us-east1`, `europe-west1`), sorted alphabetically. The set of regions can change over time.
