// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ResponseSanctionItemList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSanctionItemList{}

// ResponseSanctionItemList struct for ResponseSanctionItemList
type ResponseSanctionItemList struct {
	// The sanction list name
	Name string `json:"name"`
	// The sanction published date in the current sanctions list
	PublishedAt *string `json:"published_at,omitempty"`
}

type _ResponseSanctionItemList ResponseSanctionItemList

// NewResponseSanctionItemList instantiates a new ResponseSanctionItemList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSanctionItemList(name string) *ResponseSanctionItemList {
	this := ResponseSanctionItemList{}
	this.Name = name
	return &this
}

// NewResponseSanctionItemListWithDefaults instantiates a new ResponseSanctionItemList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSanctionItemListWithDefaults() *ResponseSanctionItemList {
	this := ResponseSanctionItemList{}
	return &this
}

// GetName returns the Name field value
func (o *ResponseSanctionItemList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionItemList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseSanctionItemList) SetName(v string) {
	o.Name = v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *ResponseSanctionItemList) GetPublishedAt() string {
	if o == nil || IsNil(o.PublishedAt) {
		var ret string
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSanctionItemList) GetPublishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *ResponseSanctionItemList) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given string and assigns it to the PublishedAt field.
func (o *ResponseSanctionItemList) SetPublishedAt(v string) {
	o.PublishedAt = &v
}

func (o ResponseSanctionItemList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSanctionItemList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.PublishedAt) {
		toSerialize["published_at"] = o.PublishedAt
	}
	return toSerialize, nil
}

func (o *ResponseSanctionItemList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varResponseSanctionItemList := _ResponseSanctionItemList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varResponseSanctionItemList)

	if err != nil {
		return err
	}

	*o = ResponseSanctionItemList(varResponseSanctionItemList)

	return err
}

type NullableResponseSanctionItemList struct {
	value *ResponseSanctionItemList
	isSet bool
}

func (v NullableResponseSanctionItemList) Get() *ResponseSanctionItemList {
	return v.value
}

func (v *NullableResponseSanctionItemList) Set(val *ResponseSanctionItemList) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSanctionItemList) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSanctionItemList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSanctionItemList(val *ResponseSanctionItemList) *NullableResponseSanctionItemList {
	return &NullableResponseSanctionItemList{value: val, isSet: true}
}

func (v NullableResponseSanctionItemList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSanctionItemList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
