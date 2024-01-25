<template>
  <div class="form-group">
    <label>{{ label }}</label>
    <input :type="type" :class="klass" v-model="val" :disabled="disabled" />
    <small class="form-text text-danger" v-if="error">{{ error }}</small>
  </div>
</template>

<script>
define(function () {
  return {
    props: ['label', 'type', 'value', 'error', 'disabled'],
    template: template,
    computed: {
      val: {
        get: function () {
          return this.value;
        },
        set: function (val) {
          this.$emit('input', val);
        }
      },
      klass: function () {
        return 'form-control' + (this.error ? ' is-invalid' : '');
      }
    }
  };
});
</script>
