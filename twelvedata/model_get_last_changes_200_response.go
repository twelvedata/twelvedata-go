// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetLastChanges200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLastChanges200Response{}

// GetLastChanges200Response struct for GetLastChanges200Response
type GetLastChanges200Response struct {
	Pagination GetLastChanges200ResponsePagination `json:"pagination"`
	// Data contains the list of last changes
	Data []LastChangeResponseItem `json:"data"`
}

type _GetLastChanges200Response GetLastChanges200Response

// NewGetLastChanges200Response instantiates a new GetLastChanges200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLastChanges200Response(pagination GetLastChanges200ResponsePagination, data []LastChangeResponseItem) *GetLastChanges200Response {
	this := GetLastChanges200Response{}
	this.Pagination = pagination
	this.Data = data
	return &this
}

// NewGetLastChanges200ResponseWithDefaults instantiates a new GetLastChanges200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLastChanges200ResponseWithDefaults() *GetLastChanges200Response {
	this := GetLastChanges200Response{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *GetLastChanges200Response) GetPagination() GetLastChanges200ResponsePagination {
	if o == nil {
		var ret GetLastChanges200ResponsePagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetLastChanges200Response) GetPaginationOk() (*GetLastChanges200ResponsePagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetLastChanges200Response) SetPagination(v GetLastChanges200ResponsePagination) {
	o.Pagination = v
}

// GetData returns the Data field value
func (o *GetLastChanges200Response) GetData() []LastChangeResponseItem {
	if o == nil {
		var ret []LastChangeResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetLastChanges200Response) GetDataOk() ([]LastChangeResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetLastChanges200Response) SetData(v []LastChangeResponseItem) {
	o.Data = v
}

func (o GetLastChanges200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLastChanges200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetLastChanges200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pagination",
		"data",
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

	varGetLastChanges200Response := _GetLastChanges200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetLastChanges200Response)

	if err != nil {
		return err
	}

	*o = GetLastChanges200Response(varGetLastChanges200Response)

	return err
}

type NullableGetLastChanges200Response struct {
	value *GetLastChanges200Response
	isSet bool
}

func (v NullableGetLastChanges200Response) Get() *GetLastChanges200Response {
	return v.value
}

func (v *NullableGetLastChanges200Response) Set(val *GetLastChanges200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLastChanges200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLastChanges200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLastChanges200Response(val *GetLastChanges200Response) *NullableGetLastChanges200Response {
	return &NullableGetLastChanges200Response{value: val, isSet: true}
}

func (v NullableGetLastChanges200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLastChanges200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
