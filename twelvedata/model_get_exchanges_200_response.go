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

// checks if the GetExchanges200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExchanges200Response{}

// GetExchanges200Response struct for GetExchanges200Response
type GetExchanges200Response struct {
	// List of exchanges
	Data []ExchangesResponseItem `json:"data"`
	// Response status
	Status string `json:"status"`
}

type _GetExchanges200Response GetExchanges200Response

// NewGetExchanges200Response instantiates a new GetExchanges200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchanges200Response(data []ExchangesResponseItem, status string) *GetExchanges200Response {
	this := GetExchanges200Response{}
	this.Data = data
	this.Status = status
	return &this
}

// NewGetExchanges200ResponseWithDefaults instantiates a new GetExchanges200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchanges200ResponseWithDefaults() *GetExchanges200Response {
	this := GetExchanges200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetExchanges200Response) GetData() []ExchangesResponseItem {
	if o == nil {
		var ret []ExchangesResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetExchanges200Response) GetDataOk() ([]ExchangesResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetExchanges200Response) SetData(v []ExchangesResponseItem) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetExchanges200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetExchanges200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetExchanges200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetExchanges200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchanges200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetExchanges200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGetExchanges200Response := _GetExchanges200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetExchanges200Response)

	if err != nil {
		return err
	}

	*o = GetExchanges200Response(varGetExchanges200Response)

	return err
}

type NullableGetExchanges200Response struct {
	value *GetExchanges200Response
	isSet bool
}

func (v NullableGetExchanges200Response) Get() *GetExchanges200Response {
	return v.value
}

func (v *NullableGetExchanges200Response) Set(val *GetExchanges200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchanges200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchanges200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchanges200Response(val *GetExchanges200Response) *NullableGetExchanges200Response {
	return &NullableGetExchanges200Response{value: val, isSet: true}
}

func (v NullableGetExchanges200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchanges200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
