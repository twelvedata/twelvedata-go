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

// checks if the GetIntervals200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIntervals200Response{}

// GetIntervals200Response struct for GetIntervals200Response
type GetIntervals200Response struct {
	// List of available intervals
	Data []string `json:"data"`
	// Status of the response
	Status string `json:"status"`
}

type _GetIntervals200Response GetIntervals200Response

// NewGetIntervals200Response instantiates a new GetIntervals200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIntervals200Response(data []string, status string) *GetIntervals200Response {
	this := GetIntervals200Response{}
	this.Data = data
	this.Status = status
	return &this
}

// NewGetIntervals200ResponseWithDefaults instantiates a new GetIntervals200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIntervals200ResponseWithDefaults() *GetIntervals200Response {
	this := GetIntervals200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetIntervals200Response) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetIntervals200Response) GetDataOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetIntervals200Response) SetData(v []string) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetIntervals200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetIntervals200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetIntervals200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetIntervals200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIntervals200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetIntervals200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetIntervals200Response := _GetIntervals200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetIntervals200Response)

	if err != nil {
		return err
	}

	*o = GetIntervals200Response(varGetIntervals200Response)

	return err
}

type NullableGetIntervals200Response struct {
	value *GetIntervals200Response
	isSet bool
}

func (v NullableGetIntervals200Response) Get() *GetIntervals200Response {
	return v.value
}

func (v *NullableGetIntervals200Response) Set(val *GetIntervals200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIntervals200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIntervals200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIntervals200Response(val *GetIntervals200Response) *NullableGetIntervals200Response {
	return &NullableGetIntervals200Response{value: val, isSet: true}
}

func (v NullableGetIntervals200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIntervals200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
