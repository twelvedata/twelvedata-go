// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PressRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PressRelease{}

// PressRelease struct for PressRelease
type PressRelease struct {
	// Press release unique identifier
	Id string `json:"id"`
	// Press release date in ISO 8601 format
	Datetime string `json:"datetime"`
	// Press release title
	Title string `json:"title"`
	// Press release body in html format
	Body string `json:"body"`
	// Custom style applied to the release
	Style *string `json:"style,omitempty"`
	// Press release language codes
	Language []string `json:"language,omitempty"`
}

type _PressRelease PressRelease

// NewPressRelease instantiates a new PressRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPressRelease(id string, datetime string, title string, body string) *PressRelease {
	this := PressRelease{}
	this.Id = id
	this.Datetime = datetime
	this.Title = title
	this.Body = body
	return &this
}

// NewPressReleaseWithDefaults instantiates a new PressRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPressReleaseWithDefaults() *PressRelease {
	this := PressRelease{}
	return &this
}

// GetId returns the Id field value
func (o *PressRelease) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PressRelease) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PressRelease) SetId(v string) {
	o.Id = v
}

// GetDatetime returns the Datetime field value
func (o *PressRelease) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *PressRelease) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *PressRelease) SetDatetime(v string) {
	o.Datetime = v
}

// GetTitle returns the Title field value
func (o *PressRelease) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PressRelease) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PressRelease) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value
func (o *PressRelease) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *PressRelease) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *PressRelease) SetBody(v string) {
	o.Body = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *PressRelease) GetStyle() string {
	if o == nil || IsNil(o.Style) {
		var ret string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PressRelease) GetStyleOk() (*string, bool) {
	if o == nil || IsNil(o.Style) {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *PressRelease) HasStyle() bool {
	if o != nil && !IsNil(o.Style) {
		return true
	}

	return false
}

// SetStyle gets a reference to the given string and assigns it to the Style field.
func (o *PressRelease) SetStyle(v string) {
	o.Style = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PressRelease) GetLanguage() []string {
	if o == nil || IsNil(o.Language) {
		var ret []string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PressRelease) GetLanguageOk() ([]string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PressRelease) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given []string and assigns it to the Language field.
func (o *PressRelease) SetLanguage(v []string) {
	o.Language = v
}

func (o PressRelease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PressRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["datetime"] = o.Datetime
	toSerialize["title"] = o.Title
	toSerialize["body"] = o.Body
	if !IsNil(o.Style) {
		toSerialize["style"] = o.Style
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

func (o *PressRelease) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"datetime",
		"title",
		"body",
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

	varPressRelease := _PressRelease{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPressRelease)

	if err != nil {
		return err
	}

	*o = PressRelease(varPressRelease)

	return err
}

type NullablePressRelease struct {
	value *PressRelease
	isSet bool
}

func (v NullablePressRelease) Get() *PressRelease {
	return v.value
}

func (v *NullablePressRelease) Set(val *PressRelease) {
	v.value = val
	v.isSet = true
}

func (v NullablePressRelease) IsSet() bool {
	return v.isSet
}

func (v *NullablePressRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePressRelease(val *PressRelease) *NullablePressRelease {
	return &NullablePressRelease{value: val, isSet: true}
}

func (v NullablePressRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePressRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
