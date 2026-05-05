// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesHtPhasor200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtPhasor200ResponseValuesInner{}

// GetTimeSeriesHtPhasor200ResponseValuesInner struct for GetTimeSeriesHtPhasor200ResponseValuesInner
type GetTimeSeriesHtPhasor200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// In_phase value
	InPhase string `json:"in_phase"`
	// Quadrature value
	Quadrature string `json:"quadrature"`
}

type _GetTimeSeriesHtPhasor200ResponseValuesInner GetTimeSeriesHtPhasor200ResponseValuesInner

// NewGetTimeSeriesHtPhasor200ResponseValuesInner instantiates a new GetTimeSeriesHtPhasor200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtPhasor200ResponseValuesInner(datetime string, inPhase string, quadrature string) *GetTimeSeriesHtPhasor200ResponseValuesInner {
	this := GetTimeSeriesHtPhasor200ResponseValuesInner{}
	this.Datetime = datetime
	this.InPhase = inPhase
	this.Quadrature = quadrature
	return &this
}

// NewGetTimeSeriesHtPhasor200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtPhasor200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtPhasor200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtPhasor200ResponseValuesInner {
	this := GetTimeSeriesHtPhasor200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetInPhase returns the InPhase field value
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetInPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InPhase
}

// GetInPhaseOk returns a tuple with the InPhase field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetInPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InPhase, true
}

// SetInPhase sets field value
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) SetInPhase(v string) {
	o.InPhase = v
}

// GetQuadrature returns the Quadrature field value
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetQuadrature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quadrature
}

// GetQuadratureOk returns a tuple with the Quadrature field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetQuadratureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quadrature, true
}

// SetQuadrature sets field value
func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) SetQuadrature(v string) {
	o.Quadrature = v
}

func (o GetTimeSeriesHtPhasor200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtPhasor200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["in_phase"] = o.InPhase
	toSerialize["quadrature"] = o.Quadrature
	return toSerialize, nil
}

func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"in_phase",
		"quadrature",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetTimeSeriesHtPhasor200ResponseValuesInner := _GetTimeSeriesHtPhasor200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesHtPhasor200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtPhasor200ResponseValuesInner(varGetTimeSeriesHtPhasor200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesHtPhasor200ResponseValuesInner struct {
	value *GetTimeSeriesHtPhasor200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesHtPhasor200ResponseValuesInner) Get() *GetTimeSeriesHtPhasor200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesHtPhasor200ResponseValuesInner) Set(val *GetTimeSeriesHtPhasor200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtPhasor200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtPhasor200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtPhasor200ResponseValuesInner(val *GetTimeSeriesHtPhasor200ResponseValuesInner) *NullableGetTimeSeriesHtPhasor200ResponseValuesInner {
	return &NullableGetTimeSeriesHtPhasor200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtPhasor200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtPhasor200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
