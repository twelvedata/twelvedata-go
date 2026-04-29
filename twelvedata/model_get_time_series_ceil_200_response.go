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

// checks if the GetTimeSeriesCeil200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCeil200Response{}

// GetTimeSeriesCeil200Response struct for GetTimeSeriesCeil200Response
type GetTimeSeriesCeil200Response struct {
	Meta GetTimeSeriesCeil200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesCeil200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesCeil200Response GetTimeSeriesCeil200Response

// NewGetTimeSeriesCeil200Response instantiates a new GetTimeSeriesCeil200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCeil200Response(meta GetTimeSeriesCeil200ResponseMeta, values []GetTimeSeriesCeil200ResponseValuesInner, status string) *GetTimeSeriesCeil200Response {
	this := GetTimeSeriesCeil200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesCeil200ResponseWithDefaults instantiates a new GetTimeSeriesCeil200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCeil200ResponseWithDefaults() *GetTimeSeriesCeil200Response {
	this := GetTimeSeriesCeil200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesCeil200Response) GetMeta() GetTimeSeriesCeil200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesCeil200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCeil200Response) GetMetaOk() (*GetTimeSeriesCeil200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesCeil200Response) SetMeta(v GetTimeSeriesCeil200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesCeil200Response) GetValues() []GetTimeSeriesCeil200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesCeil200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCeil200Response) GetValuesOk() ([]GetTimeSeriesCeil200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesCeil200Response) SetValues(v []GetTimeSeriesCeil200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesCeil200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCeil200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesCeil200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesCeil200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCeil200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesCeil200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"values",
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

	varGetTimeSeriesCeil200Response := _GetTimeSeriesCeil200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCeil200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCeil200Response(varGetTimeSeriesCeil200Response)

	return err
}

type NullableGetTimeSeriesCeil200Response struct {
	value *GetTimeSeriesCeil200Response
	isSet bool
}

func (v NullableGetTimeSeriesCeil200Response) Get() *GetTimeSeriesCeil200Response {
	return v.value
}

func (v *NullableGetTimeSeriesCeil200Response) Set(val *GetTimeSeriesCeil200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCeil200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCeil200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCeil200Response(val *GetTimeSeriesCeil200Response) *NullableGetTimeSeriesCeil200Response {
	return &NullableGetTimeSeriesCeil200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCeil200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCeil200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
