# GetKeyExecutives200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetKeyExecutives200ResponseMeta**](GetKeyExecutives200ResponseMeta.md) |  | 
**KeyExecutives** | Pointer to [**[]GetKeyExecutives200ResponseKeyExecutivesInner**](GetKeyExecutives200ResponseKeyExecutivesInner.md) | List of key executives | [optional] 

## Methods

### NewGetKeyExecutives200Response

`func NewGetKeyExecutives200Response(meta GetKeyExecutives200ResponseMeta, ) *GetKeyExecutives200Response`

NewGetKeyExecutives200Response instantiates a new GetKeyExecutives200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKeyExecutives200ResponseWithDefaults

`func NewGetKeyExecutives200ResponseWithDefaults() *GetKeyExecutives200Response`

NewGetKeyExecutives200ResponseWithDefaults instantiates a new GetKeyExecutives200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetKeyExecutives200Response) GetMeta() GetKeyExecutives200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKeyExecutives200Response) GetMetaOk() (*GetKeyExecutives200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKeyExecutives200Response) SetMeta(v GetKeyExecutives200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetKeyExecutives

`func (o *GetKeyExecutives200Response) GetKeyExecutives() []GetKeyExecutives200ResponseKeyExecutivesInner`

GetKeyExecutives returns the KeyExecutives field if non-nil, zero value otherwise.

### GetKeyExecutivesOk

`func (o *GetKeyExecutives200Response) GetKeyExecutivesOk() (*[]GetKeyExecutives200ResponseKeyExecutivesInner, bool)`

GetKeyExecutivesOk returns a tuple with the KeyExecutives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExecutives

`func (o *GetKeyExecutives200Response) SetKeyExecutives(v []GetKeyExecutives200ResponseKeyExecutivesInner)`

SetKeyExecutives sets KeyExecutives field to given value.

### HasKeyExecutives

`func (o *GetKeyExecutives200Response) HasKeyExecutives() bool`

HasKeyExecutives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


