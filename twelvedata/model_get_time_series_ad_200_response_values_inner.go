// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAd200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAd200ResponseValuesInner{}

// GetTimeSeriesAd200ResponseValuesInner struct for GetTimeSeriesAd200ResponseValuesInner
type GetTimeSeriesAd200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// AD value
	Ad string `json:"ad"`
}

type _GetTimeSeriesAd200ResponseValuesInner GetTimeSeriesAd200ResponseValuesInner

// NewGetTimeSeriesAd200ResponseValuesInner instantiates a new GetTimeSeriesAd200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAd200ResponseValuesInner(datetime string, ad string) *GetTimeSeriesAd200ResponseValuesInner {
	this := GetTimeSeriesAd200ResponseValuesInner{}
	this.Datetime = datetime
	this.Ad = ad
	return &this
}

// NewGetTimeSeriesAd200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAd200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAd200ResponseValuesInnerWithDefaults() *GetTimeSeriesAd200ResponseValuesInner {
	this := GetTimeSeriesAd200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAd200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAd200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAd200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAd returns the Ad field value
func (o *GetTimeSeriesAd200ResponseValuesInner) GetAd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ad
}

// GetAdOk returns a tuple with the Ad field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAd200ResponseValuesInner) GetAdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ad, true
}

// SetAd sets field value
func (o *GetTimeSeriesAd200ResponseValuesInner) SetAd(v string) {
	o.Ad = v
}

func (o GetTimeSeriesAd200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAd200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ad"] = o.Ad
	return toSerialize, nil
}

func (o *GetTimeSeriesAd200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ad",
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

	varGetTimeSeriesAd200ResponseValuesInner := _GetTimeSeriesAd200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAd200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAd200ResponseValuesInner(varGetTimeSeriesAd200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAd200ResponseValuesInner struct {
	value *GetTimeSeriesAd200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAd200ResponseValuesInner) Get() *GetTimeSeriesAd200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAd200ResponseValuesInner) Set(val *GetTimeSeriesAd200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAd200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAd200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAd200ResponseValuesInner(val *GetTimeSeriesAd200ResponseValuesInner) *NullableGetTimeSeriesAd200ResponseValuesInner {
	return &NullableGetTimeSeriesAd200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAd200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAd200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
