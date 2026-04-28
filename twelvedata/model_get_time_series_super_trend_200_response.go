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

// checks if the GetTimeSeriesSuperTrend200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSuperTrend200Response{}

// GetTimeSeriesSuperTrend200Response struct for GetTimeSeriesSuperTrend200Response
type GetTimeSeriesSuperTrend200Response struct {
	Meta GetTimeSeriesSuperTrend200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesSuperTrend200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesSuperTrend200Response GetTimeSeriesSuperTrend200Response

// NewGetTimeSeriesSuperTrend200Response instantiates a new GetTimeSeriesSuperTrend200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSuperTrend200Response(meta GetTimeSeriesSuperTrend200ResponseMeta, values []GetTimeSeriesSuperTrend200ResponseValuesInner, status string) *GetTimeSeriesSuperTrend200Response {
	this := GetTimeSeriesSuperTrend200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesSuperTrend200ResponseWithDefaults instantiates a new GetTimeSeriesSuperTrend200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSuperTrend200ResponseWithDefaults() *GetTimeSeriesSuperTrend200Response {
	this := GetTimeSeriesSuperTrend200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesSuperTrend200Response) GetMeta() GetTimeSeriesSuperTrend200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesSuperTrend200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200Response) GetMetaOk() (*GetTimeSeriesSuperTrend200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesSuperTrend200Response) SetMeta(v GetTimeSeriesSuperTrend200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesSuperTrend200Response) GetValues() []GetTimeSeriesSuperTrend200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesSuperTrend200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200Response) GetValuesOk() ([]GetTimeSeriesSuperTrend200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesSuperTrend200Response) SetValues(v []GetTimeSeriesSuperTrend200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesSuperTrend200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesSuperTrend200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesSuperTrend200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSuperTrend200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesSuperTrend200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesSuperTrend200Response := _GetTimeSeriesSuperTrend200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesSuperTrend200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSuperTrend200Response(varGetTimeSeriesSuperTrend200Response)

	return err
}

type NullableGetTimeSeriesSuperTrend200Response struct {
	value *GetTimeSeriesSuperTrend200Response
	isSet bool
}

func (v NullableGetTimeSeriesSuperTrend200Response) Get() *GetTimeSeriesSuperTrend200Response {
	return v.value
}

func (v *NullableGetTimeSeriesSuperTrend200Response) Set(val *GetTimeSeriesSuperTrend200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSuperTrend200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSuperTrend200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSuperTrend200Response(val *GetTimeSeriesSuperTrend200Response) *NullableGetTimeSeriesSuperTrend200Response {
	return &NullableGetTimeSeriesSuperTrend200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSuperTrend200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSuperTrend200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
