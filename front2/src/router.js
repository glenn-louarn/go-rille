import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";
import Aeroport from "./views/Aeroport.vue";
Vue.use(Router);

const route = new Router({
    routes: [
        {
            path: "/",
            redirect: "/home"
        },
        {
            path: "/home",
            name: "home",
            component: Home
        },
        {
            path: "/aeroport/:id",
            name: "aeroport",
            component: Aeroport
        }
    ]
});

export default route;