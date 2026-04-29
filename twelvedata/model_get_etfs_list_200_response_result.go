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

// checks if the GetETFsList200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsList200ResponseResult{}

// GetETFsList200ResponseResult struct for GetETFsList200ResponseResult
type GetETFsList200ResponseResult struct {
	// Total number of matching funds
	Count int64 `json:"count"`
	// List of ETFs
	List []ETFsListResponseItem `json:"list"`
}

type _GetETFsList200ResponseResult GetETFsList200ResponseResult

// NewGetETFsList200ResponseResult instantiates a new GetETFsList200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsList200ResponseResult(count int64, list []ETFsListResponseItem) *GetETFsList200ResponseResult {
	this := GetETFsList200ResponseResult{}
	this.Count = count
	this.List = list
	return &this
}

// NewGetETFsList200ResponseResultWithDefaults instantiates a new GetETFsList200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsList200ResponseResultWithDefaults() *GetETFsList200ResponseResult {
	this := GetETFsList200ResponseResult{}
	return &this
}

// GetCount returns the Count field value
func (o *GetETFsList200ResponseResult) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetETFsList200ResponseResult) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetETFsList200ResponseResult) SetCount(v int64) {
	o.Count = v
}

// GetList returns the List field value
func (o *GetETFsList200ResponseResult) GetList() []ETFsListResponseItem {
	if o == nil {
		var ret []ETFsListResponseItem
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *GetETFsList200ResponseResult) GetListOk() ([]ETFsListResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *GetETFsList200ResponseResult) SetList(v []ETFsListResponseItem) {
	o.List = v
}

func (o GetETFsList200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsList200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["list"] = o.List
	return toSerialize, nil
}

func (o *GetETFsList200ResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"list",
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

	varGetETFsList200ResponseResult := _GetETFsList200ResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetETFsList200ResponseResult)

	if err != nil {
		return err
	}

	*o = GetETFsList200ResponseResult(varGetETFsList200ResponseResult)

	return err
}

type NullableGetETFsList200ResponseResult struct {
	value *GetETFsList200ResponseResult
	isSet bool
}

func (v NullableGetETFsList200ResponseResult) Get() *GetETFsList200ResponseResult {
	return v.value
}

func (v *NullableGetETFsList200ResponseResult) Set(val *GetETFsList200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsList200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsList200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsList200ResponseResult(val *GetETFsList200ResponseResult) *NullableGetETFsList200ResponseResult {
	return &NullableGetETFsList200ResponseResult{value: val, isSet: true}
}

func (v NullableGetETFsList200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsList200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
