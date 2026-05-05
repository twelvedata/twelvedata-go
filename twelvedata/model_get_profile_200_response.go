// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetProfile200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProfile200Response{}

// GetProfile200Response struct for GetProfile200Response
type GetProfile200Response struct {
	// Ticker of the company
	Symbol string `json:"symbol"`
	// Name of the company
	Name string `json:"name"`
	// Exchange name where the company is listed
	Exchange string `json:"exchange"`
	// Market Identifier Code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Sector at which the company operates
	Sector *string `json:"sector,omitempty"`
	// Industry at which company operates
	Industry *string `json:"industry,omitempty"`
	// Number of employees in the company
	Employees *int64 `json:"employees,omitempty"`
	// Website of the company
	Website *string `json:"website,omitempty"`
	// Description of the company activities
	Description *string `json:"description,omitempty"`
	// Issue type of the stock
	Type *string `json:"type,omitempty"`
	// Name of the CEO of the company
	CEO *string `json:"CEO,omitempty"`
	// Street address of the company if presented
	Address *string `json:"address,omitempty"`
	// Secondary address of the company if presented
	Address2 *string `json:"address2,omitempty"`
	// City of the company if presented
	City *string `json:"city,omitempty"`
	// ZIP code of the company if presented
	Zip *string `json:"zip,omitempty"`
	// State of the company if presented
	State *string `json:"state,omitempty"`
	// Country of the company if presented
	Country *string `json:"country,omitempty"`
	// Phone number of the company if presented
	Phone *string `json:"phone,omitempty"`
}

type _GetProfile200Response GetProfile200Response

// NewGetProfile200Response instantiates a new GetProfile200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfile200Response(symbol string, name string, exchange string, micCode string) *GetProfile200Response {
	this := GetProfile200Response{}
	this.Symbol = symbol
	this.Name = name
	this.Exchange = exchange
	this.MicCode = micCode
	return &this
}

// NewGetProfile200ResponseWithDefaults instantiates a new GetProfile200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfile200ResponseWithDefaults() *GetProfile200Response {
	this := GetProfile200Response{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetProfile200Response) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetProfile200Response) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetProfile200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetProfile200Response) SetName(v string) {
	o.Name = v
}

// GetExchange returns the Exchange field value
func (o *GetProfile200Response) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetProfile200Response) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetProfile200Response) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetProfile200Response) SetMicCode(v string) {
	o.MicCode = v
}

// GetSector returns the Sector field value if set, zero value otherwise.
func (o *GetProfile200Response) GetSector() string {
	if o == nil || IsNil(o.Sector) {
		var ret string
		return ret
	}
	return *o.Sector
}

// GetSectorOk returns a tuple with the Sector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetSectorOk() (*string, bool) {
	if o == nil || IsNil(o.Sector) {
		return nil, false
	}
	return o.Sector, true
}

// HasSector returns a boolean if a field has been set.
func (o *GetProfile200Response) HasSector() bool {
	if o != nil && !IsNil(o.Sector) {
		return true
	}

	return false
}

// SetSector gets a reference to the given string and assigns it to the Sector field.
func (o *GetProfile200Response) SetSector(v string) {
	o.Sector = &v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *GetProfile200Response) GetIndustry() string {
	if o == nil || IsNil(o.Industry) {
		var ret string
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetIndustryOk() (*string, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *GetProfile200Response) HasIndustry() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given string and assigns it to the Industry field.
func (o *GetProfile200Response) SetIndustry(v string) {
	o.Industry = &v
}

// GetEmployees returns the Employees field value if set, zero value otherwise.
func (o *GetProfile200Response) GetEmployees() int64 {
	if o == nil || IsNil(o.Employees) {
		var ret int64
		return ret
	}
	return *o.Employees
}

// GetEmployeesOk returns a tuple with the Employees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetEmployeesOk() (*int64, bool) {
	if o == nil || IsNil(o.Employees) {
		return nil, false
	}
	return o.Employees, true
}

// HasEmployees returns a boolean if a field has been set.
func (o *GetProfile200Response) HasEmployees() bool {
	if o != nil && !IsNil(o.Employees) {
		return true
	}

	return false
}

// SetEmployees gets a reference to the given int64 and assigns it to the Employees field.
func (o *GetProfile200Response) SetEmployees(v int64) {
	o.Employees = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *GetProfile200Response) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *GetProfile200Response) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *GetProfile200Response) SetWebsite(v string) {
	o.Website = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetProfile200Response) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetProfile200Response) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetProfile200Response) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetProfile200Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetProfile200Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetProfile200Response) SetType(v string) {
	o.Type = &v
}

// GetCEO returns the CEO field value if set, zero value otherwise.
func (o *GetProfile200Response) GetCEO() string {
	if o == nil || IsNil(o.CEO) {
		var ret string
		return ret
	}
	return *o.CEO
}

// GetCEOOk returns a tuple with the CEO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetCEOOk() (*string, bool) {
	if o == nil || IsNil(o.CEO) {
		return nil, false
	}
	return o.CEO, true
}

// HasCEO returns a boolean if a field has been set.
func (o *GetProfile200Response) HasCEO() bool {
	if o != nil && !IsNil(o.CEO) {
		return true
	}

	return false
}

// SetCEO gets a reference to the given string and assigns it to the CEO field.
func (o *GetProfile200Response) SetCEO(v string) {
	o.CEO = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetProfile200Response) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetProfile200Response) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetProfile200Response) SetAddress(v string) {
	o.Address = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *GetProfile200Response) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *GetProfile200Response) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *GetProfile200Response) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *GetProfile200Response) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *GetProfile200Response) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *GetProfile200Response) SetCity(v string) {
	o.City = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *GetProfile200Response) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *GetProfile200Response) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *GetProfile200Response) SetZip(v string) {
	o.Zip = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetProfile200Response) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetProfile200Response) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetProfile200Response) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetProfile200Response) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetProfile200Response) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetProfile200Response) SetCountry(v string) {
	o.Country = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *GetProfile200Response) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfile200Response) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *GetProfile200Response) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *GetProfile200Response) SetPhone(v string) {
	o.Phone = &v
}

func (o GetProfile200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProfile200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	if !IsNil(o.Sector) {
		toSerialize["sector"] = o.Sector
	}
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	if !IsNil(o.Employees) {
		toSerialize["employees"] = o.Employees
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CEO) {
		toSerialize["CEO"] = o.CEO
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}

func (o *GetProfile200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"exchange",
		"mic_code",
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

	varGetProfile200Response := _GetProfile200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetProfile200Response)

	if err != nil {
		return err
	}

	*o = GetProfile200Response(varGetProfile200Response)

	return err
}

type NullableGetProfile200Response struct {
	value *GetProfile200Response
	isSet bool
}

func (v NullableGetProfile200Response) Get() *GetProfile200Response {
	return v.value
}

func (v *NullableGetProfile200Response) Set(val *GetProfile200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfile200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfile200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfile200Response(val *GetProfile200Response) *NullableGetProfile200Response {
	return &NullableGetProfile200Response{value: val, isSet: true}
}

func (v NullableGetProfile200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfile200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
