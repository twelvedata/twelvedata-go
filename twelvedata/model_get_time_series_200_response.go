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

// checks if the GetTimeSeries200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeries200Response{}

// GetTimeSeries200Response struct for GetTimeSeries200Response
type GetTimeSeries200Response struct {
	Meta GetTimeSeries200ResponseMeta `json:"meta"`
	// List of time series data points
	Values []TimeSeriesItem `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _GetTimeSeries200Response GetTimeSeries200Response

// NewGetTimeSeries200Response instantiates a new GetTimeSeries200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeries200Response(meta GetTimeSeries200ResponseMeta, values []TimeSeriesItem, status string) *GetTimeSeries200Response {
	this := GetTimeSeries200Response{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewGetTimeSeries200ResponseWithDefaults instantiates a new GetTimeSeries200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeries200ResponseWithDefaults() *GetTimeSeries200Response {
	this := GetTimeSeries200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTimeSeries200Response) GetMeta() GetTimeSeries200ResponseMeta {
	if o == nil {
		var ret GetTimeSeries200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeries200Response) GetMetaOk() (*GetTimeSeries200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTimeSeries200Response) SetMeta(v GetTimeSeries200ResponseMeta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *GetTimeSeries200Response) GetValues() []TimeSeriesItem {
	if o == nil {
		var ret []TimeSeriesItem
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeries200Response) GetValuesOk() ([]TimeSeriesItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTimeSeries200Response) SetValues(v []TimeSeriesItem) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *GetTimeSeries200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeries200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTimeSeries200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTimeSeries200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeries200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTimeSeries200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeries200Response := _GetTimeSeries200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeries200Response)

	if err != nil {
		return err
	}

	*o = GetTimeSeries200Response(varGetTimeSeries200Response)

	return err
}

type NullableGetTimeSeries200Response struct {
	value *GetTimeSeries200Response
	isSet bool
}

func (v NullableGetTimeSeries200Response) Get() *GetTimeSeries200Response {
	return v.value
}

func (v *NullableGetTimeSeries200Response) Set(val *GetTimeSeries200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeries200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeries200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeries200Response(val *GetTimeSeries200Response) *NullableGetTimeSeries200Response {
	return &NullableGetTimeSeries200Response{value: val, isSet: true}
}

func (v NullableGetTimeSeries200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeries200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
