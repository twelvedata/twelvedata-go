# ResponseMutualFundWorldRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolatilityMeasures** | Pointer to [**[]GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner**](GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner.md) | Volatility statistics of the fund | [optional] 
**ValuationMetrics** | Pointer to [**GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics**](GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics.md) |  | [optional] 

## Methods

### NewResponseMutualFundWorldRisk

`func NewResponseMutualFundWorldRisk() *ResponseMutualFundWorldRisk`

NewResponseMutualFundWorldRisk instantiates a new ResponseMutualFundWorldRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMutualFundWorldRiskWithDefaults

`func NewResponseMutualFundWorldRiskWithDefaults() *ResponseMutualFundWorldRisk`

NewResponseMutualFundWorldRiskWithDefaults instantiates a new ResponseMutualFundWorldRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolatilityMeasures

`func (o *ResponseMutualFundWorldRisk) GetVolatilityMeasures() []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner`

GetVolatilityMeasures returns the VolatilityMeasures field if non-nil, zero value otherwise.

### GetVolatilityMeasuresOk

`func (o *ResponseMutualFundWorldRisk) GetVolatilityMeasuresOk() (*[]GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner, bool)`

GetVolatilityMeasuresOk returns a tuple with the VolatilityMeasures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatilityMeasures

`func (o *ResponseMutualFundWorldRisk) SetVolatilityMeasures(v []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner)`

SetVolatilityMeasures sets VolatilityMeasures field to given value.

### HasVolatilityMeasures

`func (o *ResponseMutualFundWorldRisk) HasVolatilityMeasures() bool`

HasVolatilityMeasures returns a boolean if a field has been set.

### GetValuationMetrics

`func (o *ResponseMutualFundWorldRisk) GetValuationMetrics() GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics`

GetValuationMetrics returns the ValuationMetrics field if non-nil, zero value otherwise.

### GetValuationMetricsOk

`func (o *ResponseMutualFundWorldRisk) GetValuationMetricsOk() (*GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics, bool)`

GetValuationMetricsOk returns a tuple with the ValuationMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuationMetrics

`func (o *ResponseMutualFundWorldRisk) SetValuationMetrics(v GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics)`

SetValuationMetrics sets ValuationMetrics field to given value.

### HasValuationMetrics

`func (o *ResponseMutualFundWorldRisk) HasValuationMetrics() bool`

HasValuationMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


