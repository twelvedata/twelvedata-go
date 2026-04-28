# GetPriceTarget200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetPriceTarget200ResponseMeta**](GetPriceTarget200ResponseMeta.md) |  | 
**PriceTarget** | [**GetPriceTarget200ResponsePriceTarget**](GetPriceTarget200ResponsePriceTarget.md) |  | 
**Status** | **string** | Response status | 

## Methods

### NewGetPriceTarget200Response

`func NewGetPriceTarget200Response(meta GetPriceTarget200ResponseMeta, priceTarget GetPriceTarget200ResponsePriceTarget, status string, ) *GetPriceTarget200Response`

NewGetPriceTarget200Response instantiates a new GetPriceTarget200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPriceTarget200ResponseWithDefaults

`func NewGetPriceTarget200ResponseWithDefaults() *GetPriceTarget200Response`

NewGetPriceTarget200ResponseWithDefaults instantiates a new GetPriceTarget200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetPriceTarget200Response) GetMeta() GetPriceTarget200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetPriceTarget200Response) GetMetaOk() (*GetPriceTarget200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetPriceTarget200Response) SetMeta(v GetPriceTarget200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetPriceTarget

`func (o *GetPriceTarget200Response) GetPriceTarget() GetPriceTarget200ResponsePriceTarget`

GetPriceTarget returns the PriceTarget field if non-nil, zero value otherwise.

### GetPriceTargetOk

`func (o *GetPriceTarget200Response) GetPriceTargetOk() (*GetPriceTarget200ResponsePriceTarget, bool)`

GetPriceTargetOk returns a tuple with the PriceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTarget

`func (o *GetPriceTarget200Response) SetPriceTarget(v GetPriceTarget200ResponsePriceTarget)`

SetPriceTarget sets PriceTarget field to given value.


### GetStatus

`func (o *GetPriceTarget200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPriceTarget200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPriceTarget200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


