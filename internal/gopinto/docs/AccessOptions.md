# AccessOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**CredentialsId** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessOptions

`func NewAccessOptions() *AccessOptions`

NewAccessOptions instantiates a new AccessOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessOptionsWithDefaults

`func NewAccessOptionsWithDefaults() *AccessOptions`

NewAccessOptionsWithDefaults instantiates a new AccessOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AccessOptions) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AccessOptions) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AccessOptions) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AccessOptions) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetEnvironment

`func (o *AccessOptions) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AccessOptions) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AccessOptions) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AccessOptions) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentialsId

`func (o *AccessOptions) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *AccessOptions) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *AccessOptions) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.

### HasCredentialsId

`func (o *AccessOptions) HasCredentialsId() bool`

HasCredentialsId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


