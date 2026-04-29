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

// checks if the GetTimeSeriesMinusDM200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinusDM200Response{}

// GetTimeSeriesMinusDM200Response struct for GetTimeSeriesMinusDM200Response
type GetTimeSeriesMinusDM200Response struct {
	Meta GetTimeSeriesMinusDM200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMinusDM200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMinusDM200Response GetTimeSeriesMinusDM200Response

// NewGetTimeSeriesMinusDM200Response instantiates a new GetTimeSeriesMinusDM200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinusDM200Response(meta GetTimeSeriesMinusDM200ResponseMeta, values []GetTimeSeriesMinusDM200ResponseValuesInner, status string) *GetTimeSeriesMinusDM200Response {
	this := GetTimeSeriesMinusDM200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMinusDM200ResponseWithDefaults instantiates a new GetTimeSeriesMinusDM200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinusDM200ResponseWithDefaults() *GetTimeSeriesMinusDM200Response {
	this := GetTimeSeriesMinusDM200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMinusDM200Response) GetMeta() GetTimeSeriesMinusDM200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMinusDM200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDM200Response) GetMetaOk() (*GetTimeSeriesMinusDM200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMinusDM200Response) SetMeta(v GetTimeSeriesMinusDM200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMinusDM200Response) GetValues() []GetTimeSeriesMinusDM200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMinusDM200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDM200Response) GetValuesOk() ([]GetTimeSeriesMinusDM200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMinusDM200Response) SetValues(v []GetTimeSeriesMinusDM200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMinusDM200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDM200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMinusDM200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMinusDM200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinusDM200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMinusDM200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMinusDM200Response := _GetTimeSeriesMinusDM200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinusDM200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinusDM200Response(varGetTimeSeriesMinusDM200Response)

	return err
}

type NullableGetTimeSeriesMinusDM200Response struct {
	value *GetTimeSeriesMinusDM200Response
	isSet bool
}

func (v NullableGetTimeSeriesMinusDM200Response) Get() *GetTimeSeriesMinusDM200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMinusDM200Response) Set(val *GetTimeSeriesMinusDM200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinusDM200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinusDM200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinusDM200Response(val *GetTimeSeriesMinusDM200Response) *NullableGetTimeSeriesMinusDM200Response {
	return &NullableGetTimeSeriesMinusDM200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinusDM200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinusDM200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
