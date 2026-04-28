# GetSplits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetSplits200ResponseMeta**](GetSplits200ResponseMeta.md) |  | 
**Splits** | [**[]GetSplits200ResponseSplitsInner**](GetSplits200ResponseSplitsInner.md) | List of stock splits | 

## Methods

### NewGetSplits200Response

`func NewGetSplits200Response(meta GetSplits200ResponseMeta, splits []GetSplits200ResponseSplitsInner, ) *GetSplits200Response`

NewGetSplits200Response instantiates a new GetSplits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSplits200ResponseWithDefaults

`func NewGetSplits200ResponseWithDefaults() *GetSplits200Response`

NewGetSplits200ResponseWithDefaults instantiates a new GetSplits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetSplits200Response) GetMeta() GetSplits200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSplits200Response) GetMetaOk() (*GetSplits200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSplits200Response) SetMeta(v GetSplits200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetSplits

`func (o *GetSplits200Response) GetSplits() []GetSplits200ResponseSplitsInner`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *GetSplits200Response) GetSplitsOk() (*[]GetSplits200ResponseSplitsInner, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *GetSplits200Response) SetSplits(v []GetSplits200ResponseSplitsInner)`

SetSplits sets Splits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


