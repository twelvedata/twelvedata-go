# InlineObject11ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Cos** | **string** | COS value | 

## Methods

### NewInlineObject11ValuesInner

`func NewInlineObject11ValuesInner(datetime string, cos string, ) *InlineObject11ValuesInner`

NewInlineObject11ValuesInner instantiates a new InlineObject11ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject11ValuesInnerWithDefaults

`func NewInlineObject11ValuesInnerWithDefaults() *InlineObject11ValuesInner`

NewInlineObject11ValuesInnerWithDefaults instantiates a new InlineObject11ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject11ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject11ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject11ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCos

`func (o *InlineObject11ValuesInner) GetCos() string`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *InlineObject11ValuesInner) GetCosOk() (*string, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *InlineObject11ValuesInner) SetCos(v string)`

SetCos sets Cos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


