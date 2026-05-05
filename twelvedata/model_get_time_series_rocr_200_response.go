// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesRocr200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRocr200Response{}

// GetTimeSeriesRocr200Response struct for GetTimeSeriesRocr200Response
type GetTimeSeriesRocr200Response struct {
	Meta GetTimeSeriesRocr200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesRocr200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesRocr200Response GetTimeSeriesRocr200Response

// NewGetTimeSeriesRocr200Response instantiates a new GetTimeSeriesRocr200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRocr200Response(meta GetTimeSeriesRocr200ResponseMeta, values []GetTimeSeriesRocr200ResponseValuesInner, status string) *GetTimeSeriesRocr200Response {
	this := GetTimeSeriesRocr200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesRocr200ResponseWithDefaults instantiates a new GetTimeSeriesRocr200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRocr200ResponseWithDefaults() *GetTimeSeriesRocr200Response {
	this := GetTimeSeriesRocr200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesRocr200Response) GetMeta() GetTimeSeriesRocr200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesRocr200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocr200Response) GetMetaOk() (*GetTimeSeriesRocr200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesRocr200Response) SetMeta(v GetTimeSeriesRocr200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesRocr200Response) GetValues() []GetTimeSeriesRocr200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesRocr200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocr200Response) GetValuesOk() ([]GetTimeSeriesRocr200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesRocr200Response) SetValues(v []GetTimeSeriesRocr200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesRocr200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocr200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesRocr200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesRocr200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRocr200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesRocr200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesRocr200Response := _GetTimeSeriesRocr200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesRocr200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRocr200Response(varGetTimeSeriesRocr200Response)

	return err
}

type NullableGetTimeSeriesRocr200Response struct {
	value *GetTimeSeriesRocr200Response
	isSet bool
}

func (v NullableGetTimeSeriesRocr200Response) Get() *GetTimeSeriesRocr200Response {
	return v.value
}

func (v *NullableGetTimeSeriesRocr200Response) Set(val *GetTimeSeriesRocr200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRocr200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRocr200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRocr200Response(val *GetTimeSeriesRocr200Response) *NullableGetTimeSeriesRocr200Response {
	return &NullableGetTimeSeriesRocr200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRocr200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRocr200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
