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

// checks if the GetEpsTrend200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEpsTrend200Response{}

// GetEpsTrend200Response struct for GetEpsTrend200Response
type GetEpsTrend200Response struct {
	Meta GetEarningsEstimate200ResponseMeta `json:"meta"`
	// EPS trend data
	EpsTrend []GetEpsTrend200ResponseEpsTrendInner `json:"eps_trend"`
	// Status of the response
	Status string `json:"status"`
}

type _GetEpsTrend200Response GetEpsTrend200Response

// NewGetEpsTrend200Response instantiates a new GetEpsTrend200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEpsTrend200Response(meta GetEarningsEstimate200ResponseMeta, epsTrend []GetEpsTrend200ResponseEpsTrendInner, status string) *GetEpsTrend200Response {
	this := GetEpsTrend200Response{}
	this.Meta = meta
	this.EpsTrend = epsTrend
	this.Status = status
	return &this
}

// NewGetEpsTrend200ResponseWithDefaults instantiates a new GetEpsTrend200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEpsTrend200ResponseWithDefaults() *GetEpsTrend200Response {
	this := GetEpsTrend200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetEpsTrend200Response) GetMeta() GetEarningsEstimate200ResponseMeta {
	if o == nil {
		var ret GetEarningsEstimate200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetEpsTrend200Response) SetMeta(v GetEarningsEstimate200ResponseMeta) {
	o.Meta = v
}

// GetEpsTrend returns the EpsTrend field value
func (o *GetEpsTrend200Response) GetEpsTrend() []GetEpsTrend200ResponseEpsTrendInner {
	if o == nil {
		var ret []GetEpsTrend200ResponseEpsTrendInner
		return ret
	}

	return o.EpsTrend
}

// GetEpsTrendOk returns a tuple with the EpsTrend field value
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200Response) GetEpsTrendOk() ([]GetEpsTrend200ResponseEpsTrendInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EpsTrend, true
}

// SetEpsTrend sets field value
func (o *GetEpsTrend200Response) SetEpsTrend(v []GetEpsTrend200ResponseEpsTrendInner) {
	o.EpsTrend = v
}

// GetStatus returns the Status field value
func (o *GetEpsTrend200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetEpsTrend200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetEpsTrend200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEpsTrend200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["eps_trend"] = o.EpsTrend
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetEpsTrend200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"eps_trend",
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

	varGetEpsTrend200Response := _GetEpsTrend200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEpsTrend200Response)

	if err != nil {
		return err
	}

	*o = GetEpsTrend200Response(varGetEpsTrend200Response)

	return err
}

type NullableGetEpsTrend200Response struct {
	value *GetEpsTrend200Response
	isSet bool
}

func (v NullableGetEpsTrend200Response) Get() *GetEpsTrend200Response {
	return v.value
}

func (v *NullableGetEpsTrend200Response) Set(val *GetEpsTrend200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEpsTrend200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEpsTrend200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEpsTrend200Response(val *GetEpsTrend200Response) *NullableGetEpsTrend200Response {
	return &NullableGetEpsTrend200Response{value: val, isSet: true}
}

func (v NullableGetEpsTrend200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEpsTrend200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
