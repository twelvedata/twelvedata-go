/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMama200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMama200ResponseMetaIndicator{}

// GetTimeSeriesMama200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMama200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The limit for the fast moving average
	FastLimit float64 `json:"fast_limit"`
	// The limit for the slow moving average
	SlowLimit float64 `json:"slow_limit"`
}

type _GetTimeSeriesMama200ResponseMetaIndicator GetTimeSeriesMama200ResponseMetaIndicator

// NewGetTimeSeriesMama200ResponseMetaIndicator instantiates a new GetTimeSeriesMama200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMama200ResponseMetaIndicator(name string, seriesType string, fastLimit float64, slowLimit float64) *GetTimeSeriesMama200ResponseMetaIndicator {
	this := GetTimeSeriesMama200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.FastLimit = fastLimit
	this.SlowLimit = slowLimit
	return &this
}

// NewGetTimeSeriesMama200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMama200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMama200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMama200ResponseMetaIndicator {
	this := GetTimeSeriesMama200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetFastLimit returns the FastLimit field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetFastLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FastLimit
}

// GetFastLimitOk returns a tuple with the FastLimit field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetFastLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastLimit, true
}

// SetFastLimit sets field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) SetFastLimit(v float64) {
	o.FastLimit = v
}

// GetSlowLimit returns the SlowLimit field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetSlowLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SlowLimit
}

// GetSlowLimitOk returns a tuple with the SlowLimit field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMama200ResponseMetaIndicator) GetSlowLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowLimit, true
}

// SetSlowLimit sets field value
func (o *GetTimeSeriesMama200ResponseMetaIndicator) SetSlowLimit(v float64) {
	o.SlowLimit = v
}

func (o GetTimeSeriesMama200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMama200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["fast_limit"] = o.FastLimit
	toSerialize["slow_limit"] = o.SlowLimit
	return toSerialize, nil
}

func (o *GetTimeSeriesMama200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"fast_limit",
		"slow_limit",
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

	varGetTimeSeriesMama200ResponseMetaIndicator := _GetTimeSeriesMama200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMama200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMama200ResponseMetaIndicator(varGetTimeSeriesMama200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMama200ResponseMetaIndicator struct {
	value *GetTimeSeriesMama200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMama200ResponseMetaIndicator) Get() *GetTimeSeriesMama200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMama200ResponseMetaIndicator) Set(val *GetTimeSeriesMama200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMama200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMama200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMama200ResponseMetaIndicator(val *GetTimeSeriesMama200ResponseMetaIndicator) *NullableGetTimeSeriesMama200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMama200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMama200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMama200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
