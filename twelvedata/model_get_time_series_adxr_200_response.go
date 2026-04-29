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

// checks if the GetTimeSeriesAdxr200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdxr200Response{}

// GetTimeSeriesAdxr200Response struct for GetTimeSeriesAdxr200Response
type GetTimeSeriesAdxr200Response struct {
	Meta GetTimeSeriesAdxr200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAdxr200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAdxr200Response GetTimeSeriesAdxr200Response

// NewGetTimeSeriesAdxr200Response instantiates a new GetTimeSeriesAdxr200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdxr200Response(meta GetTimeSeriesAdxr200ResponseMeta, values []GetTimeSeriesAdxr200ResponseValuesInner, status string) *GetTimeSeriesAdxr200Response {
	this := GetTimeSeriesAdxr200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAdxr200ResponseWithDefaults instantiates a new GetTimeSeriesAdxr200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdxr200ResponseWithDefaults() *GetTimeSeriesAdxr200Response {
	this := GetTimeSeriesAdxr200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAdxr200Response) GetMeta() GetTimeSeriesAdxr200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAdxr200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdxr200Response) GetMetaOk() (*GetTimeSeriesAdxr200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAdxr200Response) SetMeta(v GetTimeSeriesAdxr200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAdxr200Response) GetValues() []GetTimeSeriesAdxr200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAdxr200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdxr200Response) GetValuesOk() ([]GetTimeSeriesAdxr200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAdxr200Response) SetValues(v []GetTimeSeriesAdxr200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAdxr200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdxr200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAdxr200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAdxr200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdxr200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAdxr200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAdxr200Response := _GetTimeSeriesAdxr200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAdxr200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdxr200Response(varGetTimeSeriesAdxr200Response)

	return err
}

type NullableGetTimeSeriesAdxr200Response struct {
	value *GetTimeSeriesAdxr200Response
	isSet bool
}

func (v NullableGetTimeSeriesAdxr200Response) Get() *GetTimeSeriesAdxr200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAdxr200Response) Set(val *GetTimeSeriesAdxr200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdxr200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdxr200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdxr200Response(val *GetTimeSeriesAdxr200Response) *NullableGetTimeSeriesAdxr200Response {
	return &NullableGetTimeSeriesAdxr200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdxr200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdxr200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
