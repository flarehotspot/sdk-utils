<!--
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

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
