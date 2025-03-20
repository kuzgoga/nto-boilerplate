<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import Service from './workresult.service'
import type { Scheme } from '../types/scheme.type'
import { WorkResult } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services'
import { ref } from 'vue'
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import WoodspectypeService from '../woodspectype/woodspectype.service'
const woodspectypeService = new WoodspectypeService

import WoodspecService from '../woodspec/woodspec.service'
import { GetCompMaterialSpecs, GetCompProductSpecs } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/workresultservice'
const woodspecService = new WoodspecService


const service = new Service

const items = ref<WorkResult[]>([])

const load = async () => {
  
  (scheme as any).Material.type!.nested!.values = await GetCompMaterialSpecs();
  
  (scheme as any).ProductSpec.type!.nested!.values = await GetCompProductSpecs();
  
  items.value = await service.sort(sortOptions.value) ;
  return items.value;
};

onMounted(() => {
	load()	
})

const scheme: Scheme<WorkResult> = reactive({
    entityId: "WorkResultId",

  Id: {
  russian: "Номер номенклатуры",
  type: {
    primitive: "number",
 },
},

  MaterialId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  Material: {
  russian: "Номенклатура материала",
  type: {
    nested: {
     values: [],
     field: ['Name']
   }, 
 },
},

  MaterialQuantity: {
  russian: "Количество материала",
  type: {
    primitive: "number",
 },
},

  ProductSpecId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  ProductSpec: {
  russian: "Номенклатура продукта",
  type: {
    nested: {
     values: [],
     field: ['Id']
   }, 
 },
},

  WorkDate: {
  russian: "Дата работы цеха",
 date: true,
  type: {
    primitive: "date",
 },
},

})

const getDefaults = () => getDefaultValues(scheme)

const validate: Validate<WorkResult> = entity => {
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
