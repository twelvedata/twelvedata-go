# GetDirectHolders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetDirectHolders200ResponseMeta**](GetDirectHolders200ResponseMeta.md) |  | 
**DirectHolders** | [**[]DirectHolderItem**](DirectHolderItem.md) | List of direct holders for the financial instrument | 

## Methods

### NewGetDirectHolders200Response

`func NewGetDirectHolders200Response(meta GetDirectHolders200ResponseMeta, directHolders []DirectHolderItem, ) *GetDirectHolders200Response`

NewGetDirectHolders200Response instantiates a new GetDirectHolders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDirectHolders200ResponseWithDefaults

`func NewGetDirectHolders200ResponseWithDefaults() *GetDirectHolders200Response`

NewGetDirectHolders200ResponseWithDefaults instantiates a new GetDirectHolders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDirectHolders200Response) GetMeta() GetDirectHolders200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDirectHolders200Response) GetMetaOk() (*GetDirectHolders200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDirectHolders200Response) SetMeta(v GetDirectHolders200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetDirectHolders

`func (o *GetDirectHolders200Response) GetDirectHolders() []DirectHolderItem`

GetDirectHolders returns the DirectHolders field if non-nil, zero value otherwise.

### GetDirectHoldersOk

`func (o *GetDirectHolders200Response) GetDirectHoldersOk() (*[]DirectHolderItem, bool)`

GetDirectHoldersOk returns a tuple with the DirectHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectHolders

`func (o *GetDirectHolders200Response) SetDirectHolders(v []DirectHolderItem)`

SetDirectHolders sets DirectHolders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


