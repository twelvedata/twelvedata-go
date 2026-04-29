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

// checks if the GetTimeSeriesMedPrice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMedPrice200Response{}

// GetTimeSeriesMedPrice200Response struct for GetTimeSeriesMedPrice200Response
type GetTimeSeriesMedPrice200Response struct {
	Meta GetTimeSeriesMedPrice200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesMedPrice200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesMedPrice200Response GetTimeSeriesMedPrice200Response

// NewGetTimeSeriesMedPrice200Response instantiates a new GetTimeSeriesMedPrice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMedPrice200Response(meta GetTimeSeriesMedPrice200ResponseMeta, values []GetTimeSeriesMedPrice200ResponseValuesInner, status string) *GetTimeSeriesMedPrice200Response {
	this := GetTimeSeriesMedPrice200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesMedPrice200ResponseWithDefaults instantiates a new GetTimeSeriesMedPrice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMedPrice200ResponseWithDefaults() *GetTimeSeriesMedPrice200Response {
	this := GetTimeSeriesMedPrice200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesMedPrice200Response) GetMeta() GetTimeSeriesMedPrice200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesMedPrice200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMedPrice200Response) GetMetaOk() (*GetTimeSeriesMedPrice200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesMedPrice200Response) SetMeta(v GetTimeSeriesMedPrice200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesMedPrice200Response) GetValues() []GetTimeSeriesMedPrice200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesMedPrice200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMedPrice200Response) GetValuesOk() ([]GetTimeSeriesMedPrice200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesMedPrice200Response) SetValues(v []GetTimeSeriesMedPrice200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesMedPrice200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMedPrice200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesMedPrice200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesMedPrice200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMedPrice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesMedPrice200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMedPrice200Response := _GetTimeSeriesMedPrice200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMedPrice200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMedPrice200Response(varGetTimeSeriesMedPrice200Response)

	return err
}

type NullableGetTimeSeriesMedPrice200Response struct {
	value *GetTimeSeriesMedPrice200Response
	isSet bool
}

func (v NullableGetTimeSeriesMedPrice200Response) Get() *GetTimeSeriesMedPrice200Response {
	return v.value
}

func (v *NullableGetTimeSeriesMedPrice200Response) Set(val *GetTimeSeriesMedPrice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMedPrice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMedPrice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMedPrice200Response(val *GetTimeSeriesMedPrice200Response) *NullableGetTimeSeriesMedPrice200Response {
	return &NullableGetTimeSeriesMedPrice200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMedPrice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMedPrice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
