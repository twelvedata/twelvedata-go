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

// checks if the ResponseSanctionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSanctionItem{}

// ResponseSanctionItem struct for ResponseSanctionItem
type ResponseSanctionItem struct {
	// The sanction source
	Source string `json:"source"`
	// The sanction program
	Program string `json:"program"`
	// Notes for the sanction
	Notes string `json:"notes"`
	// Sanction lists
	Lists []ResponseSanctionItemList `json:"lists"`
}

type _ResponseSanctionItem ResponseSanctionItem

// NewResponseSanctionItem instantiates a new ResponseSanctionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSanctionItem(source string, program string, notes string, lists []ResponseSanctionItemList) *ResponseSanctionItem {
	this := ResponseSanctionItem{}
	this.Source = source
	this.Program = program
	this.Notes = notes
	this.Lists = lists
	return &this
}

// NewResponseSanctionItemWithDefaults instantiates a new ResponseSanctionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSanctionItemWithDefaults() *ResponseSanctionItem {
	this := ResponseSanctionItem{}
	return &this
}

// GetSource returns the Source field value
func (o *ResponseSanctionItem) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionItem) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ResponseSanctionItem) SetSource(v string) {
	o.Source = v
}

// GetProgram returns the Program field value
func (o *ResponseSanctionItem) GetProgram() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Program
}

// GetProgramOk returns a tuple with the Program field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionItem) GetProgramOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Program, true
}

// SetProgram sets field value
func (o *ResponseSanctionItem) SetProgram(v string) {
	o.Program = v
}

// GetNotes returns the Notes field value
func (o *ResponseSanctionItem) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionItem) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *ResponseSanctionItem) SetNotes(v string) {
	o.Notes = v
}

// GetLists returns the Lists field value
func (o *ResponseSanctionItem) GetLists() []ResponseSanctionItemList {
	if o == nil {
		var ret []ResponseSanctionItemList
		return ret
	}

	return o.Lists
}

// GetListsOk returns a tuple with the Lists field value
// and a boolean to check if the value has been set.
func (o *ResponseSanctionItem) GetListsOk() ([]ResponseSanctionItemList, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lists, true
}

// SetLists sets field value
func (o *ResponseSanctionItem) SetLists(v []ResponseSanctionItemList) {
	o.Lists = v
}

func (o ResponseSanctionItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSanctionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["program"] = o.Program
	toSerialize["notes"] = o.Notes
	toSerialize["lists"] = o.Lists
	return toSerialize, nil
}

func (o *ResponseSanctionItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"program",
		"notes",
		"lists",
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

	varResponseSanctionItem := _ResponseSanctionItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varResponseSanctionItem)

	if err != nil {
		return err
	}

	*o = ResponseSanctionItem(varResponseSanctionItem)

	return err
}

type NullableResponseSanctionItem struct {
	value *ResponseSanctionItem
	isSet bool
}

func (v NullableResponseSanctionItem) Get() *ResponseSanctionItem {
	return v.value
}

func (v *NullableResponseSanctionItem) Set(val *ResponseSanctionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSanctionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSanctionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSanctionItem(val *ResponseSanctionItem) *NullableResponseSanctionItem {
	return &NullableResponseSanctionItem{value: val, isSet: true}
}

func (v NullableResponseSanctionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSanctionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
