# ApiOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessOptions** | Pointer to [**AccessOptions**](AccessOptions.md) |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApiOptions

`func NewApiOptions() *ApiOptions`

NewApiOptions instantiates a new ApiOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOptionsWithDefaults

`func NewApiOptionsWithDefaults() *ApiOptions`

NewApiOptionsWithDefaults instantiates a new ApiOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessOptions

`func (o *ApiOptions) GetAccessOptions() AccessOptions`

GetAccessOptions returns the AccessOptions field if non-nil, zero value otherwise.

### GetAccessOptionsOk

`func (o *ApiOptions) GetAccessOptionsOk() (*AccessOptions, bool)`

GetAccessOptionsOk returns a tuple with the AccessOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOptions

`func (o *ApiOptions) SetAccessOptions(v AccessOptions)`

SetAccessOptions sets AccessOptions field to given value.

### HasAccessOptions

`func (o *ApiOptions) HasAccessOptions() bool`

HasAccessOptions returns a boolean if a field has been set.

### GetMeta

`func (o *ApiOptions) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ApiOptions) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ApiOptions) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ApiOptions) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


