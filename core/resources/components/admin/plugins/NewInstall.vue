<template>
  <div>
    <h1>Install Plugin</h1>
    <hr />
    <form name="plugin-form">
      <div class="form-group">
        <label for="" class="form-label"> Source </label
        ><select name="src" id="plugin-src" class="form-control">
          <option v-for="source in sources" :value="source.value">
            {{ source.text }}
          </option>
        </select>
      </div>
    </form>
  </div>
</template>

<script>
define(function () {
  var sources = [
    { value: 'git', text: 'Git Repository' },
    { value: 'zip', text: 'Zip Archive' }
  ];
  return {
    template: template,
    data: function () {
      return {
        sources: sources
      };
    }
  };
});
</script>
