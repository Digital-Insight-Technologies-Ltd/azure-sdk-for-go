//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/FirewallPolicySignatureOverridesPatch.json
func ExampleFirewallPolicyIdpsSignaturesOverridesClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyIdpsSignaturesOverridesClient().Patch(ctx, "rg1", "firewallPolicy", armnetwork.SignaturesOverrides{
		Name: to.Ptr("default"),
		Type: to.Ptr("Microsoft.Network/firewallPolicies/signatureOverrides"),
		ID:   to.Ptr("/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default"),
		Properties: &armnetwork.SignaturesOverridesProperties{
			Signatures: map[string]*string{
				"2000105": to.Ptr("Off"),
				"2000106": to.Ptr("Deny"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SignaturesOverrides = armnetwork.SignaturesOverrides{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Network/firewallPolicies/signatureOverrides"),
	// 	ID: to.Ptr("/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default"),
	// 	Properties: &armnetwork.SignaturesOverridesProperties{
	// 		Signatures: map[string]*string{
	// 			"2000105": to.Ptr("Off"),
	// 			"2000106": to.Ptr("Deny"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/FirewallPolicySignatureOverridesPut.json
func ExampleFirewallPolicyIdpsSignaturesOverridesClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyIdpsSignaturesOverridesClient().Put(ctx, "rg1", "firewallPolicy", armnetwork.SignaturesOverrides{
		Name: to.Ptr("default"),
		Type: to.Ptr("Microsoft.Network/firewallPolicies/signatureOverrides"),
		ID:   to.Ptr("/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default"),
		Properties: &armnetwork.SignaturesOverridesProperties{
			Signatures: map[string]*string{
				"2000105": to.Ptr("Off"),
				"2000106": to.Ptr("Deny"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SignaturesOverrides = armnetwork.SignaturesOverrides{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Network/firewallPolicies/signatureOverrides"),
	// 	ID: to.Ptr("/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default"),
	// 	Properties: &armnetwork.SignaturesOverridesProperties{
	// 		Signatures: map[string]*string{
	// 			"2000105": to.Ptr("Off"),
	// 			"2000106": to.Ptr("Deny"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/FirewallPolicySignatureOverridesGet.json
func ExampleFirewallPolicyIdpsSignaturesOverridesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyIdpsSignaturesOverridesClient().Get(ctx, "rg1", "firewallPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SignaturesOverrides = armnetwork.SignaturesOverrides{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Network/firewallPolicies/signatureOverrides"),
	// 	ID: to.Ptr("/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default"),
	// 	Properties: &armnetwork.SignaturesOverridesProperties{
	// 		Signatures: map[string]*string{
	// 			"2000105": to.Ptr("Off"),
	// 			"2000106": to.Ptr("Deny"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/FirewallPolicySignatureOverridesList.json
func ExampleFirewallPolicyIdpsSignaturesOverridesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyIdpsSignaturesOverridesClient().List(ctx, "rg1", "firewallPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SignaturesOverridesList = armnetwork.SignaturesOverridesList{
	// 	Value: []*armnetwork.SignaturesOverrides{
	// 		{
	// 			Name: to.Ptr("default"),
	// 			Type: to.Ptr("Microsoft.Network/firewallPolicies/signatureOverrides"),
	// 			ID: to.Ptr("/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default"),
	// 			Properties: &armnetwork.SignaturesOverridesProperties{
	// 				Signatures: map[string]*string{
	// 					"2000105": to.Ptr("Off"),
	// 					"2000106": to.Ptr("Deny"),
	// 				},
	// 			},
	// 	}},
	// }
}
