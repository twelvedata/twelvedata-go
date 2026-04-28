# GetInstrumentType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **[]string** | List of instrument types available at Twelve Data API. | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetInstrumentType200Response

`func NewGetInstrumentType200Response(result []string, status string, ) *GetInstrumentType200Response`

NewGetInstrumentType200Response instantiates a new GetInstrumentType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstrumentType200ResponseWithDefaults

`func NewGetInstrumentType200ResponseWithDefaults() *GetInstrumentType200Response`

NewGetInstrumentType200ResponseWithDefaults instantiates a new GetInstrumentType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetInstrumentType200Response) GetResult() []string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetInstrumentType200Response) GetResultOk() (*[]string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetInstrumentType200Response) SetResult(v []string)`

SetResult sets Result field to given value.


### GetStatus

`func (o *GetInstrumentType200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInstrumentType200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInstrumentType200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


