<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import Service from './receiver.service'
import type { Scheme } from '../types/scheme.type'
import { Receiver } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services'
import { ref } from 'vue'
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import ExporterService from '../exporter/exporter.service'
const exporterService = new ExporterService

import WoodspectypeService from '../woodspectype/woodspectype.service'
const woodspectypeService = new WoodspectypeService


const service = new Service

const items = ref<Receiver[]>([])

const load = async () => {
  
  (scheme as any).Exporter.type!.nested!.values = await exporterService.readAll();
  
  (scheme as any).Spec.type!.nested!.values = await woodspectypeService.readAll();
  
  items.value = await service.sort(sortOptions.value) ;
  return items.value;
};

onMounted(() => {
	load()	
})

const scheme: Scheme<Receiver> = reactive({
    entityId: "ReceiverId",

  Id: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  ExporterId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  Exporter: {
  russian: "Поставщик",
  type: {
    nested: {
     values: [],
     field: ['Name']
   }, 
 },
},

  ExporterQuantity: {
  russian: "Количество сырья (по данным поставщика), кубометр",
  type: {
    primitive: "number",
 },
},

  FactoryQuantity: {
  russian: "Количество сырья (по данным завода), кубометр",
  type: {
    primitive: "number",
 },
},

  Description: {
  russian: "Описание",
  type: {
    primitive: "string",
 },
},

  Spec: {
  russian: "Номеклатура",
  type: {
    nested: {
     values: [],
     field: ['Name']
   }, 
 },
},

  SpecId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  CreatedAt: {
  readonly: true,
  type: {
    primitive: "date",
 },
},

})

const getDefaults = () => getDefaultValues(scheme)

const validate: Validate<Receiver> = entity => {
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
