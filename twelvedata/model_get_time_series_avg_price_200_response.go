// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAvgPrice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAvgPrice200Response{}

// GetTimeSeriesAvgPrice200Response struct for GetTimeSeriesAvgPrice200Response
type GetTimeSeriesAvgPrice200Response struct {
	Meta GetTimeSeriesAvgPrice200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesAvgPrice200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesAvgPrice200Response GetTimeSeriesAvgPrice200Response

// NewGetTimeSeriesAvgPrice200Response instantiates a new GetTimeSeriesAvgPrice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAvgPrice200Response(meta GetTimeSeriesAvgPrice200ResponseMeta, values []GetTimeSeriesAvgPrice200ResponseValuesInner, status string) *GetTimeSeriesAvgPrice200Response {
	this := GetTimeSeriesAvgPrice200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesAvgPrice200ResponseWithDefaults instantiates a new GetTimeSeriesAvgPrice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAvgPrice200ResponseWithDefaults() *GetTimeSeriesAvgPrice200Response {
	this := GetTimeSeriesAvgPrice200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesAvgPrice200Response) GetMeta() GetTimeSeriesAvgPrice200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesAvgPrice200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvgPrice200Response) GetMetaOk() (*GetTimeSeriesAvgPrice200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesAvgPrice200Response) SetMeta(v GetTimeSeriesAvgPrice200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesAvgPrice200Response) GetValues() []GetTimeSeriesAvgPrice200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesAvgPrice200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvgPrice200Response) GetValuesOk() ([]GetTimeSeriesAvgPrice200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesAvgPrice200Response) SetValues(v []GetTimeSeriesAvgPrice200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesAvgPrice200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvgPrice200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesAvgPrice200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesAvgPrice200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAvgPrice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesAvgPrice200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAvgPrice200Response := _GetTimeSeriesAvgPrice200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAvgPrice200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAvgPrice200Response(varGetTimeSeriesAvgPrice200Response)

	return err
}

type NullableGetTimeSeriesAvgPrice200Response struct {
	value *GetTimeSeriesAvgPrice200Response
	isSet bool
}

func (v NullableGetTimeSeriesAvgPrice200Response) Get() *GetTimeSeriesAvgPrice200Response {
	return v.value
}

func (v *NullableGetTimeSeriesAvgPrice200Response) Set(val *GetTimeSeriesAvgPrice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAvgPrice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAvgPrice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAvgPrice200Response(val *GetTimeSeriesAvgPrice200Response) *NullableGetTimeSeriesAvgPrice200Response {
	return &NullableGetTimeSeriesAvgPrice200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAvgPrice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAvgPrice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
