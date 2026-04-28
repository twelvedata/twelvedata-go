# GetIntervals200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **[]string** | List of available intervals | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetIntervals200Response

`func NewGetIntervals200Response(data []string, status string, ) *GetIntervals200Response`

NewGetIntervals200Response instantiates a new GetIntervals200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntervals200ResponseWithDefaults

`func NewGetIntervals200ResponseWithDefaults() *GetIntervals200Response`

NewGetIntervals200ResponseWithDefaults instantiates a new GetIntervals200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetIntervals200Response) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIntervals200Response) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIntervals200Response) SetData(v []string)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetIntervals200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIntervals200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIntervals200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


