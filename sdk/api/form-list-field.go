/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkapi

type FormListOption struct {
	Label string
	Value interface{}
}

type FormListField struct {
	Name     string
	Label    string
	Type     string
	Multiple bool
	Options  func() []FormListOption
	ValueFn  func() interface{}
}

func (f FormListField) GetName() string {
	return f.Name
}

func (f FormListField) GetLabel() string {
	return f.Label
}

func (f FormListField) GetType() string {
	return FormFieldTypeList
}

func (f FormListField) GetValue() interface{} {
	if f.ValueFn != nil {
		return f.ValueFn()
	}
	return nil
}
