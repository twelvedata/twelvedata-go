# GetETFsWorld200ResponseEtfRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolatilityMeasures** | Pointer to [**[]GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner**](GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner.md) | Risk and volatility statistics of the fund and its category over different periods | [optional] 
**ValuationMetrics** | Pointer to [**GetETFsWorld200ResponseEtfRiskValuationMetrics**](GetETFsWorld200ResponseEtfRiskValuationMetrics.md) |  | [optional] 

## Methods

### NewGetETFsWorld200ResponseEtfRisk

`func NewGetETFsWorld200ResponseEtfRisk() *GetETFsWorld200ResponseEtfRisk`

NewGetETFsWorld200ResponseEtfRisk instantiates a new GetETFsWorld200ResponseEtfRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsWorld200ResponseEtfRiskWithDefaults

`func NewGetETFsWorld200ResponseEtfRiskWithDefaults() *GetETFsWorld200ResponseEtfRisk`

NewGetETFsWorld200ResponseEtfRiskWithDefaults instantiates a new GetETFsWorld200ResponseEtfRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolatilityMeasures

`func (o *GetETFsWorld200ResponseEtfRisk) GetVolatilityMeasures() []GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner`

GetVolatilityMeasures returns the VolatilityMeasures field if non-nil, zero value otherwise.

### GetVolatilityMeasuresOk

`func (o *GetETFsWorld200ResponseEtfRisk) GetVolatilityMeasuresOk() (*[]GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner, bool)`

GetVolatilityMeasuresOk returns a tuple with the VolatilityMeasures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatilityMeasures

`func (o *GetETFsWorld200ResponseEtfRisk) SetVolatilityMeasures(v []GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner)`

SetVolatilityMeasures sets VolatilityMeasures field to given value.

### HasVolatilityMeasures

`func (o *GetETFsWorld200ResponseEtfRisk) HasVolatilityMeasures() bool`

HasVolatilityMeasures returns a boolean if a field has been set.

### GetValuationMetrics

`func (o *GetETFsWorld200ResponseEtfRisk) GetValuationMetrics() GetETFsWorld200ResponseEtfRiskValuationMetrics`

GetValuationMetrics returns the ValuationMetrics field if non-nil, zero value otherwise.

### GetValuationMetricsOk

`func (o *GetETFsWorld200ResponseEtfRisk) GetValuationMetricsOk() (*GetETFsWorld200ResponseEtfRiskValuationMetrics, bool)`

GetValuationMetricsOk returns a tuple with the ValuationMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuationMetrics

`func (o *GetETFsWorld200ResponseEtfRisk) SetValuationMetrics(v GetETFsWorld200ResponseEtfRiskValuationMetrics)`

SetValuationMetrics sets ValuationMetrics field to given value.

### HasValuationMetrics

`func (o *GetETFsWorld200ResponseEtfRisk) HasValuationMetrics() bool`

HasValuationMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


