---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_domain Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SageMaker::Domain
---

# awscc_sagemaker_domain (Data Source)

Data Source schema for AWS::SageMaker::Domain



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `app_network_access_type` (String) Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
- `app_security_group_management` (String) The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
- `auth_mode` (String) The mode of authentication that members use to access the domain.
- `default_space_settings` (Attributes) The default space settings. (see [below for nested schema](#nestedatt--default_space_settings))
- `default_user_settings` (Attributes) The default user settings. (see [below for nested schema](#nestedatt--default_user_settings))
- `domain_arn` (String) The Amazon Resource Name (ARN) of the created domain.
- `domain_id` (String) The domain name.
- `domain_name` (String) A name for the domain.
- `domain_settings` (Attributes) A collection of Domain settings. (see [below for nested schema](#nestedatt--domain_settings))
- `home_efs_file_system_id` (String) The ID of the Amazon Elastic File System (EFS) managed by this Domain.
- `kms_key_id` (String) SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.
- `security_group_id_for_domain_boundary` (String) The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
- `single_sign_on_managed_application_instance_id` (String) The SSO managed application instance ID.
- `subnet_ids` (List of String) The VPC subnets that Studio uses for communication.
- `tags` (Attributes List) A list of tags to apply to the user profile. (see [below for nested schema](#nestedatt--tags))
- `url` (String) The URL to the created domain.
- `vpc_id` (String) The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.

<a id="nestedatt--default_space_settings"></a>
### Nested Schema for `default_space_settings`

Read-Only:

- `execution_role` (String) The execution role for the space.
- `jupyter_server_app_settings` (Attributes) The Jupyter server's app settings. (see [below for nested schema](#nestedatt--default_space_settings--jupyter_server_app_settings))
- `kernel_gateway_app_settings` (Attributes) The kernel gateway app settings. (see [below for nested schema](#nestedatt--default_space_settings--kernel_gateway_app_settings))
- `security_groups` (List of String) The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.

<a id="nestedatt--default_space_settings--jupyter_server_app_settings"></a>
### Nested Schema for `default_space_settings.jupyter_server_app_settings`

Read-Only:

- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--default_space_settings--jupyter_server_app_settings--default_resource_spec))

<a id="nestedatt--default_space_settings--jupyter_server_app_settings--default_resource_spec"></a>
### Nested Schema for `default_space_settings.jupyter_server_app_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The Amazon Resource Name (ARN) of the image version created on the instance.



<a id="nestedatt--default_space_settings--kernel_gateway_app_settings"></a>
### Nested Schema for `default_space_settings.kernel_gateway_app_settings`

Read-Only:

- `custom_images` (Attributes List) A list of custom SageMaker images that are configured to run as a KernelGateway app. (see [below for nested schema](#nestedatt--default_space_settings--kernel_gateway_app_settings--custom_images))
- `default_resource_spec` (Attributes) The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app. (see [below for nested schema](#nestedatt--default_space_settings--kernel_gateway_app_settings--default_resource_spec))

<a id="nestedatt--default_space_settings--kernel_gateway_app_settings--custom_images"></a>
### Nested Schema for `default_space_settings.kernel_gateway_app_settings.custom_images`

Read-Only:

- `app_image_config_name` (String) The Name of the AppImageConfig.
- `image_name` (String) The name of the CustomImage. Must be unique to your account.
- `image_version_number` (Number) The version number of the CustomImage.


<a id="nestedatt--default_space_settings--kernel_gateway_app_settings--default_resource_spec"></a>
### Nested Schema for `default_space_settings.kernel_gateway_app_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The Amazon Resource Name (ARN) of the image version created on the instance.




<a id="nestedatt--default_user_settings"></a>
### Nested Schema for `default_user_settings`

Read-Only:

- `execution_role` (String) The execution role for the user.
- `jupyter_server_app_settings` (Attributes) The Jupyter server's app settings. (see [below for nested schema](#nestedatt--default_user_settings--jupyter_server_app_settings))
- `kernel_gateway_app_settings` (Attributes) The kernel gateway app settings. (see [below for nested schema](#nestedatt--default_user_settings--kernel_gateway_app_settings))
- `r_session_app_settings` (Attributes) A collection of settings that apply to an RSessionGateway app. (see [below for nested schema](#nestedatt--default_user_settings--r_session_app_settings))
- `r_studio_server_pro_app_settings` (Attributes) A collection of settings that configure user interaction with the RStudioServerPro app. (see [below for nested schema](#nestedatt--default_user_settings--r_studio_server_pro_app_settings))
- `security_groups` (List of String) The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
- `sharing_settings` (Attributes) The sharing settings. (see [below for nested schema](#nestedatt--default_user_settings--sharing_settings))

<a id="nestedatt--default_user_settings--jupyter_server_app_settings"></a>
### Nested Schema for `default_user_settings.jupyter_server_app_settings`

Read-Only:

- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--default_user_settings--jupyter_server_app_settings--default_resource_spec))

<a id="nestedatt--default_user_settings--jupyter_server_app_settings--default_resource_spec"></a>
### Nested Schema for `default_user_settings.jupyter_server_app_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The Amazon Resource Name (ARN) of the image version created on the instance.



<a id="nestedatt--default_user_settings--kernel_gateway_app_settings"></a>
### Nested Schema for `default_user_settings.kernel_gateway_app_settings`

Read-Only:

- `custom_images` (Attributes List) A list of custom SageMaker images that are configured to run as a KernelGateway app. (see [below for nested schema](#nestedatt--default_user_settings--kernel_gateway_app_settings--custom_images))
- `default_resource_spec` (Attributes) The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app. (see [below for nested schema](#nestedatt--default_user_settings--kernel_gateway_app_settings--default_resource_spec))

<a id="nestedatt--default_user_settings--kernel_gateway_app_settings--custom_images"></a>
### Nested Schema for `default_user_settings.kernel_gateway_app_settings.custom_images`

Read-Only:

- `app_image_config_name` (String) The Name of the AppImageConfig.
- `image_name` (String) The name of the CustomImage. Must be unique to your account.
- `image_version_number` (Number) The version number of the CustomImage.


<a id="nestedatt--default_user_settings--kernel_gateway_app_settings--default_resource_spec"></a>
### Nested Schema for `default_user_settings.kernel_gateway_app_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The Amazon Resource Name (ARN) of the image version created on the instance.



<a id="nestedatt--default_user_settings--r_session_app_settings"></a>
### Nested Schema for `default_user_settings.r_session_app_settings`

Read-Only:

- `custom_images` (Attributes List) A list of custom SageMaker images that are configured to run as a KernelGateway app. (see [below for nested schema](#nestedatt--default_user_settings--r_session_app_settings--custom_images))
- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--default_user_settings--r_session_app_settings--default_resource_spec))

<a id="nestedatt--default_user_settings--r_session_app_settings--custom_images"></a>
### Nested Schema for `default_user_settings.r_session_app_settings.custom_images`

Read-Only:

- `app_image_config_name` (String) The Name of the AppImageConfig.
- `image_name` (String) The name of the CustomImage. Must be unique to your account.
- `image_version_number` (Number) The version number of the CustomImage.


<a id="nestedatt--default_user_settings--r_session_app_settings--default_resource_spec"></a>
### Nested Schema for `default_user_settings.r_session_app_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The Amazon Resource Name (ARN) of the image version created on the instance.



<a id="nestedatt--default_user_settings--r_studio_server_pro_app_settings"></a>
### Nested Schema for `default_user_settings.r_studio_server_pro_app_settings`

Read-Only:

- `access_status` (String) Indicates whether the current user has access to the RStudioServerPro app.
- `user_group` (String) The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.


<a id="nestedatt--default_user_settings--sharing_settings"></a>
### Nested Schema for `default_user_settings.sharing_settings`

Read-Only:

- `notebook_output_option` (String) Whether to include the notebook cell output when sharing the notebook. The default is Disabled.
- `s3_kms_key_id` (String) When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.
- `s3_output_path` (String) When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.



<a id="nestedatt--domain_settings"></a>
### Nested Schema for `domain_settings`

Read-Only:

- `r_studio_server_pro_domain_settings` (Attributes) A collection of settings that update the current configuration for the RStudioServerPro Domain-level app. (see [below for nested schema](#nestedatt--domain_settings--r_studio_server_pro_domain_settings))
- `security_group_ids` (List of String) The security groups for the Amazon Virtual Private Cloud that the Domain uses for communication between Domain-level apps and user apps.

<a id="nestedatt--domain_settings--r_studio_server_pro_domain_settings"></a>
### Nested Schema for `domain_settings.r_studio_server_pro_domain_settings`

Read-Only:

- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--domain_settings--r_studio_server_pro_domain_settings--default_resource_spec))
- `domain_execution_role_arn` (String) The ARN of the execution role for the RStudioServerPro Domain-level app.
- `r_studio_connect_url` (String) A URL pointing to an RStudio Connect server.
- `r_studio_package_manager_url` (String) A URL pointing to an RStudio Package Manager server.

<a id="nestedatt--domain_settings--r_studio_server_pro_domain_settings--default_resource_spec"></a>
### Nested Schema for `domain_settings.r_studio_server_pro_domain_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The Amazon Resource Name (ARN) of the image version created on the instance.




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
