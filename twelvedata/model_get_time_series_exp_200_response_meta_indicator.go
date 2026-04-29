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

// checks if the GetTimeSeriesExp200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesExp200ResponseMetaIndicator{}

// GetTimeSeriesExp200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesExp200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _GetTimeSeriesExp200ResponseMetaIndicator GetTimeSeriesExp200ResponseMetaIndicator

// NewGetTimeSeriesExp200ResponseMetaIndicator instantiates a new GetTimeSeriesExp200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesExp200ResponseMetaIndicator(name string, seriesType string) *GetTimeSeriesExp200ResponseMetaIndicator {
	this := GetTimeSeriesExp200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewGetTimeSeriesExp200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesExp200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesExp200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesExp200ResponseMetaIndicator {
	this := GetTimeSeriesExp200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesExp200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesExp200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesExp200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesExp200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesExp200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesExp200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o GetTimeSeriesExp200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesExp200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *GetTimeSeriesExp200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
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

	varGetTimeSeriesExp200ResponseMetaIndicator := _GetTimeSeriesExp200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesExp200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesExp200ResponseMetaIndicator(varGetTimeSeriesExp200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesExp200ResponseMetaIndicator struct {
	value *GetTimeSeriesExp200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesExp200ResponseMetaIndicator) Get() *GetTimeSeriesExp200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesExp200ResponseMetaIndicator) Set(val *GetTimeSeriesExp200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesExp200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesExp200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesExp200ResponseMetaIndicator(val *GetTimeSeriesExp200ResponseMetaIndicator) *NullableGetTimeSeriesExp200ResponseMetaIndicator {
	return &NullableGetTimeSeriesExp200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesExp200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesExp200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
