// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAroon200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAroon200ResponseValuesInner{}

// GetTimeSeriesAroon200ResponseValuesInner struct for GetTimeSeriesAroon200ResponseValuesInner
type GetTimeSeriesAroon200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Aroon down value
	AroonDown string `json:"aroon_down"`
	// Aroon up value
	AroonUp string `json:"aroon_up"`
}

type _GetTimeSeriesAroon200ResponseValuesInner GetTimeSeriesAroon200ResponseValuesInner

// NewGetTimeSeriesAroon200ResponseValuesInner instantiates a new GetTimeSeriesAroon200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAroon200ResponseValuesInner(datetime string, aroonDown string, aroonUp string) *GetTimeSeriesAroon200ResponseValuesInner {
	this := GetTimeSeriesAroon200ResponseValuesInner{}
	this.Datetime = datetime
	this.AroonDown = aroonDown
	this.AroonUp = aroonUp
	return &this
}

// NewGetTimeSeriesAroon200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAroon200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAroon200ResponseValuesInnerWithDefaults() *GetTimeSeriesAroon200ResponseValuesInner {
	this := GetTimeSeriesAroon200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAroon200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAroon200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAroonDown returns the AroonDown field value
func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonDown() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AroonDown
}

// GetAroonDownOk returns a tuple with the AroonDown field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonDownOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AroonDown, true
}

// SetAroonDown sets field value
func (o *GetTimeSeriesAroon200ResponseValuesInner) SetAroonDown(v string) {
	o.AroonDown = v
}

// GetAroonUp returns the AroonUp field value
func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonUp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AroonUp
}

// GetAroonUpOk returns a tuple with the AroonUp field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200ResponseValuesInner) GetAroonUpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AroonUp, true
}

// SetAroonUp sets field value
func (o *GetTimeSeriesAroon200ResponseValuesInner) SetAroonUp(v string) {
	o.AroonUp = v
}

func (o GetTimeSeriesAroon200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAroon200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["aroon_down"] = o.AroonDown
	toSerialize["aroon_up"] = o.AroonUp
	return toSerialize, nil
}

func (o *GetTimeSeriesAroon200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"aroon_down",
		"aroon_up",
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

	varGetTimeSeriesAroon200ResponseValuesInner := _GetTimeSeriesAroon200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAroon200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAroon200ResponseValuesInner(varGetTimeSeriesAroon200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAroon200ResponseValuesInner struct {
	value *GetTimeSeriesAroon200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAroon200ResponseValuesInner) Get() *GetTimeSeriesAroon200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAroon200ResponseValuesInner) Set(val *GetTimeSeriesAroon200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAroon200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAroon200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAroon200ResponseValuesInner(val *GetTimeSeriesAroon200ResponseValuesInner) *NullableGetTimeSeriesAroon200ResponseValuesInner {
	return &NullableGetTimeSeriesAroon200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAroon200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAroon200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
