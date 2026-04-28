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

// checks if the GetLogo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLogo200Response{}

// GetLogo200Response struct for GetLogo200Response
type GetLogo200Response struct {
	Meta GetLogo200ResponseMeta `json:"meta"`
	// Link to download the logo (for stocks only)
	Url *string `json:"url,omitempty"`
	// Link to download the base currency logo (for `forex` and `crypto` only)
	LogoBase *string `json:"logo_base,omitempty"`
	// Link to download the quote currency logo (for `forex` and `crypto` only)
	LogoQuote *string `json:"logo_quote,omitempty"`
}

type _GetLogo200Response GetLogo200Response

// NewGetLogo200Response instantiates a new GetLogo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLogo200Response(meta GetLogo200ResponseMeta) *GetLogo200Response {
	this := GetLogo200Response{}
	this.Meta = meta
	return &this
}

// NewGetLogo200ResponseWithDefaults instantiates a new GetLogo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLogo200ResponseWithDefaults() *GetLogo200Response {
	this := GetLogo200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetLogo200Response) GetMeta() GetLogo200ResponseMeta {
	if o == nil {
		var ret GetLogo200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetLogo200Response) GetMetaOk() (*GetLogo200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetLogo200Response) SetMeta(v GetLogo200ResponseMeta) {
	o.Meta = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetLogo200Response) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogo200Response) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetLogo200Response) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetLogo200Response) SetUrl(v string) {
	o.Url = &v
}

// GetLogoBase returns the LogoBase field value if set, zero value otherwise.
func (o *GetLogo200Response) GetLogoBase() string {
	if o == nil || IsNil(o.LogoBase) {
		var ret string
		return ret
	}
	return *o.LogoBase
}

// GetLogoBaseOk returns a tuple with the LogoBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogo200Response) GetLogoBaseOk() (*string, bool) {
	if o == nil || IsNil(o.LogoBase) {
		return nil, false
	}
	return o.LogoBase, true
}

// HasLogoBase returns a boolean if a field has been set.
func (o *GetLogo200Response) HasLogoBase() bool {
	if o != nil && !IsNil(o.LogoBase) {
		return true
	}

	return false
}

// SetLogoBase gets a reference to the given string and assigns it to the LogoBase field.
func (o *GetLogo200Response) SetLogoBase(v string) {
	o.LogoBase = &v
}

// GetLogoQuote returns the LogoQuote field value if set, zero value otherwise.
func (o *GetLogo200Response) GetLogoQuote() string {
	if o == nil || IsNil(o.LogoQuote) {
		var ret string
		return ret
	}
	return *o.LogoQuote
}

// GetLogoQuoteOk returns a tuple with the LogoQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogo200Response) GetLogoQuoteOk() (*string, bool) {
	if o == nil || IsNil(o.LogoQuote) {
		return nil, false
	}
	return o.LogoQuote, true
}

// HasLogoQuote returns a boolean if a field has been set.
func (o *GetLogo200Response) HasLogoQuote() bool {
	if o != nil && !IsNil(o.LogoQuote) {
		return true
	}

	return false
}

// SetLogoQuote gets a reference to the given string and assigns it to the LogoQuote field.
func (o *GetLogo200Response) SetLogoQuote(v string) {
	o.LogoQuote = &v
}

func (o GetLogo200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.LogoBase) {
		toSerialize["logo_base"] = o.LogoBase
	}
	if !IsNil(o.LogoQuote) {
		toSerialize["logo_quote"] = o.LogoQuote
	}
	return toSerialize, nil
}

func (o *GetLogo200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
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

	varGetLogo200Response := _GetLogo200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetLogo200Response)

	if err != nil {
		return err
	}

	*o = GetLogo200Response(varGetLogo200Response)

	return err
}

type NullableGetLogo200Response struct {
	value *GetLogo200Response
	isSet bool
}

func (v NullableGetLogo200Response) Get() *GetLogo200Response {
	return v.value
}

func (v *NullableGetLogo200Response) Set(val *GetLogo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogo200Response(val *GetLogo200Response) *NullableGetLogo200Response {
	return &NullableGetLogo200Response{value: val, isSet: true}
}

func (v NullableGetLogo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
