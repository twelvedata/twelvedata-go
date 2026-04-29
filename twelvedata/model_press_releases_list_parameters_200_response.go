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

// checks if the PressReleasesListParameters200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PressReleasesListParameters200Response{}

// PressReleasesListParameters200Response struct for PressReleasesListParameters200Response
type PressReleasesListParameters200Response struct {
	// List of press releases
	PressReleases []PressRelease `json:"press_releases"`
	// Response status
	Status string `json:"status"`
}

type _PressReleasesListParameters200Response PressReleasesListParameters200Response

// NewPressReleasesListParameters200Response instantiates a new PressReleasesListParameters200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPressReleasesListParameters200Response(pressReleases []PressRelease, status string) *PressReleasesListParameters200Response {
	this := PressReleasesListParameters200Response{}
	this.PressReleases = pressReleases
	this.Status = status
	return &this
}

// NewPressReleasesListParameters200ResponseWithDefaults instantiates a new PressReleasesListParameters200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPressReleasesListParameters200ResponseWithDefaults() *PressReleasesListParameters200Response {
	this := PressReleasesListParameters200Response{}
	return &this
}

// GetPressReleases returns the PressReleases field value
func (o *PressReleasesListParameters200Response) GetPressReleases() []PressRelease {
	if o == nil {
		var ret []PressRelease
		return ret
	}

	return o.PressReleases
}

// GetPressReleasesOk returns a tuple with the PressReleases field value
// and a boolean to check if the value has been set.
func (o *PressReleasesListParameters200Response) GetPressReleasesOk() ([]PressRelease, bool) {
	if o == nil {
		return nil, false
	}
	return o.PressReleases, true
}

// SetPressReleases sets field value
func (o *PressReleasesListParameters200Response) SetPressReleases(v []PressRelease) {
	o.PressReleases = v
}

// GetStatus returns the Status field value
func (o *PressReleasesListParameters200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PressReleasesListParameters200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PressReleasesListParameters200Response) SetStatus(v string) {
	o.Status = v
}

func (o PressReleasesListParameters200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PressReleasesListParameters200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["press_releases"] = o.PressReleases
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *PressReleasesListParameters200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"press_releases",
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

	varPressReleasesListParameters200Response := _PressReleasesListParameters200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPressReleasesListParameters200Response)

	if err != nil {
		return err
	}

	*o = PressReleasesListParameters200Response(varPressReleasesListParameters200Response)

	return err
}

type NullablePressReleasesListParameters200Response struct {
	value *PressReleasesListParameters200Response
	isSet bool
}

func (v NullablePressReleasesListParameters200Response) Get() *PressReleasesListParameters200Response {
	return v.value
}

func (v *NullablePressReleasesListParameters200Response) Set(val *PressReleasesListParameters200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePressReleasesListParameters200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePressReleasesListParameters200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePressReleasesListParameters200Response(val *PressReleasesListParameters200Response) *NullablePressReleasesListParameters200Response {
	return &NullablePressReleasesListParameters200Response{value: val, isSet: true}
}

func (v NullablePressReleasesListParameters200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePressReleasesListParameters200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
