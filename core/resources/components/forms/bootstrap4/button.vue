<template>
  <button :type="type" :class="klass">
    <slot></slot>
  </button>
</template>

<script>
define(function () {
  return {
    props: ['type', 'class'],
    computed: {
      klass: function () {
        return 'btn ' + (this.type || 'default');
      }
    },
    template: template
  };
});
</script>
