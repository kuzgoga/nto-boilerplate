<script setup lang="ts">
import { Dialog } from "primevue";
import { useErrorStore } from "./stores/error.store";
import { RouterView } from "vue-router";

const errorStore = useErrorStore();
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
    <main class="w-screen h-screen">
       <RouterView />
    </main>
</template>
