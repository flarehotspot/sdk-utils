<template>
    <div>
        <div class="d-flex p-2 align-items-center">
            <router-link class="mr-3 btn btn-secondary" to='<% .Helpers.VueRoutePath "plugins-store" %>'>
                back to store
            </router-link>
        </div>

        <hr />

        <h3>{{ data.Name }}</h3>

        <div class="btn btn-primary w-100" @click="installRelease($event, data.Releases[0])">
            Install Plugin
        </div>

        <hr />
        <h4>Description</h4>
        <hr />

        <h4>Other Versions</h4>
        <table class="table table-bordered table-striped">
            <thead>
                <tr>
                    <th>Version</th>
                    <th>Option</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="pr in data.Releases" :key="pr.Id">
                    <td>{{ pr.Major + '.' + pr.Minor + '.' + pr.Patch }}</td>
                    <td>
                        <button type="button" class="btn btn-secondary" @click="installRelease($event, pr)">
                            Install
                        </button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
define(function () {
    return {
        template: template,
        data: function () {
            var self = this;
            return {
                data: [],
                plugin: {},
                pluginId: parseInt(self.$route.query.id) || 1,
            };
        },
        mounted: function () {
            var self = this;
            self.load();
        },
        methods: {
            load: function () {
                var self = this;

                self.pluginId = parseInt(self.$route.query.id);
                var params = {
                    id: self.pluginId
                };

                $flare.http
                    .get(
                        '<% .Helpers.UrlForRoute "admin:plugins:store:plugin"  %>',
                        params
                    )
                    .then(function (data) {
                        self.data = data;

                        self.plugin = {
                            Id: self.data.Id,
                            Name: self.data.Info.Name,
                            Package: self.data.Info.Package
                        };
                    });
            },
            installRelease: function (e, pr) {
                e.preventDefault();

                var self = this;

                // plugin def for store-based plugin
                var defParams = {
                    Src: 'store',
                    StorePackage: self.plugin.Package,
                    StoreZipUrl: pr.ZipFileUrl,
                    StorePluginVersion: self.stringifyVersion(self.data.Releases[0]),
                };

                $flare.http
                    .post('<% .Helpers.UrlForRoute "admin:plugins:install" %>', defParams)
                    .then(function (response) {
                        $flare.notify.success(`${response.Name} installed`);
                    });
            },
            stringifyVersion: function (pr) {
                return `${pr.Major}.${pr.Minor}.${pr.Patch}`;
            }
        }
    };
});
</script>
