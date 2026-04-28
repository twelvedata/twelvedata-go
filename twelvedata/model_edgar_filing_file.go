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

// checks if the EdgarFilingFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgarFilingFile{}

// EdgarFilingFile Filing file object
type EdgarFilingFile struct {
	// File name
	Name string `json:"name"`
	// File size
	Size *int64 `json:"size,omitempty"`
	// File type
	Type string `json:"type"`
	// File full url
	Url string `json:"url"`
}

type _EdgarFilingFile EdgarFilingFile

// NewEdgarFilingFile instantiates a new EdgarFilingFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgarFilingFile(name string, type_ string, url string) *EdgarFilingFile {
	this := EdgarFilingFile{}
	this.Name = name
	this.Type = type_
	this.Url = url
	return &this
}

// NewEdgarFilingFileWithDefaults instantiates a new EdgarFilingFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgarFilingFileWithDefaults() *EdgarFilingFile {
	this := EdgarFilingFile{}
	return &this
}

// GetName returns the Name field value
func (o *EdgarFilingFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgarFilingFile) SetName(v string) {
	o.Name = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *EdgarFilingFile) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgarFilingFile) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *EdgarFilingFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *EdgarFilingFile) SetSize(v int64) {
	o.Size = &v
}

// GetType returns the Type field value
func (o *EdgarFilingFile) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingFile) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EdgarFilingFile) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *EdgarFilingFile) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *EdgarFilingFile) SetUrl(v string) {
	o.Url = v
}

func (o EdgarFilingFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgarFilingFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *EdgarFilingFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"url",
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

	varEdgarFilingFile := _EdgarFilingFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgarFilingFile)

	if err != nil {
		return err
	}

	*o = EdgarFilingFile(varEdgarFilingFile)

	return err
}

type NullableEdgarFilingFile struct {
	value *EdgarFilingFile
	isSet bool
}

func (v NullableEdgarFilingFile) Get() *EdgarFilingFile {
	return v.value
}

func (v *NullableEdgarFilingFile) Set(val *EdgarFilingFile) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgarFilingFile) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgarFilingFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgarFilingFile(val *EdgarFilingFile) *NullableEdgarFilingFile {
	return &NullableEdgarFilingFile{value: val, isSet: true}
}

func (v NullableEdgarFilingFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgarFilingFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
