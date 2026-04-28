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

// checks if the GetTimeSeriesStochRsi200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStochRsi200Response{}

// GetTimeSeriesStochRsi200Response struct for GetTimeSeriesStochRsi200Response
type GetTimeSeriesStochRsi200Response struct {
	Meta GetTimeSeriesStochRsi200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesStochRsi200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesStochRsi200Response GetTimeSeriesStochRsi200Response

// NewGetTimeSeriesStochRsi200Response instantiates a new GetTimeSeriesStochRsi200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStochRsi200Response(meta GetTimeSeriesStochRsi200ResponseMeta, values []GetTimeSeriesStochRsi200ResponseValuesInner, status string) *GetTimeSeriesStochRsi200Response {
	this := GetTimeSeriesStochRsi200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesStochRsi200ResponseWithDefaults instantiates a new GetTimeSeriesStochRsi200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStochRsi200ResponseWithDefaults() *GetTimeSeriesStochRsi200Response {
	this := GetTimeSeriesStochRsi200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesStochRsi200Response) GetMeta() GetTimeSeriesStochRsi200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesStochRsi200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200Response) GetMetaOk() (*GetTimeSeriesStochRsi200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesStochRsi200Response) SetMeta(v GetTimeSeriesStochRsi200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesStochRsi200Response) GetValues() []GetTimeSeriesStochRsi200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesStochRsi200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200Response) GetValuesOk() ([]GetTimeSeriesStochRsi200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesStochRsi200Response) SetValues(v []GetTimeSeriesStochRsi200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesStochRsi200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesStochRsi200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesStochRsi200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStochRsi200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesStochRsi200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesStochRsi200Response := _GetTimeSeriesStochRsi200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesStochRsi200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStochRsi200Response(varGetTimeSeriesStochRsi200Response)

	return err
}

type NullableGetTimeSeriesStochRsi200Response struct {
	value *GetTimeSeriesStochRsi200Response
	isSet bool
}

func (v NullableGetTimeSeriesStochRsi200Response) Get() *GetTimeSeriesStochRsi200Response {
	return v.value
}

func (v *NullableGetTimeSeriesStochRsi200Response) Set(val *GetTimeSeriesStochRsi200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStochRsi200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStochRsi200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStochRsi200Response(val *GetTimeSeriesStochRsi200Response) *NullableGetTimeSeriesStochRsi200Response {
	return &NullableGetTimeSeriesStochRsi200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStochRsi200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStochRsi200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
