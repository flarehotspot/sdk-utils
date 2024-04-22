<template lang="html">
    <div>
        <h1> Logs </h1>

        <!-- search  -->
        <form action="">
            <input type="text" name="search" id="search" v-model="searchFilter">
        </form>
        <button @click="searchLogs">search</button>

        <!-- filter-->
        <form action="">
            <label for="levels">Select level: </label>
            <!-- view selector -->
            <select name="levels" id="levels" v-model="selectedLevel">
                <option value="all">All</option>
                <option value="0">INFO</option>
                <option value="1">DEBUG</option>
                <option value="2">ERROR</option>
            </select>

            <!-- plugin selector if plugin view is selected -->
            <label for="pluginselection">Plugin</label>
                <select name="pluginselection" id="pluginselection" v-model="selectedPlugin">
                    <option v-for="plugin in plugins" v-bind:value="plugin">{{ plugin }}</option>
                </select>

            <label for="dateStartFilter">From</label>
            <input type="date" name="dateStartFilter" id="dateStartFilter" :value="dateToYYYYMMDD(dateStartFilter)"
                @input="dateStartFilter = $event.target.valueAsDate">
            <label for="dateEndFilter">To</label>
            <input type="date" name="dateEndFilter" id="dateEndFilter" :value="dateToYYYYMMDD(dateEndFilter)"
                @input="dateEndFilter = $event.target.valueAsDate">


        </form>
        <button @click="filterLogs">Filter</button>

        <!-- logs list -->
        <div style="overflow-y: scroll; height: 400px;" id="logsList">

            <div v-for="log in flareView.data.logs"
                v-if="(log.level == selectedLevel || selectedLevel == 'all') &&
            (log.plugin == selectedPlugin || selectedPlugin == 'all') &&
            (log.title.toLowerCase().includes(searchFilter) || searchFilter == '') &&
            (new Date(`${log.year}-${log.month}-${log.day} ${log.hour}:${log.min}:${log.sec}`).getTime() >= dateStartFilter.getTime() && new Date(`${log.year}-${log.month}-${log.day} ${log.hour}:${log.min}`).getTime() < dateEndFilter.getTime())">

                <!-- datetime and file line-->
                <p>
                    <span>{{ log.year }}/{{ log.month }}/{{ log.day }} {{ log.hour }}:{{ log.min }}:{{ log.sec }}.{{
            log.nano }} {{ log.file }}:{{ log.line }}</span>
                </p>

                <p>
                    <span>{{ log.plugin }} </span>
                    <span>{{ log.filepluginpath }} </span>
                    <!-- <span>{{ log.filename }} </span> -->
                </p>

                <!-- level and title -->
                <p>
                    <span v-if="log.level == '0'">INFO</span>
                    <span v-if="log.level == '1'">DEBUG</span>
                    <span v-if="log.level == '2'">ERROR</span>

                    <span>{{ log.title }}</span>
                </p>

                <!-- body -->
                <div v-for="d, i in log.body">
                    <p>
                        <span v-if="i % 2 == 0">"{{ d }}": </span>
                        <span v-else>"{{ d }}"</span>
                    </p>
                </div>
            </div>
        </div>

        <p>logs per page:</p>
        <select name="perPageSelection" id="perPageSelection" v-model="perPage" @change="navigate($event, Math.ceil(this.rows / this.perPage))">
            <option value="10">10</option>
            <option value="50" selected="selected">50</option>
            <option value="100">100</option>
            <option value="200">200</option>
        </select>

        <b-pagination v-model="currentPage" :total-rows="rows" :per-page="perPage"
            @page-click="navigate"></b-pagination>
    </div>
</template>

<script setup>
define(function () {
    return {
        template: template,
        props: ['flareView'],
        data: function () {
            return {
                selectedLevel: "all",
                selectedPlugin: "all",
                dateStartFilter: new Date(),
                dateEndFilter: new Date(),
                searchFilter: "",
                lines: 50,
                plugins: ["all"],
                currentPage: 1,
                rows: 200,
                perPage: 50,
            }
        },
        methods: {
            searchLogs: function () {
                console.log(this.flareView.data);
                this.searchFilter = this.$el.querySelector('#search').value;
                console.log(this.searchFilter);
            },
            filterLogs: function () {
                this.selectedLevel = this.$el.querySelector('#levels').value;
                this.selectedPlugin = this.$el.querySelector('#pluginselection').value;

                this.dateStartFilter = new Date(this.$el.querySelector('#dateStartFilter').value);
                this.dateEndFilter = new Date(this.$el.querySelector('#dateEndFilter').value);

                this.setDateToMidnight(this.dateStartFilter);
                this.setDateToBeforeMidnight(this.dateEndFilter);
            },
            dateToYYYYMMDD: function (d) {
                var day = ("0" + d.getDate()).slice(-2);
                var month = ("0" + (d.getMonth() + 1)).slice(-2);
                var converted = d.getFullYear() + "-" + (month) + "-" + (day);
                var converted = d.getFullYear() + "-" + (month) + "-" + (day);

                return converted;
            },
            setDateToMidnight: function (d) {
                // set the time of the filter start date to 0:0:0
                d.setHours(0);
                d.setMinutes(0);
                d.setSeconds(0);
                d.setMilliseconds(0);
            },
            setDateToBeforeMidnight: function (d) {
                // set the time of the filter end date to 23:59:59
                d.setHours(23);
                d.setMinutes(59);
                d.setSeconds(59);
                d.setMilliseconds(999);
            },
            setPlugins: function () {
                this.plugins = [];
                this.plugins.push("all");

                this.flareView.data.logs.forEach((log) => {
                    if (!this.plugins.includes(log.plugin)) {
                        this.plugins.push(log.plugin);
                    }
                });
            },
            navigate: function (event, currentPage) {
                this.$router.push(
                    { path: `<% .Helpers.VueRoutePath "log-viewer" "currentPage" "${currentPage}" "perPage" "${this.perPage}" %>` }
                ).then(() => { this.$router.go(0) });
            },
            perPageChanged: function(event) {
                console.log("per page changed", this.perPage);
                console.log("new current page", this.currentPage);
            }
        },
        beforeUpdate: function () {
            this.setPlugins();
            this.filterLogs();

            // set initial start and end date to first and last logs' date
            var firstLog = this.flareView.data.logs[0];
            var lastLog = this.flareView.data.logs[this.flareView.data.logs.length - 1];

            this.dateStartFilter = new Date(`${firstLog.year}-${firstLog.month}-${firstLog.day}`);
            this.dateEndFilter = new Date(`${lastLog.year}-${lastLog.month}-${lastLog.day}`);

            this.setDateToMidnight(this.dateStartFilter);
            this.setDateToBeforeMidnight(this.dateEndFilter);
        },
        updated: function () {
            // set the scrollview of logs list to the bottom
            var logsList = this.$el.querySelector('#logsList');
            logsList.scrollTop = logsList.scrollHeight;

            // set initial pagination values
            this.perPage = this.flareView.data.perPage;
            const perPageSelection = this.$el.querySelector('#perPageSelection');
            perPageSelection.value = this.perPage;
            this.rows = this.flareView.data.rows;
            this.currentPage = this.flareView.data.currentPage;
        },
    };
});
</script>