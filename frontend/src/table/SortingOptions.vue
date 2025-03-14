<script setup lang="ts" generic="T extends IEntity">
import type {IEntity} from "../types/entity.type.ts";
import type {SortOptions} from "../types/sort-options.type.ts";
import type {Scheme} from "../types/scheme.type.ts";
import {Select} from 'primevue'
import {watch} from "vue";


const options = defineModel<SortOptions<T>>()
const optionsKeys = Object.keys(options.value as T) as (keyof T)[]

defineProps<{
  scheme: Scheme<T>
  load: () => Promise<T[]>
}>()
</script>

<template>
  <ul class="flex flex-col gap-2 native-border secondary-background  p-3 rounded-md">
    <li v-for="optionKey in optionsKeys" class="flex items-center justify-between w-full">
      <h1>{{ scheme[optionKey].russian }}</h1>
      <Select size="small" class="w-24" :options="['ASC', 'DESC']" v-model="options![optionKey]" @value-change="load"></Select>
    </li>
  </ul>
</template>
