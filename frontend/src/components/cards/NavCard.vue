<script setup lang="ts">
import { Button, Card } from 'primevue';
import { useRouter } from 'vue-router';

const router = useRouter()

const navigateTo = (path: string) => {
  localStorage.setItem('app-selection', path)
  router.push(path)
}

defineProps<{
    to: string
    title: string
    content?: string
}>()
</script>

<template>
    <Card>
      <template #title>
        <h1 class="text-md">{{ title }}</h1>
      </template>
      <template #content v-if="content">
        <p class="w-56">{{ content }}</p>
      </template>
      <template #footer>
        <Button @click="navigateTo(to)" severity="contrast" size="small">Перейти</Button>
      </template>
    </Card>
</template>