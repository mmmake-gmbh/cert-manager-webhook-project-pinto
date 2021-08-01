# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**RecordType**](RecordType.md) |  | 
**Class** | [**RecordClass**](RecordClass.md) |  | 
**Ttl** | Pointer to **int32** |  | [optional] 
**Data** | **string** |  | 

## Methods

### NewRecord

`func NewRecord(name string, type_ RecordType, class RecordClass, data string, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Record) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Record) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Record) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Record) GetType() RecordType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Record) GetTypeOk() (*RecordType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Record) SetType(v RecordType)`

SetType sets Type field to given value.


### GetClass

`func (o *Record) GetClass() RecordClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Record) GetClassOk() (*RecordClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Record) SetClass(v RecordClass)`

SetClass sets Class field to given value.


### GetTtl

`func (o *Record) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Record) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Record) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Record) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetData

`func (o *Record) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Record) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Record) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


