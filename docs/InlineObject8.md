# InlineObject8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**InlineObject8Meta**](InlineObject8Meta.md) |  | 
**Values** | [**[]InlineObject8ValuesInner**](InlineObject8ValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewInlineObject8

`func NewInlineObject8(meta InlineObject8Meta, values []InlineObject8ValuesInner, status string, ) *InlineObject8`

NewInlineObject8 instantiates a new InlineObject8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject8WithDefaults

`func NewInlineObject8WithDefaults() *InlineObject8`

NewInlineObject8WithDefaults instantiates a new InlineObject8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject8) GetMeta() InlineObject8Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject8) GetMetaOk() (*InlineObject8Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject8) SetMeta(v InlineObject8Meta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *InlineObject8) GetValues() []InlineObject8ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject8) GetValuesOk() (*[]InlineObject8ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject8) SetValues(v []InlineObject8ValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *InlineObject8) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject8) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject8) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


