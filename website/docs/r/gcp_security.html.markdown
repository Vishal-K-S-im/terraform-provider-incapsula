---
subcategory: "Cloud WAF - Site Management"
layout: "incapsula"
page_title: "Security Policies for IGC Sites"
description: |-
  Which security resources work with IGC sites.
---

# Security Policies for IGC Sites

IGC load-balancer and domain sites are ordinary V3 sites from a
security-policy perspective: policies and rules operate on the site/asset id and do not branch on
site type. The following resources therefore apply to IGC sites without changes:

* `incapsula_policy` — account-level policy definition.
* `incapsula_policy_asset_association` — associate a policy with a site asset.
* `incapsula_incap_rule` — per-site rules.
* `incapsula_security_rule_exception` — per-site security-rule exceptions.

Policies and rules apply at the site level. For a load-balancer site, the site is the single unit of
enforcement for all of its discovered domains. A domain site behaves like a single-domain Cloud WAF site.

These resources reuse the same schema as for Cloud WAF and AWS sites. See the existing Cloud WAF and AWS
security guides for the full set of arguments.

## Resources that do not apply

For IGC (`PUBLIC_CLOUD` + `GCP`) sites, the following resources are **not applicable** and should not be attached to them:

* `incapsula_data_center` / `incapsula_data_center_server` / `incapsula_data_centers_configuration` — origin data centers are managed by the connector, not by Cloud WAF DC configuration.
* CDN / caching resources (`incapsula_site_cache_configuration`, `incapsula_cache_rule`) — content delivery is not part of the IGC data path.
* Certificate resources (`incapsula_custom_certificate`, `incapsula_managed_certificate_settings`, `incapsula_site_ssl_settings`) — TLS termination for IGC sites is handled outside of Cloud WAF certificate management.

## Example

```hcl
resource "incapsula_site_v3" "gcp_lb" {
  name                  = "my-gcp-lb-site"
  type                  = "PUBLIC_CLOUD"
  cloud_type            = "GCP"
  is_load_balancer_site = true
}

resource "incapsula_incap_rule" "alert_rule" {
  name    = "gcpAlert"
  site_id = incapsula_site_v3.gcp_lb.id
  action  = "RULE_ACTION_ALERT"
  filter  = "Full-URL == \"/login\""
  enabled = true
}
```

