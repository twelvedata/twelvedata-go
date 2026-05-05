// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TimeSeriesCrossItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSeriesCrossItem{}

// TimeSeriesCrossItem TimeSeriesCrossItem represents a single data point in the time series
type TimeSeriesCrossItem struct {
	// Datetime at local exchange time referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Price at the opening of the current bar
	Open string `json:"open"`
	// Highest price which occurred during the current bar
	High string `json:"high"`
	// Lowest price which occurred during the current bar
	Low string `json:"low"`
	// Close price at the end of the bar
	Close string `json:"close"`
}

type _TimeSeriesCrossItem TimeSeriesCrossItem

// NewTimeSeriesCrossItem instantiates a new TimeSeriesCrossItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesCrossItem(datetime string, open string, high string, low string, close string) *TimeSeriesCrossItem {
	this := TimeSeriesCrossItem{}
	this.Datetime = datetime
	this.Open = open
	this.High = high
	this.Low = low
	this.Close = close
	return &this
}

// NewTimeSeriesCrossItemWithDefaults instantiates a new TimeSeriesCrossItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesCrossItemWithDefaults() *TimeSeriesCrossItem {
	this := TimeSeriesCrossItem{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *TimeSeriesCrossItem) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesCrossItem) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *TimeSeriesCrossItem) SetDatetime(v string) {
	o.Datetime = v
}

// GetOpen returns the Open field value
func (o *TimeSeriesCrossItem) GetOpen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Open
}

// GetOpenOk returns a tuple with the Open field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesCrossItem) GetOpenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Open, true
}

// SetOpen sets field value
func (o *TimeSeriesCrossItem) SetOpen(v string) {
	o.Open = v
}

// GetHigh returns the High field value
func (o *TimeSeriesCrossItem) GetHigh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesCrossItem) GetHighOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value
func (o *TimeSeriesCrossItem) SetHigh(v string) {
	o.High = v
}

// GetLow returns the Low field value
func (o *TimeSeriesCrossItem) GetLow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesCrossItem) GetLowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value
func (o *TimeSeriesCrossItem) SetLow(v string) {
	o.Low = v
}

// GetClose returns the Close field value
func (o *TimeSeriesCrossItem) GetClose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Close
}

// GetCloseOk returns a tuple with the Close field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesCrossItem) GetCloseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Close, true
}

// SetClose sets field value
func (o *TimeSeriesCrossItem) SetClose(v string) {
	o.Close = v
}

func (o TimeSeriesCrossItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSeriesCrossItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["open"] = o.Open
	toSerialize["high"] = o.High
	toSerialize["low"] = o.Low
	toSerialize["close"] = o.Close
	return toSerialize, nil
}

func (o *TimeSeriesCrossItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"open",
		"high",
		"low",
		"close",
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

	varTimeSeriesCrossItem := _TimeSeriesCrossItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTimeSeriesCrossItem)

	if err != nil {
		return err
	}

	*o = TimeSeriesCrossItem(varTimeSeriesCrossItem)

	return err
}

type NullableTimeSeriesCrossItem struct {
	value *TimeSeriesCrossItem
	isSet bool
}

func (v NullableTimeSeriesCrossItem) Get() *TimeSeriesCrossItem {
	return v.value
}

func (v *NullableTimeSeriesCrossItem) Set(val *TimeSeriesCrossItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesCrossItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesCrossItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesCrossItem(val *TimeSeriesCrossItem) *NullableTimeSeriesCrossItem {
	return &NullableTimeSeriesCrossItem{value: val, isSet: true}
}

func (v NullableTimeSeriesCrossItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesCrossItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
