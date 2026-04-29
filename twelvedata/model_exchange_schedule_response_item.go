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

// checks if the ExchangeScheduleResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeScheduleResponseItem{}

// ExchangeScheduleResponseItem ExchangeScheduleResponseItem represents details of an exchange schedule
type ExchangeScheduleResponseItem struct {
	// Official name of exchange
	Title string `json:"title"`
	// Name of exchange
	Name string `json:"name"`
	// Market identifier code (MIC) under ISO 10383 standard
	Code string `json:"code"`
	// Country to which stock exchange belongs to
	Country string `json:"country"`
	// Time zone where exchange is located
	TimeZone string `json:"time_zone"`
	// Exchange trading hours
	Sessions []ExchangeScheduleSession `json:"sessions"`
}

type _ExchangeScheduleResponseItem ExchangeScheduleResponseItem

// NewExchangeScheduleResponseItem instantiates a new ExchangeScheduleResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeScheduleResponseItem(title string, name string, code string, country string, timeZone string, sessions []ExchangeScheduleSession) *ExchangeScheduleResponseItem {
	this := ExchangeScheduleResponseItem{}
	this.Title = title
	this.Name = name
	this.Code = code
	this.Country = country
	this.TimeZone = timeZone
	this.Sessions = sessions
	return &this
}

// NewExchangeScheduleResponseItemWithDefaults instantiates a new ExchangeScheduleResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeScheduleResponseItemWithDefaults() *ExchangeScheduleResponseItem {
	this := ExchangeScheduleResponseItem{}
	return &this
}

// GetTitle returns the Title field value
func (o *ExchangeScheduleResponseItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleResponseItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ExchangeScheduleResponseItem) SetTitle(v string) {
	o.Title = v
}

// GetName returns the Name field value
func (o *ExchangeScheduleResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExchangeScheduleResponseItem) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *ExchangeScheduleResponseItem) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleResponseItem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ExchangeScheduleResponseItem) SetCode(v string) {
	o.Code = v
}

// GetCountry returns the Country field value
func (o *ExchangeScheduleResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *ExchangeScheduleResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetTimeZone returns the TimeZone field value
func (o *ExchangeScheduleResponseItem) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleResponseItem) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *ExchangeScheduleResponseItem) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetSessions returns the Sessions field value
func (o *ExchangeScheduleResponseItem) GetSessions() []ExchangeScheduleSession {
	if o == nil {
		var ret []ExchangeScheduleSession
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *ExchangeScheduleResponseItem) GetSessionsOk() ([]ExchangeScheduleSession, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *ExchangeScheduleResponseItem) SetSessions(v []ExchangeScheduleSession) {
	o.Sessions = v
}

func (o ExchangeScheduleResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeScheduleResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["name"] = o.Name
	toSerialize["code"] = o.Code
	toSerialize["country"] = o.Country
	toSerialize["time_zone"] = o.TimeZone
	toSerialize["sessions"] = o.Sessions
	return toSerialize, nil
}

func (o *ExchangeScheduleResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"name",
		"code",
		"country",
		"time_zone",
		"sessions",
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

	varExchangeScheduleResponseItem := _ExchangeScheduleResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varExchangeScheduleResponseItem)

	if err != nil {
		return err
	}

	*o = ExchangeScheduleResponseItem(varExchangeScheduleResponseItem)

	return err
}

type NullableExchangeScheduleResponseItem struct {
	value *ExchangeScheduleResponseItem
	isSet bool
}

func (v NullableExchangeScheduleResponseItem) Get() *ExchangeScheduleResponseItem {
	return v.value
}

func (v *NullableExchangeScheduleResponseItem) Set(val *ExchangeScheduleResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeScheduleResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeScheduleResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeScheduleResponseItem(val *ExchangeScheduleResponseItem) *NullableExchangeScheduleResponseItem {
	return &NullableExchangeScheduleResponseItem{value: val, isSet: true}
}

func (v NullableExchangeScheduleResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeScheduleResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
