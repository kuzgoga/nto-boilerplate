<script setup lang="ts">
import { Dialog } from "primevue";
import { useErrorStore } from "./stores/error.store";
import { RouterView } from "vue-router";
import VHeader from "./components/headers/VHeader.vue";
import { useNavModalStore } from "./stores/nav-modal.store";
import NavModal from "./components/modals/NavModal.vue";
import { routes } from "./router/router";

const errorStore = useErrorStore();
const nm = useNavModalStore()
</script>

<template>
    <Dialog
        v-model:visible="errorStore.show"
        maximizable
        class="w-[400px] h-[400px]"
    >
        <template #header> Ошибка </template>
        <h1 class="text-red-500 text-2xl">{{ errorStore.message }}</h1>
    </Dialog>
    
    <VHeader  />
    <main>
       <RouterView />
    </main>
</template>
