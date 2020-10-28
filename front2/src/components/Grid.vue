<template>
<div class="grid">
    <v-container class="container" fluid v-if="cards.length > 0">
        <v-row align="center " justify="center" class="white" width="900px">
            <template v-for="item in cards">
                <Card class="card" :key="item.aeroportInitial" :item="item"></Card>
            </template>
        </v-row>
    </v-container>
    <div fluid v-else class="pt-4 pb-8">
        <p>Aucun résultat</p>
    </div>

</div>
</template>

<script>
import Card from "@/components/Card.vue";
export default {
    components: {
        Card
    },
    data: () => ({
        cards: [{
                aeroportName: "Charle de gaule",
                aeroportInitial: "CDG",
                val: Object
            },
            {
                aeroportName: "Base aérienne d'Alverca",
                aeroportInitial: "AVR",
                val: []
            },
            {
                aeroportName: "Aéroport régional de Marana",
                aeroportInitial: "AVW",
                val: Object
            }, {
                aeroportName: "Aéroport de Malakal",
                aeroportInitial: "MAK",
                val: Object
            }
        ]

    }),
    async created() {
        this.cards[0].val = await this.getAeroport("CDG")
        this.cards[1].val = await this.getAeroport("AVR")
        this.cards[2].val = await this.getAeroport("AVW")
        this.cards[3].val = await this.getAeroport("MAK")
    },
    methods: {
        async getAeroport(name) {
            const axios = require("axios");
            let res;
            await axios
                .get('http://localhost:8081/donnees/' + name + '/2020-07-07')
                .then(response => (res = response.data))
            return res;
        }
    }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>
.card {
    margin: 20px;
}

.container {
    margin-top: 80px;
}
</style>
