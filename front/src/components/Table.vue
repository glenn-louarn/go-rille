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
        dateFin: "",
        itemsRain: [{
            Date: "10/20/2020",
            Val: 20
        }],
        itemsWind: [{
            Date: "10/20/2020",
            Val: 10
        }],
        itemsTemperature: [{
            Date: "10/20/2020",
            Val: 20
        }]
    }),
    async created() {

        // const axios = require("axios");
        // await axios
        //     .get('http://localhost:8081/date')
        //     .then(response => (this.dates = [response.dates]));
        this.dateDebut = this.dates[0];
        this.dateFin = this.dates[this.dates.length - 1];
        this.change();
    },
    methods: {
        async change() {
            const axios = require("axios");
            await axios
                .get('http://localhost:8081/donnees/' + this.$route.params.id + '/' + this.dateDebut + "/" + this.dateFin + "/RAIN")
                .then(response => (this.itemsRain = this.createlist(response.data)));
            await axios
                .get('http://localhost:8081/donnees/' + this.$route.params.id + '/' + this.dateDebut + "/" + this.dateFin + "/WIND")
                .then(response => (this.itemsWind = this.createlist(response.data)));
            await axios
                .get('http://localhost:8081/donnees/' + this.$route.params.id + '/' + this.dateDebut + "/" + this.dateFin + "/TEMPERATURE")
                .then(response => (this.itemsTemperature = this.createlist(response.data)));

        },
        createlist(resAPI) {
            let res = []
            for (let i = 0; i < resAPI.length; i++) {
                if (resAPI[i].ValeurCapteur != null) {
                    for (let j = 0; j < resAPI[i].val.length; j++) {
                        res.push({
                            date: resAPI[i].Date,
                            val: resAPI[i].ValeurCapteur[j]
                        })
                    }
                }
            }
            return res
        }
    }
};
</script>

<style>
.table {
    margin-top: 40px;
}
</style>
