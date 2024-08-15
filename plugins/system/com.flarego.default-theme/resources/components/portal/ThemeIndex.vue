<template>
  <div>
    <div class="container">
      <div class="row">
        <div class="col">
          <div
            style="height: 100px; background-color: rgba(128, 128, 128, 0)"
            class="my-2 pt-3 text-center align-middle"
          >
            Banners Here
          </div>
        </div>
      </div>

      <div v-if="data">
        <div class="row" v-for="item in data.portalItems">
          <div class="col col-md-4 offset-md-4">
            <router-link
              :to="{ name: item.route_name }"
              class="btn btn-outline-primary btn-lg btn-block mb-2"
            >
              <img :src="item.icon_url" class="portal-item-icon" />
              {{ item.label }}
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
define(function () {
  return {
    props: ["data"],
    template: template
  };
});
</script>

<style>
.portal-item-icon {
  width: 50px;
  margin-right: 10px;
  display: inline-block;
}
</style>
