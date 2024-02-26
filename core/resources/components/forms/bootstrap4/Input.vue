<!--
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

<template>
  <div class="row mb-2">
    <div class="col-xs-12 col-md-3 text-md-right pt-md-2">
      <label>{{ label }}</label>
    </div>
    <div class="col-xs-12 col-md-6">
      <input
        :type="type"
        :disabled="disabled"
        :class="'form-control' + (error ? ' is-invalid' : '')"
        v-model="val"
      />
      <small class="form-text text-danger" v-if="error">{{ error }}</small>
      <small class="form-text text-muted" v-if="help && !error">
        {{ help }}
      </small>
    </div>
  </div>
</template>

<script>
(function () {
  define(function () {
    return {
      template: template,
      props: ['name', 'label', 'type', 'value', 'help', 'error', 'disabled'],
      computed: {
        val: {
          get: function () {
            return this.value;
          },
          set: function (val) {
            this.$emit('input', val);
          }
        }
      }
    };
  });
})();
</script>
