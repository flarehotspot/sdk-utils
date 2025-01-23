/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkapi

import (
	"net/http"

	"github.com/a-h/templ"
)

type IHttpFormsApi interface {
	RegisterForms(forms ...HttpForm) (err error)
	GetForm(name string) (form IHttpForm, ok bool)
}

type IHttpForm interface {
	GetTemplate(r *http.Request) templ.Component

	GetSections() []FormSection

	GetStringValue(section string, name string) (string, error)
	GetStringValues(section string, name string) ([]string, error)

	GetIntValue(section string, name string) (int64, error)
	GetIntValues(section string, name string) ([]int64, error)

	GetFloatValue(section string, name string) (float64, error)
	GetFloatValues(section string, name string) ([]float64, error)

	GetBoolValue(section string, name string) (bool, error)
	GetBoolValues(section string, name string) ([]bool, error)

	GetMultiField(section string, name string) (IFormMultiField, error)

	ParseForm(r *http.Request) error
}
