<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetDryMaterialsList } from '../../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/dryreportservice';

const data = ref<null | {
    Materials: string[],
    MaterialsAmount: number
}>(null)

onMounted(async () => {
    data.value = await GetDryMaterialsList()
})
</script>

<template>
<div class="flex items-center justify-center">
    <div v-if="data">
        <p>Сухие пиломатериалы: {{ data.Materials.join(", ") }}</p>
        <p>Количество: {{ data.MaterialsAmount }}</p>
    </div>
    <p v-else>Нет данных</p>
</div>
</template>