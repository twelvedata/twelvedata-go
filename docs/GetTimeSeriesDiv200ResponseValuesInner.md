# GetTimeSeriesDiv200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Div** | **string** | Div value | 

## Methods

### NewGetTimeSeriesDiv200ResponseValuesInner

`func NewGetTimeSeriesDiv200ResponseValuesInner(datetime string, div string, ) *GetTimeSeriesDiv200ResponseValuesInner`

NewGetTimeSeriesDiv200ResponseValuesInner instantiates a new GetTimeSeriesDiv200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDiv200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesDiv200ResponseValuesInnerWithDefaults() *GetTimeSeriesDiv200ResponseValuesInner`

NewGetTimeSeriesDiv200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesDiv200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesDiv200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesDiv200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesDiv200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetDiv

`func (o *GetTimeSeriesDiv200ResponseValuesInner) GetDiv() string`

GetDiv returns the Div field if non-nil, zero value otherwise.

### GetDivOk

`func (o *GetTimeSeriesDiv200ResponseValuesInner) GetDivOk() (*string, bool)`

GetDivOk returns a tuple with the Div field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiv

`func (o *GetTimeSeriesDiv200ResponseValuesInner) SetDiv(v string)`

SetDiv sets Div field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


