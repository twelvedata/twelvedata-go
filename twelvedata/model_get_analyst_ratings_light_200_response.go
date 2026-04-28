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

// checks if the GetAnalystRatingsLight200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAnalystRatingsLight200Response{}

// GetAnalystRatingsLight200Response struct for GetAnalystRatingsLight200Response
type GetAnalystRatingsLight200Response struct {
	Meta GetAnalystRatingsLight200ResponseMeta `json:"meta"`
	// List of analyst ratings
	Ratings []GetAnalystRatingsLight200ResponseRatingsInner `json:"ratings,omitempty"`
	// Response status
	Status string `json:"status"`
}

type _GetAnalystRatingsLight200Response GetAnalystRatingsLight200Response

// NewGetAnalystRatingsLight200Response instantiates a new GetAnalystRatingsLight200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnalystRatingsLight200Response(meta GetAnalystRatingsLight200ResponseMeta, status string) *GetAnalystRatingsLight200Response {
	this := GetAnalystRatingsLight200Response{}
	this.Meta = meta
	this.Status = status
	return &this
}

// NewGetAnalystRatingsLight200ResponseWithDefaults instantiates a new GetAnalystRatingsLight200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnalystRatingsLight200ResponseWithDefaults() *GetAnalystRatingsLight200Response {
	this := GetAnalystRatingsLight200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetAnalystRatingsLight200Response) GetMeta() GetAnalystRatingsLight200ResponseMeta {
	if o == nil {
		var ret GetAnalystRatingsLight200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200Response) GetMetaOk() (*GetAnalystRatingsLight200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetAnalystRatingsLight200Response) SetMeta(v GetAnalystRatingsLight200ResponseMeta) {
	o.Meta = v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *GetAnalystRatingsLight200Response) GetRatings() []GetAnalystRatingsLight200ResponseRatingsInner {
	if o == nil || IsNil(o.Ratings) {
		var ret []GetAnalystRatingsLight200ResponseRatingsInner
		return ret
	}
	return o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200Response) GetRatingsOk() ([]GetAnalystRatingsLight200ResponseRatingsInner, bool) {
	if o == nil || IsNil(o.Ratings) {
		return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *GetAnalystRatingsLight200Response) HasRatings() bool {
	if o != nil && !IsNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given []GetAnalystRatingsLight200ResponseRatingsInner and assigns it to the Ratings field.
func (o *GetAnalystRatingsLight200Response) SetRatings(v []GetAnalystRatingsLight200ResponseRatingsInner) {
	o.Ratings = v
}

// GetStatus returns the Status field value
func (o *GetAnalystRatingsLight200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetAnalystRatingsLight200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetAnalystRatingsLight200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetAnalystRatingsLight200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAnalystRatingsLight200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	if !IsNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetAnalystRatingsLight200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
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

	varGetAnalystRatingsLight200Response := _GetAnalystRatingsLight200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAnalystRatingsLight200Response)

	if err != nil {
		return err
	}

	*o = GetAnalystRatingsLight200Response(varGetAnalystRatingsLight200Response)

	return err
}

type NullableGetAnalystRatingsLight200Response struct {
	value *GetAnalystRatingsLight200Response
	isSet bool
}

func (v NullableGetAnalystRatingsLight200Response) Get() *GetAnalystRatingsLight200Response {
	return v.value
}

func (v *NullableGetAnalystRatingsLight200Response) Set(val *GetAnalystRatingsLight200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnalystRatingsLight200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnalystRatingsLight200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnalystRatingsLight200Response(val *GetAnalystRatingsLight200Response) *NullableGetAnalystRatingsLight200Response {
	return &NullableGetAnalystRatingsLight200Response{value: val, isSet: true}
}

func (v NullableGetAnalystRatingsLight200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnalystRatingsLight200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
