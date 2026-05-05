// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EdgarFilingValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgarFilingValue{}

// EdgarFilingValue Filing value object
type EdgarFilingValue struct {
	// CIK code
	Cik int64 `json:"cik"`
	// Filing date in UNIX timestamp
	FiledAt int64 `json:"filed_at"`
	// Filing files
	Files []EdgarFilingFile `json:"files,omitempty"`
	// Full URL of the filing
	FilingUrl string `json:"filing_url"`
	// Filing form type
	FormType string `json:"form_type"`
	// Ticker symbols associated with the filing
	Ticker []string `json:"ticker"`
}

type _EdgarFilingValue EdgarFilingValue

// NewEdgarFilingValue instantiates a new EdgarFilingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgarFilingValue(cik int64, filedAt int64, filingUrl string, formType string, ticker []string) *EdgarFilingValue {
	this := EdgarFilingValue{}
	this.Cik = cik
	this.FiledAt = filedAt
	this.FilingUrl = filingUrl
	this.FormType = formType
	this.Ticker = ticker
	return &this
}

// NewEdgarFilingValueWithDefaults instantiates a new EdgarFilingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgarFilingValueWithDefaults() *EdgarFilingValue {
	this := EdgarFilingValue{}
	return &this
}

// GetCik returns the Cik field value
func (o *EdgarFilingValue) GetCik() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Cik
}

// GetCikOk returns a tuple with the Cik field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingValue) GetCikOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cik, true
}

// SetCik sets field value
func (o *EdgarFilingValue) SetCik(v int64) {
	o.Cik = v
}

// GetFiledAt returns the FiledAt field value
func (o *EdgarFilingValue) GetFiledAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FiledAt
}

// GetFiledAtOk returns a tuple with the FiledAt field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingValue) GetFiledAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiledAt, true
}

// SetFiledAt sets field value
func (o *EdgarFilingValue) SetFiledAt(v int64) {
	o.FiledAt = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *EdgarFilingValue) GetFiles() []EdgarFilingFile {
	if o == nil || IsNil(o.Files) {
		var ret []EdgarFilingFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgarFilingValue) GetFilesOk() ([]EdgarFilingFile, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *EdgarFilingValue) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []EdgarFilingFile and assigns it to the Files field.
func (o *EdgarFilingValue) SetFiles(v []EdgarFilingFile) {
	o.Files = v
}

// GetFilingUrl returns the FilingUrl field value
func (o *EdgarFilingValue) GetFilingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilingUrl
}

// GetFilingUrlOk returns a tuple with the FilingUrl field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingValue) GetFilingUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilingUrl, true
}

// SetFilingUrl sets field value
func (o *EdgarFilingValue) SetFilingUrl(v string) {
	o.FilingUrl = v
}

// GetFormType returns the FormType field value
func (o *EdgarFilingValue) GetFormType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormType
}

// GetFormTypeOk returns a tuple with the FormType field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingValue) GetFormTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormType, true
}

// SetFormType sets field value
func (o *EdgarFilingValue) SetFormType(v string) {
	o.FormType = v
}

// GetTicker returns the Ticker field value
func (o *EdgarFilingValue) GetTicker() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value
// and a boolean to check if the value has been set.
func (o *EdgarFilingValue) GetTickerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ticker, true
}

// SetTicker sets field value
func (o *EdgarFilingValue) SetTicker(v []string) {
	o.Ticker = v
}

func (o EdgarFilingValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgarFilingValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cik"] = o.Cik
	toSerialize["filed_at"] = o.FiledAt
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	toSerialize["filing_url"] = o.FilingUrl
	toSerialize["form_type"] = o.FormType
	toSerialize["ticker"] = o.Ticker
	return toSerialize, nil
}

func (o *EdgarFilingValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cik",
		"filed_at",
		"filing_url",
		"form_type",
		"ticker",
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

	varEdgarFilingValue := _EdgarFilingValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEdgarFilingValue)

	if err != nil {
		return err
	}

	*o = EdgarFilingValue(varEdgarFilingValue)

	return err
}

type NullableEdgarFilingValue struct {
	value *EdgarFilingValue
	isSet bool
}

func (v NullableEdgarFilingValue) Get() *EdgarFilingValue {
	return v.value
}

func (v *NullableEdgarFilingValue) Set(val *EdgarFilingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgarFilingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgarFilingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgarFilingValue(val *EdgarFilingValue) *NullableEdgarFilingValue {
	return &NullableEdgarFilingValue{value: val, isSet: true}
}

func (v NullableEdgarFilingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgarFilingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
