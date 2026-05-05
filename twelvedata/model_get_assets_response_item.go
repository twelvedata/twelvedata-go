// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetAssetsResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetsResponseItem{}

// GetAssetsResponseItem struct for GetAssetsResponseItem
type GetAssetsResponseItem struct {
	// Currency code
	Code string `json:"code"`
	// Description of the asset
	Description *string `json:"description,omitempty"`
	// Icon of the asset
	Icon *string `json:"icon,omitempty"`
	// Market identifier code, e.g. DIGITAL_CURRENCY, PHYSICAL_CURRENCY, etc.
	MicCode string `json:"mic_code"`
	// Currency symbol
	Symbol *string `json:"symbol,omitempty"`
}

type _GetAssetsResponseItem GetAssetsResponseItem

// NewGetAssetsResponseItem instantiates a new GetAssetsResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetsResponseItem(code string, micCode string) *GetAssetsResponseItem {
	this := GetAssetsResponseItem{}
	this.Code = code
	this.MicCode = micCode
	return &this
}

// NewGetAssetsResponseItemWithDefaults instantiates a new GetAssetsResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetsResponseItemWithDefaults() *GetAssetsResponseItem {
	this := GetAssetsResponseItem{}
	return &this
}

// GetCode returns the Code field value
func (o *GetAssetsResponseItem) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetAssetsResponseItem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GetAssetsResponseItem) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetAssetsResponseItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsResponseItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetAssetsResponseItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetAssetsResponseItem) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *GetAssetsResponseItem) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsResponseItem) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *GetAssetsResponseItem) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *GetAssetsResponseItem) SetIcon(v string) {
	o.Icon = &v
}

// GetMicCode returns the MicCode field value
func (o *GetAssetsResponseItem) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetAssetsResponseItem) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetAssetsResponseItem) SetMicCode(v string) {
	o.MicCode = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetAssetsResponseItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetAssetsResponseItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetAssetsResponseItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o GetAssetsResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetsResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["mic_code"] = o.MicCode
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

func (o *GetAssetsResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"mic_code",
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

	varGetAssetsResponseItem := _GetAssetsResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetAssetsResponseItem)

	if err != nil {
		return err
	}

	*o = GetAssetsResponseItem(varGetAssetsResponseItem)

	return err
}

type NullableGetAssetsResponseItem struct {
	value *GetAssetsResponseItem
	isSet bool
}

func (v NullableGetAssetsResponseItem) Get() *GetAssetsResponseItem {
	return v.value
}

func (v *NullableGetAssetsResponseItem) Set(val *GetAssetsResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetsResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetsResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetsResponseItem(val *GetAssetsResponseItem) *NullableGetAssetsResponseItem {
	return &NullableGetAssetsResponseItem{value: val, isSet: true}
}

func (v NullableGetAssetsResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetsResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
