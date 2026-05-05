// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the TechnicalIndicatorsResponseMacdOutputValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechnicalIndicatorsResponseMacdOutputValues{}

// TechnicalIndicatorsResponseMacdOutputValues An array of output values
type TechnicalIndicatorsResponseMacdOutputValues struct {
	ParameterName *TechnicalIndicatorsResponseMacdOutputValue `json:"parameter_name,omitempty"`
}

// NewTechnicalIndicatorsResponseMacdOutputValues instantiates a new TechnicalIndicatorsResponseMacdOutputValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnicalIndicatorsResponseMacdOutputValues() *TechnicalIndicatorsResponseMacdOutputValues {
	this := TechnicalIndicatorsResponseMacdOutputValues{}
	return &this
}

// NewTechnicalIndicatorsResponseMacdOutputValuesWithDefaults instantiates a new TechnicalIndicatorsResponseMacdOutputValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnicalIndicatorsResponseMacdOutputValuesWithDefaults() *TechnicalIndicatorsResponseMacdOutputValues {
	this := TechnicalIndicatorsResponseMacdOutputValues{}
	return &this
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdOutputValues) GetParameterName() TechnicalIndicatorsResponseMacdOutputValue {
	if o == nil || IsNil(o.ParameterName) {
		var ret TechnicalIndicatorsResponseMacdOutputValue
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValues) GetParameterNameOk() (*TechnicalIndicatorsResponseMacdOutputValue, bool) {
	if o == nil || IsNil(o.ParameterName) {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValues) HasParameterName() bool {
	if o != nil && !IsNil(o.ParameterName) {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given TechnicalIndicatorsResponseMacdOutputValue and assigns it to the ParameterName field.
func (o *TechnicalIndicatorsResponseMacdOutputValues) SetParameterName(v TechnicalIndicatorsResponseMacdOutputValue) {
	o.ParameterName = &v
}

func (o TechnicalIndicatorsResponseMacdOutputValues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechnicalIndicatorsResponseMacdOutputValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParameterName) {
		toSerialize["parameter_name"] = o.ParameterName
	}
	return toSerialize, nil
}

type NullableTechnicalIndicatorsResponseMacdOutputValues struct {
	value *TechnicalIndicatorsResponseMacdOutputValues
	isSet bool
}

func (v NullableTechnicalIndicatorsResponseMacdOutputValues) Get() *TechnicalIndicatorsResponseMacdOutputValues {
	return v.value
}

func (v *NullableTechnicalIndicatorsResponseMacdOutputValues) Set(val *TechnicalIndicatorsResponseMacdOutputValues) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnicalIndicatorsResponseMacdOutputValues) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnicalIndicatorsResponseMacdOutputValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnicalIndicatorsResponseMacdOutputValues(val *TechnicalIndicatorsResponseMacdOutputValues) *NullableTechnicalIndicatorsResponseMacdOutputValues {
	return &NullableTechnicalIndicatorsResponseMacdOutputValues{value: val, isSet: true}
}

func (v NullableTechnicalIndicatorsResponseMacdOutputValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechnicalIndicatorsResponseMacdOutputValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
