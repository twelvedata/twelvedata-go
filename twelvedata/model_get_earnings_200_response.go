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

// checks if the GetEarnings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarnings200Response{}

// GetEarnings200Response struct for GetEarnings200Response
type GetEarnings200Response struct {
	Meta GetEarnings200ResponseMeta `json:"meta"`
	// List of earnings data
	Earnings []EarningsItem `json:"earnings"`
	// Response status
	Status string `json:"status"`
}

type _GetEarnings200Response GetEarnings200Response

// NewGetEarnings200Response instantiates a new GetEarnings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarnings200Response(meta GetEarnings200ResponseMeta, earnings []EarningsItem, status string) *GetEarnings200Response {
	this := GetEarnings200Response{}
	this.Meta = meta
	this.Earnings = earnings
	this.Status = status
	return &this
}

// NewGetEarnings200ResponseWithDefaults instantiates a new GetEarnings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarnings200ResponseWithDefaults() *GetEarnings200Response {
	this := GetEarnings200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetEarnings200Response) GetMeta() GetEarnings200ResponseMeta {
	if o == nil {
		var ret GetEarnings200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200Response) GetMetaOk() (*GetEarnings200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetEarnings200Response) SetMeta(v GetEarnings200ResponseMeta) {
	o.Meta = v
}

// GetEarnings returns the Earnings field value
func (o *GetEarnings200Response) GetEarnings() []EarningsItem {
	if o == nil {
		var ret []EarningsItem
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200Response) GetEarningsOk() ([]EarningsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Earnings, true
}

// SetEarnings sets field value
func (o *GetEarnings200Response) SetEarnings(v []EarningsItem) {
	o.Earnings = v
}

// GetStatus returns the Status field value
func (o *GetEarnings200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetEarnings200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetEarnings200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarnings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["earnings"] = o.Earnings
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetEarnings200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
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

	varGetEarnings200Response := _GetEarnings200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEarnings200Response)

	if err != nil {
		return err
	}

	*o = GetEarnings200Response(varGetEarnings200Response)

	return err
}

type NullableGetEarnings200Response struct {
	value *GetEarnings200Response
	isSet bool
}

func (v NullableGetEarnings200Response) Get() *GetEarnings200Response {
	return v.value
}

func (v *NullableGetEarnings200Response) Set(val *GetEarnings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarnings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarnings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarnings200Response(val *GetEarnings200Response) *NullableGetEarnings200Response {
	return &NullableGetEarnings200Response{value: val, isSet: true}
}

func (v NullableGetEarnings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarnings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
