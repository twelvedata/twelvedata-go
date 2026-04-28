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

// checks if the GetTimeSeriesPercentB200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPercentB200Response{}

// GetTimeSeriesPercentB200Response struct for GetTimeSeriesPercentB200Response
type GetTimeSeriesPercentB200Response struct {
	Meta GetTimeSeriesPercentB200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesPercentB200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesPercentB200Response GetTimeSeriesPercentB200Response

// NewGetTimeSeriesPercentB200Response instantiates a new GetTimeSeriesPercentB200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPercentB200Response(meta GetTimeSeriesPercentB200ResponseMeta, values []GetTimeSeriesPercentB200ResponseValuesInner, status string) *GetTimeSeriesPercentB200Response {
	this := GetTimeSeriesPercentB200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesPercentB200ResponseWithDefaults instantiates a new GetTimeSeriesPercentB200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPercentB200ResponseWithDefaults() *GetTimeSeriesPercentB200Response {
	this := GetTimeSeriesPercentB200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesPercentB200Response) GetMeta() GetTimeSeriesPercentB200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesPercentB200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200Response) GetMetaOk() (*GetTimeSeriesPercentB200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesPercentB200Response) SetMeta(v GetTimeSeriesPercentB200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesPercentB200Response) GetValues() []GetTimeSeriesPercentB200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesPercentB200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200Response) GetValuesOk() ([]GetTimeSeriesPercentB200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesPercentB200Response) SetValues(v []GetTimeSeriesPercentB200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesPercentB200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesPercentB200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesPercentB200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPercentB200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesPercentB200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesPercentB200Response := _GetTimeSeriesPercentB200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesPercentB200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPercentB200Response(varGetTimeSeriesPercentB200Response)

	return err
}

type NullableGetTimeSeriesPercentB200Response struct {
	value *GetTimeSeriesPercentB200Response
	isSet bool
}

func (v NullableGetTimeSeriesPercentB200Response) Get() *GetTimeSeriesPercentB200Response {
	return v.value
}

func (v *NullableGetTimeSeriesPercentB200Response) Set(val *GetTimeSeriesPercentB200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPercentB200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPercentB200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPercentB200Response(val *GetTimeSeriesPercentB200Response) *NullableGetTimeSeriesPercentB200Response {
	return &NullableGetTimeSeriesPercentB200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPercentB200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPercentB200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
