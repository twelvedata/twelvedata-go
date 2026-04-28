# InlineObject12ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Cosh** | **string** | COSH value | 

## Methods

### NewInlineObject12ValuesInner

`func NewInlineObject12ValuesInner(datetime string, cosh string, ) *InlineObject12ValuesInner`

NewInlineObject12ValuesInner instantiates a new InlineObject12ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject12ValuesInnerWithDefaults

`func NewInlineObject12ValuesInnerWithDefaults() *InlineObject12ValuesInner`

NewInlineObject12ValuesInnerWithDefaults instantiates a new InlineObject12ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject12ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject12ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject12ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCosh

`func (o *InlineObject12ValuesInner) GetCosh() string`

GetCosh returns the Cosh field if non-nil, zero value otherwise.

### GetCoshOk

`func (o *InlineObject12ValuesInner) GetCoshOk() (*string, bool)`

GetCoshOk returns a tuple with the Cosh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosh

`func (o *InlineObject12ValuesInner) SetCosh(v string)`

SetCosh sets Cosh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


