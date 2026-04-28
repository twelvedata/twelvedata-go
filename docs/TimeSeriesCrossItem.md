# TimeSeriesCrossItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime at local exchange time referring to when the bar with specified interval was opened | 
**Open** | **string** | Price at the opening of the current bar | 
**High** | **string** | Highest price which occurred during the current bar | 
**Low** | **string** | Lowest price which occurred during the current bar | 
**Close** | **string** | Close price at the end of the bar | 

## Methods

### NewTimeSeriesCrossItem

`func NewTimeSeriesCrossItem(datetime string, open string, high string, low string, close string, ) *TimeSeriesCrossItem`

NewTimeSeriesCrossItem instantiates a new TimeSeriesCrossItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesCrossItemWithDefaults

`func NewTimeSeriesCrossItemWithDefaults() *TimeSeriesCrossItem`

NewTimeSeriesCrossItemWithDefaults instantiates a new TimeSeriesCrossItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *TimeSeriesCrossItem) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *TimeSeriesCrossItem) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *TimeSeriesCrossItem) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetOpen

`func (o *TimeSeriesCrossItem) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *TimeSeriesCrossItem) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *TimeSeriesCrossItem) SetOpen(v string)`

SetOpen sets Open field to given value.


### GetHigh

`func (o *TimeSeriesCrossItem) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *TimeSeriesCrossItem) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *TimeSeriesCrossItem) SetHigh(v string)`

SetHigh sets High field to given value.


### GetLow

`func (o *TimeSeriesCrossItem) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *TimeSeriesCrossItem) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *TimeSeriesCrossItem) SetLow(v string)`

SetLow sets Low field to given value.


### GetClose

`func (o *TimeSeriesCrossItem) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *TimeSeriesCrossItem) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *TimeSeriesCrossItem) SetClose(v string)`

SetClose sets Close field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


