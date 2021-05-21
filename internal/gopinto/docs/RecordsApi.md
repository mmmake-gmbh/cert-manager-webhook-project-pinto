# \RecordsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDnsRecordsDelete**](RecordsApi.md#ApiDnsRecordsDelete) | **Delete** /api/dns/Records | Deletes records which match the specified criterias
[**ApiDnsRecordsGet**](RecordsApi.md#ApiDnsRecordsGet) | **Get** /api/dns/Records | Retrieves the DNS zone&#39;s resource records
[**ApiDnsRecordsPost**](RecordsApi.md#ApiDnsRecordsPost) | **Post** /api/dns/Records | Creates a new DNS resource record



## ApiDnsRecordsDelete

> ApiDnsRecordsDelete(ctx).Provider(provider).Zone(zone).Name(name).RecordType(recordType).Environment(environment).RequestBody(requestBody).Execute()

Deletes records which match the specified criterias

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
    zone := "zone_example" // string | DNS zone to delete records of
    name := "name_example" // string | Name of the DNS record(s) to delete
    recordType := openapiclient.RecordType("A") // RecordType | RecordType of the DNS record(s) to delete
    environment := "environment_example" // string | The name of the environment to utilize (optional)
    requestBody := map[string]string{"key": "Inner_example"} // map[string]string | Meta data to pass through to the provider (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordsApi.ApiDnsRecordsDelete(context.Background()).Provider(provider).Zone(zone).Name(name).RecordType(recordType).Environment(environment).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ApiDnsRecordsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDnsRecordsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The name of the provider to utilize | 
 **zone** | **string** | DNS zone to delete records of | 
 **name** | **string** | Name of the DNS record(s) to delete | 
 **recordType** | [**RecordType**](RecordType.md) | RecordType of the DNS record(s) to delete | 
 **environment** | **string** | The name of the environment to utilize | 
 **requestBody** | **map[string]string** | Meta data to pass through to the provider | 

### Return type

 (empty response body)

### Authorization

[oidc](../README.md#oidc)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDnsRecordsGet

> []Record ApiDnsRecordsGet(ctx).Provider(provider).Zone(zone).Environment(environment).RecordType(recordType).Name(name).PageToken(pageToken).PageSize(pageSize).Execute()

Retrieves the DNS zone's resource records

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
    zone := "zone_example" // string | DNS zone to query
    environment := "environment_example" // string | The name of the environment to utilize (optional)
    recordType := openapiclient.RecordType("A") // RecordType | Filter by the record type (optional)
    name := "name_example" // string | Filter by the record name (optional)
    pageToken := "pageToken_example" // string | The token of the page to load (optional)
    pageSize := int32(56) // int32 | The size of the page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordsApi.ApiDnsRecordsGet(context.Background()).Provider(provider).Zone(zone).Environment(environment).RecordType(recordType).Name(name).PageToken(pageToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ApiDnsRecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDnsRecordsGet`: []Record
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ApiDnsRecordsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDnsRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The name of the provider to utilize | 
 **zone** | **string** | DNS zone to query | 
 **environment** | **string** | The name of the environment to utilize | 
 **recordType** | [**RecordType**](RecordType.md) | Filter by the record type | 
 **name** | **string** | Filter by the record name | 
 **pageToken** | **string** | The token of the page to load | 
 **pageSize** | **int32** | The size of the page | 

### Return type

[**[]Record**](Record.md)

### Authorization

[oidc](../README.md#oidc)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDnsRecordsPost

> Record ApiDnsRecordsPost(ctx).CreateRecordRequestModel(createRecordRequestModel).Execute()

Creates a new DNS resource record

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
    createRecordRequestModel := *openapiclient.NewCreateRecordRequestModel("Provider_example", "Zone_example", "Name_example", openapiclient.RecordType("A"), "Data_example") // CreateRecordRequestModel | Data used to create a DNS resource record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordsApi.ApiDnsRecordsPost(context.Background()).CreateRecordRequestModel(createRecordRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ApiDnsRecordsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDnsRecordsPost`: Record
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ApiDnsRecordsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDnsRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRecordRequestModel** | [**CreateRecordRequestModel**](CreateRecordRequestModel.md) | Data used to create a DNS resource record | 

### Return type

[**Record**](Record.md)

### Authorization

[oidc](../README.md#oidc)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

