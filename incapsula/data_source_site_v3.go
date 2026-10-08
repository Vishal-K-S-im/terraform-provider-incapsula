package incapsula

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIncapsulaSiteV3() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIncapsulaSiteV3Read,
		Description: "Looks up a v3 site by name and exposes its onboarding parameters, including the IGC authority header for load-balancer sites.",

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The exact name of the site to look up.",
				Required:    true,
			},
			"account_id": {
				Type:        schema.TypeString,
				Description: "The account ID (caid) to scope the lookup to. When omitted, the account of the API credentials is used.",
				Optional:    true,
			},
			"site_id": {
				Type:        schema.TypeString,
				Description: "The numeric identifier of the matched site.",
				Computed:    true,
			},
			"is_load_balancer_site": {
				Type:        schema.TypeBool,
				Description: "Whether the matched site is an IGC load-balancer site.",
				Computed:    true,
			},
			"authority_header": {
				Type:        schema.TypeString,
				Description: "The authority header assigned by Imperva. Populated for IGC load-balancer sites and for a legacy IGC default site; empty for domain sites and non-IGC sites.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "The website type of the matched site.",
				Computed:    true,
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Description: "The cloud provider type of the matched site.",
				Computed:    true,
			},
		},
	}
}

func dataSourceIncapsulaSiteV3Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	name := d.Get("name").(string)
	accountID := d.Get("account_id").(string)

	siteV3Response, diags := client.ListV3Sites(name, accountID)
	if diags != nil && diags.HasError() {
		return diags
	}
	if len(siteV3Response.Errors) > 0 {
		return diag.Errorf("Failed to look up site by name %s: %s", name, siteV3Response.Errors[0].Detail)
	}

	site, err := findSiteV3ByName(siteV3Response.Data, name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(site.Id))
	isLb := site.IsLoadBalancerSite != nil && *site.IsLoadBalancerSite
	values := map[string]interface{}{
		"site_id":               strconv.Itoa(site.Id),
		"type":                  site.SiteType,
		"cloud_type":            site.CloudType,
		"authority_header":      site.AuthorityHeader,
		"is_load_balancer_site": isLb,
	}
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	if site.SiteType == "PUBLIC_CLOUD" && site.CloudType == "GCP" && !isLb && site.AuthorityHeader == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Site is not an IGC load-balancer site",
			Detail:   "The matched site is an IGC domain site and has no authority_header. IGC onboarding parameters are only available for load-balancer sites.",
		})
	}

	return diags
}

func findSiteV3ByName(sites []SiteV3Request, name string) (*SiteV3Request, error) {
	var matches []SiteV3Request
	for _, site := range sites {
		if site.Name == name {
			matches = append(matches, site)
		}
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("no site found with name %s", name)
	}
	if len(matches) > 1 {
		ids := make([]string, 0, len(matches))
		for _, site := range matches {
			ids = append(ids, strconv.Itoa(site.Id))
		}
		return nil, fmt.Errorf("found %d sites with name %s (site ids: %s); site names must be unique for this lookup", len(matches), name, strings.Join(ids, ", "))
	}
	return &matches[0], nil
}
