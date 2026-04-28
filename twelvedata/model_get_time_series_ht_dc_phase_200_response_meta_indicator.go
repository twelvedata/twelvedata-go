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

// checks if the GetTimeSeriesHtDcPhase200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtDcPhase200ResponseMetaIndicator{}

// GetTimeSeriesHtDcPhase200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesHtDcPhase200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _GetTimeSeriesHtDcPhase200ResponseMetaIndicator GetTimeSeriesHtDcPhase200ResponseMetaIndicator

// NewGetTimeSeriesHtDcPhase200ResponseMetaIndicator instantiates a new GetTimeSeriesHtDcPhase200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtDcPhase200ResponseMetaIndicator(name string, seriesType string) *GetTimeSeriesHtDcPhase200ResponseMetaIndicator {
	this := GetTimeSeriesHtDcPhase200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewGetTimeSeriesHtDcPhase200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesHtDcPhase200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtDcPhase200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesHtDcPhase200ResponseMetaIndicator {
	this := GetTimeSeriesHtDcPhase200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o GetTimeSeriesHtDcPhase200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtDcPhase200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHtDcPhase200ResponseMetaIndicator := _GetTimeSeriesHtDcPhase200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtDcPhase200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtDcPhase200ResponseMetaIndicator(varGetTimeSeriesHtDcPhase200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator struct {
	value *GetTimeSeriesHtDcPhase200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator) Get() *GetTimeSeriesHtDcPhase200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator) Set(val *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator(val *GetTimeSeriesHtDcPhase200ResponseMetaIndicator) *NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator {
	return &NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtDcPhase200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
