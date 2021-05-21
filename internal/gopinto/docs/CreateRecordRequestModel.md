# CreateRecordRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Environment** | Pointer to **NullableString** |  | [optional] 
**Zone** | **string** |  | 
**Name** | **string** |  | 
**Class** | Pointer to [**RecordClass**](RecordClass.md) |  | [optional] 
**Type** | [**RecordType**](RecordType.md) |  | 
**Data** | **string** |  | 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateRecordRequestModel

`func NewCreateRecordRequestModel(provider string, zone string, name string, type_ RecordType, data string, ) *CreateRecordRequestModel`

NewCreateRecordRequestModel instantiates a new CreateRecordRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecordRequestModelWithDefaults

`func NewCreateRecordRequestModelWithDefaults() *CreateRecordRequestModel`

NewCreateRecordRequestModelWithDefaults instantiates a new CreateRecordRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CreateRecordRequestModel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateRecordRequestModel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateRecordRequestModel) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetEnvironment

`func (o *CreateRecordRequestModel) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateRecordRequestModel) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateRecordRequestModel) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateRecordRequestModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CreateRecordRequestModel) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CreateRecordRequestModel) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetZone

`func (o *CreateRecordRequestModel) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateRecordRequestModel) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateRecordRequestModel) SetZone(v string)`

SetZone sets Zone field to given value.


### GetName

`func (o *CreateRecordRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRecordRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRecordRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetClass

`func (o *CreateRecordRequestModel) GetClass() RecordClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *CreateRecordRequestModel) GetClassOk() (*RecordClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *CreateRecordRequestModel) SetClass(v RecordClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *CreateRecordRequestModel) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetType

`func (o *CreateRecordRequestModel) GetType() RecordType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRecordRequestModel) GetTypeOk() (*RecordType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRecordRequestModel) SetType(v RecordType)`

SetType sets Type field to given value.


### GetData

`func (o *CreateRecordRequestModel) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRecordRequestModel) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRecordRequestModel) SetData(v string)`

SetData sets Data field to given value.


### GetTtl

`func (o *CreateRecordRequestModel) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateRecordRequestModel) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateRecordRequestModel) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateRecordRequestModel) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


