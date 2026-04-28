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

// checks if the GetTimeSeriesAd200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAd200Response{}

// GetTimeSeriesAd200Response struct for GetTimeSeriesAd200Response
type GetTimeSeriesAd200Response struct {
	Meta   GetTimeSeriesAd200ResponseMeta          `json:"meta"`
	Values []GetTimeSeriesAd200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAd200Response GetTimeSeriesAd200Response

// NewGetTimeSeriesAd200Response instantiates a new GetTimeSeriesAd200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAd200Response(meta GetTimeSeriesAd200ResponseMeta, values []GetTimeSeriesAd200ResponseValuesInner, status string) *GetTimeSeriesAd200Response {
	this := GetTimeSeriesAd200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAd200ResponseWithDefaults instantiates a new GetTimeSeriesAd200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAd200ResponseWithDefaults() *GetTimeSeriesAd200Response {
	this := GetTimeSeriesAd200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAd200Response) GetMeta() GetTimeSeriesAd200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAd200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAd200Response) GetMetaOk() (*GetTimeSeriesAd200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAd200Response) SetMeta(v GetTimeSeriesAd200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAd200Response) GetValues() []GetTimeSeriesAd200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAd200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAd200Response) GetValuesOk() ([]GetTimeSeriesAd200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAd200Response) SetValues(v []GetTimeSeriesAd200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAd200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAd200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAd200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAd200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAd200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAd200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAd200Response := _GetTimeSeriesAd200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAd200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAd200Response(varGetTimeSeriesAd200Response)

	return err
}

type NullableGetTimeSeriesAd200Response struct {
	value *GetTimeSeriesAd200Response
	isSet bool
}

func (v NullableGetTimeSeriesAd200Response) Get() *GetTimeSeriesAd200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAd200Response) Set(val *GetTimeSeriesAd200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAd200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAd200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAd200Response(val *GetTimeSeriesAd200Response) *NullableGetTimeSeriesAd200Response {
	return &NullableGetTimeSeriesAd200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAd200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAd200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
