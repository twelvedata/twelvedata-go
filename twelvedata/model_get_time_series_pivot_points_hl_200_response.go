// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesPivotPointsHL200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPivotPointsHL200Response{}

// GetTimeSeriesPivotPointsHL200Response struct for GetTimeSeriesPivotPointsHL200Response
type GetTimeSeriesPivotPointsHL200Response struct {
	Meta GetTimeSeriesPivotPointsHL200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesPivotPointsHL200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesPivotPointsHL200Response GetTimeSeriesPivotPointsHL200Response

// NewGetTimeSeriesPivotPointsHL200Response instantiates a new GetTimeSeriesPivotPointsHL200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPivotPointsHL200Response(meta GetTimeSeriesPivotPointsHL200ResponseMeta, values []GetTimeSeriesPivotPointsHL200ResponseValuesInner, status string) *GetTimeSeriesPivotPointsHL200Response {
	this := GetTimeSeriesPivotPointsHL200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults instantiates a new GetTimeSeriesPivotPointsHL200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults() *GetTimeSeriesPivotPointsHL200Response {
	this := GetTimeSeriesPivotPointsHL200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesPivotPointsHL200Response) GetMeta() GetTimeSeriesPivotPointsHL200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesPivotPointsHL200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPivotPointsHL200Response) GetMetaOk() (*GetTimeSeriesPivotPointsHL200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesPivotPointsHL200Response) SetMeta(v GetTimeSeriesPivotPointsHL200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesPivotPointsHL200Response) GetValues() []GetTimeSeriesPivotPointsHL200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesPivotPointsHL200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPivotPointsHL200Response) GetValuesOk() ([]GetTimeSeriesPivotPointsHL200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesPivotPointsHL200Response) SetValues(v []GetTimeSeriesPivotPointsHL200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesPivotPointsHL200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPivotPointsHL200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesPivotPointsHL200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesPivotPointsHL200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPivotPointsHL200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesPivotPointsHL200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesPivotPointsHL200Response := _GetTimeSeriesPivotPointsHL200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesPivotPointsHL200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPivotPointsHL200Response(varGetTimeSeriesPivotPointsHL200Response)

	return err
}

type NullableGetTimeSeriesPivotPointsHL200Response struct {
	value *GetTimeSeriesPivotPointsHL200Response
	isSet bool
}

func (v NullableGetTimeSeriesPivotPointsHL200Response) Get() *GetTimeSeriesPivotPointsHL200Response {
	return v.value
}

func (v *NullableGetTimeSeriesPivotPointsHL200Response) Set(val *GetTimeSeriesPivotPointsHL200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPivotPointsHL200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPivotPointsHL200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPivotPointsHL200Response(val *GetTimeSeriesPivotPointsHL200Response) *NullableGetTimeSeriesPivotPointsHL200Response {
	return &NullableGetTimeSeriesPivotPointsHL200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPivotPointsHL200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPivotPointsHL200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
