<template>
<div>
    <v-app-bar color="deep-purple accent-4" dark absolute>
        <v-toolbar-title @click="goHome()">Méteo Aéroport</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn icon>
            <v-icon>mdi-magnify</v-icon>
        </v-btn>
    </v-app-bar>
</div>
</template>

<script>
export default {
    name: 'AppBar',
    methods: {
        goHome() {
            this.$router.push({
                path: "/home"
            })
        }
    }
}
</script>
