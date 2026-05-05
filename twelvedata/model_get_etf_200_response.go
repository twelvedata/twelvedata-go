// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetEtf200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEtf200Response{}

// GetEtf200Response struct for GetEtf200Response
type GetEtf200Response struct {
	// Count
	Count int64 `json:"count"`
	// List of ETFs
	Data []EtfResponseItem `json:"data"`
	// Response status
	Status string `json:"status"`
}

type _GetEtf200Response GetEtf200Response

// NewGetEtf200Response instantiates a new GetEtf200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEtf200Response(count int64, data []EtfResponseItem, status string) *GetEtf200Response {
	this := GetEtf200Response{}
	this.Count = count
	this.Data = data
	this.Status = status
	return &this
}

// NewGetEtf200ResponseWithDefaults instantiates a new GetEtf200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEtf200ResponseWithDefaults() *GetEtf200Response {
	this := GetEtf200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *GetEtf200Response) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetEtf200Response) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetEtf200Response) SetCount(v int64) {
	o.Count = v
}

// GetData returns the Data field value
func (o *GetEtf200Response) GetData() []EtfResponseItem {
	if o == nil {
		var ret []EtfResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEtf200Response) GetDataOk() ([]EtfResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetEtf200Response) SetData(v []EtfResponseItem) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetEtf200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetEtf200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetEtf200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetEtf200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEtf200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetEtf200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetEtf200Response := _GetEtf200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEtf200Response)

	if err != nil {
		return err
	}

	*o = GetEtf200Response(varGetEtf200Response)

	return err
}

type NullableGetEtf200Response struct {
	value *GetEtf200Response
	isSet bool
}

func (v NullableGetEtf200Response) Get() *GetEtf200Response {
	return v.value
}

func (v *NullableGetEtf200Response) Set(val *GetEtf200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEtf200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEtf200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEtf200Response(val *GetEtf200Response) *NullableGetEtf200Response {
	return &NullableGetEtf200Response{value: val, isSet: true}
}

func (v NullableGetEtf200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEtf200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
