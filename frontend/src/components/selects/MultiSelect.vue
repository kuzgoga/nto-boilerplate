<script setup lang="ts" generic="T extends IEntity">
import { structView } from '../../utils/structs/structs-view.util';
import type { IEntity } from '../../types/entity.type';
import Checkbox from '../checkboxes/Checkbox.vue';
import { onMounted, watch } from 'vue';

const { entityId } = defineProps<{
    options: T[]
    path?: string[]
    entityId: string
}>()

const selected = defineModel<T[]>({ default: [] })

const pushOrRemove = (option: T) => {
    if (selected.value.some(s => s.Id === option.Id)) {
        selected.value = selected.value.filter(s => s.Id !== option.Id)
    } else {
        selected.value.push(option)
    }
    setNullIds()
}

const setNullIds = () => {
    selected.value = selected.value.map(item => {
        item[entityId] = 0
        return item
    })
}

onMounted(setNullIds)
</script>

<template>
    <div class="relative">
        <p class="flex items-center h-8 p-3">{{ structView(selected, path) }}</p>
        <ul class="absolute max-h-20 overflow-y-auto bg-white rounded-md p-3 w-full border">
            <li v-for="option in options" :key="option.Id" class="flex items-center gap-2">
                <Checkbox :checked="selected.some(item => item.Id == option.Id)" @click="pushOrRemove(option)" />
                <label :for="option.Id.toString()">{{ structView(option, path) }}</label>
            </li>
        </ul>
    </div>
</template>