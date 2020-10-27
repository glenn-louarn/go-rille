<template>
<v-container class="container">
    <v-row>
        <v-col cols="6">
            <h2>{{item.AeroportName}} ({{item.AeroportInitial}})</h2>
            <h4 class="marginTop">Moyenne :</h4>
            <v-col>
                <v-row align="center" justify="left">
                    <v-col cols="2">
                        <v-icon>mdi-thermometer</v-icon>
                    </v-col>
                    <h4>Température : {{item.Temperature}} °C </h4>
                </v-row>
                <v-row align="center" justify="left">
                    <v-col cols="2">
                        <v-icon>mdi-weather-windy</v-icon>
                    </v-col>
                    <h4>Vent : {{item.Kmh}} km/h</h4>
                </v-row>
                <v-row align="center" justify="left">
                    <v-col cols="2">
                        <v-icon>mdi-water</v-icon>
                    </v-col>
                    <h4>Pluie : {{item.Pluie}} % </h4>
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
            Id: "0",
            AeroportName: "Default",
            AeroportInitial: "DEf",
            Kmh: 0,
            Temperature: 0,
            Pluie: 0,
        }
    }),
    async created() {
        const axios = require("axios");
        await axios
            .get('http://localhost:8081/aeroport/' + this.$route.params.id)
            .then(response => (this.item = [response.data]))
    },
    methods: {
        redirectToHome() {
            this.$router.push({
                path: "/home"
            })
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
