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

// checks if the GetTimeSeriesHtTrendline200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtTrendline200ResponseMetaIndicator{}

// GetTimeSeriesHtTrendline200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesHtTrendline200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _GetTimeSeriesHtTrendline200ResponseMetaIndicator GetTimeSeriesHtTrendline200ResponseMetaIndicator

// NewGetTimeSeriesHtTrendline200ResponseMetaIndicator instantiates a new GetTimeSeriesHtTrendline200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtTrendline200ResponseMetaIndicator(name string, seriesType string) *GetTimeSeriesHtTrendline200ResponseMetaIndicator {
	this := GetTimeSeriesHtTrendline200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewGetTimeSeriesHtTrendline200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesHtTrendline200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtTrendline200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesHtTrendline200ResponseMetaIndicator {
	this := GetTimeSeriesHtTrendline200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesHtTrendline200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendline200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesHtTrendline200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesHtTrendline200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendline200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesHtTrendline200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o GetTimeSeriesHtTrendline200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtTrendline200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *GetTimeSeriesHtTrendline200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHtTrendline200ResponseMetaIndicator := _GetTimeSeriesHtTrendline200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtTrendline200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtTrendline200ResponseMetaIndicator(varGetTimeSeriesHtTrendline200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator struct {
	value *GetTimeSeriesHtTrendline200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator) Get() *GetTimeSeriesHtTrendline200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator) Set(val *GetTimeSeriesHtTrendline200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtTrendline200ResponseMetaIndicator(val *GetTimeSeriesHtTrendline200ResponseMetaIndicator) *NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator {
	return &NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtTrendline200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
