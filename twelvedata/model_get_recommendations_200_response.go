// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetRecommendations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecommendations200Response{}

// GetRecommendations200Response struct for GetRecommendations200Response
type GetRecommendations200Response struct {
	Meta   GetRecommendations200ResponseMeta   `json:"meta"`
	Trends GetRecommendations200ResponseTrends `json:"trends"`
	// Rating from 0 to 10 represents overall analysts' recommendation. 0 to 2 - strong sell, 2 to 4 - sell, 4 to 6 - hold, 6 to 8 - buy, 8 to 10 - strong buy.
	Rating *float64 `json:"rating,omitempty"`
	// Response status
	Status string `json:"status"`
}

type _GetRecommendations200Response GetRecommendations200Response

// NewGetRecommendations200Response instantiates a new GetRecommendations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecommendations200Response(meta GetRecommendations200ResponseMeta, trends GetRecommendations200ResponseTrends, status string) *GetRecommendations200Response {
	this := GetRecommendations200Response{}
	this.Meta = meta
	this.Trends = trends
	this.Status = status
	return &this
}

// NewGetRecommendations200ResponseWithDefaults instantiates a new GetRecommendations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecommendations200ResponseWithDefaults() *GetRecommendations200Response {
	this := GetRecommendations200Response{}
	return &this
}

// GetMeta returns the Meta field value
func (o *GetRecommendations200Response) GetMeta() GetRecommendations200ResponseMeta {
	if o == nil {
		var ret GetRecommendations200ResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetRecommendations200Response) GetMetaOk() (*GetRecommendations200ResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetRecommendations200Response) SetMeta(v GetRecommendations200ResponseMeta) {
	o.Meta = v
}

// GetTrends returns the Trends field value
func (o *GetRecommendations200Response) GetTrends() GetRecommendations200ResponseTrends {
	if o == nil {
		var ret GetRecommendations200ResponseTrends
		return ret
	}

	return o.Trends
}

// GetTrendsOk returns a tuple with the Trends field value
// and a boolean to check if the value has been set.
func (o *GetRecommendations200Response) GetTrendsOk() (*GetRecommendations200ResponseTrends, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trends, true
}

// SetTrends sets field value
func (o *GetRecommendations200Response) SetTrends(v GetRecommendations200ResponseTrends) {
	o.Trends = v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *GetRecommendations200Response) GetRating() float64 {
	if o == nil || IsNil(o.Rating) {
		var ret float64
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecommendations200Response) GetRatingOk() (*float64, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *GetRecommendations200Response) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given float64 and assigns it to the Rating field.
func (o *GetRecommendations200Response) SetRating(v float64) {
	o.Rating = &v
}

// GetStatus returns the Status field value
func (o *GetRecommendations200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetRecommendations200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetRecommendations200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetRecommendations200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecommendations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["trends"] = o.Trends
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetRecommendations200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"trends",
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

	varGetRecommendations200Response := _GetRecommendations200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetRecommendations200Response)

	if err != nil {
		return err
	}

	*o = GetRecommendations200Response(varGetRecommendations200Response)

	return err
}

type NullableGetRecommendations200Response struct {
	value *GetRecommendations200Response
	isSet bool
}

func (v NullableGetRecommendations200Response) Get() *GetRecommendations200Response {
	return v.value
}

func (v *NullableGetRecommendations200Response) Set(val *GetRecommendations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecommendations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecommendations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecommendations200Response(val *GetRecommendations200Response) *NullableGetRecommendations200Response {
	return &NullableGetRecommendations200Response{value: val, isSet: true}
}

func (v NullableGetRecommendations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecommendations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
