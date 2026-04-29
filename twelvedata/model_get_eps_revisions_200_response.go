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

// checks if the GetEpsRevisions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEpsRevisions200Response{}

// GetEpsRevisions200Response struct for GetEpsRevisions200Response
type GetEpsRevisions200Response struct {
	Meta GetEarningsEstimate200ResponseMeta `json:"meta"`
	// EPS revision data
	EpsRevision []GetEpsRevisions200ResponseEpsRevisionInner `json:"eps_revision"`
	// Status of the response
	Status string `json:"status"`
}

type _GetEpsRevisions200Response GetEpsRevisions200Response

// NewGetEpsRevisions200Response instantiates a new GetEpsRevisions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEpsRevisions200Response(meta GetEarningsEstimate200ResponseMeta, epsRevision []GetEpsRevisions200ResponseEpsRevisionInner, status string) *GetEpsRevisions200Response {
	this := GetEpsRevisions200Response{}
	this.Meta = meta
	this.EpsRevision = epsRevision
	this.Status = status
	return &this
}

// NewGetEpsRevisions200ResponseWithDefaults instantiates a new GetEpsRevisions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEpsRevisions200ResponseWithDefaults() *GetEpsRevisions200Response {
	this := GetEpsRevisions200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetEpsRevisions200Response) GetMeta() GetEarningsEstimate200ResponseMeta {
	if o == nil {
		var ret GetEarningsEstimate200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetEpsRevisions200Response) SetMeta(v GetEarningsEstimate200ResponseMeta) {
	o.Meta = v
}

// GetEpsRevision returns the EpsRevision field value
func (o *GetEpsRevisions200Response) GetEpsRevision() []GetEpsRevisions200ResponseEpsRevisionInner {
	if o == nil {
		var ret []GetEpsRevisions200ResponseEpsRevisionInner
		return ret
	}

	return o.EpsRevision
}

// GetEpsRevisionOk returns a tuple with the EpsRevision field value
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200Response) GetEpsRevisionOk() ([]GetEpsRevisions200ResponseEpsRevisionInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EpsRevision, true
}

// SetEpsRevision sets field value
func (o *GetEpsRevisions200Response) SetEpsRevision(v []GetEpsRevisions200ResponseEpsRevisionInner) {
	o.EpsRevision = v
}

// GetStatus returns the Status field value
func (o *GetEpsRevisions200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetEpsRevisions200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetEpsRevisions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEpsRevisions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["eps_revision"] = o.EpsRevision
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetEpsRevisions200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"eps_revision",
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

	varGetEpsRevisions200Response := _GetEpsRevisions200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEpsRevisions200Response)

	if err != nil {
		return err
	}

	*o = GetEpsRevisions200Response(varGetEpsRevisions200Response)

	return err
}

type NullableGetEpsRevisions200Response struct {
	value *GetEpsRevisions200Response
	isSet bool
}

func (v NullableGetEpsRevisions200Response) Get() *GetEpsRevisions200Response {
	return v.value
}

func (v *NullableGetEpsRevisions200Response) Set(val *GetEpsRevisions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEpsRevisions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEpsRevisions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEpsRevisions200Response(val *GetEpsRevisions200Response) *NullableGetEpsRevisions200Response {
	return &NullableGetEpsRevisions200Response{value: val, isSet: true}
}

func (v NullableGetEpsRevisions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEpsRevisions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
