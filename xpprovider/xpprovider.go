// Package xpprovider exposes the Terraform Plugin Framework provider
// implementation of the Cisco IOS-XE Terraform provider so that it can be
// consumed in-process (Upjet "no-fork" mode) by Crossplane providers.
//
// The provider implementation itself lives under internal/, which makes it
// unimportable from other Go modules. This package is the exported shim used
// by github.com/upbound/provider-upjet-iosxe.
package xpprovider

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"

	iosxeprovider "github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider"
)

// GetProvider returns the Terraform Plugin Framework provider implementation
// for the given provider version.
func GetProvider(version string) provider.Provider {
	return iosxeprovider.New(version)()
}
