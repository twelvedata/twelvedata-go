// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TimeSeriesItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSeriesItem{}

// TimeSeriesItem struct for TimeSeriesItem
type TimeSeriesItem struct {
	// Datetime at local exchange time referring to when the bar with specified interval was opened.
	Datetime string `json:"datetime"`
	// Price at the opening of current bar
	Open string `json:"open"`
	// Highest price which occurred during the current bar.
	High string `json:"high"`
	// Lowest price which occurred during the current bar.
	Low string `json:"low"`
	// Close price at the end of the bar.
	Close string `json:"close"`
	// Trading volume which occurred during the current bar
	Volume *string `json:"volume,omitempty"`
}

type _TimeSeriesItem TimeSeriesItem

// NewTimeSeriesItem instantiates a new TimeSeriesItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesItem(datetime string, open string, high string, low string, close string) *TimeSeriesItem {
	this := TimeSeriesItem{}
	this.Datetime = datetime
	this.Open = open
	this.High = high
	this.Low = low
	this.Close = close
	return &this
}

// NewTimeSeriesItemWithDefaults instantiates a new TimeSeriesItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesItemWithDefaults() *TimeSeriesItem {
	this := TimeSeriesItem{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *TimeSeriesItem) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesItem) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *TimeSeriesItem) SetDatetime(v string) {
	o.Datetime = v
}

// GetOpen returns the Open field value
func (o *TimeSeriesItem) GetOpen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Open
}

// GetOpenOk returns a tuple with the Open field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesItem) GetOpenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Open, true
}

// SetOpen sets field value
func (o *TimeSeriesItem) SetOpen(v string) {
	o.Open = v
}

// GetHigh returns the High field value
func (o *TimeSeriesItem) GetHigh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesItem) GetHighOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value
func (o *TimeSeriesItem) SetHigh(v string) {
	o.High = v
}

// GetLow returns the Low field value
func (o *TimeSeriesItem) GetLow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesItem) GetLowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value
func (o *TimeSeriesItem) SetLow(v string) {
	o.Low = v
}

// GetClose returns the Close field value
func (o *TimeSeriesItem) GetClose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Close
}

// GetCloseOk returns a tuple with the Close field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesItem) GetCloseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Close, true
}

// SetClose sets field value
func (o *TimeSeriesItem) SetClose(v string) {
	o.Close = v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *TimeSeriesItem) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesItem) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *TimeSeriesItem) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *TimeSeriesItem) SetVolume(v string) {
	o.Volume = &v
}

func (o TimeSeriesItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSeriesItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["open"] = o.Open
	toSerialize["high"] = o.High
	toSerialize["low"] = o.Low
	toSerialize["close"] = o.Close
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

func (o *TimeSeriesItem) UnmarshalJSON(data []byte) (err error) {
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

	varTimeSeriesItem := _TimeSeriesItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTimeSeriesItem)

	if err != nil {
		return err
	}

	*o = TimeSeriesItem(varTimeSeriesItem)

	return err
}

type NullableTimeSeriesItem struct {
	value *TimeSeriesItem
	isSet bool
}

func (v NullableTimeSeriesItem) Get() *TimeSeriesItem {
	return v.value
}

func (v *NullableTimeSeriesItem) Set(val *TimeSeriesItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesItem(val *TimeSeriesItem) *NullableTimeSeriesItem {
	return &NullableTimeSeriesItem{value: val, isSet: true}
}

func (v NullableTimeSeriesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
