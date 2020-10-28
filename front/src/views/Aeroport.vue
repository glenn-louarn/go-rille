<template>
<v-container class="container">
    <v-row>
        <v-col cols="6">
            <h2>{{item.aeroportName}} ({{item.aeroportInitial}})</h2>
            <h4 class="marginTop">Moyenne :</h4>
            <v-col>
                <v-row align="center" justify="left">
                    <v-col cols="2">
                        <v-icon>mdi-thermometer</v-icon>
                    </v-col>
                    <h4>Température : {{item.val.temperature}} °C </h4>
                </v-row>
                <v-row align="center" justify="left">
                    <v-col cols="2">
                        <v-icon>mdi-weather-windy</v-icon>
                    </v-col>
                    <h4>Vent : {{item.val.wind}} km/h</h4>
                </v-row>
                <v-row align="center" justify="left">
                    <v-col cols="2">
                        <v-icon>mdi-water</v-icon>
                    </v-col>
                    <h4>Pluie : {{item.val.rain}} % </h4>
                </v-row>
            </v-col>
        </v-col>
        <v-col cols="6">
            <Table></Table>
        </v-col>
    </v-row>
</v-container>
</template>

<script>
import Table from "@/components/Table"
export default {
    components: {
        Table
    },
    data: () => ({
        item: {
            aeroportName: "Default",
            aeroportInitial: "DEf",
            val: Object
        }
    }),
    async created() {
        const axios = require("axios");
        this.item.aeroportInitial = this.$route.params.id.substr(1)
        this.initialToName()
        await axios
            .get('http://localhost:8081/donnees/' + this.item.aeroportInitial + "/2020-09-11")
            .then(response => (this.item.val = response.data))
    },
    methods: {
        redirectToHome() {
            this.$router.push({
                path: "/home"
            })
        },
        initialToName() {
            switch (this.item.aeroportInitial) {
                case "AVR":
                    this.item.aeroportName = "Base aérienne d'Alverca"
                    break;
                case "CDG":
                    this.item.aeroportName = "Charles de gaulle"
                    break;
                case "AVW":
                    this.item.aeroportName = "Aéroport régional de Marana"
                    break;
                case "MAK":
                    this.item.aeroportName = "Aéroport de Malakal"
                    break;
                default:
                    this.item.aeroportName = "def"
            }
        }
    }
};
</script>

<style>
.container {
    margin-top: 100px;
}

.marginTop {
    margin-top: 15px;
}
</style>
