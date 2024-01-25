<template>
  <form @submit="submit" class="mt-3 mt-md-5">
    <h3>{{ title }}</h3>
    <slot></slot>
  </form>
</template>

<script>
(function () {
  define(function () {
    return {
      props: ['title'],
      template: template,
      methods: {
        submit: function (e) {
          e.preventDefault();
          this.$emit('submit');
        }
      }
    };
  });
})();
</script>
