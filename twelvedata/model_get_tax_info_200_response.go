// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTaxInfo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTaxInfo200Response{}

// GetTaxInfo200Response struct for GetTaxInfo200Response
type GetTaxInfo200Response struct {
	Meta GetTaxInfo200ResponseMeta `json:"meta"`
	Data GetTaxInfo200ResponseData `json:"data"`
	// The status of the request, e.g., `ok`, `error`
	Status string `json:"status"`
}

type _GetTaxInfo200Response GetTaxInfo200Response

// NewGetTaxInfo200Response instantiates a new GetTaxInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaxInfo200Response(meta GetTaxInfo200ResponseMeta, data GetTaxInfo200ResponseData, status string) *GetTaxInfo200Response {
	this := GetTaxInfo200Response{}
	this.Meta = meta
	this.Data = data
	this.Status = status
	return &this
}

// NewGetTaxInfo200ResponseWithDefaults instantiates a new GetTaxInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaxInfo200ResponseWithDefaults() *GetTaxInfo200Response {
	this := GetTaxInfo200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetTaxInfo200Response) GetMeta() GetTaxInfo200ResponseMeta {
	if o == nil {
		var ret GetTaxInfo200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200Response) GetMetaOk() (*GetTaxInfo200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetTaxInfo200Response) SetMeta(v GetTaxInfo200ResponseMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *GetTaxInfo200Response) GetData() GetTaxInfo200ResponseData {
	if o == nil {
		var ret GetTaxInfo200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200Response) GetDataOk() (*GetTaxInfo200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetTaxInfo200Response) SetData(v GetTaxInfo200ResponseData) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetTaxInfo200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTaxInfo200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTaxInfo200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetTaxInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTaxInfo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetTaxInfo200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
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

	varGetTaxInfo200Response := _GetTaxInfo200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTaxInfo200Response)

	if err != nil {
		return err
	}

	*o = GetTaxInfo200Response(varGetTaxInfo200Response)

	return err
}

type NullableGetTaxInfo200Response struct {
	value *GetTaxInfo200Response
	isSet bool
}

func (v NullableGetTaxInfo200Response) Get() *GetTaxInfo200Response {
	return v.value
}

func (v *NullableGetTaxInfo200Response) Set(val *GetTaxInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaxInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaxInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaxInfo200Response(val *GetTaxInfo200Response) *NullableGetTaxInfo200Response {
	return &NullableGetTaxInfo200Response{value: val, isSet: true}
}

func (v NullableGetTaxInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaxInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
