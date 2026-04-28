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

// checks if the GetTechnicalIndicators200ResponseDataValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTechnicalIndicators200ResponseDataValue{}

// GetTechnicalIndicators200ResponseDataValue struct for GetTechnicalIndicators200ResponseDataValue
type GetTechnicalIndicators200ResponseDataValue struct {
	// If the indicator is tested, approved and is recommended for use returns <code>true</code>, otherwise returns <code>false</code>
	Enable bool `json:"enable"`
	// Full indicator name
	FullName string `json:"full_name"`
	// Brief description of the indicator
	Description string `json:"description"`
	// Group to which indicator belongs to
	Type string `json:"type"`
	// If indicator should be plotted over price bars returns <code>true</code>, otherwise returns <code>false</code>
	Overlay      bool                                         `json:"overlay"`
	OutputValues *TechnicalIndicatorsResponseMacdOutputValues `json:"output_values,omitempty"`
	Parameters   *TechnicalIndicatorsResponseMacdParameters   `json:"parameters,omitempty"`
	Tinting      *TechnicalIndicatorsResponseMacdTinting      `json:"tinting,omitempty"`
}

type _GetTechnicalIndicators200ResponseDataValue GetTechnicalIndicators200ResponseDataValue

// NewGetTechnicalIndicators200ResponseDataValue instantiates a new GetTechnicalIndicators200ResponseDataValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTechnicalIndicators200ResponseDataValue(enable bool, fullName string, description string, type_ string, overlay bool) *GetTechnicalIndicators200ResponseDataValue {
	this := GetTechnicalIndicators200ResponseDataValue{}
	this.Enable = enable
	this.FullName = fullName
	this.Description = description
	this.Type = type_
	this.Overlay = overlay
	return &this
}

// NewGetTechnicalIndicators200ResponseDataValueWithDefaults instantiates a new GetTechnicalIndicators200ResponseDataValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTechnicalIndicators200ResponseDataValueWithDefaults() *GetTechnicalIndicators200ResponseDataValue {
	this := GetTechnicalIndicators200ResponseDataValue{}
	return &this
}

// GetEnable returns the Enable field value
func (o *GetTechnicalIndicators200ResponseDataValue) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *GetTechnicalIndicators200ResponseDataValue) SetEnable(v bool) {
	o.Enable = v
}

// GetFullName returns the FullName field value
func (o *GetTechnicalIndicators200ResponseDataValue) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *GetTechnicalIndicators200ResponseDataValue) SetFullName(v string) {
	o.FullName = v
}

// GetDescription returns the Description field value
func (o *GetTechnicalIndicators200ResponseDataValue) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetTechnicalIndicators200ResponseDataValue) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *GetTechnicalIndicators200ResponseDataValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTechnicalIndicators200ResponseDataValue) SetType(v string) {
	o.Type = v
}

// GetOverlay returns the Overlay field value
func (o *GetTechnicalIndicators200ResponseDataValue) GetOverlay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Overlay
}

// GetOverlayOk returns a tuple with the Overlay field value
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetOverlayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overlay, true
}

// SetOverlay sets field value
func (o *GetTechnicalIndicators200ResponseDataValue) SetOverlay(v bool) {
	o.Overlay = v
}

// GetOutputValues returns the OutputValues field value if set, zero value otherwise.
func (o *GetTechnicalIndicators200ResponseDataValue) GetOutputValues() TechnicalIndicatorsResponseMacdOutputValues {
	if o == nil || IsNil(o.OutputValues) {
		var ret TechnicalIndicatorsResponseMacdOutputValues
		return ret
	}
	return *o.OutputValues
}

// GetOutputValuesOk returns a tuple with the OutputValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetOutputValuesOk() (*TechnicalIndicatorsResponseMacdOutputValues, bool) {
	if o == nil || IsNil(o.OutputValues) {
		return nil, false
	}
	return o.OutputValues, true
}

