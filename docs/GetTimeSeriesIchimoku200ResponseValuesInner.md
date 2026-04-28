# GetTimeSeriesIchimoku200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**TenkanSen** | Pointer to **string** | Tenkan-sen value | [optional] 
**KijunSen** | Pointer to **string** | Kijun-sen value | [optional] 
**SenkouSpanA** | **string** | Senkou span A value | 
**SenkouSpanB** | **string** | Senkou span B value | 
**ChikouSpan** | Pointer to **string** | Chikou span value | [optional] 

## Methods

### NewGetTimeSeriesIchimoku200ResponseValuesInner

`func NewGetTimeSeriesIchimoku200ResponseValuesInner(datetime string, senkouSpanA string, senkouSpanB string, ) *GetTimeSeriesIchimoku200ResponseValuesInner`

NewGetTimeSeriesIchimoku200ResponseValuesInner instantiates a new GetTimeSeriesIchimoku200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesIchimoku200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesIchimoku200ResponseValuesInnerWithDefaults() *GetTimeSeriesIchimoku200ResponseValuesInner`

NewGetTimeSeriesIchimoku200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesIchimoku200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTenkanSen

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetTenkanSen() string`

GetTenkanSen returns the TenkanSen field if non-nil, zero value otherwise.

### GetTenkanSenOk

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetTenkanSenOk() (*string, bool)`

GetTenkanSenOk returns a tuple with the TenkanSen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenkanSen

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetTenkanSen(v string)`

SetTenkanSen sets TenkanSen field to given value.

### HasTenkanSen

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) HasTenkanSen() bool`

HasTenkanSen returns a boolean if a field has been set.

### GetKijunSen

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetKijunSen() string`

GetKijunSen returns the KijunSen field if non-nil, zero value otherwise.

### GetKijunSenOk

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetKijunSenOk() (*string, bool)`

GetKijunSenOk returns a tuple with the KijunSen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKijunSen

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetKijunSen(v string)`

SetKijunSen sets KijunSen field to given value.

### HasKijunSen

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) HasKijunSen() bool`

HasKijunSen returns a boolean if a field has been set.

### GetSenkouSpanA

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanA() string`

GetSenkouSpanA returns the SenkouSpanA field if non-nil, zero value otherwise.

### GetSenkouSpanAOk

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanAOk() (*string, bool)`

GetSenkouSpanAOk returns a tuple with the SenkouSpanA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenkouSpanA

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetSenkouSpanA(v string)`

SetSenkouSpanA sets SenkouSpanA field to given value.


### GetSenkouSpanB

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanB() string`

GetSenkouSpanB returns the SenkouSpanB field if non-nil, zero value otherwise.

### GetSenkouSpanBOk

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetSenkouSpanBOk() (*string, bool)`

GetSenkouSpanBOk returns a tuple with the SenkouSpanB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenkouSpanB

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetSenkouSpanB(v string)`

SetSenkouSpanB sets SenkouSpanB field to given value.


### GetChikouSpan

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetChikouSpan() string`

GetChikouSpan returns the ChikouSpan field if non-nil, zero value otherwise.

### GetChikouSpanOk

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) GetChikouSpanOk() (*string, bool)`

GetChikouSpanOk returns a tuple with the ChikouSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChikouSpan

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) SetChikouSpan(v string)`

SetChikouSpan sets ChikouSpan field to given value.

### HasChikouSpan

`func (o *GetTimeSeriesIchimoku200ResponseValuesInner) HasChikouSpan() bool`

HasChikouSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


