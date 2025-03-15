import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Index from "../pages/Index.vue";


export const routes: RouteRecordRaw[] = [{
    path: "/",
    component: Index,
    name: 'Главная'
}, ] as const

export const router = createRouter({
    history: createWebHistory(),
    routes,
});

// {
//     path: '/user',
//         name: 'Пользователь',
//     component: UserPage,
//     redirect: '/user/post',
//     children: [
//     {
//         component: PostTablePage,
//         path: '/user/post',
//         name: 'Новости'
//     }, {
//         component: GrebenPage,
//         path: '/user/greben',
//         name: 'Страница гребня'
//     }
// ]
// }