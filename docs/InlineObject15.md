# InlineObject15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**InlineObject15Meta**](InlineObject15Meta.md) |  | 
**Values** | [**[]InlineObject15ValuesInner**](InlineObject15ValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewInlineObject15

`func NewInlineObject15(meta InlineObject15Meta, values []InlineObject15ValuesInner, status string, ) *InlineObject15`

NewInlineObject15 instantiates a new InlineObject15 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject15WithDefaults

`func NewInlineObject15WithDefaults() *InlineObject15`

NewInlineObject15WithDefaults instantiates a new InlineObject15 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject15) GetMeta() InlineObject15Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject15) GetMetaOk() (*InlineObject15Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject15) SetMeta(v InlineObject15Meta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *InlineObject15) GetValues() []InlineObject15ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject15) GetValuesOk() (*[]InlineObject15ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject15) SetValues(v []InlineObject15ValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *InlineObject15) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject15) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject15) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


