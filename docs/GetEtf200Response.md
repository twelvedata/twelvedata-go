# GetEtf200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Count | 
**Data** | [**[]EtfResponseItem**](EtfResponseItem.md) | List of ETFs | 
**Status** | **string** | Response status | 

## Methods

### NewGetEtf200Response

`func NewGetEtf200Response(count int64, data []EtfResponseItem, status string, ) *GetEtf200Response`

NewGetEtf200Response instantiates a new GetEtf200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEtf200ResponseWithDefaults

`func NewGetEtf200ResponseWithDefaults() *GetEtf200Response`

NewGetEtf200ResponseWithDefaults instantiates a new GetEtf200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetEtf200Response) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetEtf200Response) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetEtf200Response) SetCount(v int64)`

SetCount sets Count field to given value.


### GetData

`func (o *GetEtf200Response) GetData() []EtfResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEtf200Response) GetDataOk() (*[]EtfResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEtf200Response) SetData(v []EtfResponseItem)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetEtf200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEtf200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEtf200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


