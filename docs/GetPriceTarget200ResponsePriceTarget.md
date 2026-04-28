# GetPriceTarget200ResponsePriceTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**High** | Pointer to **float64** | Highest price target given by an analyst | [optional] 
**Median** | Pointer to **float64** | Median price target given across analysts | [optional] 
**Low** | Pointer to **float64** | Lowest price target given by an analyst | [optional] 
**Average** | Pointer to **float64** | Average price target given across analysts | [optional] 
**Current** | Pointer to **float64** | Current price from of a security | [optional] 
**Currency** | **string** | Currency in which the price targets values are quoted | 

## Methods

### NewGetPriceTarget200ResponsePriceTarget

`func NewGetPriceTarget200ResponsePriceTarget(currency string, ) *GetPriceTarget200ResponsePriceTarget`

NewGetPriceTarget200ResponsePriceTarget instantiates a new GetPriceTarget200ResponsePriceTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPriceTarget200ResponsePriceTargetWithDefaults

`func NewGetPriceTarget200ResponsePriceTargetWithDefaults() *GetPriceTarget200ResponsePriceTarget`

NewGetPriceTarget200ResponsePriceTargetWithDefaults instantiates a new GetPriceTarget200ResponsePriceTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHigh

`func (o *GetPriceTarget200ResponsePriceTarget) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *GetPriceTarget200ResponsePriceTarget) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *GetPriceTarget200ResponsePriceTarget) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *GetPriceTarget200ResponsePriceTarget) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetMedian

`func (o *GetPriceTarget200ResponsePriceTarget) GetMedian() float64`

GetMedian returns the Median field if non-nil, zero value otherwise.

### GetMedianOk

`func (o *GetPriceTarget200ResponsePriceTarget) GetMedianOk() (*float64, bool)`

GetMedianOk returns a tuple with the Median field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedian

`func (o *GetPriceTarget200ResponsePriceTarget) SetMedian(v float64)`

SetMedian sets Median field to given value.

### HasMedian

`func (o *GetPriceTarget200ResponsePriceTarget) HasMedian() bool`

HasMedian returns a boolean if a field has been set.

### GetLow

`func (o *GetPriceTarget200ResponsePriceTarget) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *GetPriceTarget200ResponsePriceTarget) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *GetPriceTarget200ResponsePriceTarget) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *GetPriceTarget200ResponsePriceTarget) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetAverage

`func (o *GetPriceTarget200ResponsePriceTarget) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *GetPriceTarget200ResponsePriceTarget) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *GetPriceTarget200ResponsePriceTarget) SetAverage(v float64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *GetPriceTarget200ResponsePriceTarget) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetCurrent

`func (o *GetPriceTarget200ResponsePriceTarget) GetCurrent() float64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetPriceTarget200ResponsePriceTarget) GetCurrentOk() (*float64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetPriceTarget200ResponsePriceTarget) SetCurrent(v float64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *GetPriceTarget200ResponsePriceTarget) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetCurrency

`func (o *GetPriceTarget200ResponsePriceTarget) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetPriceTarget200ResponsePriceTarget) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetPriceTarget200ResponsePriceTarget) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


