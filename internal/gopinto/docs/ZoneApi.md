# \ZoneApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsApiZonesDelete**](ZoneApi.md#DnsApiZonesDelete) | **Delete** /dns/api/Zones | Deletes a DNS zone from the passed provider
[**DnsApiZonesGet**](ZoneApi.md#DnsApiZonesGet) | **Get** /dns/api/Zones | Retrieves the DNS zones assigned to the account
[**DnsApiZonesPost**](ZoneApi.md#DnsApiZonesPost) | **Post** /dns/api/Zones | Creates a new DNS zone
[**DnsApiZonesZoneGet**](ZoneApi.md#DnsApiZonesZoneGet) | **Get** /dns/api/Zones/{zone} | Loads the specified DNS zone



## DnsApiZonesDelete

> DnsApiZonesDelete(ctx).Name(name).XApiOptions(xApiOptions).Execute()

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
    name := "name_example" // string | Name of the DNS zone to delete.
    xApiOptions := "xApiOptions_example" // string | Data used to access the API <br /><br />  Schema: <br />  { <br />  \"access_options\": { <br />  \"provider\": \"provider\", <br />  \"environment\": \"environment\", <br />  \"credentials_id\": \"00000000-0000-0000-0000-000000000000\" <br />  }, <br />  \"meta\": { <br />  \"additionalKey\": \"keyValue\" <br />  } <br />  }

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.DnsApiZonesDelete(context.Background()).Name(name).XApiOptions(xApiOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.DnsApiZonesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsApiZonesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of the DNS zone to delete. | 
 **xApiOptions** | **string** | Data used to access the API &lt;br /&gt;&lt;br /&gt;  Schema: &lt;br /&gt;  { &lt;br /&gt;  \&quot;access_options\&quot;: { &lt;br /&gt;  \&quot;provider\&quot;: \&quot;provider\&quot;, &lt;br /&gt;  \&quot;environment\&quot;: \&quot;environment\&quot;, &lt;br /&gt;  \&quot;credentials_id\&quot;: \&quot;00000000-0000-0000-0000-000000000000\&quot; &lt;br /&gt;  }, &lt;br /&gt;  \&quot;meta\&quot;: { &lt;br /&gt;  \&quot;additionalKey\&quot;: \&quot;keyValue\&quot; &lt;br /&gt;  } &lt;br /&gt;  } | 

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


## DnsApiZonesGet

> []Zone DnsApiZonesGet(ctx).XApiOptions(xApiOptions).Execute()

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
    xApiOptions := "xApiOptions_example" // string | Data used to access the API <br /><br />  Schema: <br />  { <br />  \"access_options\": { <br />  \"provider\": \"provider\", <br />  \"environment\": \"environment\", <br />  \"credentials_id\": \"00000000-0000-0000-0000-000000000000\" <br />  }, <br />  \"meta\": { <br />  \"additionalKey\": \"keyValue\" <br />  } <br />  }

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.DnsApiZonesGet(context.Background()).XApiOptions(xApiOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.DnsApiZonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsApiZonesGet`: []Zone
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.DnsApiZonesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsApiZonesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiOptions** | **string** | Data used to access the API &lt;br /&gt;&lt;br /&gt;  Schema: &lt;br /&gt;  { &lt;br /&gt;  \&quot;access_options\&quot;: { &lt;br /&gt;  \&quot;provider\&quot;: \&quot;provider\&quot;, &lt;br /&gt;  \&quot;environment\&quot;: \&quot;environment\&quot;, &lt;br /&gt;  \&quot;credentials_id\&quot;: \&quot;00000000-0000-0000-0000-000000000000\&quot; &lt;br /&gt;  }, &lt;br /&gt;  \&quot;meta\&quot;: { &lt;br /&gt;  \&quot;additionalKey\&quot;: \&quot;keyValue\&quot; &lt;br /&gt;  } &lt;br /&gt;  } | 

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


## DnsApiZonesPost

> Zone DnsApiZonesPost(ctx).XApiOptions(xApiOptions).CreateZoneRequestModel(createZoneRequestModel).Execute()

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
    xApiOptions := "xApiOptions_example" // string | Data used to access the API <br /><br />  Schema: <br />  { <br />  \"access_options\": { <br />  \"provider\": \"provider\", <br />  \"environment\": \"environment\", <br />  \"credentials_id\": \"00000000-0000-0000-0000-000000000000\" <br />  }, <br />  \"meta\": { <br />  \"additionalKey\": \"keyValue\" <br />  } <br />  }
    createZoneRequestModel := *openapiclient.NewCreateZoneRequestModel("Name_example") // CreateZoneRequestModel | Data used to create a DNS zone

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.DnsApiZonesPost(context.Background()).XApiOptions(xApiOptions).CreateZoneRequestModel(createZoneRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.DnsApiZonesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsApiZonesPost`: Zone
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.DnsApiZonesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsApiZonesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiOptions** | **string** | Data used to access the API &lt;br /&gt;&lt;br /&gt;  Schema: &lt;br /&gt;  { &lt;br /&gt;  \&quot;access_options\&quot;: { &lt;br /&gt;  \&quot;provider\&quot;: \&quot;provider\&quot;, &lt;br /&gt;  \&quot;environment\&quot;: \&quot;environment\&quot;, &lt;br /&gt;  \&quot;credentials_id\&quot;: \&quot;00000000-0000-0000-0000-000000000000\&quot; &lt;br /&gt;  }, &lt;br /&gt;  \&quot;meta\&quot;: { &lt;br /&gt;  \&quot;additionalKey\&quot;: \&quot;keyValue\&quot; &lt;br /&gt;  } &lt;br /&gt;  } | 
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


## DnsApiZonesZoneGet

> Zone DnsApiZonesZoneGet(ctx, zone).XApiOptions(xApiOptions).Execute()

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
    zone := "zone_example" // string | The name of the zone to query.
    xApiOptions := "xApiOptions_example" // string | Data used to access the API <br /><br />  Schema: <br />  { <br />  \"access_options\": { <br />  \"provider\": \"provider\", <br />  \"environment\": \"environment\", <br />  \"credentials_id\": \"00000000-0000-0000-0000-000000000000\" <br />  }, <br />  \"meta\": { <br />  \"additionalKey\": \"keyValue\" <br />  } <br />  }

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.DnsApiZonesZoneGet(context.Background(), zone).XApiOptions(xApiOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.DnsApiZonesZoneGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsApiZonesZoneGet`: Zone
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.DnsApiZonesZoneGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | The name of the zone to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsApiZonesZoneGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiOptions** | **string** | Data used to access the API &lt;br /&gt;&lt;br /&gt;  Schema: &lt;br /&gt;  { &lt;br /&gt;  \&quot;access_options\&quot;: { &lt;br /&gt;  \&quot;provider\&quot;: \&quot;provider\&quot;, &lt;br /&gt;  \&quot;environment\&quot;: \&quot;environment\&quot;, &lt;br /&gt;  \&quot;credentials_id\&quot;: \&quot;00000000-0000-0000-0000-000000000000\&quot; &lt;br /&gt;  }, &lt;br /&gt;  \&quot;meta\&quot;: { &lt;br /&gt;  \&quot;additionalKey\&quot;: \&quot;keyValue\&quot; &lt;br /&gt;  } &lt;br /&gt;  } | 

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

