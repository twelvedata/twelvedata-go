// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetLastChanges200ResponsePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLastChanges200ResponsePagination{}

// GetLastChanges200ResponsePagination Pagination information
type GetLastChanges200ResponsePagination struct {
	// Current page number
	CurrentPage int64 `json:"current_page"`
	// Records per page
	PerPage int64 `json:"per_page"`
}

type _GetLastChanges200ResponsePagination GetLastChanges200ResponsePagination

// NewGetLastChanges200ResponsePagination instantiates a new GetLastChanges200ResponsePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLastChanges200ResponsePagination(currentPage int64, perPage int64) *GetLastChanges200ResponsePagination {
	this := GetLastChanges200ResponsePagination{}
	this.CurrentPage = currentPage
	this.PerPage = perPage
	return &this
}

// NewGetLastChanges200ResponsePaginationWithDefaults instantiates a new GetLastChanges200ResponsePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLastChanges200ResponsePaginationWithDefaults() *GetLastChanges200ResponsePagination {
	this := GetLastChanges200ResponsePagination{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value
func (o *GetLastChanges200ResponsePagination) GetCurrentPage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *GetLastChanges200ResponsePagination) GetCurrentPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value
func (o *GetLastChanges200ResponsePagination) SetCurrentPage(v int64) {
	o.CurrentPage = v
}

// GetPerPage returns the PerPage field value
func (o *GetLastChanges200ResponsePagination) GetPerPage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *GetLastChanges200ResponsePagination) GetPerPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *GetLastChanges200ResponsePagination) SetPerPage(v int64) {
	o.PerPage = v
}

func (o GetLastChanges200ResponsePagination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLastChanges200ResponsePagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current_page"] = o.CurrentPage
	toSerialize["per_page"] = o.PerPage
	return toSerialize, nil
}

func (o *GetLastChanges200ResponsePagination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_page",
		"per_page",
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

	varGetLastChanges200ResponsePagination := _GetLastChanges200ResponsePagination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetLastChanges200ResponsePagination)

	if err != nil {
		return err
	}

	*o = GetLastChanges200ResponsePagination(varGetLastChanges200ResponsePagination)

	return err
}

type NullableGetLastChanges200ResponsePagination struct {
	value *GetLastChanges200ResponsePagination
	isSet bool
}

func (v NullableGetLastChanges200ResponsePagination) Get() *GetLastChanges200ResponsePagination {
	return v.value
}

func (v *NullableGetLastChanges200ResponsePagination) Set(val *GetLastChanges200ResponsePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLastChanges200ResponsePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLastChanges200ResponsePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLastChanges200ResponsePagination(val *GetLastChanges200ResponsePagination) *NullableGetLastChanges200ResponsePagination {
	return &NullableGetLastChanges200ResponsePagination{value: val, isSet: true}
}

func (v NullableGetLastChanges200ResponsePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLastChanges200ResponsePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
