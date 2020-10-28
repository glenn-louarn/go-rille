<template>
<v-container>
    <v-row>
        <v-col cols="6">
            <v-select v-model="dateDebut" :items="dates" filled label="Filled style" @change="change()"></v-select>
        </v-col>
        <v-col cols="6">
            <v-select v-model="dateFin" :items="dates" filled label="Filled style" @change="change()"></v-select>
        </v-col>
    </v-row>
    <v-simple-table class="table">
        <template v-slot:default>
            <thead>
                <tr>
                    <th class="text-left">
                        Date
                    </th>
                    <th class="text-left">
                        Capteur
                    </th>
                    <th class="text-left">
                        Valeur
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in itemsRain" :key="item.Val">
                    <td>{{ item.Date }}</td>
                    <td>Pluie</td>
                    <td>{{ item.Val }}</td>
                </tr>
                <tr v-for="item in itemsTemperature" :key="item.Val">
                    <td>{{ item.Date }}</td>
                    <td>Temperature</td>
                    <td>{{ item.Val }}</td>
                </tr>
                <tr v-for="item in itemsWind" :key="item.Val">
                    <td>{{ item.Date }}</td>
                    <td>Vents</td>
                    <td>{{ item.Val }}</td>
                </tr>
            </tbody>
        </template>
    </v-simple-table>
</v-container>
</template>

<script>
export default {
    data: () => ({
        dates: [
            '2020-03-15', '2020-03-16', '2020-03-17', '2020-03-18'
        ],
        dateDebut: "",
        dateFin: ""
    }),
    props: {
        itemsRain: Object,
        itemsWind: Object,
        itemsTemperature: Object
    },

    async created() {

        // const axios = require("axios");
        // await axios
        //     .get('http://localhost:8081/date')
        //     .then(response => (this.dates = [response.dates]));
        this.dateDebut = this.dates[0];
        this.dateFin = this.dates[this.dates.length - 1];
    },
    methods: {
        change() {
            this.$emit("change");
        }

    }
};
</script>

<style>
.table {
    margin-top: 40px;
}
</style>
