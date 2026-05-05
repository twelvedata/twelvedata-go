// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the TechnicalIndicatorsResponseMacdOutputValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechnicalIndicatorsResponseMacdOutputValue{}

// TechnicalIndicatorsResponseMacdOutputValue Output parameter name. Example values: <code>ad</code>, <code>add</code>, <code>adxr</code>, <code>aroonosc</code>,  <code>macd</code>, <code>macd_signal</code>, <code>macd_hist</code>, etc
type TechnicalIndicatorsResponseMacdOutputValue struct {
	// Suggested color for displaying returns hex color code
	DefaultColor *string `json:"default_color,omitempty"`
	// How output value should be rendered
	Display *string `json:"display,omitempty"`
	// If output value has minimum bound
	MinRange *int64 `json:"min_range,omitempty"`
	// If output value has maximum bound
	MaxRange *int64 `json:"max_range,omitempty"`
}

// NewTechnicalIndicatorsResponseMacdOutputValue instantiates a new TechnicalIndicatorsResponseMacdOutputValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnicalIndicatorsResponseMacdOutputValue() *TechnicalIndicatorsResponseMacdOutputValue {
	this := TechnicalIndicatorsResponseMacdOutputValue{}
	return &this
}

// NewTechnicalIndicatorsResponseMacdOutputValueWithDefaults instantiates a new TechnicalIndicatorsResponseMacdOutputValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnicalIndicatorsResponseMacdOutputValueWithDefaults() *TechnicalIndicatorsResponseMacdOutputValue {
	this := TechnicalIndicatorsResponseMacdOutputValue{}
	return &this
}

// GetDefaultColor returns the DefaultColor field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetDefaultColor() string {
	if o == nil || IsNil(o.DefaultColor) {
		var ret string
		return ret
	}
	return *o.DefaultColor
}

// GetDefaultColorOk returns a tuple with the DefaultColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetDefaultColorOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultColor) {
		return nil, false
	}
	return o.DefaultColor, true
}

// HasDefaultColor returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) HasDefaultColor() bool {
	if o != nil && !IsNil(o.DefaultColor) {
		return true
	}

	return false
}

// SetDefaultColor gets a reference to the given string and assigns it to the DefaultColor field.
func (o *TechnicalIndicatorsResponseMacdOutputValue) SetDefaultColor(v string) {
	o.DefaultColor = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *TechnicalIndicatorsResponseMacdOutputValue) SetDisplay(v string) {
	o.Display = &v
}

// GetMinRange returns the MinRange field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetMinRange() int64 {
	if o == nil || IsNil(o.MinRange) {
		var ret int64
		return ret
	}
	return *o.MinRange
}

// GetMinRangeOk returns a tuple with the MinRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetMinRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinRange) {
		return nil, false
	}
	return o.MinRange, true
}

// HasMinRange returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) HasMinRange() bool {
	if o != nil && !IsNil(o.MinRange) {
		return true
	}

	return false
}

// SetMinRange gets a reference to the given int64 and assigns it to the MinRange field.
func (o *TechnicalIndicatorsResponseMacdOutputValue) SetMinRange(v int64) {
	o.MinRange = &v
}

// GetMaxRange returns the MaxRange field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetMaxRange() int64 {
	if o == nil || IsNil(o.MaxRange) {
		var ret int64
		return ret
	}
	return *o.MaxRange
}

// GetMaxRangeOk returns a tuple with the MaxRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) GetMaxRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRange) {
		return nil, false
	}
	return o.MaxRange, true
}

// HasMaxRange returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdOutputValue) HasMaxRange() bool {
	if o != nil && !IsNil(o.MaxRange) {
		return true
	}

	return false
}

// SetMaxRange gets a reference to the given int64 and assigns it to the MaxRange field.
func (o *TechnicalIndicatorsResponseMacdOutputValue) SetMaxRange(v int64) {
	o.MaxRange = &v
}

func (o TechnicalIndicatorsResponseMacdOutputValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechnicalIndicatorsResponseMacdOutputValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultColor) {
		toSerialize["default_color"] = o.DefaultColor
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.MinRange) {
		toSerialize["min_range"] = o.MinRange
	}
	if !IsNil(o.MaxRange) {
		toSerialize["max_range"] = o.MaxRange
	}
	return toSerialize, nil
}

type NullableTechnicalIndicatorsResponseMacdOutputValue struct {
	value *TechnicalIndicatorsResponseMacdOutputValue
	isSet bool
}

func (v NullableTechnicalIndicatorsResponseMacdOutputValue) Get() *TechnicalIndicatorsResponseMacdOutputValue {
	return v.value
}

func (v *NullableTechnicalIndicatorsResponseMacdOutputValue) Set(val *TechnicalIndicatorsResponseMacdOutputValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnicalIndicatorsResponseMacdOutputValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnicalIndicatorsResponseMacdOutputValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnicalIndicatorsResponseMacdOutputValue(val *TechnicalIndicatorsResponseMacdOutputValue) *NullableTechnicalIndicatorsResponseMacdOutputValue {
	return &NullableTechnicalIndicatorsResponseMacdOutputValue{value: val, isSet: true}
}

func (v NullableTechnicalIndicatorsResponseMacdOutputValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechnicalIndicatorsResponseMacdOutputValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
