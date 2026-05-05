// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInner{}

// GetBalanceSheet200ResponseBalanceSheetInner struct for GetBalanceSheet200ResponseBalanceSheetInner
type GetBalanceSheet200ResponseBalanceSheetInner struct {
	// Date of fiscal period ending
	FiscalDate string `json:"fiscal_date"`
	// Fiscal year
	Year               *int64                                                         `json:"year,omitempty"`
	Assets             *GetBalanceSheet200ResponseBalanceSheetInnerAssets             `json:"assets,omitempty"`
	Liabilities        *GetBalanceSheet200ResponseBalanceSheetInnerLiabilities        `json:"liabilities,omitempty"`
	ShareholdersEquity *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity `json:"shareholders_equity,omitempty"`
}

type _GetBalanceSheet200ResponseBalanceSheetInner GetBalanceSheet200ResponseBalanceSheetInner

// NewGetBalanceSheet200ResponseBalanceSheetInner instantiates a new GetBalanceSheet200ResponseBalanceSheetInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInner(fiscalDate string) *GetBalanceSheet200ResponseBalanceSheetInner {
	this := GetBalanceSheet200ResponseBalanceSheetInner{}
	this.FiscalDate = fiscalDate
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInner {
	this := GetBalanceSheet200ResponseBalanceSheetInner{}
	return &this
}

// GetFiscalDate returns the FiscalDate field value
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetFiscalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiscalDate
}

// GetFiscalDateOk returns a tuple with the FiscalDate field value
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetFiscalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiscalDate, true
}

// SetFiscalDate sets field value
func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetFiscalDate(v string) {
	o.FiscalDate = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetYear(v int64) {
	o.Year = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetAssets() GetBalanceSheet200ResponseBalanceSheetInnerAssets {
	if o == nil || IsNil(o.Assets) {
		var ret GetBalanceSheet200ResponseBalanceSheetInnerAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetAssetsOk() (*GetBalanceSheet200ResponseBalanceSheetInnerAssets, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given GetBalanceSheet200ResponseBalanceSheetInnerAssets and assigns it to the Assets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetAssets(v GetBalanceSheet200ResponseBalanceSheetInnerAssets) {
	o.Assets = &v
}

// GetLiabilities returns the Liabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetLiabilities() GetBalanceSheet200ResponseBalanceSheetInnerLiabilities {
	if o == nil || IsNil(o.Liabilities) {
		var ret GetBalanceSheet200ResponseBalanceSheetInnerLiabilities
		return ret
	}
	return *o.Liabilities
}

// GetLiabilitiesOk returns a tuple with the Liabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetLiabilitiesOk() (*GetBalanceSheet200ResponseBalanceSheetInnerLiabilities, bool) {
	if o == nil || IsNil(o.Liabilities) {
		return nil, false
	}
	return o.Liabilities, true
}

// HasLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasLiabilities() bool {
	if o != nil && !IsNil(o.Liabilities) {
		return true
	}

	return false
}

// SetLiabilities gets a reference to the given GetBalanceSheet200ResponseBalanceSheetInnerLiabilities and assigns it to the Liabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetLiabilities(v GetBalanceSheet200ResponseBalanceSheetInnerLiabilities) {
	o.Liabilities = &v
}

// GetShareholdersEquity returns the ShareholdersEquity field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetShareholdersEquity() GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity {
	if o == nil || IsNil(o.ShareholdersEquity) {
		var ret GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity
		return ret
	}
	return *o.ShareholdersEquity
}

// GetShareholdersEquityOk returns a tuple with the ShareholdersEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetShareholdersEquityOk() (*GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity, bool) {
	if o == nil || IsNil(o.ShareholdersEquity) {
		return nil, false
	}
	return o.ShareholdersEquity, true
}

// HasShareholdersEquity returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasShareholdersEquity() bool {
	if o != nil && !IsNil(o.ShareholdersEquity) {
		return true
	}

	return false
}

// SetShareholdersEquity gets a reference to the given GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity and assigns it to the ShareholdersEquity field.
func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetShareholdersEquity(v GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) {
	o.ShareholdersEquity = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fiscal_date"] = o.FiscalDate
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.Liabilities) {
		toSerialize["liabilities"] = o.Liabilities
	}
	if !IsNil(o.ShareholdersEquity) {
		toSerialize["shareholders_equity"] = o.ShareholdersEquity
	}
	return toSerialize, nil
}

func (o *GetBalanceSheet200ResponseBalanceSheetInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fiscal_date",
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

	varGetBalanceSheet200ResponseBalanceSheetInner := _GetBalanceSheet200ResponseBalanceSheetInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetBalanceSheet200ResponseBalanceSheetInner)

	if err != nil {
		return err
	}

	*o = GetBalanceSheet200ResponseBalanceSheetInner(varGetBalanceSheet200ResponseBalanceSheetInner)

	return err
}

type NullableGetBalanceSheet200ResponseBalanceSheetInner struct {
	value *GetBalanceSheet200ResponseBalanceSheetInner
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInner) Get() *GetBalanceSheet200ResponseBalanceSheetInner {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInner) Set(val *GetBalanceSheet200ResponseBalanceSheetInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInner(val *GetBalanceSheet200ResponseBalanceSheetInner) *NullableGetBalanceSheet200ResponseBalanceSheetInner {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInner{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
