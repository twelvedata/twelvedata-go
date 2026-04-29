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

// checks if the GetTimeSeriesFloor200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesFloor200ResponseMetaIndicator{}

// GetTimeSeriesFloor200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesFloor200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _GetTimeSeriesFloor200ResponseMetaIndicator GetTimeSeriesFloor200ResponseMetaIndicator

// NewGetTimeSeriesFloor200ResponseMetaIndicator instantiates a new GetTimeSeriesFloor200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesFloor200ResponseMetaIndicator(name string, seriesType string) *GetTimeSeriesFloor200ResponseMetaIndicator {
	this := GetTimeSeriesFloor200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewGetTimeSeriesFloor200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesFloor200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesFloor200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesFloor200ResponseMetaIndicator {
	this := GetTimeSeriesFloor200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesFloor200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesFloor200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesFloor200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesFloor200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesFloor200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesFloor200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o GetTimeSeriesFloor200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesFloor200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *GetTimeSeriesFloor200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesFloor200ResponseMetaIndicator := _GetTimeSeriesFloor200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesFloor200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesFloor200ResponseMetaIndicator(varGetTimeSeriesFloor200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesFloor200ResponseMetaIndicator struct {
	value *GetTimeSeriesFloor200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesFloor200ResponseMetaIndicator) Get() *GetTimeSeriesFloor200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesFloor200ResponseMetaIndicator) Set(val *GetTimeSeriesFloor200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesFloor200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesFloor200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesFloor200ResponseMetaIndicator(val *GetTimeSeriesFloor200ResponseMetaIndicator) *NullableGetTimeSeriesFloor200ResponseMetaIndicator {
	return &NullableGetTimeSeriesFloor200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesFloor200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesFloor200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
