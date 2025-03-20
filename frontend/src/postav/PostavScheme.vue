<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import Service from './postav.service'
import type { Scheme } from '../types/scheme.type'
import { Postav } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services'
import { ref } from 'vue'
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import WoodspecService from '../woodspec/woodspec.service'
const woodspecService = new WoodspecService


const service = new Service

const items = ref<Postav[]>([])

const load = async () => {
  
  (scheme as any).CenterDoskaSpec.type!.nested!.values = await woodspecService.readAll();
  
  items.value = await service.sort(sortOptions.value) ;
  return items.value;
};

onMounted(() => {
	load()	
})

const scheme: Scheme<Postav> = reactive({
    entityId: "PostavId",

  Id: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  Name: {
  russian: "Наименование постава",
  type: {
    primitive: "string",
 },
},

  CenterDoskaSpecId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  CenterDoskaSpec: {
  russian: "Номеклатура центральной доски",
  type: {
    nested: {
     values: [],
     field: ['Id']
   }, 
 },
},

  CenterOutPercent: {
  russian: "Процент выхода центральной доски",
  type: {
    primitive: "number",
 },
},

  BacksideDoskaSpecId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  BacksideDoskaSpec: {
  russian: "Номенклатура боковой доски",
  type: {
    nested: {
     values: [],
     field: ['Id']
   }, 
 },
},

  BacksideOutPercent: {
  russian: "Процент выхода боковой доски",
  type: {
    primitive: "number",
 },
},

  OpilkiSpecId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  OpilkiSpec: {
  russian: "Номеклатура опилок",
  type: {
    nested: {
     values: [],
     field: ['Id']
   }, 
 },
},

  OpilkiOutPercent: {
  russian: "Процент выхода опилок",
  type: {
    primitive: "number",
 },
},

})

const getDefaults = () => getDefaultValues(scheme)

const validate: Validate<Postav> = entity => {
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
