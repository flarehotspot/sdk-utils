<template>
  <div class="row mb-2">
    <div class="hidden-sm col-md-3"></div>
    <div class="col col-md-6">
      <button
        :type="type || 'button'"
        :class="'btn btn-' + (variant || 'primary')"
        :disabled="disabled"
      >
        <slot></slot>
      </button>
    </div>
    <div class="col"></div>
  </div>
</template>

<script>
define(function () {
  return {
    props: ['type', 'variant', 'disabled', 'to'],
    inheritAttrs: false,
    template: template
  };
});
</script>
