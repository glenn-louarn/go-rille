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
                        Température
                    </th>
                    <th class="text-left">
                        Vent
                    </th>
                    <th class="text-left">
                        Pluie
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in items" :key="item.Date">
                    <td>{{ item.Date }}</td>
                    <td>{{ item.Kmh }}</td>
                    <td>{{ item.Temperature }}</td>
                    <td>{{ item.Pluie }}</td>
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
            '10/11/2020', '11/11/2020', '12/11/2020', '13/11/2020'
        ],
        dateDebut: "",
        dateFin: "",
        items: [{
            Date: "10/20/2020",
            Kmh: 20,
            Temperature: 20,
            Pluie: 20
        }, {
            Date: "11/20/2020",
            Kmh: 20,
            Temperature: 20,
            Pluie: 20
        }, {
            Date: "13/20/2020",
            Kmh: 20,
            Temperature: 20,
            Pluie: 20
        }, {
            Date: "12/20/2020",
            Kmh: 20,
            Temperature: 20,
            Pluie: 20
        }, {
            Date: "14/20/2020",
            Kmh: 20,
            Temperature: 20,
            Pluie: 20
        }, {
            Date: "15/20/2020",
            Kmh: 20,
            Temperature: 20,
            Pluie: 20
        }]
    }),
    async created() {

        const axios = require("axios");
        await axios
            .get('http://localhost:8081/date')
            .then(response => (this.dates = [response.dates]));
        this.dateDebut = this.dates[0];
        this.dateFin = this.dates[this.dates.length - 1];
        this.change();
    },
    methods: {
        async change() {
            const axios = require("axios");
            await axios
                .get('http://localhost:8081/date/filter/' + this.dateDebut + "/" + this.dateFin)
                .then(response => (this.cards = [response.data]));
            //console.log("je change debut ", this.dateDebut, " fin ", this.dateFin)
        }
    }
};
</script>

<style>
.table {
    margin-top: 40px;
}
</style>
