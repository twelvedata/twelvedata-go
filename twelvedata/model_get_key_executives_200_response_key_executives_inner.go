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

// checks if the GetKeyExecutives200ResponseKeyExecutivesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetKeyExecutives200ResponseKeyExecutivesInner{}

// GetKeyExecutives200ResponseKeyExecutivesInner struct for GetKeyExecutives200ResponseKeyExecutivesInner
type GetKeyExecutives200ResponseKeyExecutivesInner struct {
	// Full name of an executive, including first name, middle name, last name, and suffix
	Name string `json:"name"`
	// Refers to job title
	Title string `json:"title"`
	// Current age of an executive if available
	Age *int64 `json:"age,omitempty"`
	// Year of birth of an executive if available
	YearBorn *int64 `json:"year_born,omitempty"`
	// Total salary of an executive if available
	Pay *int64 `json:"pay,omitempty"`
}

type _GetKeyExecutives200ResponseKeyExecutivesInner GetKeyExecutives200ResponseKeyExecutivesInner

// NewGetKeyExecutives200ResponseKeyExecutivesInner instantiates a new GetKeyExecutives200ResponseKeyExecutivesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKeyExecutives200ResponseKeyExecutivesInner(name string, title string) *GetKeyExecutives200ResponseKeyExecutivesInner {
	this := GetKeyExecutives200ResponseKeyExecutivesInner{}
	this.Name = name
	this.Title = title
	return &this
}

// NewGetKeyExecutives200ResponseKeyExecutivesInnerWithDefaults instantiates a new GetKeyExecutives200ResponseKeyExecutivesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKeyExecutives200ResponseKeyExecutivesInnerWithDefaults() *GetKeyExecutives200ResponseKeyExecutivesInner {
	this := GetKeyExecutives200ResponseKeyExecutivesInner{}
	return &this
}

// GetName returns the Name field value
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetTitle(v string) {
	o.Title = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetAge() int64 {
	if o == nil || IsNil(o.Age) {
		var ret int64
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetAgeOk() (*int64, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given int64 and assigns it to the Age field.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetAge(v int64) {
	o.Age = &v
}

// GetYearBorn returns the YearBorn field value if set, zero value otherwise.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetYearBorn() int64 {
	if o == nil || IsNil(o.YearBorn) {
		var ret int64
		return ret
	}
	return *o.YearBorn
}

// GetYearBornOk returns a tuple with the YearBorn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetYearBornOk() (*int64, bool) {
	if o == nil || IsNil(o.YearBorn) {
		return nil, false
	}
	return o.YearBorn, true
}

// HasYearBorn returns a boolean if a field has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) HasYearBorn() bool {
	if o != nil && !IsNil(o.YearBorn) {
		return true
	}

	return false
}

// SetYearBorn gets a reference to the given int64 and assigns it to the YearBorn field.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetYearBorn(v int64) {
	o.YearBorn = &v
}

// GetPay returns the Pay field value if set, zero value otherwise.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetPay() int64 {
	if o == nil || IsNil(o.Pay) {
		var ret int64
		return ret
	}
	return *o.Pay
}

// GetPayOk returns a tuple with the Pay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) GetPayOk() (*int64, bool) {
	if o == nil || IsNil(o.Pay) {
		return nil, false
	}
	return o.Pay, true
}

// HasPay returns a boolean if a field has been set.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) HasPay() bool {
	if o != nil && !IsNil(o.Pay) {
		return true
	}

	return false
}

// SetPay gets a reference to the given int64 and assigns it to the Pay field.
func (o *GetKeyExecutives200ResponseKeyExecutivesInner) SetPay(v int64) {
	o.Pay = &v
}

func (o GetKeyExecutives200ResponseKeyExecutivesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetKeyExecutives200ResponseKeyExecutivesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !IsNil(o.YearBorn) {
		toSerialize["year_born"] = o.YearBorn
	}
	if !IsNil(o.Pay) {
		toSerialize["pay"] = o.Pay
	}
	return toSerialize, nil
}

func (o *GetKeyExecutives200ResponseKeyExecutivesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"title",
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

	varGetKeyExecutives200ResponseKeyExecutivesInner := _GetKeyExecutives200ResponseKeyExecutivesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetKeyExecutives200ResponseKeyExecutivesInner)

	if err != nil {
		return err
	}

	*o = GetKeyExecutives200ResponseKeyExecutivesInner(varGetKeyExecutives200ResponseKeyExecutivesInner)

	return err
}

type NullableGetKeyExecutives200ResponseKeyExecutivesInner struct {
	value *GetKeyExecutives200ResponseKeyExecutivesInner
	isSet bool
}

func (v NullableGetKeyExecutives200ResponseKeyExecutivesInner) Get() *GetKeyExecutives200ResponseKeyExecutivesInner {
	return v.value
}

func (v *NullableGetKeyExecutives200ResponseKeyExecutivesInner) Set(val *GetKeyExecutives200ResponseKeyExecutivesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKeyExecutives200ResponseKeyExecutivesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKeyExecutives200ResponseKeyExecutivesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKeyExecutives200ResponseKeyExecutivesInner(val *GetKeyExecutives200ResponseKeyExecutivesInner) *NullableGetKeyExecutives200ResponseKeyExecutivesInner {
	return &NullableGetKeyExecutives200ResponseKeyExecutivesInner{value: val, isSet: true}
}

func (v NullableGetKeyExecutives200ResponseKeyExecutivesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKeyExecutives200ResponseKeyExecutivesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