// HasOutputValues returns a boolean if a field has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) HasOutputValues() bool {
	if o != nil && !IsNil(o.OutputValues) {
		return true
	}

	return false
}

// SetOutputValues gets a reference to the given TechnicalIndicatorsResponseMacdOutputValues and assigns it to the OutputValues field.
func (o *GetTechnicalIndicators200ResponseDataValue) SetOutputValues(v TechnicalIndicatorsResponseMacdOutputValues) {
	o.OutputValues = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *GetTechnicalIndicators200ResponseDataValue) GetParameters() TechnicalIndicatorsResponseMacdParameters {
	if o == nil || IsNil(o.Parameters) {
		var ret TechnicalIndicatorsResponseMacdParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetParametersOk() (*TechnicalIndicatorsResponseMacdParameters, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given TechnicalIndicatorsResponseMacdParameters and assigns it to the Parameters field.
func (o *GetTechnicalIndicators200ResponseDataValue) SetParameters(v TechnicalIndicatorsResponseMacdParameters) {
	o.Parameters = &v
}

// GetTinting returns the Tinting field value if set, zero value otherwise.
func (o *GetTechnicalIndicators200ResponseDataValue) GetTinting() TechnicalIndicatorsResponseMacdTinting {
	if o == nil || IsNil(o.Tinting) {
		var ret TechnicalIndicatorsResponseMacdTinting
		return ret
	}
	return *o.Tinting
}

// GetTintingOk returns a tuple with the Tinting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) GetTintingOk() (*TechnicalIndicatorsResponseMacdTinting, bool) {
	if o == nil || IsNil(o.Tinting) {
		return nil, false
	}
	return o.Tinting, true
}

// HasTinting returns a boolean if a field has been set.
func (o *GetTechnicalIndicators200ResponseDataValue) HasTinting() bool {
	if o != nil && !IsNil(o.Tinting) {
		return true
	}

	return false
}

// SetTinting gets a reference to the given TechnicalIndicatorsResponseMacdTinting and assigns it to the Tinting field.
func (o *GetTechnicalIndicators200ResponseDataValue) SetTinting(v TechnicalIndicatorsResponseMacdTinting) {
	o.Tinting = &v
}

func (o GetTechnicalIndicators200ResponseDataValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTechnicalIndicators200ResponseDataValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable"] = o.Enable
	toSerialize["full_name"] = o.FullName
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	toSerialize["overlay"] = o.Overlay
	if !IsNil(o.OutputValues) {
		toSerialize["output_values"] = o.OutputValues
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Tinting) {
		toSerialize["tinting"] = o.Tinting
	}
	return toSerialize, nil
}

func (o *GetTechnicalIndicators200ResponseDataValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enable",
		"full_name",
		"description",
		"type",
		"overlay",
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

	varGetTechnicalIndicators200ResponseDataValue := _GetTechnicalIndicators200ResponseDataValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTechnicalIndicators200ResponseDataValue)

	if err != nil {
		return err
	}

	*o = GetTechnicalIndicators200ResponseDataValue(varGetTechnicalIndicators200ResponseDataValue)

	return err
}

type NullableGetTechnicalIndicators200ResponseDataValue struct {
	value *GetTechnicalIndicators200ResponseDataValue
	isSet bool
}

func (v NullableGetTechnicalIndicators200ResponseDataValue) Get() *GetTechnicalIndicators200ResponseDataValue {
	return v.value
}

func (v *NullableGetTechnicalIndicators200ResponseDataValue) Set(val *GetTechnicalIndicators200ResponseDataValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTechnicalIndicators200ResponseDataValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTechnicalIndicators200ResponseDataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTechnicalIndicators200ResponseDataValue(val *GetTechnicalIndicators200ResponseDataValue) *NullableGetTechnicalIndicators200ResponseDataValue {
	return &NullableGetTechnicalIndicators200ResponseDataValue{value: val, isSet: true}
}

func (v NullableGetTechnicalIndicators200ResponseDataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTechnicalIndicators200ResponseDataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
