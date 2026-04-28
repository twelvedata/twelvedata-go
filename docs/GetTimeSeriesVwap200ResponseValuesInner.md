# GetTimeSeriesVwap200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**VwapLower** | Pointer to **string** | VWAP lower value | [optional] 
**Vwap** | **string** | VWAP value | 
**VwapUpper** | Pointer to **string** | VWAP upper value | [optional] 

## Methods

### NewGetTimeSeriesVwap200ResponseValuesInner

`func NewGetTimeSeriesVwap200ResponseValuesInner(datetime string, vwap string, ) *GetTimeSeriesVwap200ResponseValuesInner`

NewGetTimeSeriesVwap200ResponseValuesInner instantiates a new GetTimeSeriesVwap200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesVwap200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesVwap200ResponseValuesInnerWithDefaults() *GetTimeSeriesVwap200ResponseValuesInner`

NewGetTimeSeriesVwap200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesVwap200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesVwap200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetVwapLower

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapLower() string`

GetVwapLower returns the VwapLower field if non-nil, zero value otherwise.

### GetVwapLowerOk

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapLowerOk() (*string, bool)`

GetVwapLowerOk returns a tuple with the VwapLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVwapLower

`func (o *GetTimeSeriesVwap200ResponseValuesInner) SetVwapLower(v string)`

SetVwapLower sets VwapLower field to given value.

### HasVwapLower

`func (o *GetTimeSeriesVwap200ResponseValuesInner) HasVwapLower() bool`

HasVwapLower returns a boolean if a field has been set.

### GetVwap

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwap() string`

GetVwap returns the Vwap field if non-nil, zero value otherwise.

### GetVwapOk

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapOk() (*string, bool)`

GetVwapOk returns a tuple with the Vwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVwap

`func (o *GetTimeSeriesVwap200ResponseValuesInner) SetVwap(v string)`

SetVwap sets Vwap field to given value.


### GetVwapUpper

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapUpper() string`

GetVwapUpper returns the VwapUpper field if non-nil, zero value otherwise.

### GetVwapUpperOk

`func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapUpperOk() (*string, bool)`

GetVwapUpperOk returns a tuple with the VwapUpper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVwapUpper

`func (o *GetTimeSeriesVwap200ResponseValuesInner) SetVwapUpper(v string)`

SetVwapUpper sets VwapUpper field to given value.

### HasVwapUpper

`func (o *GetTimeSeriesVwap200ResponseValuesInner) HasVwapUpper() bool`

HasVwapUpper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


