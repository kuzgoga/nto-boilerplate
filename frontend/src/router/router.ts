import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Index from "../pages/Index.vue";
import PsIndex from "../pages/ps/PsIndex.vue";
import DevIndex from "../pages/dev/DevIndex.vue";
import DelIndex from "../pages/del/DelIndex.vue";
import WoodspecScheme from "../woodspec/WoodspecScheme.vue";
import ExporterScheme from "../exporter/ExporterScheme.vue";
import ReceiverScheme from "../receiver/ReceiverScheme.vue";
import WoodspectypeScheme from "../woodspectype/WoodspectypeScheme.vue";
import { WorkResult } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services";
import WorkresultScheme from "../workresult/WorkresultScheme.vue";
import PostavScheme from "../postav/PostavScheme.vue";
import DevReport from "../pages/dev/DevReport.vue";
import SushkaresultScheme from "../sushkaresult/SushkaresultScheme.vue";
import SushkaReport from "../pages/dev/SushkaReport.vue";
import DrymodeScheme from "../drymode/DrymodeScheme.vue";


export const routes: RouteRecordRaw[] = [
    {
        path: "/",
        component: Index,
        name: 'Главная',
    },
    {
        path: '/ps',
        component: PsIndex,
        name: 'Приемка сырья',
        children: [
            {
                name: "Номенклатура",
                path: '/ps/n',
                component: WoodspecScheme
            },
            {
                path: '/ps/ex',
                component: ExporterScheme,
                name: 'Экспортер'
            }, {
                path: '/ps/re',
                name: "Приемка",
                component: ReceiverScheme
            }, {
                name: "Вид номенклатуры",
                component: WoodspectypeScheme,
                path: '/ps/ws'
            }
        ]
    }, {
        path: '/dev',
        component: DevIndex,
        name: 'Переработка',
        children: [
            {
                name: "Номенклатура",
                path: '/ps/n',
                component: WoodspecScheme
            },
            {
                name: "Вид номенклатуры",
                component: WoodspectypeScheme,
                path: '/ps/ws'
            }, {
                name: "Результат работы цеха",
                path: '/dev/workresult',
                component: WorkresultScheme
            }, {
                name: "Постав",
                component: PostavScheme,
                path: '/dev/postav'
            }, {
                name: "Отчет",
                component: DevReport,
                path: '/dev/report'
            }, {
                name: "Сушка",
                component: SushkaresultScheme,
                path: '/dev/sushka'
            }, {
                name: "Отчет по сушильному комплексу",
                component: SushkaReport,
                path: '/dev/sushkareport'
            }, {
                name: 'Режим сушки',
                component: DrymodeScheme,
                path: '/dev/drymode'
            }
        ]
    }, {
        path: '/del',
        component: DelIndex,
        name: 'Отгрузка',children: [
            {
                name: "Номенклатура",
                path: '/ps/n',
                component: WoodspecScheme
            },
            {
                name: "Вид номенклатуры",
                component: WoodspectypeScheme,
                path: '/ps/ws'
            }
        ]
    }] as const

export const router = createRouter({
    history: createWebHistory(),
    routes,
});