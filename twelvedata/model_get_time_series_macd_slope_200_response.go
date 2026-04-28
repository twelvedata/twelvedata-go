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

// checks if the GetTimeSeriesMacdSlope200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacdSlope200Response{}

// GetTimeSeriesMacdSlope200Response struct for GetTimeSeriesMacdSlope200Response
type GetTimeSeriesMacdSlope200Response struct {
	Meta GetTimeSeriesMacdSlope200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMacdSlope200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMacdSlope200Response GetTimeSeriesMacdSlope200Response

// NewGetTimeSeriesMacdSlope200Response instantiates a new GetTimeSeriesMacdSlope200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacdSlope200Response(meta GetTimeSeriesMacdSlope200ResponseMeta, values []GetTimeSeriesMacdSlope200ResponseValuesInner, status string) *GetTimeSeriesMacdSlope200Response {
	this := GetTimeSeriesMacdSlope200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMacdSlope200ResponseWithDefaults instantiates a new GetTimeSeriesMacdSlope200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacdSlope200ResponseWithDefaults() *GetTimeSeriesMacdSlope200Response {
	this := GetTimeSeriesMacdSlope200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMacdSlope200Response) GetMeta() GetTimeSeriesMacdSlope200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMacdSlope200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200Response) GetMetaOk() (*GetTimeSeriesMacdSlope200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMacdSlope200Response) SetMeta(v GetTimeSeriesMacdSlope200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMacdSlope200Response) GetValues() []GetTimeSeriesMacdSlope200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMacdSlope200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200Response) GetValuesOk() ([]GetTimeSeriesMacdSlope200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMacdSlope200Response) SetValues(v []GetTimeSeriesMacdSlope200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMacdSlope200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMacdSlope200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMacdSlope200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacdSlope200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMacdSlope200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMacdSlope200Response := _GetTimeSeriesMacdSlope200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMacdSlope200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacdSlope200Response(varGetTimeSeriesMacdSlope200Response)

	return err
}

type NullableGetTimeSeriesMacdSlope200Response struct {
	value *GetTimeSeriesMacdSlope200Response
	isSet bool
}

func (v NullableGetTimeSeriesMacdSlope200Response) Get() *GetTimeSeriesMacdSlope200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMacdSlope200Response) Set(val *GetTimeSeriesMacdSlope200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacdSlope200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacdSlope200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacdSlope200Response(val *GetTimeSeriesMacdSlope200Response) *NullableGetTimeSeriesMacdSlope200Response {
	return &NullableGetTimeSeriesMacdSlope200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacdSlope200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacdSlope200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
