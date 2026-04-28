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

// checks if the GetTimeSeriesMcGinleyDynamic200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMcGinleyDynamic200Response{}

// GetTimeSeriesMcGinleyDynamic200Response struct for GetTimeSeriesMcGinleyDynamic200Response
type GetTimeSeriesMcGinleyDynamic200Response struct {
	Meta GetTimeSeriesMcGinleyDynamic200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMcGinleyDynamic200Response GetTimeSeriesMcGinleyDynamic200Response

// NewGetTimeSeriesMcGinleyDynamic200Response instantiates a new GetTimeSeriesMcGinleyDynamic200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMcGinleyDynamic200Response(meta GetTimeSeriesMcGinleyDynamic200ResponseMeta, values []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner, status string) *GetTimeSeriesMcGinleyDynamic200Response {
	this := GetTimeSeriesMcGinleyDynamic200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMcGinleyDynamic200ResponseWithDefaults instantiates a new GetTimeSeriesMcGinleyDynamic200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMcGinleyDynamic200ResponseWithDefaults() *GetTimeSeriesMcGinleyDynamic200Response {
	this := GetTimeSeriesMcGinleyDynamic200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMcGinleyDynamic200Response) GetMeta() GetTimeSeriesMcGinleyDynamic200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMcGinleyDynamic200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMcGinleyDynamic200Response) GetMetaOk() (*GetTimeSeriesMcGinleyDynamic200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMcGinleyDynamic200Response) SetMeta(v GetTimeSeriesMcGinleyDynamic200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMcGinleyDynamic200Response) GetValues() []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMcGinleyDynamic200Response) GetValuesOk() ([]GetTimeSeriesMcGinleyDynamic200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMcGinleyDynamic200Response) SetValues(v []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMcGinleyDynamic200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMcGinleyDynamic200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMcGinleyDynamic200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMcGinleyDynamic200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMcGinleyDynamic200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMcGinleyDynamic200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMcGinleyDynamic200Response := _GetTimeSeriesMcGinleyDynamic200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMcGinleyDynamic200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMcGinleyDynamic200Response(varGetTimeSeriesMcGinleyDynamic200Response)

	return err
}

type NullableGetTimeSeriesMcGinleyDynamic200Response struct {
	value *GetTimeSeriesMcGinleyDynamic200Response
	isSet bool
}

func (v NullableGetTimeSeriesMcGinleyDynamic200Response) Get() *GetTimeSeriesMcGinleyDynamic200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200Response) Set(val *GetTimeSeriesMcGinleyDynamic200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMcGinleyDynamic200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMcGinleyDynamic200Response(val *GetTimeSeriesMcGinleyDynamic200Response) *NullableGetTimeSeriesMcGinleyDynamic200Response {
	return &NullableGetTimeSeriesMcGinleyDynamic200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMcGinleyDynamic200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
