# InlineObject16ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Tan** | **string** | TAN value | 

## Methods

### NewInlineObject16ValuesInner

`func NewInlineObject16ValuesInner(datetime string, tan string, ) *InlineObject16ValuesInner`

NewInlineObject16ValuesInner instantiates a new InlineObject16ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject16ValuesInnerWithDefaults

`func NewInlineObject16ValuesInnerWithDefaults() *InlineObject16ValuesInner`

NewInlineObject16ValuesInnerWithDefaults instantiates a new InlineObject16ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject16ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject16ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject16ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTan

`func (o *InlineObject16ValuesInner) GetTan() string`

GetTan returns the Tan field if non-nil, zero value otherwise.

### GetTanOk

`func (o *InlineObject16ValuesInner) GetTanOk() (*string, bool)`

GetTanOk returns a tuple with the Tan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTan

`func (o *InlineObject16ValuesInner) SetTan(v string)`

SetTan sets Tan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


