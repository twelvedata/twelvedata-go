# GetTimeSeriesCross200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**CrossMeta**](CrossMeta.md) |  | 
**Values** | [**[]TimeSeriesCrossItem**](TimeSeriesCrossItem.md) | Array of time series data points | 

## Methods

### NewGetTimeSeriesCross200Response

`func NewGetTimeSeriesCross200Response(meta CrossMeta, values []TimeSeriesCrossItem, ) *GetTimeSeriesCross200Response`

NewGetTimeSeriesCross200Response instantiates a new GetTimeSeriesCross200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCross200ResponseWithDefaults

`func NewGetTimeSeriesCross200ResponseWithDefaults() *GetTimeSeriesCross200Response`

NewGetTimeSeriesCross200ResponseWithDefaults instantiates a new GetTimeSeriesCross200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesCross200Response) GetMeta() CrossMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesCross200Response) GetMetaOk() (*CrossMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesCross200Response) SetMeta(v CrossMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesCross200Response) GetValues() []TimeSeriesCrossItem`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesCross200Response) GetValuesOk() (*[]TimeSeriesCrossItem, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesCross200Response) SetValues(v []TimeSeriesCrossItem)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


