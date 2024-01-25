(function (Vue, $flare) {
  var form = '<% .Data.ThemesApi.GetFormFieldPath "form.vue" %>';
  var input = '<% .Data.ThemesApi.GetFormFieldPath "input.vue" %>';
  var button = '<% .Data.ThemesApi.GetFormFieldPath "button.vue" %>';

  Vue.component('flare-form', $flare.vueLazyLoad(form));
  Vue.component('flare-form-input', $flare.vueLazyLoad(input));
  Vue.component('flare-button', $flare.vueLazyLoad(button));
})(window.Vue, window.$flare);
