<template>
    <div>
        <h1>System Updates</h1>

        <hr />

        <div class="" v-if="hasUpdates(current, latest)">
            <h1>Update Available</h1>
            <div class="btn btn-primary" @click="installUpdate">Install Update</div>
        </div>
        <div v-else>
            <h1>Latest</h1>
        </div>

        <p>Latest Version: {{ stringifyVersion(latest) }}</p>
        <p>Installed Version: {{ stringifyVersion(current) }}</p>

        <div class="btn btn-primary" @click="fetchLatest">Check Updates</div>
    </div>
</template>

<script>
define(function () {
    return {
        template: template,
        data: function () {
            return {
                latest: {},
                current: {},
                isUpToDate: false,
                localCoreFilesPath: '',
                localArchBinFilesPath: ''
            };
        },
        mounted: async function () {
            var self = this;

            await self.getCurrent();
            await self.fetchLatest();

            console.log('latest: ', self.latest);
        },
        methods: {
            fetchLatest: async function () {
                var self = this;

                await $flare.http
                    .get('<% .Helpers.UrlForRoute "admin:core:fetch" %>')
                    .then(function (response) {
                        console.log(response);
                        self.latest = response;
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            getCurrent: async function () {
                var self = this;

                await $flare.http
                    .get('<% .Helpers.UrlForRoute "admin:core:current" %>')
                    .then(function (response) {
                        console.log(response);
                        self.current = response;
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            downloadUpdate: async function () {
                var self = this;
                await $flare.http
                    .post('<% .Helpers.UrlForRoute "admin:core:download" %>', {
                        Major: self.latest.Major,
                        Minor: self.latest.Minor,
                        Patch: self.latest.Patch,
                        CoreZipFileUrl: self.latest.CoreZipFileUrl,
                        ArchBinFileUrl: self.latest.ArchBinFileUrl
                    })
                    .then(function (response) {
                        console.log("local core files path: ", response);
                        self.localCoreFilesPath = response.LocalCoreFilesPath;
                        self.localArchBinFilesPath = response.LocalArchBinFilesPath;

                        console.log(self.localCoreFilesPath);
                        console.log(self.localArchBinFilesPath);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            installUpdate: async function () {
                var self = this;

                if (
                    self.latest.ZipFileUrl == '' ||
                    self.latest.ZipFileUrl === undefined
                ) {
                    await self.downloadUpdate();
                }

                await $flare.http
                    .post('<% .Helpers.UrlForRoute "admin:core:update" %>', {
                        LocalCoreFilesPath: self.localCoreFilesPath,
                        LocalArchBinFilesPath: self.localArchBinFilesPath,
                        Major: self.latest.Major,
                        Minor: self.latest.Minor,
                        Patch: self.latest.Patch,
                    })
                    .then(function (response) {
                        console.log('update response: ', response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            stringifyVersion: function (version) {
                return 'v' + version.Major + '.' + version.Minor + '.' + version.Patch;
            },
            hasUpdates: function (current, latest) {
                if (current.Major < latest.Major) {
                    return true;
                }
                if (current.Minor < latest.Minor) {
                    return true;
                }
                if (current.Patch < latest.Patch) {
                    return true;
                }
                return false;
            }
        }
    };
});
</script>
