<script setup lang="ts">
import { Button, Drawer } from 'primevue';
import { useNavModalStore } from '../../stores/nav-modal.store';
import { useRoute, useRouter } from 'vue-router';

const navModalStore = useNavModalStore()

const router = useRouter()

const logout = () => {
    navModalStore.changeVisibility()
    router.replace('/')
}

export interface Route {
    name: string
    path: string
}

defineProps<{
    routes: Route[]
}>()

const route = useRoute()

</script>

<template>
    <Drawer v-model:visible="navModalStore.visible">
        <div class="flex flex-col gap-2">
            <RouterLink :style="{ color: route.fullPath.endsWith(r.path) && 'var(--p-primary-color)' }" @click="() => {
                navModalStore.changeVisibility()
            }" v-for="r in routes" :to="r.path">{{ r.name }}
            </RouterLink>
        </div>
        <template #footer>
            <Button severity="danger" :size="'small'" class="w-full" @click="logout">Выйти</Button>
        </template>
    </Drawer>
</template>