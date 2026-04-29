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

// checks if the GetTimeSeriesKeltner200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKeltner200ResponseValuesInner{}

// GetTimeSeriesKeltner200ResponseValuesInner struct for GetTimeSeriesKeltner200ResponseValuesInner
type GetTimeSeriesKeltner200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Upper line value
	UpperLine string `json:"upper_line"`
	// Middle line value
	MiddleLine string `json:"middle_line"`
	// Lower line value
	LowerLine string `json:"lower_line"`
}

type _GetTimeSeriesKeltner200ResponseValuesInner GetTimeSeriesKeltner200ResponseValuesInner

// NewGetTimeSeriesKeltner200ResponseValuesInner instantiates a new GetTimeSeriesKeltner200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKeltner200ResponseValuesInner(datetime string, upperLine string, middleLine string, lowerLine string) *GetTimeSeriesKeltner200ResponseValuesInner {
	this := GetTimeSeriesKeltner200ResponseValuesInner{}
	this.Datetime = datetime
	this.UpperLine = upperLine
	this.MiddleLine = middleLine
	this.LowerLine = lowerLine
	return &this
}

// NewGetTimeSeriesKeltner200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesKeltner200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKeltner200ResponseValuesInnerWithDefaults() *GetTimeSeriesKeltner200ResponseValuesInner {
	this := GetTimeSeriesKeltner200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetUpperLine returns the UpperLine field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetUpperLine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpperLine
}

// GetUpperLineOk returns a tuple with the UpperLine field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetUpperLineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpperLine, true
}

// SetUpperLine sets field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetUpperLine(v string) {
	o.UpperLine = v
}

// GetMiddleLine returns the MiddleLine field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetMiddleLine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MiddleLine
}

// GetMiddleLineOk returns a tuple with the MiddleLine field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetMiddleLineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MiddleLine, true
}

// SetMiddleLine sets field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetMiddleLine(v string) {
	o.MiddleLine = v
}

// GetLowerLine returns the LowerLine field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetLowerLine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LowerLine
}

// GetLowerLineOk returns a tuple with the LowerLine field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseValuesInner) GetLowerLineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowerLine, true
}

// SetLowerLine sets field value
func (o *GetTimeSeriesKeltner200ResponseValuesInner) SetLowerLine(v string) {
	o.LowerLine = v
}

func (o GetTimeSeriesKeltner200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKeltner200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["upper_line"] = o.UpperLine
	toSerialize["middle_line"] = o.MiddleLine
	toSerialize["lower_line"] = o.LowerLine
	return toSerialize, nil
}

func (o *GetTimeSeriesKeltner200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"upper_line",
		"middle_line",
		"lower_line",
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

	varGetTimeSeriesKeltner200ResponseValuesInner := _GetTimeSeriesKeltner200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesKeltner200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKeltner200ResponseValuesInner(varGetTimeSeriesKeltner200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesKeltner200ResponseValuesInner struct {
	value *GetTimeSeriesKeltner200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesKeltner200ResponseValuesInner) Get() *GetTimeSeriesKeltner200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesKeltner200ResponseValuesInner) Set(val *GetTimeSeriesKeltner200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKeltner200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKeltner200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKeltner200ResponseValuesInner(val *GetTimeSeriesKeltner200ResponseValuesInner) *NullableGetTimeSeriesKeltner200ResponseValuesInner {
	return &NullableGetTimeSeriesKeltner200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKeltner200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKeltner200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
