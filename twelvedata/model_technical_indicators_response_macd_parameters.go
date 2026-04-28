/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the TechnicalIndicatorsResponseMacdParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechnicalIndicatorsResponseMacdParameters{}

// TechnicalIndicatorsResponseMacdParameters An array of input parameters for the indicator
type TechnicalIndicatorsResponseMacdParameters struct {
	ParameterName *TechnicalIndicatorsResponseMacdParameter `json:"parameter_name,omitempty"`
}

// NewTechnicalIndicatorsResponseMacdParameters instantiates a new TechnicalIndicatorsResponseMacdParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnicalIndicatorsResponseMacdParameters() *TechnicalIndicatorsResponseMacdParameters {
	this := TechnicalIndicatorsResponseMacdParameters{}
	return &this
}

// NewTechnicalIndicatorsResponseMacdParametersWithDefaults instantiates a new TechnicalIndicatorsResponseMacdParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnicalIndicatorsResponseMacdParametersWithDefaults() *TechnicalIndicatorsResponseMacdParameters {
	this := TechnicalIndicatorsResponseMacdParameters{}
	return &this
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdParameters) GetParameterName() TechnicalIndicatorsResponseMacdParameter {
	if o == nil || IsNil(o.ParameterName) {
		var ret TechnicalIndicatorsResponseMacdParameter
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdParameters) GetParameterNameOk() (*TechnicalIndicatorsResponseMacdParameter, bool) {
	if o == nil || IsNil(o.ParameterName) {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdParameters) HasParameterName() bool {
	if o != nil && !IsNil(o.ParameterName) {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given TechnicalIndicatorsResponseMacdParameter and assigns it to the ParameterName field.
func (o *TechnicalIndicatorsResponseMacdParameters) SetParameterName(v TechnicalIndicatorsResponseMacdParameter) {
	o.ParameterName = &v
}

func (o TechnicalIndicatorsResponseMacdParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechnicalIndicatorsResponseMacdParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParameterName) {
		toSerialize["parameter_name"] = o.ParameterName
	}
	return toSerialize, nil
}

type NullableTechnicalIndicatorsResponseMacdParameters struct {
	value *TechnicalIndicatorsResponseMacdParameters
	isSet bool
}

func (v NullableTechnicalIndicatorsResponseMacdParameters) Get() *TechnicalIndicatorsResponseMacdParameters {
	return v.value
}

func (v *NullableTechnicalIndicatorsResponseMacdParameters) Set(val *TechnicalIndicatorsResponseMacdParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnicalIndicatorsResponseMacdParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnicalIndicatorsResponseMacdParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnicalIndicatorsResponseMacdParameters(val *TechnicalIndicatorsResponseMacdParameters) *NullableTechnicalIndicatorsResponseMacdParameters {
	return &NullableTechnicalIndicatorsResponseMacdParameters{value: val, isSet: true}
}

func (v NullableTechnicalIndicatorsResponseMacdParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechnicalIndicatorsResponseMacdParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
