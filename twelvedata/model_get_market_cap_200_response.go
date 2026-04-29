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

// checks if the GetMarketCap200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketCap200Response{}

// GetMarketCap200Response struct for GetMarketCap200Response
type GetMarketCap200Response struct {
	Meta GetMarketCap200ResponseMeta `json:"meta"`
	// Market capitalization values
	MarketCap []GetMarketCap200ResponseMarketCapInner `json:"market_cap"`
}

type _GetMarketCap200Response GetMarketCap200Response

// NewGetMarketCap200Response instantiates a new GetMarketCap200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketCap200Response(meta GetMarketCap200ResponseMeta, marketCap []GetMarketCap200ResponseMarketCapInner) *GetMarketCap200Response {
	this := GetMarketCap200Response{}
	this.Meta = meta
	this.MarketCap = marketCap
	return &this
}

// NewGetMarketCap200ResponseWithDefaults instantiates a new GetMarketCap200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketCap200ResponseWithDefaults() *GetMarketCap200Response {
	this := GetMarketCap200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetMarketCap200Response) GetMeta() GetMarketCap200ResponseMeta {
	if o == nil {
		var ret GetMarketCap200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetMarketCap200Response) GetMetaOk() (*GetMarketCap200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetMarketCap200Response) SetMeta(v GetMarketCap200ResponseMeta) {
	o.Meta = v
}

// GetMarketCap returns the MarketCap field value
func (o *GetMarketCap200Response) GetMarketCap() []GetMarketCap200ResponseMarketCapInner {
	if o == nil {
		var ret []GetMarketCap200ResponseMarketCapInner
		return ret
	}

	return o.MarketCap
}

// GetMarketCapOk returns a tuple with the MarketCap field value
// and a boolean to check if the value has been set.
func (o *GetMarketCap200Response) GetMarketCapOk() ([]GetMarketCap200ResponseMarketCapInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketCap, true
}

// SetMarketCap sets field value
func (o *GetMarketCap200Response) SetMarketCap(v []GetMarketCap200ResponseMarketCapInner) {
	o.MarketCap = v
}

func (o GetMarketCap200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketCap200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["market_cap"] = o.MarketCap
	return toSerialize, nil
}

func (o *GetMarketCap200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"market_cap",
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

	varGetMarketCap200Response := _GetMarketCap200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetMarketCap200Response)

	if err != nil {
		return err
	}

	*o = GetMarketCap200Response(varGetMarketCap200Response)

	return err
}

type NullableGetMarketCap200Response struct {
	value *GetMarketCap200Response
	isSet bool
}

func (v NullableGetMarketCap200Response) Get() *GetMarketCap200Response {
	return v.value
}

func (v *NullableGetMarketCap200Response) Set(val *GetMarketCap200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketCap200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketCap200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketCap200Response(val *GetMarketCap200Response) *NullableGetMarketCap200Response {
	return &NullableGetMarketCap200Response{value: val, isSet: true}
}

func (v NullableGetMarketCap200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketCap200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
