// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ExchangeScheduleSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeScheduleSession{}

// ExchangeScheduleSession ExchangeScheduleSession represents a trading session
type ExchangeScheduleSession struct {
	// Opening time of the session
	OpenTime string `json:"open_time"`
	// Closing time of the session
	CloseTime string `json:"close_time"`
	// Name of the session
	SessionName string `json:"session_name"`
	// Type of the session
	SessionType string `json:"session_type"`
}

type _ExchangeScheduleSession ExchangeScheduleSession

// NewExchangeScheduleSession instantiates a new ExchangeScheduleSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeScheduleSession(openTime string, closeTime string, sessionName string, sessionType string) *ExchangeScheduleSession {
	this := ExchangeScheduleSession{}
	this.OpenTime = openTime
	this.CloseTime = closeTime
	this.SessionName = sessionName
	this.SessionType = sessionType
	return &this
}

// NewExchangeScheduleSessionWithDefaults instantiates a new ExchangeScheduleSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeScheduleSessionWithDefaults() *ExchangeScheduleSession {
	this := ExchangeScheduleSession{}
	return &this
}

// GetOpenTime returns the OpenTime field value
func (o *ExchangeScheduleSession) GetOpenTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleSession) GetOpenTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenTime, true
}

// SetOpenTime sets field value
func (o *ExchangeScheduleSession) SetOpenTime(v string) {
	o.OpenTime = v
}

// GetCloseTime returns the CloseTime field value
func (o *ExchangeScheduleSession) GetCloseTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleSession) GetCloseTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloseTime, true
}

// SetCloseTime sets field value
func (o *ExchangeScheduleSession) SetCloseTime(v string) {
	o.CloseTime = v
}

// GetSessionName returns the SessionName field value
func (o *ExchangeScheduleSession) GetSessionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleSession) GetSessionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionName, true
}

// SetSessionName sets field value
func (o *ExchangeScheduleSession) SetSessionName(v string) {
	o.SessionName = v
}

// GetSessionType returns the SessionType field value
func (o *ExchangeScheduleSession) GetSessionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionType
}

// GetSessionTypeOk returns a tuple with the SessionType field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleSession) GetSessionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionType, true
}

// SetSessionType sets field value
func (o *ExchangeScheduleSession) SetSessionType(v string) {
	o.SessionType = v
}

func (o ExchangeScheduleSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeScheduleSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["open_time"] = o.OpenTime
	toSerialize["close_time"] = o.CloseTime
	toSerialize["session_name"] = o.SessionName
	toSerialize["session_type"] = o.SessionType
	return toSerialize, nil
}

func (o *ExchangeScheduleSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"open_time",
		"close_time",
		"session_name",
		"session_type",
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

	varExchangeScheduleSession := _ExchangeScheduleSession{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varExchangeScheduleSession)

	if err != nil {
		return err
	}

	*o = ExchangeScheduleSession(varExchangeScheduleSession)

	return err
}

type NullableExchangeScheduleSession struct {
	value *ExchangeScheduleSession
	isSet bool
}

func (v NullableExchangeScheduleSession) Get() *ExchangeScheduleSession {
	return v.value
}

func (v *NullableExchangeScheduleSession) Set(val *ExchangeScheduleSession) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeScheduleSession) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeScheduleSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeScheduleSession(val *ExchangeScheduleSession) *NullableExchangeScheduleSession {
	return &NullableExchangeScheduleSession{value: val, isSet: true}
}

func (v NullableExchangeScheduleSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeScheduleSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
