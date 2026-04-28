# InlineObject14ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Sin** | **string** | SIN value | 

## Methods

### NewInlineObject14ValuesInner

`func NewInlineObject14ValuesInner(datetime string, sin string, ) *InlineObject14ValuesInner`

NewInlineObject14ValuesInner instantiates a new InlineObject14ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject14ValuesInnerWithDefaults

`func NewInlineObject14ValuesInnerWithDefaults() *InlineObject14ValuesInner`

NewInlineObject14ValuesInnerWithDefaults instantiates a new InlineObject14ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject14ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject14ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject14ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSin

`func (o *InlineObject14ValuesInner) GetSin() string`

GetSin returns the Sin field if non-nil, zero value otherwise.

### GetSinOk

`func (o *InlineObject14ValuesInner) GetSinOk() (*string, bool)`

GetSinOk returns a tuple with the Sin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSin

`func (o *InlineObject14ValuesInner) SetSin(v string)`

SetSin sets Sin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


