<template>
    <div>
        <h1>Plugins Manager</h1>

        <hr />

        <h2>Install Plugins</h2>
        <div class="mb-2">
            <router-link class="" to='<% .Helpers.VueRoutePath "plugins-new" %>'>Install a plugin locally</router-link>
            or
            <router-link class="btn btn-primary" to='<% .Helpers.VueRoutePath "plugins-store" %>'>Visit plugins
                store</router-link>
        </div>

        <br />

        <div class="d-flex">
            <h2 class="flex-fill">Installed Plugins</h2>
            <button class="btn btn-primary" v-on:click="checkUpdates">
                Check for updates
            </button>
        </div>
        <table class="table table-bordered table-striped">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Description</th>
                    <th>Version</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="p in plugins">
                    <td>{{ p.Info.Name }}</td>
                    <td>{{ p.Info.Description }}</td>
                    <td>{{ p.Info.Version }}</td>
                    <td>
                        <button type="button" class="btn btn-danger" v-on:click="uninstall(p.Info.Package)">
                            Uninstall
                        </button>
                        <button type="button" class="btn btn-info" v-on:click="update(p.Info.Package)"
                            v-if="p.HasUpdates">
                            Update
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
            return {
                plugins: []
            };
        },
        mounted: function () {
            var self = this;
            $flare.http
                .get('<% .Helpers.UrlForRoute "admin:plugins:index"  %>')
                .then(function (plugins) {
                    self.plugins = plugins;
                });
        },
        methods: {
            uninstall: function (pkg) {
                var self = this;
                var yes = confirm('Are you sure you want to uninstall this plugin?');
                if (!yes) {
                    return;
                }

                $flare.http
                    .post('<% .Helpers.UrlForRoute "admin:plugins:uninstall" %>', {
                        pkg: pkg
                    })
                    .then(function () {
                        console.log('Uninstalled plugin: ' + pkg);
                    });
            },
            update: function (pkg) {
                var self = this;
                var yes = confirm('Are you sure you want to update this plugin?');
                if (!yes) {
                    return;
                }

                $flare.http
                    .post('<% .Helpers.UrlForRoute "admin:plugins:update" %>', {
                        pkg: pkg
                    })
                    .then(function () {
                        console.log('Updated plugin: ' + pkg);
                    });
            },
            checkUpdates: async function () {
                var self = this;
                console.log("Checking plugin updates..");

                await $flare.http.get(
                    '<% .Helpers.UrlForRoute "admin:plugins:checkupdates" %>'
                ).then(function (response) {
                    self.plugins = response;
                });

                console.log("Checking plugin updates complete!");
            }
        }
    };
});
</script>
