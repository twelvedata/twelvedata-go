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

// checks if the GetTimeSeriesAdx200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdx200Response{}

// GetTimeSeriesAdx200Response struct for GetTimeSeriesAdx200Response
type GetTimeSeriesAdx200Response struct {
	Meta GetTimeSeriesAdx200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAdx200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAdx200Response GetTimeSeriesAdx200Response

// NewGetTimeSeriesAdx200Response instantiates a new GetTimeSeriesAdx200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdx200Response(meta GetTimeSeriesAdx200ResponseMeta, values []GetTimeSeriesAdx200ResponseValuesInner, status string) *GetTimeSeriesAdx200Response {
	this := GetTimeSeriesAdx200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAdx200ResponseWithDefaults instantiates a new GetTimeSeriesAdx200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdx200ResponseWithDefaults() *GetTimeSeriesAdx200Response {
	this := GetTimeSeriesAdx200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAdx200Response) GetMeta() GetTimeSeriesAdx200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAdx200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdx200Response) GetMetaOk() (*GetTimeSeriesAdx200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAdx200Response) SetMeta(v GetTimeSeriesAdx200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAdx200Response) GetValues() []GetTimeSeriesAdx200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAdx200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdx200Response) GetValuesOk() ([]GetTimeSeriesAdx200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAdx200Response) SetValues(v []GetTimeSeriesAdx200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAdx200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdx200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAdx200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAdx200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdx200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAdx200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAdx200Response := _GetTimeSeriesAdx200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAdx200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdx200Response(varGetTimeSeriesAdx200Response)

	return err
}

type NullableGetTimeSeriesAdx200Response struct {
	value *GetTimeSeriesAdx200Response
	isSet bool
}

func (v NullableGetTimeSeriesAdx200Response) Get() *GetTimeSeriesAdx200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAdx200Response) Set(val *GetTimeSeriesAdx200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdx200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdx200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdx200Response(val *GetTimeSeriesAdx200Response) *NullableGetTimeSeriesAdx200Response {
	return &NullableGetTimeSeriesAdx200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdx200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdx200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
