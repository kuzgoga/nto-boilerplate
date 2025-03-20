<script setup lang="ts">
import NavModal from '../../components/modals/NavModal.vue';
import { routes } from '../../router/router';
import WoodspecScheme from '../../woodspec/WoodspecScheme.vue';

</script>

<template>
    <NavModal :routes="<any>routes[routes.findIndex(r => r.path == '/dev')].children" />
            <RouterView/>
    
</template>