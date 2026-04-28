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

// checks if the GetTimeSeriesAdd200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdd200ResponseMetaIndicator{}

// GetTimeSeriesAdd200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesAdd200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type used as the first part of technical indicator
	SeriesType1 string `json:"series_type_1"`
	// Price type used as the second part of technical indicator
	SeriesType2 string `json:"series_type_2"`
}

type _GetTimeSeriesAdd200ResponseMetaIndicator GetTimeSeriesAdd200ResponseMetaIndicator

// NewGetTimeSeriesAdd200ResponseMetaIndicator instantiates a new GetTimeSeriesAdd200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdd200ResponseMetaIndicator(name string, seriesType1 string, seriesType2 string) *GetTimeSeriesAdd200ResponseMetaIndicator {
	this := GetTimeSeriesAdd200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType1 = seriesType1
	this.SeriesType2 = seriesType2
	return &this
}

// NewGetTimeSeriesAdd200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesAdd200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdd200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesAdd200ResponseMetaIndicator {
	this := GetTimeSeriesAdd200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType1 returns the SeriesType1 field value
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) GetSeriesType1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType1
}

// GetSeriesType1Ok returns a tuple with the SeriesType1 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) GetSeriesType1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType1, true
}

// SetSeriesType1 sets field value
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) SetSeriesType1(v string) {
	o.SeriesType1 = v
}

// GetSeriesType2 returns the SeriesType2 field value
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) GetSeriesType2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType2
}

// GetSeriesType2Ok returns a tuple with the SeriesType2 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) GetSeriesType2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType2, true
}

// SetSeriesType2 sets field value
func (o *GetTimeSeriesAdd200ResponseMetaIndicator) SetSeriesType2(v string) {
	o.SeriesType2 = v
}

func (o GetTimeSeriesAdd200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdd200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type_1"] = o.SeriesType1
	toSerialize["series_type_2"] = o.SeriesType2
	return toSerialize, nil
}

func (o *GetTimeSeriesAdd200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type_1",
		"series_type_2",
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

	varGetTimeSeriesAdd200ResponseMetaIndicator := _GetTimeSeriesAdd200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAdd200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdd200ResponseMetaIndicator(varGetTimeSeriesAdd200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesAdd200ResponseMetaIndicator struct {
	value *GetTimeSeriesAdd200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesAdd200ResponseMetaIndicator) Get() *GetTimeSeriesAdd200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesAdd200ResponseMetaIndicator) Set(val *GetTimeSeriesAdd200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdd200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdd200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdd200ResponseMetaIndicator(val *GetTimeSeriesAdd200ResponseMetaIndicator) *NullableGetTimeSeriesAdd200ResponseMetaIndicator {
	return &NullableGetTimeSeriesAdd200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdd200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdd200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
