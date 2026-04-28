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

// checks if the MarketStateResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketStateResponseItem{}

// MarketStateResponseItem struct for MarketStateResponseItem
type MarketStateResponseItem struct {
	// The full name of exchange
	Name string `json:"name"`
	// Market identifier code (MIC) under ISO 10383 standard
	Code string `json:"code"`
	// Country where the exchange is located
	Country string `json:"country"`
	// Indicates if the market is currently open
	IsMarketOpen bool `json:"is_market_open"`
	// Time after market opening in <code>HH:MM:SS</code> format; if currently closed - returns <code>00:00:00</code>
	TimeAfterOpen string `json:"time_after_open"`
	// Time to market opening in <code>HH:MM:SS</code> format; if currently open - returns <code>00:00:00</code>
	TimeToOpen string `json:"time_to_open"`
	// Time to market closing in <code>HH:MM:SS</code> format; if currently closed - returns <code>00:00:00</code>
	TimeToClose string `json:"time_to_close"`
}

type _MarketStateResponseItem MarketStateResponseItem

// NewMarketStateResponseItem instantiates a new MarketStateResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketStateResponseItem(name string, code string, country string, isMarketOpen bool, timeAfterOpen string, timeToOpen string, timeToClose string) *MarketStateResponseItem {
	this := MarketStateResponseItem{}
	this.Name = name
	this.Code = code
	this.Country = country
	this.IsMarketOpen = isMarketOpen
	this.TimeAfterOpen = timeAfterOpen
	this.TimeToOpen = timeToOpen
	this.TimeToClose = timeToClose
	return &this
}

// NewMarketStateResponseItemWithDefaults instantiates a new MarketStateResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketStateResponseItemWithDefaults() *MarketStateResponseItem {
	this := MarketStateResponseItem{}
	return &this
}

// GetName returns the Name field value
func (o *MarketStateResponseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MarketStateResponseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MarketStateResponseItem) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *MarketStateResponseItem) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *MarketStateResponseItem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *MarketStateResponseItem) SetCode(v string) {
	o.Code = v
}

// GetCountry returns the Country field value
func (o *MarketStateResponseItem) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *MarketStateResponseItem) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *MarketStateResponseItem) SetCountry(v string) {
	o.Country = v
}

// GetIsMarketOpen returns the IsMarketOpen field value
func (o *MarketStateResponseItem) GetIsMarketOpen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMarketOpen
}

// GetIsMarketOpenOk returns a tuple with the IsMarketOpen field value
// and a boolean to check if the value has been set.
func (o *MarketStateResponseItem) GetIsMarketOpenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMarketOpen, true
}

// SetIsMarketOpen sets field value
func (o *MarketStateResponseItem) SetIsMarketOpen(v bool) {
	o.IsMarketOpen = v
}

// GetTimeAfterOpen returns the TimeAfterOpen field value
func (o *MarketStateResponseItem) GetTimeAfterOpen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeAfterOpen
}

// GetTimeAfterOpenOk returns a tuple with the TimeAfterOpen field value
// and a boolean to check if the value has been set.
func (o *MarketStateResponseItem) GetTimeAfterOpenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeAfterOpen, true
}

// SetTimeAfterOpen sets field value
func (o *MarketStateResponseItem) SetTimeAfterOpen(v string) {
	o.TimeAfterOpen = v
}

// GetTimeToOpen returns the TimeToOpen field value
func (o *MarketStateResponseItem) GetTimeToOpen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeToOpen
}

// GetTimeToOpenOk returns a tuple with the TimeToOpen field value
// and a boolean to check if the value has been set.
func (o *MarketStateResponseItem) GetTimeToOpenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeToOpen, true
}

// SetTimeToOpen sets field value
func (o *MarketStateResponseItem) SetTimeToOpen(v string) {
	o.TimeToOpen = v
}

// GetTimeToClose returns the TimeToClose field value
func (o *MarketStateResponseItem) GetTimeToClose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeToClose
}

// GetTimeToCloseOk returns a tuple with the TimeToClose field value
// and a boolean to check if the value has been set.
func (o *MarketStateResponseItem) GetTimeToCloseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeToClose, true
}

// SetTimeToClose sets field value
func (o *MarketStateResponseItem) SetTimeToClose(v string) {
	o.TimeToClose = v
}

func (o MarketStateResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketStateResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["code"] = o.Code
	toSerialize["country"] = o.Country
	toSerialize["is_market_open"] = o.IsMarketOpen
	toSerialize["time_after_open"] = o.TimeAfterOpen
	toSerialize["time_to_open"] = o.TimeToOpen
	toSerialize["time_to_close"] = o.TimeToClose
	return toSerialize, nil
}

func (o *MarketStateResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"code",
		"country",
		"is_market_open",
		"time_after_open",
		"time_to_open",
		"time_to_close",
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

	varMarketStateResponseItem := _MarketStateResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketStateResponseItem)

	if err != nil {
		return err
	}

	*o = MarketStateResponseItem(varMarketStateResponseItem)

	return err
}

type NullableMarketStateResponseItem struct {
	value *MarketStateResponseItem
	isSet bool
}

func (v NullableMarketStateResponseItem) Get() *MarketStateResponseItem {
	return v.value
}

func (v *NullableMarketStateResponseItem) Set(val *MarketStateResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketStateResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketStateResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketStateResponseItem(val *MarketStateResponseItem) *NullableMarketStateResponseItem {
	return &NullableMarketStateResponseItem{value: val, isSet: true}
}

func (v NullableMarketStateResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketStateResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
