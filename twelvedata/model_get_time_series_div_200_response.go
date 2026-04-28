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

// checks if the GetTimeSeriesDiv200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesDiv200Response{}

// GetTimeSeriesDiv200Response struct for GetTimeSeriesDiv200Response
type GetTimeSeriesDiv200Response struct {
	Meta GetTimeSeriesDiv200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesDiv200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesDiv200Response GetTimeSeriesDiv200Response

// NewGetTimeSeriesDiv200Response instantiates a new GetTimeSeriesDiv200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesDiv200Response(meta GetTimeSeriesDiv200ResponseMeta, values []GetTimeSeriesDiv200ResponseValuesInner, status string) *GetTimeSeriesDiv200Response {
	this := GetTimeSeriesDiv200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesDiv200ResponseWithDefaults instantiates a new GetTimeSeriesDiv200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesDiv200ResponseWithDefaults() *GetTimeSeriesDiv200Response {
	this := GetTimeSeriesDiv200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesDiv200Response) GetMeta() GetTimeSeriesDiv200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesDiv200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDiv200Response) GetMetaOk() (*GetTimeSeriesDiv200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesDiv200Response) SetMeta(v GetTimeSeriesDiv200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesDiv200Response) GetValues() []GetTimeSeriesDiv200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesDiv200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDiv200Response) GetValuesOk() ([]GetTimeSeriesDiv200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesDiv200Response) SetValues(v []GetTimeSeriesDiv200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesDiv200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDiv200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesDiv200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesDiv200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesDiv200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesDiv200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesDiv200Response := _GetTimeSeriesDiv200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesDiv200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesDiv200Response(varGetTimeSeriesDiv200Response)

	return err
}

type NullableGetTimeSeriesDiv200Response struct {
	value *GetTimeSeriesDiv200Response
	isSet bool
}

func (v NullableGetTimeSeriesDiv200Response) Get() *GetTimeSeriesDiv200Response {
	return v.value
}

func (v *NullableGetTimeSeriesDiv200Response) Set(val *GetTimeSeriesDiv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesDiv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesDiv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesDiv200Response(val *GetTimeSeriesDiv200Response) *NullableGetTimeSeriesDiv200Response {
	return &NullableGetTimeSeriesDiv200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesDiv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesDiv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
