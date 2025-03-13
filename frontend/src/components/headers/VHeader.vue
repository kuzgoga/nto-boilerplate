<script lang="ts" setup>
import { Button, Toolbar } from 'primevue';
import { useRoute } from 'vue-router';
import { useNavModalStore } from '../../stores/nav-modal.store';

const route = useRoute()

const navModalStore = useNavModalStore()
</script>

<template>
    <header >
        <Toolbar style="border-radius: 0; border-top: none; border-left: none; border-right: none">
            <template #start>
                <p class="font-bold">{{ $route.matched?.[$route.matched.length - 2]?.name }} / {{ route.name }}</p>
            </template>
            <template #end>
                <Button severity="secondary" size="small" @click="navModalStore.changeVisibility">
                    <span class="pi pi-list"></span>
                </Button>
            </template>
        </Toolbar>
    </header>
</template>