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

// checks if the GetTimeSeriesSar200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSar200Response{}

// GetTimeSeriesSar200Response struct for GetTimeSeriesSar200Response
type GetTimeSeriesSar200Response struct {
	Meta GetTimeSeriesSar200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesSar200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesSar200Response GetTimeSeriesSar200Response

// NewGetTimeSeriesSar200Response instantiates a new GetTimeSeriesSar200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSar200Response(meta GetTimeSeriesSar200ResponseMeta, values []GetTimeSeriesSar200ResponseValuesInner, status string) *GetTimeSeriesSar200Response {
	this := GetTimeSeriesSar200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesSar200ResponseWithDefaults instantiates a new GetTimeSeriesSar200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSar200ResponseWithDefaults() *GetTimeSeriesSar200Response {
	this := GetTimeSeriesSar200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesSar200Response) GetMeta() GetTimeSeriesSar200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesSar200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200Response) GetMetaOk() (*GetTimeSeriesSar200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesSar200Response) SetMeta(v GetTimeSeriesSar200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesSar200Response) GetValues() []GetTimeSeriesSar200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesSar200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200Response) GetValuesOk() ([]GetTimeSeriesSar200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesSar200Response) SetValues(v []GetTimeSeriesSar200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesSar200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesSar200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesSar200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSar200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesSar200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesSar200Response := _GetTimeSeriesSar200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesSar200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSar200Response(varGetTimeSeriesSar200Response)

	return err
}

type NullableGetTimeSeriesSar200Response struct {
	value *GetTimeSeriesSar200Response
	isSet bool
}

func (v NullableGetTimeSeriesSar200Response) Get() *GetTimeSeriesSar200Response {
	return v.value
}

func (v *NullableGetTimeSeriesSar200Response) Set(val *GetTimeSeriesSar200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSar200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSar200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSar200Response(val *GetTimeSeriesSar200Response) *NullableGetTimeSeriesSar200Response {
	return &NullableGetTimeSeriesSar200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSar200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSar200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
