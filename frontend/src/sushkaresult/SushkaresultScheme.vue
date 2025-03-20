<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import Service from './sushkaresult.service'
import type { Scheme } from '../types/scheme.type'
import { SushkaResult } from '../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services'
import { ref } from 'vue'
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import WoodspecService from '../woodspec/woodspec.service'
import SushkaResultService from './sushkaresult.service'
const woodspecService = new WoodspecService


const service = new Service

const items = ref<SushkaResult[]>([])

const load = async () => {

  (scheme as any).MaterialSpec.type!.nested!.values = (await woodspecService.readAll()).filter(el => el.WoodSpecType.Name.includes('Сырые пиломатериалы'));
  (scheme as any).ProductSpec.type!.nested!.values = (await woodspecService.readAll()).filter(el => el.WoodSpecType.Name.includes('Сухие пиломатериалы'));
  items.value = (await service.sort(sortOptions.value));
  return items.value;
};

onMounted(() => {
  load()
})

const scheme: Scheme<SushkaResult> = reactive({
  entityId: "SushkaResultId",

  Id: {
    hidden: true,
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

  MaterialSpecId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  MaterialSpec: {
    russian: "Номер номенклатуры материала",
    type: {
      nested: {
        values: [],
        field: ['Id']
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
    russian: "Номер номенклатуры продукта",
    type: {
      nested: {
        values: [],
        field: ['Id']
      },
    },
  },

  ProductAmount: {
    type: {
      primitive: 'number',
    },
    russian: "Количество"
  },

  WorkDate: {
    russian: "Дата работы",
    date: true,
    type: {
      primitive: "date",
    },
  },

})

const getDefaults = () => getDefaultValues(scheme)

const validate: Validate<SushkaResult> = (entity: SushkaResult) => {
  if ((!entity.Description || entity.MaterialQuantity == undefined || entity.ProductAmount == undefined)) {
    return {
      status: 'error',
      message: 'Заполните все поля!'
    }
  }
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
    <Table :scheme :service :get-defaults :load :items :validate @on-search="search" v-model:sort-options="sortOptions">
    </Table>
  </main>
</template>
