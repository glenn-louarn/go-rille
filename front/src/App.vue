<template>
<v-app>
    <AppBar></AppBar>
    <router-view />
</v-app>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
export default {
    name: 'App',

    components: {
        AppBar,
    }
};
</script>
