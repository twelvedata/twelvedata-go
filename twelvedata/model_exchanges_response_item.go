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

// checks if the ExchangesResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangesResponseItem{}

// ExchangesResponseItem ExchangesResponseItem represents details of an exchange
type ExchangesResponseItem struct {
	// Title of exchange
	Title string `json:"title"`
	// Name of exchange
	Name string `json:"name"`
	// Market identifier code (MIC) under ISO 10383 standard
	Code string `json:"code"`
	// Country to which stock exchange belongs to
	Country string `json:"country"`
	// Time zone where exchange is located
	Timezone string                       `json:"timezone"`
	Access   *ExchangesResponseItemAccess `json:"access,omitempty"`
}

type _ExchangesResponseItem ExchangesResponseItem

// NewExchangesResponseItem instantiates a new ExchangesResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangesResponseItem(title string, name string, code string, country string, timezone string) *ExchangesResponseItem {
	this := ExchangesResponseItem{}
	this.Title = title
	this.Name = name
	this.Code = code
	this.Country = country
	this.Timezone = timezone
	return &this
}

// NewExchangesResponseItemWithDefaults instantiates a new ExchangesResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangesResponseItemWithDefaults() *ExchangesResponseItem {
	this := ExchangesResponseItem{}
	return &this
}

// GetTitle returns the Title field value
func (o *ExchangesResponseItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ExchangesResponseItem) SetTitle(v string) {
	o.Title = v
}

// GetName returns the Name field value
func (o *ExchangesResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExchangesResponseItem) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *ExchangesResponseItem) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ExchangesResponseItem) SetCode(v string) {
	o.Code = v
}

// GetCountry returns the Country field value
func (o *ExchangesResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *ExchangesResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetTimezone returns the Timezone field value
func (o *ExchangesResponseItem) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItem) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *ExchangesResponseItem) SetTimezone(v string) {
	o.Timezone = v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ExchangesResponseItem) GetAccess() ExchangesResponseItemAccess {
	if o == nil || IsNil(o.Access) {
		var ret ExchangesResponseItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesResponseItem) GetAccessOk() (*ExchangesResponseItemAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ExchangesResponseItem) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given ExchangesResponseItemAccess and assigns it to the Access field.
func (o *ExchangesResponseItem) SetAccess(v ExchangesResponseItemAccess) {
	o.Access = &v
}

func (o ExchangesResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangesResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["name"] = o.Name
	toSerialize["code"] = o.Code
	toSerialize["country"] = o.Country
	toSerialize["timezone"] = o.Timezone
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

func (o *ExchangesResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"name",
		"code",
		"country",
		"timezone",
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

	varExchangesResponseItem := _ExchangesResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangesResponseItem)

	if err != nil {
		return err
	}

	*o = ExchangesResponseItem(varExchangesResponseItem)

	return err
}

type NullableExchangesResponseItem struct {
	value *ExchangesResponseItem
	isSet bool
}

func (v NullableExchangesResponseItem) Get() *ExchangesResponseItem {
	return v.value
}

func (v *NullableExchangesResponseItem) Set(val *ExchangesResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangesResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangesResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangesResponseItem(val *ExchangesResponseItem) *NullableExchangesResponseItem {
	return &NullableExchangesResponseItem{value: val, isSet: true}
}

func (v NullableExchangesResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangesResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
