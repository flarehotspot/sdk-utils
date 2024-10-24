<template>
    <div>
        <h1>System Updates</h1>

        <hr />

        <div class="" v-if="hasUpdates(current, latest.Version)">
            <h1>Update Available</h1>
            <div class="btn btn-primary" @click="installUpdate">Install Update</div>
        </div>
        <div v-else>
            <h1>Latest</h1>
        </div>

        <p>Latest Version: {{ getLatestVersion() }}</p>
        <p>Installed Version: {{ getCurrentVersion() }}</p>

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
        },
        methods: {
            getCurrent: async function () {
                var self = this;

                await $flare.http
                    .get('<% .Helpers.UrlForRoute "admin:core:current" %>')
                    .then(function (response) {
                        // TODO: remove logs
                        console.log("get current response: ", response);
                        self.current = {
                            Major: response.Major,
                            Minor: response.Minor,
                            Patch: response.Patch,
                        };
                        console.log("self current: ", self.current);
                    })
                    .catch(function (error) {
                        console.log("error getting current version: ", error);
                    });
            },
            fetchLatest: async function () {
                var self = this;

                await $flare.http
                    .get('<% .Helpers.UrlForRoute "admin:core:fetch" %>')
                    .then(function (response) {
                        // TODO: remove logs
                        console.log("fetch latest response: ", response);
                        self.latest = {
                            Version: {
                                Major: response.Version.Major,
                                Minor: response.Version.Minor,
                                Patch: response.Version.Patch,
                            },
                            CoreZipFileUrl: response.CoreZipFileUrl,
                            ArchBinFileUrl: response.ArchBinFileUrl,
                        };
                        console.log("self latest: ", self.latest);
                    })
                    .catch(function (error) {
                        console.log("error fetching latest release: ", error);
                    });
            },
            downloadUpdate: async function () {
                var self = this;

                var payload = {
                    Major: self.latest.Major,
                    Minor: self.latest.Minor,
                    Patch: self.latest.Patch,
                    CoreZipFileUrl: self.latest.CoreZipFileUrl,
                    ArchBinFileUrl: self.latest.ArchBinFileUrl
                }

                // TODO: remove console log
                console.log("download update payload: ", payload);

                await $flare.http
                    .post('<% .Helpers.UrlForRoute "admin:core:download" %>', payload)
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

                var payload = {
                    LocalCoreFilesPath: self.localCoreFilesPath,
                    LocalArchBinFilesPath: self.localArchBinFilesPath,
                    Major: self.latest.Major,
                    Minor: self.latest.Minor,
                    Patch: self.latest.Patch,
                }

                // TODO: remove console log
                console.log("install update payload: ", payload);

                await $flare.http
                    .post('<% .Helpers.UrlForRoute "admin:core:update" %>', payload)
                    .then(function (response) {
                        console.log('update response: ', response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            getCurrentVersion: function () {
                var self = this;
                return self.stringifyVersion(self.current);
            },
            getLatestVersion: function () {
                var self = this;
                return self.stringifyVersion(self.latest.Version);
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
