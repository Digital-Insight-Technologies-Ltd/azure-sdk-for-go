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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/CustomIpPrefixDelete.json
func ExampleCustomIPPrefixesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomIPPrefixesClient().BeginDelete(ctx, "rg1", "test-customipprefix", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/CustomIpPrefixGet.json
func ExampleCustomIPPrefixesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomIPPrefixesClient().Get(ctx, "rg1", "test-customipprefix", &armnetwork.CustomIPPrefixesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomIPPrefix = armnetwork.CustomIPPrefix{
	// 	Name: to.Ptr("test-customipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
	// 		AuthorizationMessage: to.Ptr("authorizationMessage"),
	// 		ChildCustomIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		Cidr: to.Ptr("0.0.0.0/24"),
	// 		CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
	// 		ExpressRouteAdvertise: to.Ptr(false),
	// 		FailedReason: to.Ptr(""),
	// 		NoInternetAdvertise: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 		SignedMessage: to.Ptr("signedMessage"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/CustomIpPrefixCreateCustomizedValues.json
func ExampleCustomIPPrefixesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomIPPrefixesClient().BeginCreateOrUpdate(ctx, "rg1", "test-customipprefix", armnetwork.CustomIPPrefix{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
			Cidr: to.Ptr("0.0.0.0/24"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomIPPrefix = armnetwork.CustomIPPrefix{
	// 	Name: to.Ptr("test-customipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
	// 		AuthorizationMessage: to.Ptr("authorizationMessage"),
	// 		ChildCustomIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		Cidr: to.Ptr("192.168.254.2/24"),
	// 		CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioning),
	// 		FailedReason: to.Ptr(""),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 		SignedMessage: to.Ptr("signedMessage"),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/CustomIpPrefixUpdateTags.json
func ExampleCustomIPPrefixesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomIPPrefixesClient().UpdateTags(ctx, "rg1", "test-customipprefix", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomIPPrefix = armnetwork.CustomIPPrefix{
	// 	Name: to.Ptr("test-customipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
	// 		AuthorizationMessage: to.Ptr("authorizationMessage"),
	// 		ChildCustomIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		Cidr: to.Ptr("192.168.254.2/24"),
	// 		CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioning),
	// 		FailedReason: to.Ptr(""),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 		SignedMessage: to.Ptr("signedMessage"),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/CustomIpPrefixListAll.json
func ExampleCustomIPPrefixesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomIPPrefixesClient().NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CustomIPPrefixListResult = armnetwork.CustomIPPrefixListResult{
		// 	Value: []*armnetwork.CustomIPPrefix{
		// 		{
		// 			Name: to.Ptr("test-customipprefix"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.0.0/24"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix2"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/customIpPrefixes/test-customipprefix2"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.2.0/23"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix3"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix3"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.4.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix4"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix4"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("2607:f0d1:1002:0001::/64"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				CustomIPPrefixParent: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix5"),
		// 				},
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix5"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix5"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix4"),
		// 				}},
		// 				Cidr: to.Ptr("2607:f0d1:1002::/48"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix6"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.5.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedState("ProvisionFailed")),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr("CustomerSignatureNotVerified"),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix7"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.6.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				Asn: to.Ptr("11"),
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				Geo: to.Ptr(armnetwork.GeoGLOBAL),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(true),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(true),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				PrefixType: to.Ptr(armnetwork.CustomIPPrefixTypeParent),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/CustomIpPrefixList.json
func ExampleCustomIPPrefixesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomIPPrefixesClient().NewListPager("rg1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CustomIPPrefixListResult = armnetwork.CustomIPPrefixListResult{
		// 	Value: []*armnetwork.CustomIPPrefix{
		// 		{
		// 			Name: to.Ptr("test-customipprefix"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.0.0/24"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix2"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix2"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.1.0/30"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix4"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix4"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("2607:f0d1:1002:0001::/64"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
		// 				CustomIPPrefixParent: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix5"),
		// 				},
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix5"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix5"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix4"),
		// 				}},
		// 				Cidr: to.Ptr("2607:f0d1:1002::/48"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioned),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-customipprefix6"),
		// 			Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/customIpPrefixes/test-customipprefix8"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
		// 				AuthorizationMessage: to.Ptr("authorizationMessage"),
		// 				ChildCustomIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				Cidr: to.Ptr("0.0.7.0/22"),
		// 				CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioning),
		// 				ExpressRouteAdvertise: to.Ptr(false),
		// 				FailedReason: to.Ptr(""),
		// 				NoInternetAdvertise: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 				},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 				SignedMessage: to.Ptr("signedMessage"),
		// 			},
		// 	}},
		// }
	}
}
