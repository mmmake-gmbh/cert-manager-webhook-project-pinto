# \ZonesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDnsZonesGet**](ZonesApi.md#ApiDnsZonesGet) | **Get** /api/dns/Zones | Retrieves the DNS zones assigned to the account
[**ApiDnsZonesPost**](ZonesApi.md#ApiDnsZonesPost) | **Post** /api/dns/Zones | Creates a new DNS zone
[**ApiDnsZonesZoneDelete**](ZonesApi.md#ApiDnsZonesZoneDelete) | **Delete** /api/dns/Zones/{zone} | Deletes a DNS zone from the passed provider
[**ApiDnsZonesZoneGet**](ZonesApi.md#ApiDnsZonesZoneGet) | **Get** /api/dns/Zones/{zone} | Loads the specified DNS zone



## ApiDnsZonesGet

> []Zone ApiDnsZonesGet(ctx).Provider(provider).Environment(environment).PageToken(pageToken).PageSize(pageSize).Execute()

Retrieves the DNS zones assigned to the account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    provider := "provider_example" // string | The name of the provider to utilize
    environment := "environment_example" // string | The name of the environment to utilize (optional)
    pageToken := "pageToken_example" // string | The token of the page to load (optional)
    pageSize := int32(56) // int32 | The size of the page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.ApiDnsZonesGet(context.Background()).Provider(provider).Environment(environment).PageToken(pageToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ApiDnsZonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDnsZonesGet`: []Zone
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ApiDnsZonesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDnsZonesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The name of the provider to utilize | 
 **environment** | **string** | The name of the environment to utilize | 
 **pageToken** | **string** | The token of the page to load | 
 **pageSize** | **int32** | The size of the page | 

### Return type

[**[]Zone**](Zone.md)

### Authorization

[oidc](../README.md#oidc)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDnsZonesPost

> Zone ApiDnsZonesPost(ctx).CreateZoneRequestModel(createZoneRequestModel).Execute()

Creates a new DNS zone

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createZoneRequestModel := *openapiclient.NewCreateZoneRequestModel("Provider_example", "Name_example") // CreateZoneRequestModel | Data used to create a DNS zone

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.ApiDnsZonesPost(context.Background()).CreateZoneRequestModel(createZoneRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ApiDnsZonesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDnsZonesPost`: Zone
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ApiDnsZonesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDnsZonesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createZoneRequestModel** | [**CreateZoneRequestModel**](CreateZoneRequestModel.md) | Data used to create a DNS zone | 

### Return type

[**Zone**](Zone.md)

### Authorization

[oidc](../README.md#oidc)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDnsZonesZoneDelete

> ApiDnsZonesZoneDelete(ctx, zone).Provider(provider).Environment(environment).Execute()

Deletes a DNS zone from the passed provider

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    provider := "provider_example" // string | The name of the provider to utilize
    zone := "zone_example" // string | Name of the DNS zone to delete
    environment := "environment_example" // string | The name of the environment to utilize (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.ApiDnsZonesZoneDelete(context.Background(), zone).Provider(provider).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ApiDnsZonesZoneDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Name of the DNS zone to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDnsZonesZoneDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The name of the provider to utilize | 

 **environment** | **string** | The name of the environment to utilize | 

### Return type

 (empty response body)

### Authorization

[oidc](../README.md#oidc)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDnsZonesZoneGet

> Zone ApiDnsZonesZoneGet(ctx, zone).Provider(provider).Environment(environment).Execute()

Loads the specified DNS zone

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    provider := "provider_example" // string | The name of the provider to utilize
    zone := "zone_example" // string | Name of the DNS zone to load
    environment := "environment_example" // string | The name of the environment to utilize (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZonesApi.ApiDnsZonesZoneGet(context.Background(), zone).Provider(provider).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.ApiDnsZonesZoneGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDnsZonesZoneGet`: Zone
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.ApiDnsZonesZoneGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Name of the DNS zone to load | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDnsZonesZoneGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The name of the provider to utilize | 

 **environment** | **string** | The name of the environment to utilize | 

### Return type

[**Zone**](Zone.md)

### Authorization

[oidc](../README.md#oidc)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

