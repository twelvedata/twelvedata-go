# InlineObject17

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**InlineObject17Meta**](InlineObject17Meta.md) |  | 
**Values** | [**[]InlineObject17ValuesInner**](InlineObject17ValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewInlineObject17

`func NewInlineObject17(meta InlineObject17Meta, values []InlineObject17ValuesInner, status string, ) *InlineObject17`

NewInlineObject17 instantiates a new InlineObject17 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject17WithDefaults

`func NewInlineObject17WithDefaults() *InlineObject17`

NewInlineObject17WithDefaults instantiates a new InlineObject17 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject17) GetMeta() InlineObject17Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject17) GetMetaOk() (*InlineObject17Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject17) SetMeta(v InlineObject17Meta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *InlineObject17) GetValues() []InlineObject17ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject17) GetValuesOk() (*[]InlineObject17ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject17) SetValues(v []InlineObject17ValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *InlineObject17) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject17) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject17) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


