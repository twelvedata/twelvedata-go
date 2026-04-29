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

// checks if the GetTimeSeriesMaxIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMaxIndex200Response{}

// GetTimeSeriesMaxIndex200Response struct for GetTimeSeriesMaxIndex200Response
type GetTimeSeriesMaxIndex200Response struct {
	Meta GetTimeSeriesMaxIndex200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMaxIndex200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMaxIndex200Response GetTimeSeriesMaxIndex200Response

// NewGetTimeSeriesMaxIndex200Response instantiates a new GetTimeSeriesMaxIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMaxIndex200Response(meta GetTimeSeriesMaxIndex200ResponseMeta, values []GetTimeSeriesMaxIndex200ResponseValuesInner, status string) *GetTimeSeriesMaxIndex200Response {
	this := GetTimeSeriesMaxIndex200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMaxIndex200ResponseWithDefaults instantiates a new GetTimeSeriesMaxIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMaxIndex200ResponseWithDefaults() *GetTimeSeriesMaxIndex200Response {
	this := GetTimeSeriesMaxIndex200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMaxIndex200Response) GetMeta() GetTimeSeriesMaxIndex200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMaxIndex200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMaxIndex200Response) GetMetaOk() (*GetTimeSeriesMaxIndex200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMaxIndex200Response) SetMeta(v GetTimeSeriesMaxIndex200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMaxIndex200Response) GetValues() []GetTimeSeriesMaxIndex200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMaxIndex200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMaxIndex200Response) GetValuesOk() ([]GetTimeSeriesMaxIndex200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMaxIndex200Response) SetValues(v []GetTimeSeriesMaxIndex200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMaxIndex200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMaxIndex200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMaxIndex200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMaxIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMaxIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMaxIndex200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMaxIndex200Response := _GetTimeSeriesMaxIndex200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMaxIndex200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMaxIndex200Response(varGetTimeSeriesMaxIndex200Response)

	return err
}

type NullableGetTimeSeriesMaxIndex200Response struct {
	value *GetTimeSeriesMaxIndex200Response
	isSet bool
}

func (v NullableGetTimeSeriesMaxIndex200Response) Get() *GetTimeSeriesMaxIndex200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMaxIndex200Response) Set(val *GetTimeSeriesMaxIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMaxIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMaxIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMaxIndex200Response(val *GetTimeSeriesMaxIndex200Response) *NullableGetTimeSeriesMaxIndex200Response {
	return &NullableGetTimeSeriesMaxIndex200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMaxIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMaxIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
