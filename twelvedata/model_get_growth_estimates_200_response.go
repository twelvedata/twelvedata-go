// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetGrowthEstimates200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGrowthEstimates200Response{}

// GetGrowthEstimates200Response struct for GetGrowthEstimates200Response
type GetGrowthEstimates200Response struct {
	Meta            GetEarningsEstimate200ResponseMeta            `json:"meta"`
	GrowthEstimates *GetGrowthEstimates200ResponseGrowthEstimates `json:"growth_estimates,omitempty"`
	// Status of the request
	Status string `json:"status"`
}

type _GetGrowthEstimates200Response GetGrowthEstimates200Response

// NewGetGrowthEstimates200Response instantiates a new GetGrowthEstimates200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGrowthEstimates200Response(meta GetEarningsEstimate200ResponseMeta, status string) *GetGrowthEstimates200Response {
	this := GetGrowthEstimates200Response{}
	this.Meta = meta
	this.Status = status
	return &this
}

// NewGetGrowthEstimates200ResponseWithDefaults instantiates a new GetGrowthEstimates200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGrowthEstimates200ResponseWithDefaults() *GetGrowthEstimates200Response {
	this := GetGrowthEstimates200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetGrowthEstimates200Response) GetMeta() GetEarningsEstimate200ResponseMeta {
	if o == nil {
		var ret GetEarningsEstimate200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetGrowthEstimates200Response) SetMeta(v GetEarningsEstimate200ResponseMeta) {
	o.Meta = v
}

// GetGrowthEstimates returns the GrowthEstimates field value if set, zero value otherwise.
func (o *GetGrowthEstimates200Response) GetGrowthEstimates() GetGrowthEstimates200ResponseGrowthEstimates {
	if o == nil || IsNil(o.GrowthEstimates) {
		var ret GetGrowthEstimates200ResponseGrowthEstimates
		return ret
	}
	return *o.GrowthEstimates
}

// GetGrowthEstimatesOk returns a tuple with the GrowthEstimates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200Response) GetGrowthEstimatesOk() (*GetGrowthEstimates200ResponseGrowthEstimates, bool) {
	if o == nil || IsNil(o.GrowthEstimates) {
		return nil, false
	}
	return o.GrowthEstimates, true
}

// HasGrowthEstimates returns a boolean if a field has been set.
func (o *GetGrowthEstimates200Response) HasGrowthEstimates() bool {
	if o != nil && !IsNil(o.GrowthEstimates) {
		return true
	}

	return false
}

// SetGrowthEstimates gets a reference to the given GetGrowthEstimates200ResponseGrowthEstimates and assigns it to the GrowthEstimates field.
func (o *GetGrowthEstimates200Response) SetGrowthEstimates(v GetGrowthEstimates200ResponseGrowthEstimates) {
	o.GrowthEstimates = &v
}

// GetStatus returns the Status field value
func (o *GetGrowthEstimates200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetGrowthEstimates200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetGrowthEstimates200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGrowthEstimates200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	if !IsNil(o.GrowthEstimates) {
		toSerialize["growth_estimates"] = o.GrowthEstimates
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetGrowthEstimates200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetGrowthEstimates200Response := _GetGrowthEstimates200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetGrowthEstimates200Response)

	if err != nil {
		return err
	}

	*o = GetGrowthEstimates200Response(varGetGrowthEstimates200Response)

	return err
}

type NullableGetGrowthEstimates200Response struct {
	value *GetGrowthEstimates200Response
	isSet bool
}

func (v NullableGetGrowthEstimates200Response) Get() *GetGrowthEstimates200Response {
	return v.value
}

func (v *NullableGetGrowthEstimates200Response) Set(val *GetGrowthEstimates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGrowthEstimates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGrowthEstimates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGrowthEstimates200Response(val *GetGrowthEstimates200Response) *NullableGetGrowthEstimates200Response {
	return &NullableGetGrowthEstimates200Response{value: val, isSet: true}
}

func (v NullableGetGrowthEstimates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGrowthEstimates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
