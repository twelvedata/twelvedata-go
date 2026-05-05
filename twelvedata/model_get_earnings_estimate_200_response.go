// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetEarningsEstimate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarningsEstimate200Response{}

// GetEarningsEstimate200Response struct for GetEarningsEstimate200Response
type GetEarningsEstimate200Response struct {
	Meta GetEarningsEstimate200ResponseMeta `json:"meta"`
	// List of earnings estimates
	EarningsEstimate []GetEarningsEstimate200ResponseEarningsEstimateInner `json:"earnings_estimate"`
	// Response status
	Status string `json:"status"`
}

type _GetEarningsEstimate200Response GetEarningsEstimate200Response

// NewGetEarningsEstimate200Response instantiates a new GetEarningsEstimate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarningsEstimate200Response(meta GetEarningsEstimate200ResponseMeta, earningsEstimate []GetEarningsEstimate200ResponseEarningsEstimateInner, status string) *GetEarningsEstimate200Response {
	this := GetEarningsEstimate200Response{}
	this.Meta = meta
	this.EarningsEstimate = earningsEstimate
	this.Status = status
	return &this
}

// NewGetEarningsEstimate200ResponseWithDefaults instantiates a new GetEarningsEstimate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarningsEstimate200ResponseWithDefaults() *GetEarningsEstimate200Response {
	this := GetEarningsEstimate200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetEarningsEstimate200Response) GetMeta() GetEarningsEstimate200ResponseMeta {
	if o == nil {
		var ret GetEarningsEstimate200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetEarningsEstimate200Response) SetMeta(v GetEarningsEstimate200ResponseMeta) {
	o.Meta = v
}

// GetEarningsEstimate returns the EarningsEstimate field value
func (o *GetEarningsEstimate200Response) GetEarningsEstimate() []GetEarningsEstimate200ResponseEarningsEstimateInner {
	if o == nil {
		var ret []GetEarningsEstimate200ResponseEarningsEstimateInner
		return ret
	}

	return o.EarningsEstimate
}

// GetEarningsEstimateOk returns a tuple with the EarningsEstimate field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200Response) GetEarningsEstimateOk() ([]GetEarningsEstimate200ResponseEarningsEstimateInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EarningsEstimate, true
}

// SetEarningsEstimate sets field value
func (o *GetEarningsEstimate200Response) SetEarningsEstimate(v []GetEarningsEstimate200ResponseEarningsEstimateInner) {
	o.EarningsEstimate = v
}

// GetStatus returns the Status field value
func (o *GetEarningsEstimate200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetEarningsEstimate200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetEarningsEstimate200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarningsEstimate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["earnings_estimate"] = o.EarningsEstimate
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetEarningsEstimate200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"earnings_estimate",
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

	varGetEarningsEstimate200Response := _GetEarningsEstimate200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEarningsEstimate200Response)

	if err != nil {
		return err
	}

	*o = GetEarningsEstimate200Response(varGetEarningsEstimate200Response)

	return err
}

type NullableGetEarningsEstimate200Response struct {
	value *GetEarningsEstimate200Response
	isSet bool
}

func (v NullableGetEarningsEstimate200Response) Get() *GetEarningsEstimate200Response {
	return v.value
}

func (v *NullableGetEarningsEstimate200Response) Set(val *GetEarningsEstimate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarningsEstimate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarningsEstimate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarningsEstimate200Response(val *GetEarningsEstimate200Response) *NullableGetEarningsEstimate200Response {
	return &NullableGetEarningsEstimate200Response{value: val, isSet: true}
}

func (v NullableGetEarningsEstimate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarningsEstimate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
