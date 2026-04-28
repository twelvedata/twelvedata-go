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

// checks if the GetTimeSeriesKeltner200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKeltner200Response{}

// GetTimeSeriesKeltner200Response struct for GetTimeSeriesKeltner200Response
type GetTimeSeriesKeltner200Response struct {
	Meta GetTimeSeriesKeltner200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesKeltner200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesKeltner200Response GetTimeSeriesKeltner200Response

// NewGetTimeSeriesKeltner200Response instantiates a new GetTimeSeriesKeltner200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKeltner200Response(meta GetTimeSeriesKeltner200ResponseMeta, values []GetTimeSeriesKeltner200ResponseValuesInner, status string) *GetTimeSeriesKeltner200Response {
	this := GetTimeSeriesKeltner200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesKeltner200ResponseWithDefaults instantiates a new GetTimeSeriesKeltner200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKeltner200ResponseWithDefaults() *GetTimeSeriesKeltner200Response {
	this := GetTimeSeriesKeltner200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesKeltner200Response) GetMeta() GetTimeSeriesKeltner200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesKeltner200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200Response) GetMetaOk() (*GetTimeSeriesKeltner200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesKeltner200Response) SetMeta(v GetTimeSeriesKeltner200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesKeltner200Response) GetValues() []GetTimeSeriesKeltner200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesKeltner200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200Response) GetValuesOk() ([]GetTimeSeriesKeltner200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesKeltner200Response) SetValues(v []GetTimeSeriesKeltner200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesKeltner200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesKeltner200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesKeltner200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKeltner200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesKeltner200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesKeltner200Response := _GetTimeSeriesKeltner200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesKeltner200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKeltner200Response(varGetTimeSeriesKeltner200Response)

	return err
}

type NullableGetTimeSeriesKeltner200Response struct {
	value *GetTimeSeriesKeltner200Response
	isSet bool
}

func (v NullableGetTimeSeriesKeltner200Response) Get() *GetTimeSeriesKeltner200Response {
	return v.value
}

func (v *NullableGetTimeSeriesKeltner200Response) Set(val *GetTimeSeriesKeltner200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKeltner200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKeltner200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKeltner200Response(val *GetTimeSeriesKeltner200Response) *NullableGetTimeSeriesKeltner200Response {
	return &NullableGetTimeSeriesKeltner200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKeltner200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKeltner200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
