<template lang="html">
    <div>
        <h1> Logs </h1>

        <button @click="console.log(document.querySelecor('#levels'))">Test</button>

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
                <option value="defaultTheme">Core</option>
                <option value="defaultTheme">Default Theme</option>
                <option value="themePicker">Theme Picker</option>
            </select>

            <label for="datestart">From</label>
            <input type="date" name="datestart">
            <label for="dateend">To</label>
            <input type="date" name="dateend">

        </form>
        <button @click="filterLogs">Filter</button>

        <!-- logs list -->
        <div v-for="log in flareView.data" v-if="log.level == level || level == 'all'">
            <!-- datetime and file line-->
            <p>
                <span>{{ log.year }}/{{ log.month }}/{{ log.day }} {{ log.hour }}:{{ log.min }}:{{ log.sec }}.{{
            log.nano }} {{ log.file }}:{{ log.line }}</span>
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

<script>
define(function () {
    return {
        template: template,
        props: ['flareView'],
        data: function () {
            return {
                level: "all",
                plugin: "all",
            }
        },
        methods: {
            filterLogs() {
                this.level = this.$el.querySelector('#levels').value;
                console.log(this.level);
            }
        }
    };
});
</script>