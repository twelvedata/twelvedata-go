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

// checks if the GetTimeSeriesMinMaxIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinMaxIndex200Response{}

// GetTimeSeriesMinMaxIndex200Response struct for GetTimeSeriesMinMaxIndex200Response
type GetTimeSeriesMinMaxIndex200Response struct {
	Meta GetTimeSeriesMinMaxIndex200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMinMaxIndex200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMinMaxIndex200Response GetTimeSeriesMinMaxIndex200Response

// NewGetTimeSeriesMinMaxIndex200Response instantiates a new GetTimeSeriesMinMaxIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinMaxIndex200Response(meta GetTimeSeriesMinMaxIndex200ResponseMeta, values []GetTimeSeriesMinMaxIndex200ResponseValuesInner, status string) *GetTimeSeriesMinMaxIndex200Response {
	this := GetTimeSeriesMinMaxIndex200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMinMaxIndex200ResponseWithDefaults instantiates a new GetTimeSeriesMinMaxIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinMaxIndex200ResponseWithDefaults() *GetTimeSeriesMinMaxIndex200Response {
	this := GetTimeSeriesMinMaxIndex200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMinMaxIndex200Response) GetMeta() GetTimeSeriesMinMaxIndex200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMinMaxIndex200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200Response) GetMetaOk() (*GetTimeSeriesMinMaxIndex200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMinMaxIndex200Response) SetMeta(v GetTimeSeriesMinMaxIndex200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMinMaxIndex200Response) GetValues() []GetTimeSeriesMinMaxIndex200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMinMaxIndex200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200Response) GetValuesOk() ([]GetTimeSeriesMinMaxIndex200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMinMaxIndex200Response) SetValues(v []GetTimeSeriesMinMaxIndex200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMinMaxIndex200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMinMaxIndex200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMinMaxIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinMaxIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMinMaxIndex200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMinMaxIndex200Response := _GetTimeSeriesMinMaxIndex200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMinMaxIndex200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinMaxIndex200Response(varGetTimeSeriesMinMaxIndex200Response)

	return err
}

type NullableGetTimeSeriesMinMaxIndex200Response struct {
	value *GetTimeSeriesMinMaxIndex200Response
	isSet bool
}

func (v NullableGetTimeSeriesMinMaxIndex200Response) Get() *GetTimeSeriesMinMaxIndex200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMinMaxIndex200Response) Set(val *GetTimeSeriesMinMaxIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinMaxIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinMaxIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinMaxIndex200Response(val *GetTimeSeriesMinMaxIndex200Response) *NullableGetTimeSeriesMinMaxIndex200Response {
	return &NullableGetTimeSeriesMinMaxIndex200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinMaxIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinMaxIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
