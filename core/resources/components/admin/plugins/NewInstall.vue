<template>
  <div>
    <h1>Install Plugin</h1>
    <hr />
    <form name="plugin-form">
      <p>
        {{ def }}
      </p>
      <div class="form-group">
        <label for="" class="form-label"> Source </label
        ><select
          v-model="def.Src"
          name="src"
          id="plugin-src"
          class="form-control"
        >
          <option v-for="source in sources" :value="source.value">
            {{ source.text }}
          </option>
        </select>
      </div>
      <div v-if="def.Src == 'git'">
        <div class="form-group">
          <label for="" class="form-label">Git Repository:</label
          ><input
            v-model="def.GitURL"
            type="text"
            class="form-control"
            placeholder="https://github.com/user/repo"
          />
        </div>
        <div class="form-group">
          <label for="" class="form-label">Branch/Commit/Hash(optional):</label>
          <input
            v-model="def.GitRef"
            type="text"
            class="form-control"
            placeholder="main"
          />
        </div>
        <div class="mt-3">
          <button type="submit" v-on:click="install" class="btn btn-primary">
            Install
          </button>
        </div>
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
        sources: sources,
        def: {
          Src: 'git',
          GitURL: '',
          GitRef: ''
        }
      };
    },
    methods: {
      install: function (e) {
        e.preventDefault();
        var self = this;
        $flare.http.post(
          '<% .Helpers.UrlForRoute "admin:plugins:install" %>',
          self.def
        );
      }
    }
  };
});
</script>
