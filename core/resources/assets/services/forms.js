/**
 * @file             : forms.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.ph>
 * Date              : Jan 27, 2024
 * Last Modified Date: Feb 27, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
(function (Vue, $flare) {
  var form = '<% .Data.ThemesApi.GetFormFieldPath "Form.vue" %>';
  var input = '<% .Data.ThemesApi.GetFormFieldPath "Input.vue" %>';
  var button = '<% .Data.ThemesApi.GetFormFieldPath "Button.vue" %>';

  Vue.component('flare-form', $flare.vueLazyLoad(form));
  Vue.component('flare-form-input', $flare.vueLazyLoad(input));
  Vue.component('flare-button', $flare.vueLazyLoad(button));
})(window.Vue, window.$flare);
