<template lang="html">
    <div>
        <h1> Logs </h1>

        <!-- <button @click="console.log(document.querySelecor('#levels'))">Test</button> -->

        <!-- search and filter-->
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

            <label for="datestart">From</label>
            <input type="date" name="datestart" id="datestart" :value="dateToYYYYMMDD(datestart)"
                @input="datestart = $event.target.valueAsDate">
            <label for="dateend">To</label>
            <input type="date" name="dateend" id="dateend" :value="dateToYYYYMMDD(dateend)"
                @input="dateend = $event.target.valueAsDate">

        </form>
        <button @click="filterLogs">Filter</button>

        <!-- logs list -->
        <div v-for="log in flareView.data"
            v-if="(log.level == level || level == 'all') &&
                (log.plugin == plugin || plugin == 'all') &&
                (new Date(`${log.year}-${log.month}-${log.day} ${log.hour}:${log.min}:${log.sec}`).getTime() >= datestart.getTime() && new Date(`${log.year}-${log.month}-${log.day} ${log.hour}:${log.min}`).getTime() < dateend.getTime())">

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

</template>

<script setup>
define(function () {
    return {
        template: template,
        props: ['flareView'],
        data: function () {
            return {
                level: "all",
                plugin: "all",
                datestart: new Date(),
                dateend: new Date(),
            }
        },
        methods: {
            filterLogs() {
                this.level = this.$el.querySelector('#levels').value;
                this.plugin = this.$el.querySelector('#pluginselection').value;

                this.datestart = new Date(this.$el.querySelector('#datestart').value);
                this.dateend = new Date(this.$el.querySelector('#dateend').value);

                this.setDatestartTimeToMidnight();
                this.setDateendTimeBeforeMidnight();
            },
            dateToYYYYMMDD(d) {
                var day = ("0" + d.getDate()).slice(-2);
                var month = ("0" + (d.getMonth() + 1)).slice(-2);
                var converted = d.getFullYear() + "-" + (month) + "-" + (day);

                return converted;
            },
            setInitialDates() {
                this.datestart = new Date();
                this.datestart.setHours(0);
                this.datestart.setMinutes(0);
                this.datestart.setSeconds(0);
                this.datestart.setMilliseconds(0);

                this.dateend = new Date();
                this.dateend.setHours(23);
                this.dateend.setMinutes(59);
                this.dateend.setSeconds(59);
                this.dateend.setMilliseconds(999);
            },
            setDatestartTimeToMidnight() {
                // set the time of the filter start date to 0:0:0
                this.datestart.setHours(0);
                this.datestart.setMinutes(0);
                this.datestart.setSeconds(0);
                this.datestart.setMilliseconds(0);
            },
            setDateendTimeBeforeMidnight() {
                // set the time of the filter end date to 23:59:59
                this.dateend.setHours(23);
                this.dateend.setMinutes(59);
                this.dateend.setSeconds(59);
                this.dateend.setMilliseconds(999);
            }
        },
        beforeMount() {
            this.setInitialDates();
        }
    };
});
</script>