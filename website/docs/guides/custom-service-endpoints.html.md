---
subcategory: ""
layout: "ibm"
page_title: "IBM Cloud Provider plugin for Terraform Custom Service Endpoint Configuration"
description: |-
  Configuring the IBM Cloud Provider plugin for Terraform to connect to custom IBM service endpoints.
---

# Custom Service Endpoint Configuration

The IBM Cloud Provider plugin for Terraform configuration can be customized to connect to non-default IBM service endpoints  This may be useful for environments with specific compliance requirements, such as Classic, VPE or local testing.

This guide outlines how to get started with customizing endpoints, the available endpoint configurations, and offers example configurations for working with certain local development and testing solutions.

~> **NOTE:** Support for connecting the IBM Cloud Provider plugin for Terraform with custom endpoints is offered as best effort. Individual Terraform resources may require compatibility updates to work in certain environments. Integration testing by IBM-Cloud during provider changes is exclusively done against default IBM endpoints.

<!-- TOC depthFrom:2 -->

- [Getting Started with Custom Endpoints](#getting-started-with-custom-endpoints)
- [Available Endpoint Customizations](#available-endpoint-customizations)
- [File Structuring](#file-structuring)
- [Prioritisation of Endpoints](#prioritisation-of-endpoints)
<!-- /TOC -->

## Getting Started with Custom Endpoints

To configure the IBM Cloud Provider plugin for Terraform to use customized endpoints, it can be done within `provider` declarations using the `visibility` and `endpoints_file` attributes.   e.g.

```terraform
provider "ibm" {
  # ... potentially other provider configuration ...

  visiblity="private"
  endpoints_file= "endpoints.json"
}
```

If multiple, different IBM Cloud Provider plugin for Terraform configurations are required, see the [Terraform documentation on multiple provider instances](https://www.terraform.io/docs/configuration/providers.html#alias-multiple-provider-instances) for additional information about the `alias` provider configuration and its usage.

## Available Endpoint Customizations

The IBM Cloud Provider plugin for Terraform allows the following endpoints to be customized:

| Service | Endpoint Variable |
|---------|-----------------|
|API Gateway|IBMCLOUD_API_GATEWAY_ENDPOINT|
|Account Management|IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT|
|Catalog Management|IBMCLOUD_CATALOG_MANAGEMENT_API_ENDPOINT|
|Certificate Manager|IBMCLOUD_CERTIFICATE_MANAGER_API_ENDPOINT|
|Cloud Object Storage|IBMCLOUD_COS_CONFIG_ENDPOINT|
|Internet Services|IBMCLOUD_CIS_API_ENDPOINT|
|Container Registry|IBMCLOUD_CR_API_ENDPOINT|
|Kubernetes Service|IBMCLOUD_CS_API_ENDPOINT|
|Direct Link|IBMCLOUD_DL_API_ENDPOINT|
|Direct Link Provider|IBMCLOUD_DL_PROVIDER_API_ENDPOINT|
|Enterprise Management|IBMCLOUD_ENTERPRISE_API_ENDPOINT|
|Cloud Functions|IBMCLOUD_FUNCTIONS_API_ENDPOINT|
|Global Tagging|IBMCLOUD_GT_API_ENDPOINT|
|Global Search|IBMCLOUD_GS_API_ENDPOINT|
|Hper Protect Crypto Services|IBMCLOUD_HPCS_API_ENDPOINT|
|Identity and Access Management|IBMCLOUD_IAM_API_ENDPOINT|
|Identity and Access Management(PAP)|IBMCLOUD_IAMPAP_API_ENDPOINT|
|Cloud Databases|IBMCLOUD_ICD_API_ENDPOINT|
|Virtual Private Cloud (VPC)|IBMCLOUD_IS_NG_API_ENDPOINT|
|Key Management Services|IBMCLOUD_KP_API_ENDPOINT|
|Cloud Foundry|IBMCLOUD_MCCP_API_ENDPOINT|
|Push Notifications|IBMCLOUD_PUSH_API_ENDPOINT|
|Private DNS|IBMCLOUD_PRIVATE_DNS_API_ENDPOINT|
|Resource Controller|IBMCLOUD_RESOURCE_CONTROLLER_API_ENDPOINT|
|Resource Manager|IBMCLOUD_RESOURCE_MANAGEMENT_API_ENDPOINT|
|Global Catalog|IBMCLOUD_RESOURCE_CATALOG_API_ENDPOINT|
|Satellite|IBMCLOUD_SATELLITE_API_ENDPOINT|
|Schematics|IBMCLOUD_SCHEMATICS_API_ENDPOINT|
|Secrets Manager|IBMCLOUD_SECRETS_MANAGER_API_ENDPOINT|
|Transit Gateway|IBMCLOUD_TG_API_ENDPOINT|
|UAA|IBMCLOUD_UAA_ENDPOINT|
|User Management|IBMCLOUD_USER_MANAGEMENT_ENDPOINT|

## File Structuring

Public and private regional endpoints of a service should be provided in respective `visibility` blocks for a given endpoint point. Allowed visibility values in a file are `public` and `private`

```json
{
    "ENDPOINT":{
        "VISIBILITY":{
            "REGION":"<service endpoint>"
        }
    }
}
```

 For eg:

```json
{
    "IBMCLOUD_API_GATEWAY_ENDPOINT":{
        "public":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        },
        "private":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        }
    },
    "IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT":{
        "public":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        },
        "private":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        }
    }
}
```

## Prioritisation of Endpoints

1. Endpoints defined externally using Environment variable
2. Endpoints defined using `endpoints_file` in the provider block
3. Default private or public endpoints based on `visibility` in the provider block 
### 1. Endpoints defined externally using Environment variable

The provider gives highest priority to the exported environment variables (refer to the table for variable name); the provider use this endpoint URL to connect to the IBM Cloud Service, it will ignore the visibility & endpoint_file settings defined in the provider block.


Eg:

Define your provider block with/without `visibility` and `endpoint_file` attributes and export the endpoints environment variable before running any of terrfaorm commands.
IBM Terraform provider initialises respective service client with the exported endpoint.

~> **NOTE:**  These endpoints can't be defined as provider level arguments but, only declared as environment variables.

```terraform
provider "ibm" {
    # ... potentially other provider configuration ...
}
```

```text
export IBMCLOUD_API_GATEWAY_ENDPOINT="<endpoint_value>" 
```

### 2. Endpoints defined using `endpoints_file` in the provider block

The provider will use this endpoint_file value to fetch the endpoint URL to connect to the IBM Cloud Service, depending on the `region` and  `visibility` settings in the provider block. 

**NOTE:**  

- Provider level argument for Endpoints File - `endpoints_file`
- Environment variable for Endpoints File    - `IBMCLOUD_ENDPOINTS_FILE`
- `visibility` argument should also be passed along with the `endpoints_file` for the provider to determine `public` and `private` endpoints.
- Allowable values for `visibility` when `endpoints_file` attribute is given are `public` and `private`. Default: `public`

Eg:

```terraform
    provider "ibm" {
        # ... potentially other provider configuration ...
        endpoints_file = "endpoints file path"
        visibility     = "private"
    }
```

These arguments can also be exported as environment variables instead of defining at provider level

Eg:

```terraform
    provider "ibm" {
        # ... potentially other provider configuration ...
    }

```

```text
export IC_ENDPOINTS_FILE="<endpoint_value>"
export IC_VISIBILITY="private"
```

### 3. Default private or public endpoints based on `visibility` in the provider block 

For a given `region` and  `visibility` settings, if there is no endpoint environment variable and no endpoint_file in the provider block, then the provider will use the default (or pre-defined) endpoint URL settings implemented by the `IBM Cloud Provider plugin for Terraform`

~> **NOTE:**  In order to use the private endpoint from an IBM Cloud resource (such as, a classic VM instance), one must have VRF-enabled account.  If the Cloud service does not support private endpoint, the terraform resource or datasource will log an error.

- Allowable values for `visibility` are `public`, `private`, `public-and-private`. Default: `public`.
  - If visibility is set to `public`, provider uses regional public endpoint or global public endpoint. The regional public endpoints has higher precedence.
  - If visibility is set to `private`, provider uses regional private endpoint or global private endpoint. The regional private endpoint is given higher precedence.  
  - If visibility is set to `public-and-private`, provider uses regional private endpoints or global private endpoint. If service doesn't support regional or global private endpoints it will use the regional or global public endpoint.
- This can also be sourced from the `IC_VISIBILITY` (higher precedence) or `IBMCLOUD_VISIBILITY` environment variable.

Eg:

```terraform
    provider "ibm" {
        # ... potentially other provider configuration ...
        visibility     = "private"
    }
```

These arguments can also be exported as environment variable instead of defining at provider level

Eg:

```terraform
    provider "ibm" {
        # ... potentially other provider configuration ...
    }
```

```text
export IC_VISIBILITY="private" or export IC_VISIBILITY="public-and-private"
```
