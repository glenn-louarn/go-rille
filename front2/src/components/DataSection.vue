<template>
<div class="dataSection">
    <!-- Search field -->
    <div class="row justify-content-center">
        <div class="col-11">
            <div class="input-group mb-3">
                <input type="text" class="form-control" placeholder="Rechercher un aéroport">
                <div class="input-group-append">
                    <button class="btn btn-outline-secondary" type="button">Rechercher <i class="fas fa-plane"></i></button>
                </div>
            </div>
        </div>
    </div>
    <!-- Data block -->

    <!-- Data section -->
    <v-container fluid v-if="cards.length > 0">
        <v-row align="center" justify="center" class="white">
            <template v-for="item in cards">
                <Card :key="item.id" :aeroportName="item.aeroportName" :aeroportInitial="item.aeroportInitial" :kmh="item.kmh" :temperature="item.temperature" :pluie="item.pluie"></Card>
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
            id: "1",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "2",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "3",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "4",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "5",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "6",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "7",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "8",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "9",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "10",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "11",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "12",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "13",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "14",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, {
            id: "15",
            aeroportName: "Charle de gaule",
            aeroportInitial: "CDG",
            kmh: 20,
            temperature: 20,
            pluie: 20
        }, ]

    }),
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>
h3 {
    margin: 40px 0 0;
}

ul {
    list-style-type: none;
    padding: 0;
}

li {
    display: inline-block;
    margin: 0 10px;
}

a {
    color: #42b983;
}
</style>
