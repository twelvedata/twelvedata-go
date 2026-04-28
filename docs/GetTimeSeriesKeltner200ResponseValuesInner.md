# GetTimeSeriesKeltner200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**UpperLine** | **string** | Upper line value | 
**MiddleLine** | **string** | Middle line value | 
**LowerLine** | **string** | Lower line value | 

## Methods

### NewGetTimeSeriesKeltner200ResponseValuesInner

`func NewGetTimeSeriesKeltner200ResponseValuesInner(datetime string, upperLine string, middleLine string, lowerLine string, ) *GetTimeSeriesKeltner200ResponseValuesInner`

NewGetTimeSeriesKeltner200ResponseValuesInner instantiates a new GetTimeSeriesKeltner200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesKeltner200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesKeltner200ResponseValuesInnerWithDefaults() *GetTimeSeriesKeltner200ResponseValuesInner`

NewGetTimeSeriesKeltner200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesKeltner200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetUpperLine

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetUpperLine() string`

GetUpperLine returns the UpperLine field if non-nil, zero value otherwise.

### GetUpperLineOk

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetUpperLineOk() (*string, bool)`

GetUpperLineOk returns a tuple with the UpperLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperLine

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetUpperLine(v string)`

SetUpperLine sets UpperLine field to given value.


### GetMiddleLine

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetMiddleLine() string`

GetMiddleLine returns the MiddleLine field if non-nil, zero value otherwise.

### GetMiddleLineOk

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetMiddleLineOk() (*string, bool)`

GetMiddleLineOk returns a tuple with the MiddleLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleLine

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetMiddleLine(v string)`

SetMiddleLine sets MiddleLine field to given value.


### GetLowerLine

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetLowerLine() string`

GetLowerLine returns the LowerLine field if non-nil, zero value otherwise.

### GetLowerLineOk

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetLowerLineOk() (*string, bool)`

GetLowerLineOk returns a tuple with the LowerLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerLine

`func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetLowerLine(v string)`

SetLowerLine sets LowerLine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


