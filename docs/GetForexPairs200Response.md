# GetForexPairs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Count | 
**Data** | [**[]ForexResponseItem**](ForexResponseItem.md) | List of forex pairs | 
**Status** | **string** | Response status | 

## Methods

### NewGetForexPairs200Response

`func NewGetForexPairs200Response(count int64, data []ForexResponseItem, status string, ) *GetForexPairs200Response`

NewGetForexPairs200Response instantiates a new GetForexPairs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetForexPairs200ResponseWithDefaults

`func NewGetForexPairs200ResponseWithDefaults() *GetForexPairs200Response`

NewGetForexPairs200ResponseWithDefaults instantiates a new GetForexPairs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetForexPairs200Response) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetForexPairs200Response) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetForexPairs200Response) SetCount(v int64)`

SetCount sets Count field to given value.


### GetData

`func (o *GetForexPairs200Response) GetData() []ForexResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetForexPairs200Response) GetDataOk() (*[]ForexResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetForexPairs200Response) SetData(v []ForexResponseItem)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetForexPairs200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetForexPairs200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetForexPairs200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


