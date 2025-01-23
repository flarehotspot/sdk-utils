/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkapi

const (
	FormFieldTypeText    string = "text"
	FormFieldTypeDecimal string = "decimal"
	FormFieldTypeInteger string = "int"
	FormFieldTypeBoolean string = "bool"
	FormFieldTypeList    string = "list"
	FormFieldTypeMulti   string = "multi"
)

type IFormField interface {
	GetName() string
	GetLabel() string
	GetType() string
	GetValue() interface{}
}

type HttpForm struct {
	Name          string
	CallbackRoute string
	Sections      []FormSection
	SubmitLabel   string
}

type FormSection struct {
	Name   string
	Fields []IFormField
}

type SectionData struct {
	Name   string          `json:"name"`
	Fields []FormFieldData `json:"fields"`
}

type FormFieldData struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
