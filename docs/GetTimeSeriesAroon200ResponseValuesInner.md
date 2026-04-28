# GetTimeSeriesAroon200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**AroonDown** | **string** | Aroon down value | 
**AroonUp** | **string** | Aroon up value | 

## Methods

### NewGetTimeSeriesAroon200ResponseValuesInner

`func NewGetTimeSeriesAroon200ResponseValuesInner(datetime string, aroonDown string, aroonUp string, ) *GetTimeSeriesAroon200ResponseValuesInner`

NewGetTimeSeriesAroon200ResponseValuesInner instantiates a new GetTimeSeriesAroon200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAroon200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAroon200ResponseValuesInnerWithDefaults() *GetTimeSeriesAroon200ResponseValuesInner`

NewGetTimeSeriesAroon200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAroon200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAroon200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAroon200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAroon200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAroonDown

`func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonDown() string`

GetAroonDown returns the AroonDown field if non-nil, zero value otherwise.

### GetAroonDownOk

`func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonDownOk() (*string, bool)`

GetAroonDownOk returns a tuple with the AroonDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAroonDown

`func (o *GetTimeSeriesAroon200ResponseValuesInner) SetAroonDown(v string)`

SetAroonDown sets AroonDown field to given value.


### GetAroonUp

`func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonUp() string`

GetAroonUp returns the AroonUp field if non-nil, zero value otherwise.

### GetAroonUpOk

`func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonUpOk() (*string, bool)`

GetAroonUpOk returns a tuple with the AroonUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAroonUp

`func (o *GetTimeSeriesAroon200ResponseValuesInner) SetAroonUp(v string)`

SetAroonUp sets AroonUp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


