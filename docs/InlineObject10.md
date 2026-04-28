# InlineObject10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**InlineObject10Meta**](InlineObject10Meta.md) |  | 
**Values** | [**[]InlineObject10ValuesInner**](InlineObject10ValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewInlineObject10

`func NewInlineObject10(meta InlineObject10Meta, values []InlineObject10ValuesInner, status string, ) *InlineObject10`

NewInlineObject10 instantiates a new InlineObject10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject10WithDefaults

`func NewInlineObject10WithDefaults() *InlineObject10`

NewInlineObject10WithDefaults instantiates a new InlineObject10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject10) GetMeta() InlineObject10Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject10) GetMetaOk() (*InlineObject10Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject10) SetMeta(v InlineObject10Meta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *InlineObject10) GetValues() []InlineObject10ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject10) GetValuesOk() (*[]InlineObject10ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject10) SetValues(v []InlineObject10ValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *InlineObject10) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject10) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject10) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


