# InlineObject15ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Sinh** | **string** | SINH value | 

## Methods

### NewInlineObject15ValuesInner

`func NewInlineObject15ValuesInner(datetime string, sinh string, ) *InlineObject15ValuesInner`

NewInlineObject15ValuesInner instantiates a new InlineObject15ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject15ValuesInnerWithDefaults

`func NewInlineObject15ValuesInnerWithDefaults() *InlineObject15ValuesInner`

NewInlineObject15ValuesInnerWithDefaults instantiates a new InlineObject15ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject15ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject15ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject15ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSinh

`func (o *InlineObject15ValuesInner) GetSinh() string`

GetSinh returns the Sinh field if non-nil, zero value otherwise.

### GetSinhOk

`func (o *InlineObject15ValuesInner) GetSinhOk() (*string, bool)`

GetSinhOk returns a tuple with the Sinh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinh

`func (o *InlineObject15ValuesInner) SetSinh(v string)`

SetSinh sets Sinh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


