# InlineObject9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**InlineObject9Meta**](InlineObject9Meta.md) |  | 
**Values** | [**[]InlineObject9ValuesInner**](InlineObject9ValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewInlineObject9

`func NewInlineObject9(meta InlineObject9Meta, values []InlineObject9ValuesInner, status string, ) *InlineObject9`

NewInlineObject9 instantiates a new InlineObject9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject9WithDefaults

`func NewInlineObject9WithDefaults() *InlineObject9`

NewInlineObject9WithDefaults instantiates a new InlineObject9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject9) GetMeta() InlineObject9Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject9) GetMetaOk() (*InlineObject9Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject9) SetMeta(v InlineObject9Meta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *InlineObject9) GetValues() []InlineObject9ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject9) GetValuesOk() (*[]InlineObject9ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject9) SetValues(v []InlineObject9ValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *InlineObject9) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject9) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject9) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


