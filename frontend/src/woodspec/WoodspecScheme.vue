<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import Service from './woodspec.service'
import type { Scheme } from '../types/scheme.type'
import { WoodSpec } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services'
import { ref } from 'vue'
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";
import { Button, DatePicker, Dialog, InputNumber, InputText, Select, Textarea, ToggleSwitch } from 'primevue';
import WoodspectypeService from '../woodspectype/woodspectype.service'
const woodspectypeService = new WoodspectypeService


const service = new Service

const items = ref<WoodSpec[]>([])

const load = async () => {
  
  (scheme as any).WoodSpecType.type!.nested!.values = await woodspectypeService.readAll();
  
  items.value = await service.sort(sortOptions.value) ;
  return items.value;
};

onMounted(() => {
	load()	
})

const scheme: Scheme<WoodSpec> = reactive({
    entityId: "WoodSpecId",

  Id: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  WoodSpecTypeId: {
  hidden: true,
  type: {
    primitive: "number",
 },
},

  WoodSpecType: {
  russian: "Вид номеклатуры",
  type: {
    nested: {
     values: [],
     field: ['Name']
   }, 
 },
},

  Poroda: {
  russian: "Порода древесины",
  type: {
    primitive: "string",
 },
 customWindow: {
  common: true
 }
},

  Sort: {
  russian: "Сорт",
  type: {
    primitive: "string",
 },
 customWindow: {
  common: true
 }
},

  DlinaBrevna: {
  russian: "Длина бревна",
  type: {
    primitive: "string",
 },
 customWindow: {
  common: true
 }
},

  DiameterBrevna: {
  russian: "Диаметр бревна",
  type: {
    primitive: "string",
 },
 customWindow: {
  common: true
 }
},

  SechenieDoski: {
  russian: "Сечение доски",
  type: {
    primitive: "string",
 },
 customWindow: {
  common: true
 }
},

  ProcentVlajnosti: {
  russian: "Процент влажности",
  type: {
    primitive: "number",
 },
 customWindow: {
  common: true
 }
},

  DiameterGranyl: {
  russian: "Диаметр гранул",
  type: {
    primitive: "string",
 },
 customWindow: {
  common: true
 }
},

})

const getDefaults = () => getDefaultValues(scheme)

const validate: Validate<WoodSpec> = entity => {
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
		<Table name="Номенклатура" :scheme :service :get-defaults :load :items :validate @on-search="search" v-model:sort-options="sortOptions">
      <template #Poroda="{data}">
        <InputText v-model="data.Poroda" v-if="['Круглый лес', 'Сырые пиломатериалы', 'Сухие пиломатериалы', 'Опилки'].includes(data!.WoodSpecType!.Name as any)"/>
        <p v-else>Недоступно</p>
      </template>
      <template #Sort="{data}">
        <InputText v-model="data.Sort" v-if="['Сырые пиломатериалы', 'Сухие пиломатериалы', 'Круглый лес', 'Пеллеты'].includes(data.WoodSpecType.Name)" />
        <p v-else>Недоступно</p>
      </template>
      <template #DlinaBrevna="{data}" >
        <InputText v-model="data.DlinaBrevna" v-if="['Крглый лес'].includes(data.WoodSpecType.Name)" />
        <p v-else>Недоступно</p>

      </template>
      <template #ProcentVlajnosti="{data}" >
        <InputText v-model="data.ProcentVlajnosti" v-if="['Сухие пиломатериалы', 'Сырые пиломатериалы'].includes(data.WoodSpecType.Name)" />
        <p v-else>Недоступно</p>

      </template>
      <template #DiameterBrevna="{data}">
        <InputText v-model="data.DiameterBrevna" v-if="['Круглый лес'].includes(data.WoodSpecType.Name)" />
        <p v-else>Недоступно</p>

      </template>
      <template #DiameterGranyl="{data}">
        <InputText v-model="data.DiameterGranyl" v-if="['Пеллеты'].includes(data.WoodSpecType.Name)"/>
        <p v-else>Недоступно</p>
      </template>
      <template #SechenieDoski="{data}">
        <InputText v-model="data.SechenieDoski" v-if="['Сырые пиломатериалы', 'Сухие пиломатериалы'].includes(data.WoodSpecType.Name)" />
        <p v-else>Недоступно</p>
      </template>

    </Table>
	</main>
</template>
