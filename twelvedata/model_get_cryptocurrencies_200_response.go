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

// checks if the GetCryptocurrencies200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCryptocurrencies200Response{}

// GetCryptocurrencies200Response struct for GetCryptocurrencies200Response
type GetCryptocurrencies200Response struct {
	// Count
	Count int64 `json:"count"`
	// List of cryptocurrencies
	Data []CryptocurrencyResponseItem `json:"data"`
	// Response status
	Status string `json:"status"`
}

type _GetCryptocurrencies200Response GetCryptocurrencies200Response

// NewGetCryptocurrencies200Response instantiates a new GetCryptocurrencies200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCryptocurrencies200Response(count int64, data []CryptocurrencyResponseItem, status string) *GetCryptocurrencies200Response {
	this := GetCryptocurrencies200Response{}
	this.Count = count
	this.Data = data
	this.Status = status
	return &this
}

// NewGetCryptocurrencies200ResponseWithDefaults instantiates a new GetCryptocurrencies200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCryptocurrencies200ResponseWithDefaults() *GetCryptocurrencies200Response {
	this := GetCryptocurrencies200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *GetCryptocurrencies200Response) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetCryptocurrencies200Response) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetCryptocurrencies200Response) SetCount(v int64) {
	o.Count = v
}

// GetData returns the Data field value
func (o *GetCryptocurrencies200Response) GetData() []CryptocurrencyResponseItem {
	if o == nil {
		var ret []CryptocurrencyResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCryptocurrencies200Response) GetDataOk() ([]CryptocurrencyResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCryptocurrencies200Response) SetData(v []CryptocurrencyResponseItem) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetCryptocurrencies200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetCryptocurrencies200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetCryptocurrencies200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetCryptocurrencies200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCryptocurrencies200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetCryptocurrencies200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"data",
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

	varGetCryptocurrencies200Response := _GetCryptocurrencies200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCryptocurrencies200Response)

	if err != nil {
		return err
	}

	*o = GetCryptocurrencies200Response(varGetCryptocurrencies200Response)

	return err
}

type NullableGetCryptocurrencies200Response struct {
	value *GetCryptocurrencies200Response
	isSet bool
}

func (v NullableGetCryptocurrencies200Response) Get() *GetCryptocurrencies200Response {
	return v.value
}

func (v *NullableGetCryptocurrencies200Response) Set(val *GetCryptocurrencies200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCryptocurrencies200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCryptocurrencies200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCryptocurrencies200Response(val *GetCryptocurrencies200Response) *NullableGetCryptocurrencies200Response {
	return &NullableGetCryptocurrencies200Response{value: val, isSet: true}
}

func (v NullableGetCryptocurrencies200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCryptocurrencies200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
