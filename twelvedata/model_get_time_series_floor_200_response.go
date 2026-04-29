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

// checks if the GetTimeSeriesFloor200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesFloor200Response{}

// GetTimeSeriesFloor200Response struct for GetTimeSeriesFloor200Response
type GetTimeSeriesFloor200Response struct {
	Meta GetTimeSeriesFloor200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesFloor200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesFloor200Response GetTimeSeriesFloor200Response

// NewGetTimeSeriesFloor200Response instantiates a new GetTimeSeriesFloor200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesFloor200Response(meta GetTimeSeriesFloor200ResponseMeta, values []GetTimeSeriesFloor200ResponseValuesInner, status string) *GetTimeSeriesFloor200Response {
	this := GetTimeSeriesFloor200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesFloor200ResponseWithDefaults instantiates a new GetTimeSeriesFloor200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesFloor200ResponseWithDefaults() *GetTimeSeriesFloor200Response {
	this := GetTimeSeriesFloor200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesFloor200Response) GetMeta() GetTimeSeriesFloor200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesFloor200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesFloor200Response) GetMetaOk() (*GetTimeSeriesFloor200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesFloor200Response) SetMeta(v GetTimeSeriesFloor200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesFloor200Response) GetValues() []GetTimeSeriesFloor200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesFloor200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesFloor200Response) GetValuesOk() ([]GetTimeSeriesFloor200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesFloor200Response) SetValues(v []GetTimeSeriesFloor200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesFloor200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesFloor200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesFloor200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesFloor200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesFloor200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesFloor200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesFloor200Response := _GetTimeSeriesFloor200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesFloor200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesFloor200Response(varGetTimeSeriesFloor200Response)

	return err
}

type NullableGetTimeSeriesFloor200Response struct {
	value *GetTimeSeriesFloor200Response
	isSet bool
}

func (v NullableGetTimeSeriesFloor200Response) Get() *GetTimeSeriesFloor200Response {
	return v.value
}

func (v *NullableGetTimeSeriesFloor200Response) Set(val *GetTimeSeriesFloor200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesFloor200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesFloor200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesFloor200Response(val *GetTimeSeriesFloor200Response) *NullableGetTimeSeriesFloor200Response {
	return &NullableGetTimeSeriesFloor200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesFloor200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesFloor200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
