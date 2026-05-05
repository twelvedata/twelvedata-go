// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesBBands200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesBBands200ResponseValuesInner{}

// GetTimeSeriesBBands200ResponseValuesInner struct for GetTimeSeriesBBands200ResponseValuesInner
type GetTimeSeriesBBands200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Upper band value
	UpperBand string `json:"upper_band"`
	// Middle band value
	MiddleBand string `json:"middle_band"`
	// Lower band value
	LowerBand string `json:"lower_band"`
}

type _GetTimeSeriesBBands200ResponseValuesInner GetTimeSeriesBBands200ResponseValuesInner

// NewGetTimeSeriesBBands200ResponseValuesInner instantiates a new GetTimeSeriesBBands200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesBBands200ResponseValuesInner(datetime string, upperBand string, middleBand string, lowerBand string) *GetTimeSeriesBBands200ResponseValuesInner {
	this := GetTimeSeriesBBands200ResponseValuesInner{}
	this.Datetime = datetime
	this.UpperBand = upperBand
	this.MiddleBand = middleBand
	this.LowerBand = lowerBand
	return &this
}

// NewGetTimeSeriesBBands200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesBBands200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesBBands200ResponseValuesInnerWithDefaults() *GetTimeSeriesBBands200ResponseValuesInner {
	this := GetTimeSeriesBBands200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetUpperBand returns the UpperBand field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetUpperBand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpperBand
}

// GetUpperBandOk returns a tuple with the UpperBand field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetUpperBandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpperBand, true
}

// SetUpperBand sets field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) SetUpperBand(v string) {
	o.UpperBand = v
}

// GetMiddleBand returns the MiddleBand field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetMiddleBand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MiddleBand
}

// GetMiddleBandOk returns a tuple with the MiddleBand field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetMiddleBandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MiddleBand, true
}

// SetMiddleBand sets field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) SetMiddleBand(v string) {
	o.MiddleBand = v
}

// GetLowerBand returns the LowerBand field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetLowerBand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LowerBand
}

// GetLowerBandOk returns a tuple with the LowerBand field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseValuesInner) GetLowerBandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowerBand, true
}

// SetLowerBand sets field value
func (o *GetTimeSeriesBBands200ResponseValuesInner) SetLowerBand(v string) {
	o.LowerBand = v
}

func (o GetTimeSeriesBBands200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesBBands200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["upper_band"] = o.UpperBand
	toSerialize["middle_band"] = o.MiddleBand
	toSerialize["lower_band"] = o.LowerBand
	return toSerialize, nil
}

func (o *GetTimeSeriesBBands200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"upper_band",
		"middle_band",
		"lower_band",
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

	varGetTimeSeriesBBands200ResponseValuesInner := _GetTimeSeriesBBands200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesBBands200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesBBands200ResponseValuesInner(varGetTimeSeriesBBands200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesBBands200ResponseValuesInner struct {
	value *GetTimeSeriesBBands200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesBBands200ResponseValuesInner) Get() *GetTimeSeriesBBands200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesBBands200ResponseValuesInner) Set(val *GetTimeSeriesBBands200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesBBands200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesBBands200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesBBands200ResponseValuesInner(val *GetTimeSeriesBBands200ResponseValuesInner) *NullableGetTimeSeriesBBands200ResponseValuesInner {
	return &NullableGetTimeSeriesBBands200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesBBands200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesBBands200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
