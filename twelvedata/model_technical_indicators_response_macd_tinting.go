// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the TechnicalIndicatorsResponseMacdTinting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechnicalIndicatorsResponseMacdTinting{}

// TechnicalIndicatorsResponseMacdTinting An array of tinting values used for proper indicator coloring
type TechnicalIndicatorsResponseMacdTinting struct {
	// How the tinting should be rendered
	Display *string `json:"display,omitempty"`
	// Hex color code for the tinting
	Color *string `json:"color,omitempty"`
	// Transparency level, float value from <code>0</code> to <code>1</code>
	Transparency *float64 `json:"transparency,omitempty"`
	// Lower bound of tinting, can be a number or a return parameter name
	LowerBound map[string]interface{} `json:"lower_bound,omitempty"`
	// Upper bound of tinting, can be a number or a return parameter name
	UpperBound map[string]interface{} `json:"upper_bound,omitempty"`
}

// NewTechnicalIndicatorsResponseMacdTinting instantiates a new TechnicalIndicatorsResponseMacdTinting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnicalIndicatorsResponseMacdTinting() *TechnicalIndicatorsResponseMacdTinting {
	this := TechnicalIndicatorsResponseMacdTinting{}
	return &this
}

// NewTechnicalIndicatorsResponseMacdTintingWithDefaults instantiates a new TechnicalIndicatorsResponseMacdTinting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnicalIndicatorsResponseMacdTintingWithDefaults() *TechnicalIndicatorsResponseMacdTinting {
	this := TechnicalIndicatorsResponseMacdTinting{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdTinting) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *TechnicalIndicatorsResponseMacdTinting) SetDisplay(v string) {
	o.Display = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdTinting) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TechnicalIndicatorsResponseMacdTinting) SetColor(v string) {
	o.Color = &v
}

// GetTransparency returns the Transparency field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdTinting) GetTransparency() float64 {
	if o == nil || IsNil(o.Transparency) {
		var ret float64
		return ret
	}
	return *o.Transparency
}

// GetTransparencyOk returns a tuple with the Transparency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) GetTransparencyOk() (*float64, bool) {
	if o == nil || IsNil(o.Transparency) {
		return nil, false
	}
	return o.Transparency, true
}

// HasTransparency returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) HasTransparency() bool {
	if o != nil && !IsNil(o.Transparency) {
		return true
	}

	return false
}

// SetTransparency gets a reference to the given float64 and assigns it to the Transparency field.
func (o *TechnicalIndicatorsResponseMacdTinting) SetTransparency(v float64) {
	o.Transparency = &v
}

// GetLowerBound returns the LowerBound field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdTinting) GetLowerBound() map[string]interface{} {
	if o == nil || IsNil(o.LowerBound) {
		var ret map[string]interface{}
		return ret
	}
	return o.LowerBound
}

// GetLowerBoundOk returns a tuple with the LowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) GetLowerBoundOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LowerBound) {
		return map[string]interface{}{}, false
	}
	return o.LowerBound, true
}

// HasLowerBound returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) HasLowerBound() bool {
	if o != nil && !IsNil(o.LowerBound) {
		return true
	}

	return false
}

// SetLowerBound gets a reference to the given map[string]interface{} and assigns it to the LowerBound field.
func (o *TechnicalIndicatorsResponseMacdTinting) SetLowerBound(v map[string]interface{}) {
	o.LowerBound = v
}

// GetUpperBound returns the UpperBound field value if set, zero value otherwise.
func (o *TechnicalIndicatorsResponseMacdTinting) GetUpperBound() map[string]interface{} {
	if o == nil || IsNil(o.UpperBound) {
		var ret map[string]interface{}
		return ret
	}
	return o.UpperBound
}

// GetUpperBoundOk returns a tuple with the UpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) GetUpperBoundOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UpperBound) {
		return map[string]interface{}{}, false
	}
	return o.UpperBound, true
}

// HasUpperBound returns a boolean if a field has been set.
func (o *TechnicalIndicatorsResponseMacdTinting) HasUpperBound() bool {
	if o != nil && !IsNil(o.UpperBound) {
		return true
	}

	return false
}

// SetUpperBound gets a reference to the given map[string]interface{} and assigns it to the UpperBound field.
func (o *TechnicalIndicatorsResponseMacdTinting) SetUpperBound(v map[string]interface{}) {
	o.UpperBound = v
}

func (o TechnicalIndicatorsResponseMacdTinting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechnicalIndicatorsResponseMacdTinting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Transparency) {
		toSerialize["transparency"] = o.Transparency
	}
	if !IsNil(o.LowerBound) {
		toSerialize["lower_bound"] = o.LowerBound
	}
	if !IsNil(o.UpperBound) {
		toSerialize["upper_bound"] = o.UpperBound
	}
	return toSerialize, nil
}

type NullableTechnicalIndicatorsResponseMacdTinting struct {
	value *TechnicalIndicatorsResponseMacdTinting
	isSet bool
}

func (v NullableTechnicalIndicatorsResponseMacdTinting) Get() *TechnicalIndicatorsResponseMacdTinting {
	return v.value
}

func (v *NullableTechnicalIndicatorsResponseMacdTinting) Set(val *TechnicalIndicatorsResponseMacdTinting) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnicalIndicatorsResponseMacdTinting) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnicalIndicatorsResponseMacdTinting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnicalIndicatorsResponseMacdTinting(val *TechnicalIndicatorsResponseMacdTinting) *NullableTechnicalIndicatorsResponseMacdTinting {
	return &NullableTechnicalIndicatorsResponseMacdTinting{value: val, isSet: true}
}

func (v NullableTechnicalIndicatorsResponseMacdTinting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechnicalIndicatorsResponseMacdTinting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
