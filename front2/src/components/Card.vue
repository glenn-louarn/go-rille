<template>
<v-hover>
    <template v-slot="{ hover }">
        <div :class="`elevation-${hover ? 12 : 0}`">
            <v-card width="260px" height="280px" @click="getAeroport()">
                <v-card-title>

                    {{item.AeroportName}}
                </v-card-title>

                <v-card-subtitle>
                    ({{item.AeroportInitial}})
                </v-card-subtitle>
                <v-list dense>
                    <v-list-item-group color="primary">
                        <v-list-item>
                            <v-list-item-icon>
                                <v-icon>mdi-water</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title v-text="item.Pluie+'%'"></v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                        <v-list-item>
                            <v-list-item-icon>
                            
                                <v-icon>mdi-weather-windy</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title v-text="item.Kmh+'km/h'"></v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                        <v-list-item>
                            <v-list-item-icon>
                                <v-icon>mdi-thermometer</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title v-text="item.Temperature+'°C'"></v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                    </v-list-item-group>
                </v-list>
            </v-card>
        </div>
    </template>
</v-hover>
</template>

<script>
export default {
    name: 'Card',
    props: {
        item: Object,
        id: String,
        aeroportName: String,
        aeroportInitial: String,
        kmh: Number,
        temperature: Number,
        pluie: Number
    },
    methods: {

        getAeroport() {
            this.$router.push({
                path: "/aeroport/:" + this.item.id
            })
        }
    }
}
</script>
