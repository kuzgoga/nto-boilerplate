<script lang="ts" setup>
import { Button, Toolbar } from 'primevue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()

const logout = () => {
    localStorage.removeItem('app-selection')
    router.replace('/')
}

</script>

<template>
    <header >
        <Toolbar>
            <template #start>
                <p class="font-bold">{{ $route.matched?.[$route.matched.length - 2]?.name }} / {{ route.name }}</p>
            </template>
            <template #end>
                <Button severity="danger" :size="'small'" @click="logout">Выйти</Button>
            </template>
        </Toolbar>
    </header>
</template>