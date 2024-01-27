(function (Vue, $flare) {
  var form = '<% .Data.ThemesApi.GetFormFieldPath "Form.vue" %>';
  var input = '<% .Data.ThemesApi.GetFormFieldPath "Input.vue" %>';
  var button = '<% .Data.ThemesApi.GetFormFieldPath "Button.vue" %>';

  Vue.component('flare-form', $flare.vueLazyLoad(form));
  Vue.component('flare-form-input', $flare.vueLazyLoad(input));
  Vue.component('flare-button', $flare.vueLazyLoad(button));
})(window.Vue, window.$flare);
