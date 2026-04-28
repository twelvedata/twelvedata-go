# GetCommodities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Count | 
**Data** | [**[]CommoditiesResponseItem**](CommoditiesResponseItem.md) | List of commodities | 
**Status** | **string** | Response status | 

## Methods

### NewGetCommodities200Response

`func NewGetCommodities200Response(count int64, data []CommoditiesResponseItem, status string, ) *GetCommodities200Response`

NewGetCommodities200Response instantiates a new GetCommodities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommodities200ResponseWithDefaults

`func NewGetCommodities200ResponseWithDefaults() *GetCommodities200Response`

NewGetCommodities200ResponseWithDefaults instantiates a new GetCommodities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetCommodities200Response) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetCommodities200Response) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetCommodities200Response) SetCount(v int64)`

SetCount sets Count field to given value.


### GetData

`func (o *GetCommodities200Response) GetData() []CommoditiesResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCommodities200Response) GetDataOk() (*[]CommoditiesResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCommodities200Response) SetData(v []CommoditiesResponseItem)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetCommodities200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommodities200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommodities200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


