# GetTaxInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTaxInfo200ResponseMeta**](GetTaxInfo200ResponseMeta.md) |  | 
**Data** | [**GetTaxInfo200ResponseData**](GetTaxInfo200ResponseData.md) |  | 
**Status** | **string** | The status of the request, e.g., &#x60;ok&#x60;, &#x60;error&#x60; | 

## Methods

### NewGetTaxInfo200Response

`func NewGetTaxInfo200Response(meta GetTaxInfo200ResponseMeta, data GetTaxInfo200ResponseData, status string, ) *GetTaxInfo200Response`

NewGetTaxInfo200Response instantiates a new GetTaxInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxInfo200ResponseWithDefaults

`func NewGetTaxInfo200ResponseWithDefaults() *GetTaxInfo200Response`

NewGetTaxInfo200ResponseWithDefaults instantiates a new GetTaxInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTaxInfo200Response) GetMeta() GetTaxInfo200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTaxInfo200Response) GetMetaOk() (*GetTaxInfo200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTaxInfo200Response) SetMeta(v GetTaxInfo200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *GetTaxInfo200Response) GetData() GetTaxInfo200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTaxInfo200Response) GetDataOk() (*GetTaxInfo200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTaxInfo200Response) SetData(v GetTaxInfo200ResponseData)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetTaxInfo200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTaxInfo200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTaxInfo200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


