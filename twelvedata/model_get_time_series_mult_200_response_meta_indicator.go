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

// checks if the GetTimeSeriesMult200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMult200ResponseMetaIndicator{}

// GetTimeSeriesMult200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMult200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Specifies the first price data type
	SeriesType1 string `json:"series_type_1"`
	// Specifies the second price data type
	SeriesType2 string `json:"series_type_2"`
}

type _GetTimeSeriesMult200ResponseMetaIndicator GetTimeSeriesMult200ResponseMetaIndicator

// NewGetTimeSeriesMult200ResponseMetaIndicator instantiates a new GetTimeSeriesMult200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMult200ResponseMetaIndicator(name string, seriesType1 string, seriesType2 string) *GetTimeSeriesMult200ResponseMetaIndicator {
	this := GetTimeSeriesMult200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType1 = seriesType1
	this.SeriesType2 = seriesType2
	return &this
}

// NewGetTimeSeriesMult200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMult200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMult200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMult200ResponseMetaIndicator {
	this := GetTimeSeriesMult200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMult200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMult200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType1 returns the SeriesType1 field value
func (o *GetTimeSeriesMult200ResponseMetaIndicator) GetSeriesType1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType1
}

// GetSeriesType1Ok returns a tuple with the SeriesType1 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200ResponseMetaIndicator) GetSeriesType1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType1, true
}

// SetSeriesType1 sets field value
func (o *GetTimeSeriesMult200ResponseMetaIndicator) SetSeriesType1(v string) {
	o.SeriesType1 = v
}

// GetSeriesType2 returns the SeriesType2 field value
func (o *GetTimeSeriesMult200ResponseMetaIndicator) GetSeriesType2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType2
}

// GetSeriesType2Ok returns a tuple with the SeriesType2 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200ResponseMetaIndicator) GetSeriesType2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType2, true
}

// SetSeriesType2 sets field value
func (o *GetTimeSeriesMult200ResponseMetaIndicator) SetSeriesType2(v string) {
	o.SeriesType2 = v
}

func (o GetTimeSeriesMult200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMult200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type_1"] = o.SeriesType1
	toSerialize["series_type_2"] = o.SeriesType2
	return toSerialize, nil
}

func (o *GetTimeSeriesMult200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMult200ResponseMetaIndicator := _GetTimeSeriesMult200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMult200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMult200ResponseMetaIndicator(varGetTimeSeriesMult200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMult200ResponseMetaIndicator struct {
	value *GetTimeSeriesMult200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMult200ResponseMetaIndicator) Get() *GetTimeSeriesMult200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMult200ResponseMetaIndicator) Set(val *GetTimeSeriesMult200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMult200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMult200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMult200ResponseMetaIndicator(val *GetTimeSeriesMult200ResponseMetaIndicator) *NullableGetTimeSeriesMult200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMult200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMult200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMult200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
