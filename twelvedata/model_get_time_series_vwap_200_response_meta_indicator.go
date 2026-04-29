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

// checks if the GetTimeSeriesVwap200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesVwap200ResponseMetaIndicator{}

// GetTimeSeriesVwap200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesVwap200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Standard deviation time period
	SdTimePeriod *int64 `json:"sd_time_period,omitempty"`
	// Standard deviation value
	Sd *float64 `json:"sd,omitempty"`
}

type _GetTimeSeriesVwap200ResponseMetaIndicator GetTimeSeriesVwap200ResponseMetaIndicator

// NewGetTimeSeriesVwap200ResponseMetaIndicator instantiates a new GetTimeSeriesVwap200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesVwap200ResponseMetaIndicator(name string) *GetTimeSeriesVwap200ResponseMetaIndicator {
	this := GetTimeSeriesVwap200ResponseMetaIndicator{}
	this.Name = name
	return &this
}

// NewGetTimeSeriesVwap200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesVwap200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesVwap200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesVwap200ResponseMetaIndicator {
	this := GetTimeSeriesVwap200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSdTimePeriod returns the SdTimePeriod field value if set, zero value otherwise.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSdTimePeriod() int64 {
	if o == nil || IsNil(o.SdTimePeriod) {
		var ret int64
		return ret
	}
	return *o.SdTimePeriod
}

// GetSdTimePeriodOk returns a tuple with the SdTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSdTimePeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.SdTimePeriod) {
		return nil, false
	}
	return o.SdTimePeriod, true
}

// HasSdTimePeriod returns a boolean if a field has been set.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) HasSdTimePeriod() bool {
	if o != nil && !IsNil(o.SdTimePeriod) {
		return true
	}

	return false
}

// SetSdTimePeriod gets a reference to the given int64 and assigns it to the SdTimePeriod field.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) SetSdTimePeriod(v int64) {
	o.SdTimePeriod = &v
}

// GetSd returns the Sd field value if set, zero value otherwise.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSd() float64 {
	if o == nil || IsNil(o.Sd) {
		var ret float64
		return ret
	}
	return *o.Sd
}

// GetSdOk returns a tuple with the Sd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSdOk() (*float64, bool) {
	if o == nil || IsNil(o.Sd) {
		return nil, false
	}
	return o.Sd, true
}

// HasSd returns a boolean if a field has been set.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) HasSd() bool {
	if o != nil && !IsNil(o.Sd) {
		return true
	}

	return false
}

// SetSd gets a reference to the given float64 and assigns it to the Sd field.
func (o *GetTimeSeriesVwap200ResponseMetaIndicator) SetSd(v float64) {
	o.Sd = &v
}

func (o GetTimeSeriesVwap200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesVwap200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.SdTimePeriod) {
		toSerialize["sd_time_period"] = o.SdTimePeriod
	}
	if !IsNil(o.Sd) {
		toSerialize["sd"] = o.Sd
	}
	return toSerialize, nil
}

func (o *GetTimeSeriesVwap200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGetTimeSeriesVwap200ResponseMetaIndicator := _GetTimeSeriesVwap200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesVwap200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesVwap200ResponseMetaIndicator(varGetTimeSeriesVwap200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesVwap200ResponseMetaIndicator struct {
	value *GetTimeSeriesVwap200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesVwap200ResponseMetaIndicator) Get() *GetTimeSeriesVwap200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesVwap200ResponseMetaIndicator) Set(val *GetTimeSeriesVwap200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesVwap200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesVwap200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesVwap200ResponseMetaIndicator(val *GetTimeSeriesVwap200ResponseMetaIndicator) *NullableGetTimeSeriesVwap200ResponseMetaIndicator {
	return &NullableGetTimeSeriesVwap200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesVwap200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesVwap200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
