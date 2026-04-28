# GetTechnicalIndicators200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**map[string]GetTechnicalIndicators200ResponseDataValue**](GetTechnicalIndicators200ResponseDataValue.md) | Map of technical indicators available at Twelve Data API | 
**Status** | **string** | Response status | 

## Methods

### NewGetTechnicalIndicators200Response

`func NewGetTechnicalIndicators200Response(data map[string]GetTechnicalIndicators200ResponseDataValue, status string, ) *GetTechnicalIndicators200Response`

NewGetTechnicalIndicators200Response instantiates a new GetTechnicalIndicators200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTechnicalIndicators200ResponseWithDefaults

`func NewGetTechnicalIndicators200ResponseWithDefaults() *GetTechnicalIndicators200Response`

NewGetTechnicalIndicators200ResponseWithDefaults instantiates a new GetTechnicalIndicators200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTechnicalIndicators200Response) GetData() map[string]GetTechnicalIndicators200ResponseDataValue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTechnicalIndicators200Response) GetDataOk() (*map[string]GetTechnicalIndicators200ResponseDataValue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTechnicalIndicators200Response) SetData(v map[string]GetTechnicalIndicators200ResponseDataValue)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetTechnicalIndicators200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTechnicalIndicators200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTechnicalIndicators200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


