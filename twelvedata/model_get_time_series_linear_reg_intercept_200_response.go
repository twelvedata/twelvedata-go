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

// checks if the GetTimeSeriesLinearRegIntercept200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLinearRegIntercept200Response{}

// GetTimeSeriesLinearRegIntercept200Response struct for GetTimeSeriesLinearRegIntercept200Response
type GetTimeSeriesLinearRegIntercept200Response struct {
	Meta GetTimeSeriesLinearRegIntercept200ResponseMeta `json:"meta"`
	// Array of time series data points
	Values []GetTimeSeriesLinearRegIntercept200ResponseValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeriesLinearRegIntercept200Response GetTimeSeriesLinearRegIntercept200Response

// NewGetTimeSeriesLinearRegIntercept200Response instantiates a new GetTimeSeriesLinearRegIntercept200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLinearRegIntercept200Response(meta GetTimeSeriesLinearRegIntercept200ResponseMeta, values []GetTimeSeriesLinearRegIntercept200ResponseValuesInner, status string) *GetTimeSeriesLinearRegIntercept200Response {
	this := GetTimeSeriesLinearRegIntercept200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeriesLinearRegIntercept200ResponseWithDefaults instantiates a new GetTimeSeriesLinearRegIntercept200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLinearRegIntercept200ResponseWithDefaults() *GetTimeSeriesLinearRegIntercept200Response {
	this := GetTimeSeriesLinearRegIntercept200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeriesLinearRegIntercept200Response) GetMeta() GetTimeSeriesLinearRegIntercept200ResponseMeta {
	if o == nil {
		var ret GetTimeSeriesLinearRegIntercept200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegIntercept200Response) GetMetaOk() (*GetTimeSeriesLinearRegIntercept200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeriesLinearRegIntercept200Response) SetMeta(v GetTimeSeriesLinearRegIntercept200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeriesLinearRegIntercept200Response) GetValues() []GetTimeSeriesLinearRegIntercept200ResponseValuesInner {
	if o == nil {
		var ret []GetTimeSeriesLinearRegIntercept200ResponseValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegIntercept200Response) GetValuesOk() ([]GetTimeSeriesLinearRegIntercept200ResponseValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeriesLinearRegIntercept200Response) SetValues(v []GetTimeSeriesLinearRegIntercept200ResponseValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeriesLinearRegIntercept200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegIntercept200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeriesLinearRegIntercept200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeriesLinearRegIntercept200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLinearRegIntercept200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeriesLinearRegIntercept200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesLinearRegIntercept200Response := _GetTimeSeriesLinearRegIntercept200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesLinearRegIntercept200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLinearRegIntercept200Response(varGetTimeSeriesLinearRegIntercept200Response)

	return err
}

type NullableGetTimeSeriesLinearRegIntercept200Response struct {
	value *GetTimeSeriesLinearRegIntercept200Response
	isSet bool
}

func (v NullableGetTimeSeriesLinearRegIntercept200Response) Get() *GetTimeSeriesLinearRegIntercept200Response {
	return v.value
}

func (v *NullableGetTimeSeriesLinearRegIntercept200Response) Set(val *GetTimeSeriesLinearRegIntercept200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLinearRegIntercept200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLinearRegIntercept200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLinearRegIntercept200Response(val *GetTimeSeriesLinearRegIntercept200Response) *NullableGetTimeSeriesLinearRegIntercept200Response {
	return &NullableGetTimeSeriesLinearRegIntercept200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLinearRegIntercept200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLinearRegIntercept200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
