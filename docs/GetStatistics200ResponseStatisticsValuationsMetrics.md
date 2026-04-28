# GetStatistics200ResponseStatisticsValuationsMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketCapitalization** | Pointer to **float64** | Refers to the market value of the company&#39;s outstanding shares | [optional] 
**EnterpriseValue** | Pointer to **float64** | Refers to enterprise value (EV) of the company, often used as a more comprehensive alternative to market capitalization | [optional] 
**TrailingPe** | Pointer to **float64** | Refers to the trailing price-to-earnings (P/E). It is calculated by taking the current stock price and dividing it by the trailing earnings per share (EPS) for the past 12 months | [optional] 
**ForwardPe** | Pointer to **float64** | Refers to the forward price-to-earnings ratio. It is calculated by dividing the current share price by the estimated future earnings per share | [optional] 
**PegRatio** | Pointer to **float64** | The price/earnings to growth (PEG) ratio is a price-to-earnings ratio divided by the growth rate of the earnings | [optional] 
**PriceToSalesTtm** | Pointer to **float64** | The price-to-sales (P/S) ratio is a valuation ratio that compares the market capitalization to its revenues over the last 12 months | [optional] 
**PriceToBookMrq** | Pointer to **float64** | The price-to-book (P/B) ratio is equal to the current share price divided by the book value of all shares (BVPS) over the last quarter | [optional] 
**EnterpriseToRevenue** | Pointer to **float64** | The enterprise value-to-revenue multiple (EV/R) is a measure that compares enterprise value to revenue | [optional] 
**EnterpriseToEbitda** | Pointer to **float64** | The enterprise value-to-ebitda multiple (EV/EBITDA) is a measure that compares enterprise value to EBITDA | [optional] 

## Methods

### NewGetStatistics200ResponseStatisticsValuationsMetrics

`func NewGetStatistics200ResponseStatisticsValuationsMetrics() *GetStatistics200ResponseStatisticsValuationsMetrics`

NewGetStatistics200ResponseStatisticsValuationsMetrics instantiates a new GetStatistics200ResponseStatisticsValuationsMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatistics200ResponseStatisticsValuationsMetricsWithDefaults

`func NewGetStatistics200ResponseStatisticsValuationsMetricsWithDefaults() *GetStatistics200ResponseStatisticsValuationsMetrics`

NewGetStatistics200ResponseStatisticsValuationsMetricsWithDefaults instantiates a new GetStatistics200ResponseStatisticsValuationsMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketCapitalization

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetMarketCapitalization() float64`

GetMarketCapitalization returns the MarketCapitalization field if non-nil, zero value otherwise.

### GetMarketCapitalizationOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetMarketCapitalizationOk() (*float64, bool)`

GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapitalization

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetMarketCapitalization(v float64)`

SetMarketCapitalization sets MarketCapitalization field to given value.

### HasMarketCapitalization

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasMarketCapitalization() bool`

HasMarketCapitalization returns a boolean if a field has been set.

### GetEnterpriseValue

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseValue() float64`

GetEnterpriseValue returns the EnterpriseValue field if non-nil, zero value otherwise.

### GetEnterpriseValueOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseValueOk() (*float64, bool)`

GetEnterpriseValueOk returns a tuple with the EnterpriseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseValue

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseValue(v float64)`

SetEnterpriseValue sets EnterpriseValue field to given value.

### HasEnterpriseValue

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseValue() bool`

HasEnterpriseValue returns a boolean if a field has been set.

### GetTrailingPe

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetTrailingPe() float64`

GetTrailingPe returns the TrailingPe field if non-nil, zero value otherwise.

### GetTrailingPeOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetTrailingPeOk() (*float64, bool)`

GetTrailingPeOk returns a tuple with the TrailingPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingPe

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetTrailingPe(v float64)`

SetTrailingPe sets TrailingPe field to given value.

### HasTrailingPe

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasTrailingPe() bool`

HasTrailingPe returns a boolean if a field has been set.

### GetForwardPe

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetForwardPe() float64`

GetForwardPe returns the ForwardPe field if non-nil, zero value otherwise.

### GetForwardPeOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetForwardPeOk() (*float64, bool)`

GetForwardPeOk returns a tuple with the ForwardPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPe

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetForwardPe(v float64)`

SetForwardPe sets ForwardPe field to given value.

### HasForwardPe

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasForwardPe() bool`

HasForwardPe returns a boolean if a field has been set.

### GetPegRatio

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPegRatio() float64`

GetPegRatio returns the PegRatio field if non-nil, zero value otherwise.

### GetPegRatioOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPegRatioOk() (*float64, bool)`

GetPegRatioOk returns a tuple with the PegRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegRatio

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPegRatio(v float64)`

SetPegRatio sets PegRatio field to given value.

### HasPegRatio

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPegRatio() bool`

HasPegRatio returns a boolean if a field has been set.

### GetPriceToSalesTtm

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToSalesTtm() float64`

GetPriceToSalesTtm returns the PriceToSalesTtm field if non-nil, zero value otherwise.

### GetPriceToSalesTtmOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToSalesTtmOk() (*float64, bool)`

GetPriceToSalesTtmOk returns a tuple with the PriceToSalesTtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToSalesTtm

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPriceToSalesTtm(v float64)`

SetPriceToSalesTtm sets PriceToSalesTtm field to given value.

### HasPriceToSalesTtm

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPriceToSalesTtm() bool`

HasPriceToSalesTtm returns a boolean if a field has been set.

### GetPriceToBookMrq

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToBookMrq() float64`

GetPriceToBookMrq returns the PriceToBookMrq field if non-nil, zero value otherwise.

### GetPriceToBookMrqOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToBookMrqOk() (*float64, bool)`

GetPriceToBookMrqOk returns a tuple with the PriceToBookMrq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToBookMrq

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPriceToBookMrq(v float64)`

SetPriceToBookMrq sets PriceToBookMrq field to given value.

### HasPriceToBookMrq

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPriceToBookMrq() bool`

HasPriceToBookMrq returns a boolean if a field has been set.

### GetEnterpriseToRevenue

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToRevenue() float64`

GetEnterpriseToRevenue returns the EnterpriseToRevenue field if non-nil, zero value otherwise.

### GetEnterpriseToRevenueOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToRevenueOk() (*float64, bool)`

GetEnterpriseToRevenueOk returns a tuple with the EnterpriseToRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseToRevenue

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseToRevenue(v float64)`

SetEnterpriseToRevenue sets EnterpriseToRevenue field to given value.

### HasEnterpriseToRevenue

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseToRevenue() bool`

HasEnterpriseToRevenue returns a boolean if a field has been set.

### GetEnterpriseToEbitda

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToEbitda() float64`

GetEnterpriseToEbitda returns the EnterpriseToEbitda field if non-nil, zero value otherwise.

### GetEnterpriseToEbitdaOk

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToEbitdaOk() (*float64, bool)`

GetEnterpriseToEbitdaOk returns a tuple with the EnterpriseToEbitda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseToEbitda

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseToEbitda(v float64)`

SetEnterpriseToEbitda sets EnterpriseToEbitda field to given value.

### HasEnterpriseToEbitda

`func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseToEbitda() bool`

HasEnterpriseToEbitda returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


