/*
Search API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: latest
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Response{}

// Response struct for Response
type Response struct {
	// The payload
	Data interface{} `json:"data,omitempty"`
	Cursor *ResponseCursor `json:"cursor,omitempty"`
	// The kind of the object, either \"TRANSACTION\" or \"META\"
	Kind interface{} `json:"kind,omitempty"`
	// The ledger
	Ledger interface{} `json:"ledger,omitempty"`
}

// NewResponse instantiates a new Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponse() *Response {
	this := Response{}
	return &this
}

// NewResponseWithDefaults instantiates a new Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithDefaults() *Response {
	this := Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Response) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetDataOk() (*interface{}, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Response) HasData() bool {
	if o != nil && isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *Response) SetData(v interface{}) {
	o.Data = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *Response) GetCursor() ResponseCursor {
	if o == nil || isNil(o.Cursor) {
		var ret ResponseCursor
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetCursorOk() (*ResponseCursor, bool) {
	if o == nil || isNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *Response) HasCursor() bool {
	if o != nil && !isNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given ResponseCursor and assigns it to the Cursor field.
func (o *Response) SetCursor(v ResponseCursor) {
	o.Cursor = &v
}

// GetKind returns the Kind field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Response) GetKind() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetKindOk() (*interface{}, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return &o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Response) HasKind() bool {
	if o != nil && isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given interface{} and assigns it to the Kind field.
func (o *Response) SetKind(v interface{}) {
	o.Kind = v
}

// GetLedger returns the Ledger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Response) GetLedger() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Ledger
}

// GetLedgerOk returns a tuple with the Ledger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetLedgerOk() (*interface{}, bool) {
	if o == nil || isNil(o.Ledger) {
		return nil, false
	}
	return &o.Ledger, true
}

// HasLedger returns a boolean if a field has been set.
func (o *Response) HasLedger() bool {
	if o != nil && isNil(o.Ledger) {
		return true
	}

	return false
}

// SetLedger gets a reference to the given interface{} and assigns it to the Ledger field.
func (o *Response) SetLedger(v interface{}) {
	o.Ledger = v
}

func (o Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Ledger != nil {
		toSerialize["ledger"] = o.Ledger
	}
	return toSerialize, nil
}

type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


