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

// checks if the GetTimeSeriesRvol200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRvol200Response{}

// GetTimeSeriesRvol200Response struct for GetTimeSeriesRvol200Response
type GetTimeSeriesRvol200Response struct {
	Meta GetTimeSeriesRvol200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesRvol200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesRvol200Response GetTimeSeriesRvol200Response

// NewGetTimeSeriesRvol200Response instantiates a new GetTimeSeriesRvol200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRvol200Response(meta GetTimeSeriesRvol200ResponseMeta, values []GetTimeSeriesRvol200ResponseValuesInner, status string) *GetTimeSeriesRvol200Response {
	this := GetTimeSeriesRvol200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesRvol200ResponseWithDefaults instantiates a new GetTimeSeriesRvol200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRvol200ResponseWithDefaults() *GetTimeSeriesRvol200Response {
	this := GetTimeSeriesRvol200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesRvol200Response) GetMeta() GetTimeSeriesRvol200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesRvol200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRvol200Response) GetMetaOk() (*GetTimeSeriesRvol200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesRvol200Response) SetMeta(v GetTimeSeriesRvol200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesRvol200Response) GetValues() []GetTimeSeriesRvol200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesRvol200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRvol200Response) GetValuesOk() ([]GetTimeSeriesRvol200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesRvol200Response) SetValues(v []GetTimeSeriesRvol200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesRvol200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRvol200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesRvol200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesRvol200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRvol200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesRvol200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesRvol200Response := _GetTimeSeriesRvol200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesRvol200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRvol200Response(varGetTimeSeriesRvol200Response)

	return err
}

type NullableGetTimeSeriesRvol200Response struct {
	value *GetTimeSeriesRvol200Response
	isSet bool
}

func (v NullableGetTimeSeriesRvol200Response) Get() *GetTimeSeriesRvol200Response {
	return v.value
}

func (v *NullableGetTimeSeriesRvol200Response) Set(val *GetTimeSeriesRvol200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRvol200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRvol200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRvol200Response(val *GetTimeSeriesRvol200Response) *NullableGetTimeSeriesRvol200Response {
	return &NullableGetTimeSeriesRvol200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRvol200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRvol200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
