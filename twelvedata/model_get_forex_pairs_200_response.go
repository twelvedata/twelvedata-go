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

// checks if the GetForexPairs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetForexPairs200Response{}

// GetForexPairs200Response struct for GetForexPairs200Response
type GetForexPairs200Response struct {
	// List of forex pairs
	Data []ForexResponseItem `json:"data"`
	// Response status
	Status string `json:"status"`
}

type _GetForexPairs200Response GetForexPairs200Response

// NewGetForexPairs200Response instantiates a new GetForexPairs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetForexPairs200Response(data []ForexResponseItem, status string) *GetForexPairs200Response {
	this := GetForexPairs200Response{}
	this.Data = data
	this.Status = status
	return &this
}

// NewGetForexPairs200ResponseWithDefaults instantiates a new GetForexPairs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetForexPairs200ResponseWithDefaults() *GetForexPairs200Response {
	this := GetForexPairs200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetForexPairs200Response) GetData() []ForexResponseItem {
	if o == nil {
		var ret []ForexResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetForexPairs200Response) GetDataOk() ([]ForexResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetForexPairs200Response) SetData(v []ForexResponseItem) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetForexPairs200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetForexPairs200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetForexPairs200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetForexPairs200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetForexPairs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetForexPairs200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetForexPairs200Response := _GetForexPairs200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetForexPairs200Response)

	if err != nil {
		return err
	}

	*o = GetForexPairs200Response(varGetForexPairs200Response)

	return err
}

type NullableGetForexPairs200Response struct {
	value *GetForexPairs200Response
	isSet bool
}

func (v NullableGetForexPairs200Response) Get() *GetForexPairs200Response {
	return v.value
}

func (v *NullableGetForexPairs200Response) Set(val *GetForexPairs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetForexPairs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetForexPairs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetForexPairs200Response(val *GetForexPairs200Response) *NullableGetForexPairs200Response {
	return &NullableGetForexPairs200Response{value: val, isSet: true}
}

func (v NullableGetForexPairs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetForexPairs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
