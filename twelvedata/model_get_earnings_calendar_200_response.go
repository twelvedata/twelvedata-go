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

// checks if the GetEarningsCalendar200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarningsCalendar200Response{}

// GetEarningsCalendar200Response struct for GetEarningsCalendar200Response
type GetEarningsCalendar200Response struct {
	// Map of dates to earnings data
	Earnings map[string][]GetEarningsCalendar200ResponseEarningsValueInner `json:"earnings"`
	// Response status
	Status string `json:"status"`
}

type _GetEarningsCalendar200Response GetEarningsCalendar200Response

// NewGetEarningsCalendar200Response instantiates a new GetEarningsCalendar200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarningsCalendar200Response(earnings map[string][]GetEarningsCalendar200ResponseEarningsValueInner, status string) *GetEarningsCalendar200Response {
	this := GetEarningsCalendar200Response{}
	this.Earnings = earnings
	this.Status = status
	return &this
}

// NewGetEarningsCalendar200ResponseWithDefaults instantiates a new GetEarningsCalendar200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarningsCalendar200ResponseWithDefaults() *GetEarningsCalendar200Response {
	this := GetEarningsCalendar200Response{}
	return &this
}

// GetEarnings returns the Earnings field value
func (o *GetEarningsCalendar200Response) GetEarnings() map[string][]GetEarningsCalendar200ResponseEarningsValueInner {
	if o == nil {
		var ret map[string][]GetEarningsCalendar200ResponseEarningsValueInner
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200Response) GetEarningsOk() (map[string][]GetEarningsCalendar200ResponseEarningsValueInner, bool) {
	if o == nil {
		return map[string][]GetEarningsCalendar200ResponseEarningsValueInner{}, false
	}
	return o.Earnings, true
}

// SetEarnings sets field value
func (o *GetEarningsCalendar200Response) SetEarnings(v map[string][]GetEarningsCalendar200ResponseEarningsValueInner) {
	o.Earnings = v
}

// GetStatus returns the Status field value
func (o *GetEarningsCalendar200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetEarningsCalendar200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetEarningsCalendar200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetEarningsCalendar200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarningsCalendar200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["earnings"] = o.Earnings
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetEarningsCalendar200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"earnings",
		"status",
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

	varGetEarningsCalendar200Response := _GetEarningsCalendar200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEarningsCalendar200Response)

	if err != nil {
		return err
	}

	*o = GetEarningsCalendar200Response(varGetEarningsCalendar200Response)

	return err
}

type NullableGetEarningsCalendar200Response struct {
	value *GetEarningsCalendar200Response
	isSet bool
}

func (v NullableGetEarningsCalendar200Response) Get() *GetEarningsCalendar200Response {
	return v.value
}

func (v *NullableGetEarningsCalendar200Response) Set(val *GetEarningsCalendar200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarningsCalendar200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarningsCalendar200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarningsCalendar200Response(val *GetEarningsCalendar200Response) *NullableGetEarningsCalendar200Response {
	return &NullableGetEarningsCalendar200Response{value: val, isSet: true}
}

func (v NullableGetEarningsCalendar200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarningsCalendar200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
