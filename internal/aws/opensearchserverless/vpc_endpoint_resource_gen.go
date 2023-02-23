// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_opensearchserverless_vpc_endpoint", vpcEndpointResource)
}

// vpcEndpointResource returns the Terraform awscc_opensearchserverless_vpc_endpoint resource.
// This Terraform resource corresponds to the CloudFormation AWS::OpenSearchServerless::VpcEndpoint resource.
func vpcEndpointResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the VPC Endpoint",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^vpce-[0-9a-z]*$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the VPC Endpoint",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the VPC Endpoint",
		//	  "maxLength": 32,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z][a-z0-9-]{2,31}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the VPC Endpoint",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 32),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z][a-z0-9-]{2,31}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of one or more security groups to associate with the endpoint network interface",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "minLength": 1,
		//	    "pattern": "^[\\w+\\-]+$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The ID of one or more security groups to associate with the endpoint network interface",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 5),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[\\w+\\-]+$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of one or more subnets in which to create an endpoint network interface",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 32,
		//	    "minLength": 1,
		//	    "pattern": "^subnet-([0-9a-f]{8}|[0-9a-f]{17})$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 6,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The ID of one or more subnets in which to create an endpoint network interface",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 6),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^subnet-([0-9a-f]{8}|[0-9a-f]{17})$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC in which the endpoint will be used.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^vpc-[0-9a-z]*$",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC in which the endpoint will be used.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^vpc-[0-9a-z]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Amazon OpenSearchServerless vpc endpoint resource",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::OpenSearchServerless::VpcEndpoint").WithTerraformTypeName("awscc_opensearchserverless_vpc_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":                 "Id",
		"name":               "Name",
		"security_group_ids": "SecurityGroupIds",
		"subnet_ids":         "SubnetIds",
		"vpc_id":             "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
