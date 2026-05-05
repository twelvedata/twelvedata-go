// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesCoppock200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCoppock200Response{}

// GetTimeSeriesCoppock200Response struct for GetTimeSeriesCoppock200Response
type GetTimeSeriesCoppock200Response struct {
	Meta GetTimeSeriesCoppock200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesCoppock200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesCoppock200Response GetTimeSeriesCoppock200Response

// NewGetTimeSeriesCoppock200Response instantiates a new GetTimeSeriesCoppock200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCoppock200Response(meta GetTimeSeriesCoppock200ResponseMeta, values []GetTimeSeriesCoppock200ResponseValuesInner, status string) *GetTimeSeriesCoppock200Response {
	this := GetTimeSeriesCoppock200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesCoppock200ResponseWithDefaults instantiates a new GetTimeSeriesCoppock200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCoppock200ResponseWithDefaults() *GetTimeSeriesCoppock200Response {
	this := GetTimeSeriesCoppock200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesCoppock200Response) GetMeta() GetTimeSeriesCoppock200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesCoppock200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200Response) GetMetaOk() (*GetTimeSeriesCoppock200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesCoppock200Response) SetMeta(v GetTimeSeriesCoppock200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesCoppock200Response) GetValues() []GetTimeSeriesCoppock200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesCoppock200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200Response) GetValuesOk() ([]GetTimeSeriesCoppock200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesCoppock200Response) SetValues(v []GetTimeSeriesCoppock200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesCoppock200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesCoppock200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesCoppock200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCoppock200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesCoppock200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesCoppock200Response := _GetTimeSeriesCoppock200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCoppock200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCoppock200Response(varGetTimeSeriesCoppock200Response)

	return err
}

type NullableGetTimeSeriesCoppock200Response struct {
	value *GetTimeSeriesCoppock200Response
	isSet bool
}

func (v NullableGetTimeSeriesCoppock200Response) Get() *GetTimeSeriesCoppock200Response {
	return v.value
}

func (v *NullableGetTimeSeriesCoppock200Response) Set(val *GetTimeSeriesCoppock200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCoppock200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCoppock200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCoppock200Response(val *GetTimeSeriesCoppock200Response) *NullableGetTimeSeriesCoppock200Response {
	return &NullableGetTimeSeriesCoppock200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCoppock200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCoppock200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
