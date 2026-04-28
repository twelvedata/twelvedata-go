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

// checks if the GetTimeSeriesVar200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesVar200ResponseValuesInner{}

// GetTimeSeriesVar200ResponseValuesInner struct for GetTimeSeriesVar200ResponseValuesInner
type GetTimeSeriesVar200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// VAR value
	Var string `json:"var"`
}

type _GetTimeSeriesVar200ResponseValuesInner GetTimeSeriesVar200ResponseValuesInner

// NewGetTimeSeriesVar200ResponseValuesInner instantiates a new GetTimeSeriesVar200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesVar200ResponseValuesInner(datetime string, var_ string) *GetTimeSeriesVar200ResponseValuesInner {
	this := GetTimeSeriesVar200ResponseValuesInner{}
	this.Datetime = datetime
	this.Var = var_
	return &this
}

// NewGetTimeSeriesVar200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesVar200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesVar200ResponseValuesInnerWithDefaults() *GetTimeSeriesVar200ResponseValuesInner {
	this := GetTimeSeriesVar200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesVar200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVar200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesVar200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetVar returns the Var field value
func (o *GetTimeSeriesVar200ResponseValuesInner) GetVar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Var
}

// GetVarOk returns a tuple with the Var field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVar200ResponseValuesInner) GetVarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var, true
}

// SetVar sets field value
func (o *GetTimeSeriesVar200ResponseValuesInner) SetVar(v string) {
	o.Var = v
}

func (o GetTimeSeriesVar200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesVar200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["var"] = o.Var
	return toSerialize, nil
}

func (o *GetTimeSeriesVar200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"var",
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

	varGetTimeSeriesVar200ResponseValuesInner := _GetTimeSeriesVar200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesVar200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesVar200ResponseValuesInner(varGetTimeSeriesVar200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesVar200ResponseValuesInner struct {
	value *GetTimeSeriesVar200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesVar200ResponseValuesInner) Get() *GetTimeSeriesVar200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesVar200ResponseValuesInner) Set(val *GetTimeSeriesVar200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesVar200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesVar200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesVar200ResponseValuesInner(val *GetTimeSeriesVar200ResponseValuesInner) *NullableGetTimeSeriesVar200ResponseValuesInner {
	return &NullableGetTimeSeriesVar200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesVar200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesVar200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
