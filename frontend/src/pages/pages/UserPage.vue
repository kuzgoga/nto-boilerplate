<script setup lang="ts">
import { RouterView, useRoute, type RouteRecordRaw } from 'vue-router';
import NavModal, { type Route } from '../../components/modals/NavModal.vue';
import { routes } from '../../router/router';

const route = useRoute()

const lastRouteName = route.matched[route.matched.length - 2].path
</script>

<template>
    <NavModal :routes="<Route[]>routes[routes.findIndex(r => r.path == lastRouteName)].children" />
    <RouterView />
</template>