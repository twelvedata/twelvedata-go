# GetStatistics200ResponseStatisticsStockStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharesOutstanding** | Pointer to **float64** | Refers for the shares outstanding currently held by all its shareholders | [optional] 
**FloatShares** | Pointer to **float64** | Refers to floating stock is the number of public shares a company has available for trading on the open market | [optional] 
**Avg10Volume** | Pointer to **int64** | Refers to the average 10 days volume | [optional] 
**Avg90Volume** | Pointer to **int64** | Refers to the average 90 days volume | [optional] 
**SharesShort** | Pointer to **int64** | Refers to the number of shares that have been shorted | [optional] 
**ShortRatio** | Pointer to **float64** | Refers to short ratio measure | [optional] 
**ShortPercentOfSharesOutstanding** | Pointer to **float64** | Refers to the number of shorted shares divided by the number of shares outstanding | [optional] 
**PercentHeldByInsiders** | Pointer to **float64** | Refers to the percentage of shares held by the company insiders | [optional] 
**PercentHeldByInstitutions** | Pointer to **float64** | Refers to the percentage of shares held by the institutions | [optional] 

## Methods

### NewGetStatistics200ResponseStatisticsStockStatistics

`func NewGetStatistics200ResponseStatisticsStockStatistics() *GetStatistics200ResponseStatisticsStockStatistics`

NewGetStatistics200ResponseStatisticsStockStatistics instantiates a new GetStatistics200ResponseStatisticsStockStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatistics200ResponseStatisticsStockStatisticsWithDefaults

`func NewGetStatistics200ResponseStatisticsStockStatisticsWithDefaults() *GetStatistics200ResponseStatisticsStockStatistics`

NewGetStatistics200ResponseStatisticsStockStatisticsWithDefaults instantiates a new GetStatistics200ResponseStatisticsStockStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharesOutstanding

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesOutstanding() float64`

GetSharesOutstanding returns the SharesOutstanding field if non-nil, zero value otherwise.

### GetSharesOutstandingOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesOutstandingOk() (*float64, bool)`

GetSharesOutstandingOk returns a tuple with the SharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstanding

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetSharesOutstanding(v float64)`

SetSharesOutstanding sets SharesOutstanding field to given value.

### HasSharesOutstanding

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasSharesOutstanding() bool`

HasSharesOutstanding returns a boolean if a field has been set.

### GetFloatShares

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetFloatShares() float64`

GetFloatShares returns the FloatShares field if non-nil, zero value otherwise.

### GetFloatSharesOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetFloatSharesOk() (*float64, bool)`

GetFloatSharesOk returns a tuple with the FloatShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatShares

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetFloatShares(v float64)`

SetFloatShares sets FloatShares field to given value.

### HasFloatShares

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasFloatShares() bool`

HasFloatShares returns a boolean if a field has been set.

### GetAvg10Volume

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg10Volume() int64`

GetAvg10Volume returns the Avg10Volume field if non-nil, zero value otherwise.

### GetAvg10VolumeOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg10VolumeOk() (*int64, bool)`

GetAvg10VolumeOk returns a tuple with the Avg10Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg10Volume

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetAvg10Volume(v int64)`

SetAvg10Volume sets Avg10Volume field to given value.

### HasAvg10Volume

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasAvg10Volume() bool`

HasAvg10Volume returns a boolean if a field has been set.

### GetAvg90Volume

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg90Volume() int64`

GetAvg90Volume returns the Avg90Volume field if non-nil, zero value otherwise.

### GetAvg90VolumeOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg90VolumeOk() (*int64, bool)`

GetAvg90VolumeOk returns a tuple with the Avg90Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg90Volume

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetAvg90Volume(v int64)`

SetAvg90Volume sets Avg90Volume field to given value.

### HasAvg90Volume

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasAvg90Volume() bool`

HasAvg90Volume returns a boolean if a field has been set.

### GetSharesShort

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesShort() int64`

GetSharesShort returns the SharesShort field if non-nil, zero value otherwise.

### GetSharesShortOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesShortOk() (*int64, bool)`

GetSharesShortOk returns a tuple with the SharesShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesShort

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetSharesShort(v int64)`

SetSharesShort sets SharesShort field to given value.

### HasSharesShort

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasSharesShort() bool`

HasSharesShort returns a boolean if a field has been set.

### GetShortRatio

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortRatio() float64`

GetShortRatio returns the ShortRatio field if non-nil, zero value otherwise.

### GetShortRatioOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortRatioOk() (*float64, bool)`

GetShortRatioOk returns a tuple with the ShortRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortRatio

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetShortRatio(v float64)`

SetShortRatio sets ShortRatio field to given value.

### HasShortRatio

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasShortRatio() bool`

HasShortRatio returns a boolean if a field has been set.

### GetShortPercentOfSharesOutstanding

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortPercentOfSharesOutstanding() float64`

GetShortPercentOfSharesOutstanding returns the ShortPercentOfSharesOutstanding field if non-nil, zero value otherwise.

### GetShortPercentOfSharesOutstandingOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortPercentOfSharesOutstandingOk() (*float64, bool)`

GetShortPercentOfSharesOutstandingOk returns a tuple with the ShortPercentOfSharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortPercentOfSharesOutstanding

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetShortPercentOfSharesOutstanding(v float64)`

SetShortPercentOfSharesOutstanding sets ShortPercentOfSharesOutstanding field to given value.

### HasShortPercentOfSharesOutstanding

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasShortPercentOfSharesOutstanding() bool`

HasShortPercentOfSharesOutstanding returns a boolean if a field has been set.

### GetPercentHeldByInsiders

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInsiders() float64`

GetPercentHeldByInsiders returns the PercentHeldByInsiders field if non-nil, zero value otherwise.

### GetPercentHeldByInsidersOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInsidersOk() (*float64, bool)`

GetPercentHeldByInsidersOk returns a tuple with the PercentHeldByInsiders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentHeldByInsiders

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetPercentHeldByInsiders(v float64)`

SetPercentHeldByInsiders sets PercentHeldByInsiders field to given value.

### HasPercentHeldByInsiders

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasPercentHeldByInsiders() bool`

HasPercentHeldByInsiders returns a boolean if a field has been set.

### GetPercentHeldByInstitutions

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInstitutions() float64`

GetPercentHeldByInstitutions returns the PercentHeldByInstitutions field if non-nil, zero value otherwise.

### GetPercentHeldByInstitutionsOk

`func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInstitutionsOk() (*float64, bool)`

GetPercentHeldByInstitutionsOk returns a tuple with the PercentHeldByInstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentHeldByInstitutions

`func (o *GetStatistics200ResponseStatisticsStockStatistics) SetPercentHeldByInstitutions(v float64)`

SetPercentHeldByInstitutions sets PercentHeldByInstitutions field to given value.

### HasPercentHeldByInstitutions

`func (o *GetStatistics200ResponseStatisticsStockStatistics) HasPercentHeldByInstitutions() bool`

HasPercentHeldByInstitutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


