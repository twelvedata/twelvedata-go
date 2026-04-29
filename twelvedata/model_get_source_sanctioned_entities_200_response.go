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

// checks if the GetSourceSanctionedEntities200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSourceSanctionedEntities200Response{}

// GetSourceSanctionedEntities200Response struct for GetSourceSanctionedEntities200Response
type GetSourceSanctionedEntities200Response struct {
	// List of sanctioned entities
	Sanctions []ResponseSanctionedEntitiy `json:"sanctions"`
	// Total number of sanctioned entities
	Count int64 `json:"count"`
	// Response status
	Status string `json:"status"`
}

type _GetSourceSanctionedEntities200Response GetSourceSanctionedEntities200Response

// NewGetSourceSanctionedEntities200Response instantiates a new GetSourceSanctionedEntities200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourceSanctionedEntities200Response(sanctions []ResponseSanctionedEntitiy, count int64, status string) *GetSourceSanctionedEntities200Response {
	this := GetSourceSanctionedEntities200Response{}
	this.Sanctions = sanctions
	this.Count = count
	this.Status = status
	return &this
}

// NewGetSourceSanctionedEntities200ResponseWithDefaults instantiates a new GetSourceSanctionedEntities200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourceSanctionedEntities200ResponseWithDefaults() *GetSourceSanctionedEntities200Response {
	this := GetSourceSanctionedEntities200Response{}
	return &this
}

// GetSanctions returns the Sanctions field value
func (o *GetSourceSanctionedEntities200Response) GetSanctions() []ResponseSanctionedEntitiy {
	if o == nil {
		var ret []ResponseSanctionedEntitiy
		return ret
	}

	return o.Sanctions
}

// GetSanctionsOk returns a tuple with the Sanctions field value
// and a boolean to check if the value has been set.
func (o *GetSourceSanctionedEntities200Response) GetSanctionsOk() ([]ResponseSanctionedEntitiy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sanctions, true
}

// SetSanctions sets field value
func (o *GetSourceSanctionedEntities200Response) SetSanctions(v []ResponseSanctionedEntitiy) {
	o.Sanctions = v
}

// GetCount returns the Count field value
func (o *GetSourceSanctionedEntities200Response) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetSourceSanctionedEntities200Response) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetSourceSanctionedEntities200Response) SetCount(v int64) {
	o.Count = v
}

// GetStatus returns the Status field value
func (o *GetSourceSanctionedEntities200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetSourceSanctionedEntities200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetSourceSanctionedEntities200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetSourceSanctionedEntities200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSourceSanctionedEntities200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sanctions"] = o.Sanctions
	toSerialize["count"] = o.Count
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetSourceSanctionedEntities200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sanctions",
		"count",
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

	varGetSourceSanctionedEntities200Response := _GetSourceSanctionedEntities200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetSourceSanctionedEntities200Response)

	if err != nil {
		return err
	}

	*o = GetSourceSanctionedEntities200Response(varGetSourceSanctionedEntities200Response)

	return err
}

type NullableGetSourceSanctionedEntities200Response struct {
	value *GetSourceSanctionedEntities200Response
	isSet bool
}

func (v NullableGetSourceSanctionedEntities200Response) Get() *GetSourceSanctionedEntities200Response {
	return v.value
}

func (v *NullableGetSourceSanctionedEntities200Response) Set(val *GetSourceSanctionedEntities200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourceSanctionedEntities200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourceSanctionedEntities200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourceSanctionedEntities200Response(val *GetSourceSanctionedEntities200Response) *NullableGetSourceSanctionedEntities200Response {
	return &NullableGetSourceSanctionedEntities200Response{value: val, isSet: true}
}

func (v NullableGetSourceSanctionedEntities200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourceSanctionedEntities200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
