// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package systemsmanagersap

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
	registry.AddResourceFactory("awscc_systemsmanagersap_application", applicationResource)
}

// applicationResource returns the Terraform awscc_systemsmanagersap_application resource.
// This Terraform resource corresponds to the CloudFormation AWS::SystemsManagerSAP::Application resource.
func applicationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "[\\w\\d]{1,50}",
		//	  "type": "string"
		//	}
		"application_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[\\w\\d]{1,50}"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ApplicationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "HANA"
		//	  ],
		//	  "type": "string"
		//	}
		"application_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"HANA",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Helix application",
		//	  "pattern": "^arn:(.+:){2,4}.+$|^arn:(.+:){1,3}.+\\/.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Helix application",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Credentials
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "CredentialType": {
		//	        "enum": [
		//	          "ADMIN"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "DatabaseName": {
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "SecretId": {
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"credentials": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CredentialType
					"credential_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"ADMIN",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DatabaseName
					"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SecretId
					"secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Credentials is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Instances
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "pattern": "^i-[\\w\\d]{8}$|^i-[\\w\\d]{17}$",
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"instances": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^i-[\\w\\d]{8}$|^i-[\\w\\d]{17}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Instances is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SapInstanceNumber
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "[0-9]{2}",
		//	  "type": "string"
		//	}
		"sap_instance_number": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9]{2}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// SapInstanceNumber is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Sid
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "[A-Z][A-Z0-9]{2}",
		//	  "type": "string"
		//	}
		"sid": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[A-Z][A-Z0-9]{2}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Sid is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags of a SystemsManagerSAP application.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags of a SystemsManagerSAP application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::SystemsManagerSAP::Application",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SystemsManagerSAP::Application").WithTerraformTypeName("awscc_systemsmanagersap_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_id":      "ApplicationId",
		"application_type":    "ApplicationType",
		"arn":                 "Arn",
		"credential_type":     "CredentialType",
		"credentials":         "Credentials",
		"database_name":       "DatabaseName",
		"instances":           "Instances",
		"key":                 "Key",
		"sap_instance_number": "SapInstanceNumber",
		"secret_id":           "SecretId",
		"sid":                 "Sid",
		"tags":                "Tags",
		"value":               "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Credentials",
		"/properties/Instances",
		"/properties/SapInstanceNumber",
		"/properties/Sid",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
