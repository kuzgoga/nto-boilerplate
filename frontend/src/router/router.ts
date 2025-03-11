import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Index from "../pages/Index.vue";

export const routes: RouteRecordRaw[] = [{
    path: "/",
    component: Index,
    name: 'Главная',
}]

export const router = createRouter({
    history: createWebHistory(),
    routes,
});