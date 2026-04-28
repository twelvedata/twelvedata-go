# GetEarningsCalendar200ResponseEarningsValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Instrument symbol (ticker) | [optional] 
**Name** | Pointer to **string** | Full name of instrument | [optional] 
**Currency** | Pointer to **string** | Currency in which instrument is traded by ISO 4217 standard | [optional] 
**Exchange** | Pointer to **string** | Exchange where instrument is traded | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Country** | Pointer to **string** | Country where exchange is located | [optional] 
**Time** | Pointer to **string** | Can be either of the following values: &#x60;Pre Market&#x60;, &#x60;After Hours&#x60;, &#x60;Time Not Supplied&#x60; | [optional] 
**EpsEstimate** | Pointer to **float64** | Analyst estimate of the future company earning | [optional] 
**EpsActual** | Pointer to **float64** | Actual value of reported earning | [optional] 
**Difference** | Pointer to **float64** | Delta between &#x60;eps_actual&#x60; and &#x60;eps_estimate&#x60; | [optional] 
**SurprisePrc** | Pointer to **float64** | Surprise in percentage of the &#x60;eps_actual&#x60; related to &#x60;eps_estimate&#x60; | [optional] 

## Methods

### NewGetEarningsCalendar200ResponseEarningsValueInner

`func NewGetEarningsCalendar200ResponseEarningsValueInner() *GetEarningsCalendar200ResponseEarningsValueInner`

NewGetEarningsCalendar200ResponseEarningsValueInner instantiates a new GetEarningsCalendar200ResponseEarningsValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEarningsCalendar200ResponseEarningsValueInnerWithDefaults

`func NewGetEarningsCalendar200ResponseEarningsValueInnerWithDefaults() *GetEarningsCalendar200ResponseEarningsValueInner`

NewGetEarningsCalendar200ResponseEarningsValueInnerWithDefaults instantiates a new GetEarningsCalendar200ResponseEarningsValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetCountry

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetTime

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetEpsEstimate

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsEstimate() float64`

GetEpsEstimate returns the EpsEstimate field if non-nil, zero value otherwise.

### GetEpsEstimateOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsEstimateOk() (*float64, bool)`

GetEpsEstimateOk returns a tuple with the EpsEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsEstimate

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetEpsEstimate(v float64)`

SetEpsEstimate sets EpsEstimate field to given value.

### HasEpsEstimate

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasEpsEstimate() bool`

HasEpsEstimate returns a boolean if a field has been set.

### GetEpsActual

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsActual() float64`

GetEpsActual returns the EpsActual field if non-nil, zero value otherwise.

### GetEpsActualOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetEpsActualOk() (*float64, bool)`

GetEpsActualOk returns a tuple with the EpsActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsActual

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetEpsActual(v float64)`

SetEpsActual sets EpsActual field to given value.

### HasEpsActual

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasEpsActual() bool`

HasEpsActual returns a boolean if a field has been set.

### GetDifference

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetDifference() float64`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetDifferenceOk() (*float64, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetDifference(v float64)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### GetSurprisePrc

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSurprisePrc() float64`

GetSurprisePrc returns the SurprisePrc field if non-nil, zero value otherwise.

### GetSurprisePrcOk

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) GetSurprisePrcOk() (*float64, bool)`

GetSurprisePrcOk returns a tuple with the SurprisePrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurprisePrc

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) SetSurprisePrc(v float64)`

SetSurprisePrc sets SurprisePrc field to given value.

### HasSurprisePrc

`func (o *GetEarningsCalendar200ResponseEarningsValueInner) HasSurprisePrc() bool`

HasSurprisePrc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


