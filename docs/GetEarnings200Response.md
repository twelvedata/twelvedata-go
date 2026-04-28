# GetEarnings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetEarnings200ResponseMeta**](GetEarnings200ResponseMeta.md) |  | 
**Earnings** | [**[]EarningsItem**](EarningsItem.md) | List of earnings data | 
**Status** | **string** | Response status | 

## Methods

### NewGetEarnings200Response

`func NewGetEarnings200Response(meta GetEarnings200ResponseMeta, earnings []EarningsItem, status string, ) *GetEarnings200Response`

NewGetEarnings200Response instantiates a new GetEarnings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEarnings200ResponseWithDefaults

`func NewGetEarnings200ResponseWithDefaults() *GetEarnings200Response`

NewGetEarnings200ResponseWithDefaults instantiates a new GetEarnings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetEarnings200Response) GetMeta() GetEarnings200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetEarnings200Response) GetMetaOk() (*GetEarnings200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetEarnings200Response) SetMeta(v GetEarnings200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetEarnings

`func (o *GetEarnings200Response) GetEarnings() []EarningsItem`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *GetEarnings200Response) GetEarningsOk() (*[]EarningsItem, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *GetEarnings200Response) SetEarnings(v []EarningsItem)`

SetEarnings sets Earnings field to given value.


### GetStatus

`func (o *GetEarnings200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEarnings200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEarnings200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


