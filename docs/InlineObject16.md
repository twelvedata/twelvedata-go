# InlineObject16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**InlineObject16Meta**](InlineObject16Meta.md) |  | 
**Values** | [**[]InlineObject16ValuesInner**](InlineObject16ValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewInlineObject16

`func NewInlineObject16(meta InlineObject16Meta, values []InlineObject16ValuesInner, status string, ) *InlineObject16`

NewInlineObject16 instantiates a new InlineObject16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject16WithDefaults

`func NewInlineObject16WithDefaults() *InlineObject16`

NewInlineObject16WithDefaults instantiates a new InlineObject16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject16) GetMeta() InlineObject16Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject16) GetMetaOk() (*InlineObject16Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject16) SetMeta(v InlineObject16Meta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *InlineObject16) GetValues() []InlineObject16ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject16) GetValuesOk() (*[]InlineObject16ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject16) SetValues(v []InlineObject16ValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *InlineObject16) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject16) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject16) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


