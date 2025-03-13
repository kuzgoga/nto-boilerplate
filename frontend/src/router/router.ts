import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Index from "../pages/Index.vue";
import UserPage from "../pages/user/UserPage.vue";
import PostTablePage from "../pages/tables/PostTablePage.vue";

export const routes: RouteRecordRaw[] = [{
    path: "/",
    component: Index,
    name: 'Главная'
}, {
    path: '/user',
    name: 'Пользователь',
    component: UserPage,
    redirect: '/user/post',
    children: [
        {
            component: PostTablePage,
            path: '/user/post',
            name: 'Новости'
        }
    ]
}]

export const router = createRouter({
    history: createWebHistory(),
    routes,
});