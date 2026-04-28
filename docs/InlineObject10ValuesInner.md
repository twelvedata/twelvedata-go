# InlineObject10ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Atan** | **string** | ATAN value | 

## Methods

### NewInlineObject10ValuesInner

`func NewInlineObject10ValuesInner(datetime string, atan string, ) *InlineObject10ValuesInner`

NewInlineObject10ValuesInner instantiates a new InlineObject10ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject10ValuesInnerWithDefaults

`func NewInlineObject10ValuesInnerWithDefaults() *InlineObject10ValuesInner`

NewInlineObject10ValuesInnerWithDefaults instantiates a new InlineObject10ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject10ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject10ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject10ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAtan

`func (o *InlineObject10ValuesInner) GetAtan() string`

GetAtan returns the Atan field if non-nil, zero value otherwise.

### GetAtanOk

`func (o *InlineObject10ValuesInner) GetAtanOk() (*string, bool)`

GetAtanOk returns a tuple with the Atan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtan

`func (o *InlineObject10ValuesInner) SetAtan(v string)`

SetAtan sets Atan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


