<!--
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

<template>
    <div>
        <ul>
            <li v-for="(url, name) in flareView.data">
                <router-link :to="url">{{ name }}</router-link>
            </li>
        </ul>
    </div>
</template>

<script>
(function () {
    define(function () {
        return {
            props: ['flareView'],
            template: template
        };
    });
})();
</script>
