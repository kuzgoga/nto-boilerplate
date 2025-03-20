<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import Service from './drymode.service'
import type { Scheme } from '../types/scheme.type'
import { DryMode } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services'
import { ref } from 'vue'
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import WoodspecService from '../woodspec/woodspec.service'
const woodspecService = new WoodspecService


const service = new Service

const items = ref<DryMode[]>([])

const load = async () => {
  
  (scheme as any).WetMaterials.type!.nested!.values = await woodspecService.readAll();
  
  items.value = await service.sort(sortOptions.value) ;
  return items.value;
};

onMounted(() => {
	load()	
})

const scheme: Scheme<DryMode> = reactive({
    entityId: "DryModeId",

  Id: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  Name: {
  russian: "Имя",
  type: {
    primitive: "string",
 },
},

  WetMaterialsId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  WetMaterials: {
  russian: "Номенклатура сырых пиломатериалов",
  type: {
    nested: {
     values: [],
     field: ['']
   }, 
 },
},

  HumidityPercent: {
  russian: "Итоговый процент влажности сухих пиломатериалов",
  type: {
    primitive: "number",
 },
},

  ProcentYsyshki: {
  russian: "Процент усушки",
  type: {
    primitive: "number",
 },
},

})

const getDefaults = () => getDefaultValues(scheme)

const validate: Validate<DryMode> = entity => {
  return {
    status: 'success'
  }
};

const search = async (input: string) => {
    items.value = await service.search(input)
}

const sortOptions = ref(getDefaultSortOptions(scheme))

</script>

<template>
	<main class="w-screen h-screen">
		<Table :scheme :service :get-defaults :load :items :validate @on-search="search" v-model:sort-options="sortOptions"></Table>
	</main>
</template>
