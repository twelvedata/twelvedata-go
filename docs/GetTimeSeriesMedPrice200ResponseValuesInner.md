# GetTimeSeriesMedPrice200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Medprice** | **string** | Medprice value | 

## Methods

### NewGetTimeSeriesMedPrice200ResponseValuesInner

`func NewGetTimeSeriesMedPrice200ResponseValuesInner(datetime string, medprice string, ) *GetTimeSeriesMedPrice200ResponseValuesInner`

NewGetTimeSeriesMedPrice200ResponseValuesInner instantiates a new GetTimeSeriesMedPrice200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMedPrice200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMedPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesMedPrice200ResponseValuesInner`

NewGetTimeSeriesMedPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMedPrice200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMedPrice200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMedprice

`func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetMedprice() string`

GetMedprice returns the Medprice field if non-nil, zero value otherwise.

### GetMedpriceOk

`func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetMedpriceOk() (*string, bool)`

GetMedpriceOk returns a tuple with the Medprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedprice

`func (o *GetTimeSeriesMedPrice200ResponseValuesInner) SetMedprice(v string)`

SetMedprice sets Medprice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


