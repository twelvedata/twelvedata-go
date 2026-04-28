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

// checks if the GetTimeSeriesWma200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesWma200Response{}

// GetTimeSeriesWma200Response struct for GetTimeSeriesWma200Response
type GetTimeSeriesWma200Response struct {
	Meta GetTimeSeriesWma200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesWma200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesWma200Response GetTimeSeriesWma200Response

// NewGetTimeSeriesWma200Response instantiates a new GetTimeSeriesWma200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesWma200Response(meta GetTimeSeriesWma200ResponseMeta, values []GetTimeSeriesWma200ResponseValuesInner, status string) *GetTimeSeriesWma200Response {
	this := GetTimeSeriesWma200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesWma200ResponseWithDefaults instantiates a new GetTimeSeriesWma200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesWma200ResponseWithDefaults() *GetTimeSeriesWma200Response {
	this := GetTimeSeriesWma200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesWma200Response) GetMeta() GetTimeSeriesWma200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesWma200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWma200Response) GetMetaOk() (*GetTimeSeriesWma200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesWma200Response) SetMeta(v GetTimeSeriesWma200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesWma200Response) GetValues() []GetTimeSeriesWma200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesWma200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWma200Response) GetValuesOk() ([]GetTimeSeriesWma200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesWma200Response) SetValues(v []GetTimeSeriesWma200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesWma200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWma200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesWma200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesWma200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesWma200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesWma200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesWma200Response := _GetTimeSeriesWma200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesWma200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesWma200Response(varGetTimeSeriesWma200Response)

	return err
}

type NullableGetTimeSeriesWma200Response struct {
	value *GetTimeSeriesWma200Response
	isSet bool
}

func (v NullableGetTimeSeriesWma200Response) Get() *GetTimeSeriesWma200Response {
	return v.value
}

func (v *NullableGetTimeSeriesWma200Response) Set(val *GetTimeSeriesWma200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesWma200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesWma200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesWma200Response(val *GetTimeSeriesWma200Response) *NullableGetTimeSeriesWma200Response {
	return &NullableGetTimeSeriesWma200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesWma200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesWma200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
