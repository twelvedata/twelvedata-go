// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetPriceTarget200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPriceTarget200Response{}

// GetPriceTarget200Response struct for GetPriceTarget200Response
type GetPriceTarget200Response struct {
	Meta        GetPriceTarget200ResponseMeta        `json:"meta"`
	PriceTarget GetPriceTarget200ResponsePriceTarget `json:"price_target"`
	// Response status
	Status string `json:"status"`
}

type _GetPriceTarget200Response GetPriceTarget200Response

// NewGetPriceTarget200Response instantiates a new GetPriceTarget200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPriceTarget200Response(meta GetPriceTarget200ResponseMeta, priceTarget GetPriceTarget200ResponsePriceTarget, status string) *GetPriceTarget200Response {
	this := GetPriceTarget200Response{}
	this.Meta = meta
	this.PriceTarget = priceTarget
	this.Status = status
	return &this
}

// NewGetPriceTarget200ResponseWithDefaults instantiates a new GetPriceTarget200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPriceTarget200ResponseWithDefaults() *GetPriceTarget200Response {
	this := GetPriceTarget200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetPriceTarget200Response) GetMeta() GetPriceTarget200ResponseMeta {
	if o == nil {
		var ret GetPriceTarget200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200Response) GetMetaOk() (*GetPriceTarget200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetPriceTarget200Response) SetMeta(v GetPriceTarget200ResponseMeta) {
	o.Meta = v
}

// GetPriceTarget returns the PriceTarget field value
func (o *GetPriceTarget200Response) GetPriceTarget() GetPriceTarget200ResponsePriceTarget {
	if o == nil {
		var ret GetPriceTarget200ResponsePriceTarget
		return ret
	}

	return o.PriceTarget
}

// GetPriceTargetOk returns a tuple with the PriceTarget field value
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200Response) GetPriceTargetOk() (*GetPriceTarget200ResponsePriceTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceTarget, true
}

// SetPriceTarget sets field value
func (o *GetPriceTarget200Response) SetPriceTarget(v GetPriceTarget200ResponsePriceTarget) {
	o.PriceTarget = v
}

// GetStatus returns the Status field value
func (o *GetPriceTarget200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetPriceTarget200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetPriceTarget200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPriceTarget200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["price_target"] = o.PriceTarget
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetPriceTarget200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"price_target",
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

	varGetPriceTarget200Response := _GetPriceTarget200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetPriceTarget200Response)

	if err != nil {
		return err
	}

	*o = GetPriceTarget200Response(varGetPriceTarget200Response)

	return err
}

type NullableGetPriceTarget200Response struct {
	value *GetPriceTarget200Response
	isSet bool
}

func (v NullableGetPriceTarget200Response) Get() *GetPriceTarget200Response {
	return v.value
}

func (v *NullableGetPriceTarget200Response) Set(val *GetPriceTarget200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPriceTarget200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPriceTarget200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPriceTarget200Response(val *GetPriceTarget200Response) *NullableGetPriceTarget200Response {
	return &NullableGetPriceTarget200Response{value: val, isSet: true}
}

func (v NullableGetPriceTarget200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPriceTarget200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
