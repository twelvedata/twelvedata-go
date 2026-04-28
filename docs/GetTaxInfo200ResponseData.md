# GetTaxInfo200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxIndicator** | **string** | The instrument tax indicator, can be &#x60;null&#x60;, &#x60;us_1446f&#x60;, &#x60;spanish_ftt&#x60;, &#x60;uk_stamp_duty&#x60;, &#x60;hk_stamp_duty&#x60;, &#x60;french_ftt&#x60; or &#x60;italian_ftt&#x60; | 

## Methods

### NewGetTaxInfo200ResponseData

`func NewGetTaxInfo200ResponseData(taxIndicator string, ) *GetTaxInfo200ResponseData`

NewGetTaxInfo200ResponseData instantiates a new GetTaxInfo200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxInfo200ResponseDataWithDefaults

`func NewGetTaxInfo200ResponseDataWithDefaults() *GetTaxInfo200ResponseData`

NewGetTaxInfo200ResponseDataWithDefaults instantiates a new GetTaxInfo200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxIndicator

`func (o *GetTaxInfo200ResponseData) GetTaxIndicator() string`

GetTaxIndicator returns the TaxIndicator field if non-nil, zero value otherwise.

### GetTaxIndicatorOk

`func (o *GetTaxInfo200ResponseData) GetTaxIndicatorOk() (*string, bool)`

GetTaxIndicatorOk returns a tuple with the TaxIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIndicator

`func (o *GetTaxInfo200ResponseData) SetTaxIndicator(v string)`

SetTaxIndicator sets TaxIndicator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


