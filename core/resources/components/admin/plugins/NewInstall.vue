<template>
    <div>
        <h1>Install Plugin</h1>
        <hr />
        <form name="plugin-form">
            <p>
                {{ def }}
            </p>
            <div class="form-group">
                <label for="" class="form-label"> Source </label>
                <select v-model="def.Src" name="src" id="plugin-src" class="form-control">
                    <option v-for="source in sources" :value="source.value">
                        {{ source.text }}
                    </option>
                </select>
            </div>
            <div v-if="def.Src == 'git'">
                <div class="form-group">
                    <label for="" class="form-label">Git Repository:</label>
                    <input v-model="def.GitURL" type="text" class="form-control"
                        placeholder="https://github.com/user/repo" />
                </div>
                <div class="form-group">
                    <label for="" class="form-label">Branch/Commit/Hash(optional):</label>
                    <input v-model="def.GitRef" type="text" class="form-control" placeholder="main" />
                </div>
                <div class="mt-3">
                    <button type="submit" v-on:click="install" class="btn btn-primary">
                        Install
                    </button>
                </div>
            </div>
            <div v-if="def.Src == 'zip'">
                <div class="form-group">
                    <label for="" class="form-label">Zip File:</label>
                    <input type="file" @change="handleFileUpload" class="form-control"
                        placeholder="plugin release zip file" />
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
                    Src: 'zip',
                    GitURL: '',
                    GitRef: '',
                    LocalZipFile: '',
                },
            };
        },
        methods: {
            handleFileUpload: function (e) {
                e.preventDefault();
                var self = this;

                let formData = new FormData();
                formData.append('file', e.target.files[0]);

                // upload file post request
                fetch(
                    '<% .Helpers.UrlForRoute "admin:upload:file" %>', {
                    method: 'POST',
                    body: formData
                })
                    .then(function (response) {
                        if (!response.ok) {
                            throw new Error("Network response was not ok: " + response.statusText);
                        }
                        var responseJson = response.json();
                        return responseJson;
                    })
                    .then(function (data) {
                        self.def.LocalZipFile = data;
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            install: function (e) {
                e.preventDefault();
                var self = this;

                console.log("installing plugin");

                if (self.def.Src == 'zip') {
                    $flare.http.post(
                        '<% .Helpers.UrlForRoute "admin:plugins:install" %>',
                        self.def
                    ).then(function (data) {
                        $flare.notify.success(`${data.Name} installed`);
                    }).catch(function (error) {
                        console.log(error);
                    });

                    return
                }

                $flare.http.post(
                    '<% .Helpers.UrlForRoute "admin:plugins:install" %>',
                    self.def
                ).then(function (data) {
                    $flare.notify.success(`${data.Name} installed`);
                });
            }
        }
    };
});
</script>
