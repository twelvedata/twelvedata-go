# GetTechnicalIndicators200ResponseDataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | If the indicator is tested, approved and is recommended for use returns &lt;code&gt;true&lt;/code&gt;, otherwise returns &lt;code&gt;false&lt;/code&gt; | 
**FullName** | **string** | Full indicator name | 
**Description** | **string** | Brief description of the indicator | 
**Type** | **string** | Group to which indicator belongs to | 
**Overlay** | **bool** | If indicator should be plotted over price bars returns &lt;code&gt;true&lt;/code&gt;, otherwise returns &lt;code&gt;false&lt;/code&gt; | 
**OutputValues** | Pointer to [**TechnicalIndicatorsResponseMacdOutputValues**](TechnicalIndicatorsResponseMacdOutputValues.md) |  | [optional] 
**Parameters** | Pointer to [**TechnicalIndicatorsResponseMacdParameters**](TechnicalIndicatorsResponseMacdParameters.md) |  | [optional] 
**Tinting** | Pointer to [**TechnicalIndicatorsResponseMacdTinting**](TechnicalIndicatorsResponseMacdTinting.md) |  | [optional] 

## Methods

### NewGetTechnicalIndicators200ResponseDataValue

`func NewGetTechnicalIndicators200ResponseDataValue(enable bool, fullName string, description string, type_ string, overlay bool, ) *GetTechnicalIndicators200ResponseDataValue`

NewGetTechnicalIndicators200ResponseDataValue instantiates a new GetTechnicalIndicators200ResponseDataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTechnicalIndicators200ResponseDataValueWithDefaults

`func NewGetTechnicalIndicators200ResponseDataValueWithDefaults() *GetTechnicalIndicators200ResponseDataValue`

NewGetTechnicalIndicators200ResponseDataValueWithDefaults instantiates a new GetTechnicalIndicators200ResponseDataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GetTechnicalIndicators200ResponseDataValue) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GetTechnicalIndicators200ResponseDataValue) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetFullName

`func (o *GetTechnicalIndicators200ResponseDataValue) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetTechnicalIndicators200ResponseDataValue) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetDescription

`func (o *GetTechnicalIndicators200ResponseDataValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTechnicalIndicators200ResponseDataValue) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *GetTechnicalIndicators200ResponseDataValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTechnicalIndicators200ResponseDataValue) SetType(v string)`

SetType sets Type field to given value.


### GetOverlay

`func (o *GetTechnicalIndicators200ResponseDataValue) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *GetTechnicalIndicators200ResponseDataValue) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.


### GetOutputValues

`func (o *GetTechnicalIndicators200ResponseDataValue) GetOutputValues() TechnicalIndicatorsResponseMacdOutputValues`

GetOutputValues returns the OutputValues field if non-nil, zero value otherwise.

### GetOutputValuesOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetOutputValuesOk() (*TechnicalIndicatorsResponseMacdOutputValues, bool)`

GetOutputValuesOk returns a tuple with the OutputValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputValues

`func (o *GetTechnicalIndicators200ResponseDataValue) SetOutputValues(v TechnicalIndicatorsResponseMacdOutputValues)`

SetOutputValues sets OutputValues field to given value.

### HasOutputValues

`func (o *GetTechnicalIndicators200ResponseDataValue) HasOutputValues() bool`

HasOutputValues returns a boolean if a field has been set.

### GetParameters

`func (o *GetTechnicalIndicators200ResponseDataValue) GetParameters() TechnicalIndicatorsResponseMacdParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetParametersOk() (*TechnicalIndicatorsResponseMacdParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetTechnicalIndicators200ResponseDataValue) SetParameters(v TechnicalIndicatorsResponseMacdParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GetTechnicalIndicators200ResponseDataValue) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTinting

`func (o *GetTechnicalIndicators200ResponseDataValue) GetTinting() TechnicalIndicatorsResponseMacdTinting`

GetTinting returns the Tinting field if non-nil, zero value otherwise.

### GetTintingOk

`func (o *GetTechnicalIndicators200ResponseDataValue) GetTintingOk() (*TechnicalIndicatorsResponseMacdTinting, bool)`

GetTintingOk returns a tuple with the Tinting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinting

`func (o *GetTechnicalIndicators200ResponseDataValue) SetTinting(v TechnicalIndicatorsResponseMacdTinting)`

SetTinting sets Tinting field to given value.

### HasTinting

`func (o *GetTechnicalIndicators200ResponseDataValue) HasTinting() bool`

HasTinting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


