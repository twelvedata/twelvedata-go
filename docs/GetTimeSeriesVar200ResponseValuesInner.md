# GetTimeSeriesVar200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Var** | **string** | VAR value | 

## Methods

### NewGetTimeSeriesVar200ResponseValuesInner

`func NewGetTimeSeriesVar200ResponseValuesInner(datetime string, var_ string, ) *GetTimeSeriesVar200ResponseValuesInner`

NewGetTimeSeriesVar200ResponseValuesInner instantiates a new GetTimeSeriesVar200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesVar200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesVar200ResponseValuesInnerWithDefaults() *GetTimeSeriesVar200ResponseValuesInner`

NewGetTimeSeriesVar200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesVar200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesVar200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesVar200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesVar200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetVar

`func (o *GetTimeSeriesVar200ResponseValuesInner) GetVar() string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *GetTimeSeriesVar200ResponseValuesInner) GetVarOk() (*string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *GetTimeSeriesVar200ResponseValuesInner) SetVar(v string)`

SetVar sets Var field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


