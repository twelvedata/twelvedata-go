# InlineObject14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**InlineObject14Meta**](InlineObject14Meta.md) |  | 
**Values** | [**[]InlineObject14ValuesInner**](InlineObject14ValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewInlineObject14

`func NewInlineObject14(meta InlineObject14Meta, values []InlineObject14ValuesInner, status string, ) *InlineObject14`

NewInlineObject14 instantiates a new InlineObject14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject14WithDefaults

`func NewInlineObject14WithDefaults() *InlineObject14`

NewInlineObject14WithDefaults instantiates a new InlineObject14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject14) GetMeta() InlineObject14Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject14) GetMetaOk() (*InlineObject14Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject14) SetMeta(v InlineObject14Meta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *InlineObject14) GetValues() []InlineObject14ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject14) GetValuesOk() (*[]InlineObject14ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject14) SetValues(v []InlineObject14ValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *InlineObject14) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject14) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject14) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


