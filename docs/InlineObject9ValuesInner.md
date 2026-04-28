# InlineObject9ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Asin** | **string** | ASIN value | 

## Methods

### NewInlineObject9ValuesInner

`func NewInlineObject9ValuesInner(datetime string, asin string, ) *InlineObject9ValuesInner`

NewInlineObject9ValuesInner instantiates a new InlineObject9ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject9ValuesInnerWithDefaults

`func NewInlineObject9ValuesInnerWithDefaults() *InlineObject9ValuesInner`

NewInlineObject9ValuesInnerWithDefaults instantiates a new InlineObject9ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject9ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject9ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject9ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAsin

`func (o *InlineObject9ValuesInner) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *InlineObject9ValuesInner) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *InlineObject9ValuesInner) SetAsin(v string)`

SetAsin sets Asin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


