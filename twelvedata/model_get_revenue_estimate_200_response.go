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

// checks if the GetRevenueEstimate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRevenueEstimate200Response{}

// GetRevenueEstimate200Response struct for GetRevenueEstimate200Response
type GetRevenueEstimate200Response struct {
	Meta GetEarningsEstimate200ResponseMeta `json:"meta"`
	// Revenue estimate data
	RevenueEstimate []GetRevenueEstimate200ResponseRevenueEstimateInner `json:"revenue_estimate"`
	// Status of the response
	Status string `json:"status"`
}

type _GetRevenueEstimate200Response GetRevenueEstimate200Response

// NewGetRevenueEstimate200Response instantiates a new GetRevenueEstimate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRevenueEstimate200Response(meta GetEarningsEstimate200ResponseMeta, revenueEstimate []GetRevenueEstimate200ResponseRevenueEstimateInner, status string) *GetRevenueEstimate200Response {
	this := GetRevenueEstimate200Response{}
	this.Meta = meta
	this.RevenueEstimate = revenueEstimate
	this.Status = status
	return &this
}

// NewGetRevenueEstimate200ResponseWithDefaults instantiates a new GetRevenueEstimate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRevenueEstimate200ResponseWithDefaults() *GetRevenueEstimate200Response {
	this := GetRevenueEstimate200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetRevenueEstimate200Response) GetMeta() GetEarningsEstimate200ResponseMeta {
	if o == nil {
		var ret GetEarningsEstimate200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetRevenueEstimate200Response) SetMeta(v GetEarningsEstimate200ResponseMeta) {
	o.Meta = v
}

// GetRevenueEstimate returns the RevenueEstimate field value
func (o *GetRevenueEstimate200Response) GetRevenueEstimate() []GetRevenueEstimate200ResponseRevenueEstimateInner {
	if o == nil {
		var ret []GetRevenueEstimate200ResponseRevenueEstimateInner
		return ret
	}

	return o.RevenueEstimate
}

// GetRevenueEstimateOk returns a tuple with the RevenueEstimate field value
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200Response) GetRevenueEstimateOk() ([]GetRevenueEstimate200ResponseRevenueEstimateInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevenueEstimate, true
}

// SetRevenueEstimate sets field value
func (o *GetRevenueEstimate200Response) SetRevenueEstimate(v []GetRevenueEstimate200ResponseRevenueEstimateInner) {
	o.RevenueEstimate = v
}

// GetStatus returns the Status field value
func (o *GetRevenueEstimate200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetRevenueEstimate200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetRevenueEstimate200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRevenueEstimate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["revenue_estimate"] = o.RevenueEstimate
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetRevenueEstimate200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"revenue_estimate",
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

	varGetRevenueEstimate200Response := _GetRevenueEstimate200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetRevenueEstimate200Response)

	if err != nil {
		return err
	}

	*o = GetRevenueEstimate200Response(varGetRevenueEstimate200Response)

	return err
}

type NullableGetRevenueEstimate200Response struct {
	value *GetRevenueEstimate200Response
	isSet bool
}

func (v NullableGetRevenueEstimate200Response) Get() *GetRevenueEstimate200Response {
	return v.value
}

func (v *NullableGetRevenueEstimate200Response) Set(val *GetRevenueEstimate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRevenueEstimate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRevenueEstimate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRevenueEstimate200Response(val *GetRevenueEstimate200Response) *NullableGetRevenueEstimate200Response {
	return &NullableGetRevenueEstimate200Response{value: val, isSet: true}
}

func (v NullableGetRevenueEstimate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRevenueEstimate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
