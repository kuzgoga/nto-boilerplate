<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import Service from './exporter.service'
import type { Scheme } from '../types/scheme.type'
import { Exporter } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services'
import { ref } from 'vue'
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import ReceiverService from '../receiver/receiver.service'
const receiverService = new ReceiverService


const service = new Service

const items = ref<Exporter[]>([])

const load = async () => {
  
  (scheme as any).Receivers.type!.nested!.values = await receiverService.readAll();
  
  items.value = await service.sort(sortOptions.value) ;
  return items.value;
};

onMounted(() => {
	load()	
})

const scheme: Scheme<Exporter> = reactive({
    entityId: "ExporterId",

  Id: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  Name: {
  type: {
    primitive: 'string'
  },
  russian: 'Название'
},

  Description: {
    type: {
    primitive: 'string'
  },
  russian: 'Описание' 
 },
 Receivers: {
  hidden: true,
  type: {
    nested: {
      field: ['Id'],
      values: []
    }
  }
 }
})

const getDefaults = () => getDefaultValues(scheme)

const validate: Validate<Exporter> = entity => {
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
