<template>
  <form v-on="$listeners" class="mt-3" novalidate>
    <div class="text-center">
      <h3>{{ heading }}</h3>
    </div>
    <hr />
    <slot></slot>
  </form>
</template>

<script>
(function () {
  define(function () {
    return {
      props: ['heading'],
      template: template
    };
  });
})();
</script>
