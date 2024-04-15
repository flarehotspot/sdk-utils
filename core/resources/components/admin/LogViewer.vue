<template lang="html">
    <div>
        <h1> Logs </h1>

        <!-- search  -->
        <form action="">
            <input type="text" name="search" id="search">
        </form>
        <button @click="searchLogs">search</button>

        <!-- filter-->
        <form action="">
            <label for="levels">Select level: </label>

            <!-- view selector -->
            <select name="levels" id="levels">
                <option value="all">All</option>
                <option value="0">INFO</option>
                <option value="1">DEBUG</option>
                <option value="2">ERROR</option>
            </select>

            <!-- plugin selector if plugin view is selected -->
            <label for="pluginselection">Plugin</label>
            <select name="pluginselection" id="pluginselection">
                <option value="all">All</option>
                <option value="core">Core</option>
                <option value="defaultTheme">Default Theme</option>
                <option value="themePicker">Theme Picker</option>
            </select>

            <label for="dateStartFilter">From</label>
            <input type="date" name="dateStartFilter" id="dateStartFilter" :value="dateToYYYYMMDD(dateStartFilter)"
                @input="dateStartFilter = $event.target.valueAsDate">
            <label for="dateEndFilter">To</label>
            <input type="date" name="dateEndFilter" id="dateEndFilter" :value="dateToYYYYMMDD(dateEndFilter)"
                @input="dateEndFilter = $event.target.valueAsDate">

        </form>
        <button @click="filterLogs">Filter</button>

        <!-- <p>{{ flareView.data.logs }}</p> -->

        <!-- logs list -->
        <div style="overflow-y: scroll; height: 400px;" id="logsList">

            <div v-for="log in flareView.data.logs"
                v-if="(log.level == levelFilter || levelFilter == 'all') &&
            (log.plugin == pluginFilter || pluginFilter == 'all') &&
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

        <!-- pagination -->
        <p>pages:</p>
        <!-- <div v-for="page in (flareView.data.logSize)/ 100">
            <a href='<% .Helpers.VueRoutePath "log-viewer" "page" "{{ page }}" "lines" "50"%>'>{{ page }}</a>
        </div> -->

        <a href='<% .Helpers.VueRoutePath "log-viewer" "page" "1" "lines" "200" %>'> 1 </a>
        <router-link to='<% .Helpers.VueRoutePath "log-viewer" "page" "1" "lines" "200"%>'>
            view logs
        </router-link>

        <p>logs per page:</p>
        <select name="linesSelection" id="linesId">
            <option value="10">10</option>
            <option value="50">50</option>
            <option value="100">100</option>
            <option value="200">200</option>
            <option value="all">all</option>
        </select>
    </div>
</template>

<script setup>
define(function () {
    return {
        template: template,
        props: ['flareView'],
        data: function () {
            return {
                levelFilter: "all",
                pluginFilter: "all",
                dateStartFilter: new Date(),
                dateEndFilter: new Date(),
                searchFilter: "",
                lines: 50,
                plugins: ["all"],
            }
        },
        methods: {
            searchLogs: function () {
                console.log(this.flareView.data);
                this.searchFilter = this.$el.querySelector('#search').value;
                console.log(this.searchFilter);
            },
            filterLogs: function () {
                this.levelFilter = this.$el.querySelector('#levels').value;
                this.pluginFilter = this.$el.querySelector('#pluginselection').value;

                this.dateStartFilter = new Date(this.$el.querySelector('#dateStartFilter').value);
                this.dateEndFilter = new Date(this.$el.querySelector('#dateEndFilter').value);

                this.setDatestartTimeToMidnight();
                this.setDateendTimeBeforeMidnight();
            },
            dateToYYYYMMDD: function (d) {
                var day = ("0" + d.getDate()).slice(-2);
                var month = ("0" + (d.getMonth() + 1)).slice(-2);
                var converted = d.getFullYear() + "-" + (month) + "-" + (day);
                var converted = d.getFullYear() + "-" + (month) + "-" + (day);

                return converted;
            },
            setInitialDates: function () {
                this.dateStartFilter = new Date();
                this.dateStartFilter.setHours(0);
                this.dateStartFilter.setMinutes(0);
                this.dateStartFilter.setSeconds(0);
                this.dateStartFilter.setMilliseconds(0);

                this.dateEndFilter = new Date();
                this.dateEndFilter.setHours(23);
                this.dateEndFilter.setMinutes(59);
                this.dateEndFilter.setSeconds(59);
                this.dateEndFilter.setMilliseconds(999);
            },
            setDatestartTimeToMidnight: function () {
                // set the time of the filter start date to 0:0:0
                this.dateStartFilter.setHours(0);
                this.dateStartFilter.setMinutes(0);
                this.dateStartFilter.setSeconds(0);
                this.dateStartFilter.setMilliseconds(0);
            },
            setDateendTimeBeforeMidnight: function () {
                // set the time of the filter end date to 23:59:59
                this.dateEndFilter.setHours(23);
                this.dateEndFilter.setMinutes(59);
                this.dateEndFilter.setSeconds(59);
                this.dateEndFilter.setMilliseconds(999);
            },
            setPlugins: function () {
                console.log("setting plugins");
            }
        },
        beforeMount: function () {
            this.setInitialDates();
            this.setPlugins();
        },
        mounted: function() {
        },
        updated: function() {
            console.log("Logging inside the updated function");
            var logsList = this.$el.querySelector('#logsList');
            logsList.scrollTop = logsList.scrollHeight;
        }
    };
});
</script>