/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the TechnicalIndicatorsResponseMacdParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechnicalIndicatorsResponseMacdParameter{}

// TechnicalIndicatorsResponseMacdParameter Input parameter name. Example values: <code>series_type</code>, <code>fast_period</code>, <code>slow_period</code>, <code>time_period</code>, <code>signal_period</code>
type TechnicalIndicatorsResponseMacdParameter struct {
	// Specifies parameter value set by default
	Default *int64 `json:"default,omitempty"`
	// If the parameter has upper bound in order to ensure correct calculation
	MaxRange *int64 `json:"max_range,omitempty"`
	// If the parameter has lower bound in order to ensure correct calculation
	MinRange *int64 `json:"min_range,omitempty"`
	// An array of available parameter values
	Range []string `json:"range,omitempty"`
	// Type of parameter might be <code>string</code>, <code>int</code>, <code>float</code> or <code>array</code>
	Type *string `json:"type,omitempty"`
}

// NewTechnicalIndicatorsResponseMacdParameter instantiates a new TechnicalIndicatorsResponseMacdParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnicalIndicatorsResponseMacdParameter() *TechnicalIndicatorsResponseMacdParameter {
	this := TechnicalIndicatorsResponseMacdParameter{}
	return &this
}

// NewTechnicalIndicatorsResponseMacdParameterWithDefaults instantiates a new TechnicalIndicatorsResponseMacdParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnicalIndicatorsResponseMacdParameterWithDefaults() *TechnicalIndicatorsResponseMacdParameter {
	this := TechnicalIndicatorsResponseMacdParameter{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdParameter) GetDefault() int64 {
	if o == nil || IsNil(o.Default) {
		var ret int64
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) GetDefaultOk() (*int64, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given int64 and assigns it to the Default field.
func (o *TechnicalIndicatorsResponseMacdParameter) SetDefault(v int64) {
	o.Default = &v
}

// GetMaxRange returns the MaxRange field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdParameter) GetMaxRange() int64 {
	if o == nil || IsNil(o.MaxRange) {
		var ret int64
		return ret
	}
	return *o.MaxRange
}

// GetMaxRangeOk returns a tuple with the MaxRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) GetMaxRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRange) {
		return nil, false
	}
	return o.MaxRange, true
}

// HasMaxRange returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) HasMaxRange() bool {
	if o != nil && !IsNil(o.MaxRange) {
		return true
	}

	return false
}

// SetMaxRange gets a reference to the given int64 and assigns it to the MaxRange field.
func (o *TechnicalIndicatorsResponseMacdParameter) SetMaxRange(v int64) {
	o.MaxRange = &v
}

// GetMinRange returns the MinRange field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdParameter) GetMinRange() int64 {
	if o == nil || IsNil(o.MinRange) {
		var ret int64
		return ret
	}
	return *o.MinRange
}

// GetMinRangeOk returns a tuple with the MinRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) GetMinRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinRange) {
		return nil, false
	}
	return o.MinRange, true
}

// HasMinRange returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) HasMinRange() bool {
	if o != nil && !IsNil(o.MinRange) {
		return true
	}

	return false
}

// SetMinRange gets a reference to the given int64 and assigns it to the MinRange field.
func (o *TechnicalIndicatorsResponseMacdParameter) SetMinRange(v int64) {
	o.MinRange = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdParameter) GetRange() []string {
	if o == nil || IsNil(o.Range) {
		var ret []string
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) GetRangeOk() ([]string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given []string and assigns it to the Range field.
func (o *TechnicalIndicatorsResponseMacdParameter) SetRange(v []string) {
	o.Range = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdParameter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdParameter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TechnicalIndicatorsResponseMacdParameter) SetType(v string) {
	o.Type = &v
}

func (o TechnicalIndicatorsResponseMacdParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechnicalIndicatorsResponseMacdParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.MaxRange) {
		toSerialize["max_range"] = o.MaxRange
	}
	if !IsNil(o.MinRange) {
		toSerialize["min_range"] = o.MinRange
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTechnicalIndicatorsResponseMacdParameter struct {
	value *TechnicalIndicatorsResponseMacdParameter
	isSet bool
}

func (v NullableTechnicalIndicatorsResponseMacdParameter) Get() *TechnicalIndicatorsResponseMacdParameter {
	return v.value
}

func (v *NullableTechnicalIndicatorsResponseMacdParameter) Set(val *TechnicalIndicatorsResponseMacdParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnicalIndicatorsResponseMacdParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnicalIndicatorsResponseMacdParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnicalIndicatorsResponseMacdParameter(val *TechnicalIndicatorsResponseMacdParameter) *NullableTechnicalIndicatorsResponseMacdParameter {
	return &NullableTechnicalIndicatorsResponseMacdParameter{value: val, isSet: true}
}

func (v NullableTechnicalIndicatorsResponseMacdParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechnicalIndicatorsResponseMacdParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
