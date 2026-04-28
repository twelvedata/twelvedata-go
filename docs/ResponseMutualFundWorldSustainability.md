# ResponseMutualFundWorldSustainability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int64** | Sustainability score: asset-weighted average of normalized company-level ESG Scores for the covered holdings in the portfolio from &#x60;0&#x60; to &#x60;100&#x60; | [optional] 
**CorporateEsgPillars** | Pointer to [**GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars**](GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars.md) |  | [optional] 
**SustainableInvestment** | Pointer to **bool** | Indication that the fund discloses in their prospectus that they employ socially responsible or ESG principles in their investment selection processes | [optional] 
**CorporateAum** | Pointer to **float64** | Percentage of AUM used to calculate sustainability score | [optional] 

## Methods

### NewResponseMutualFundWorldSustainability

`func NewResponseMutualFundWorldSustainability() *ResponseMutualFundWorldSustainability`

NewResponseMutualFundWorldSustainability instantiates a new ResponseMutualFundWorldSustainability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMutualFundWorldSustainabilityWithDefaults

`func NewResponseMutualFundWorldSustainabilityWithDefaults() *ResponseMutualFundWorldSustainability`

NewResponseMutualFundWorldSustainabilityWithDefaults instantiates a new ResponseMutualFundWorldSustainability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *ResponseMutualFundWorldSustainability) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ResponseMutualFundWorldSustainability) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ResponseMutualFundWorldSustainability) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *ResponseMutualFundWorldSustainability) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetCorporateEsgPillars

`func (o *ResponseMutualFundWorldSustainability) GetCorporateEsgPillars() GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars`

GetCorporateEsgPillars returns the CorporateEsgPillars field if non-nil, zero value otherwise.

### GetCorporateEsgPillarsOk

`func (o *ResponseMutualFundWorldSustainability) GetCorporateEsgPillarsOk() (*GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars, bool)`

GetCorporateEsgPillarsOk returns a tuple with the CorporateEsgPillars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateEsgPillars

`func (o *ResponseMutualFundWorldSustainability) SetCorporateEsgPillars(v GetMutualFundsWorld200ResponseMutualFundSustainabilityCorporateEsgPillars)`

SetCorporateEsgPillars sets CorporateEsgPillars field to given value.

### HasCorporateEsgPillars

`func (o *ResponseMutualFundWorldSustainability) HasCorporateEsgPillars() bool`

HasCorporateEsgPillars returns a boolean if a field has been set.

### GetSustainableInvestment

`func (o *ResponseMutualFundWorldSustainability) GetSustainableInvestment() bool`

GetSustainableInvestment returns the SustainableInvestment field if non-nil, zero value otherwise.

### GetSustainableInvestmentOk

`func (o *ResponseMutualFundWorldSustainability) GetSustainableInvestmentOk() (*bool, bool)`

GetSustainableInvestmentOk returns a tuple with the SustainableInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainableInvestment

`func (o *ResponseMutualFundWorldSustainability) SetSustainableInvestment(v bool)`

SetSustainableInvestment sets SustainableInvestment field to given value.

### HasSustainableInvestment

`func (o *ResponseMutualFundWorldSustainability) HasSustainableInvestment() bool`

HasSustainableInvestment returns a boolean if a field has been set.

### GetCorporateAum

`func (o *ResponseMutualFundWorldSustainability) GetCorporateAum() float64`

GetCorporateAum returns the CorporateAum field if non-nil, zero value otherwise.

### GetCorporateAumOk

`func (o *ResponseMutualFundWorldSustainability) GetCorporateAumOk() (*float64, bool)`

GetCorporateAumOk returns a tuple with the CorporateAum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateAum

`func (o *ResponseMutualFundWorldSustainability) SetCorporateAum(v float64)`

SetCorporateAum sets CorporateAum field to given value.

### HasCorporateAum

`func (o *ResponseMutualFundWorldSustainability) HasCorporateAum() bool`

HasCorporateAum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


