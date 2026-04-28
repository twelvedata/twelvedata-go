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

// checks if the GetSymbolSearch200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSymbolSearch200Response{}

// GetSymbolSearch200Response struct for GetSymbolSearch200Response
type GetSymbolSearch200Response struct {
	// List of symbols matching the search criteria
	Data []SymbolSearchResponseItem `json:"data"`
	// Status of the response
	Status string `json:"status"`
}

type _GetSymbolSearch200Response GetSymbolSearch200Response

// NewGetSymbolSearch200Response instantiates a new GetSymbolSearch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSymbolSearch200Response(data []SymbolSearchResponseItem, status string) *GetSymbolSearch200Response {
	this := GetSymbolSearch200Response{}
	this.Data = data
	this.Status = status
	return &this
}

// NewGetSymbolSearch200ResponseWithDefaults instantiates a new GetSymbolSearch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSymbolSearch200ResponseWithDefaults() *GetSymbolSearch200Response {
	this := GetSymbolSearch200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetSymbolSearch200Response) GetData() []SymbolSearchResponseItem {
	if o == nil {
		var ret []SymbolSearchResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSymbolSearch200Response) GetDataOk() ([]SymbolSearchResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetSymbolSearch200Response) SetData(v []SymbolSearchResponseItem) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetSymbolSearch200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetSymbolSearch200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetSymbolSearch200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetSymbolSearch200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSymbolSearch200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetSymbolSearch200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetSymbolSearch200Response := _GetSymbolSearch200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSymbolSearch200Response)

	if err != nil {
		return err
	}

	*o = GetSymbolSearch200Response(varGetSymbolSearch200Response)

	return err
}

type NullableGetSymbolSearch200Response struct {
	value *GetSymbolSearch200Response
	isSet bool
}

func (v NullableGetSymbolSearch200Response) Get() *GetSymbolSearch200Response {
	return v.value
}

func (v *NullableGetSymbolSearch200Response) Set(val *GetSymbolSearch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSymbolSearch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSymbolSearch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSymbolSearch200Response(val *GetSymbolSearch200Response) *NullableGetSymbolSearch200Response {
	return &NullableGetSymbolSearch200Response{value: val, isSet: true}
}

func (v NullableGetSymbolSearch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSymbolSearch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
