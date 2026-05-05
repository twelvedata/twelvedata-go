// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetStocks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStocks200Response{}

// GetStocks200Response struct for GetStocks200Response
type GetStocks200Response struct {
	// Count
	Count int64 `json:"count"`
	// List of stock instruments
	Data []StocksResponseItem `json:"data"`
	// Response status
	Status string `json:"status"`
}

type _GetStocks200Response GetStocks200Response

// NewGetStocks200Response instantiates a new GetStocks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStocks200Response(count int64, data []StocksResponseItem, status string) *GetStocks200Response {
	this := GetStocks200Response{}
	this.Count = count
	this.Data = data
	this.Status = status
	return &this
}

// NewGetStocks200ResponseWithDefaults instantiates a new GetStocks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStocks200ResponseWithDefaults() *GetStocks200Response {
	this := GetStocks200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *GetStocks200Response) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetStocks200Response) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetStocks200Response) SetCount(v int64) {
	o.Count = v
}

// GetData returns the Data field value
func (o *GetStocks200Response) GetData() []StocksResponseItem {
	if o == nil {
		var ret []StocksResponseItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetStocks200Response) GetDataOk() ([]StocksResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetStocks200Response) SetData(v []StocksResponseItem) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *GetStocks200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetStocks200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetStocks200Response) SetStatus(v string) {
	o.Status = v
}

func (o GetStocks200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStocks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *GetStocks200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
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

	varGetStocks200Response := _GetStocks200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetStocks200Response)

	if err != nil {
		return err
	}

	*o = GetStocks200Response(varGetStocks200Response)

	return err
}

type NullableGetStocks200Response struct {
	value *GetStocks200Response
	isSet bool
}

func (v NullableGetStocks200Response) Get() *GetStocks200Response {
	return v.value
}

func (v *NullableGetStocks200Response) Set(val *GetStocks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStocks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStocks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStocks200Response(val *GetStocks200Response) *NullableGetStocks200Response {
	return &NullableGetStocks200Response{value: val, isSet: true}
}

func (v NullableGetStocks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStocks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
