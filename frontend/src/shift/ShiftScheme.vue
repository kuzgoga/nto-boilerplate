<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./shift.service";
import type { Scheme } from "../types/scheme.type";
import { Shift } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import ProducttypeService from "../producttype/producttype.service";
const producttypeService = new ProducttypeService();

import WorkareaService from "../workarea/workarea.service";
const workareaService = new WorkareaService();

const service = new Service();

const items = ref<Shift[]>([]);

const load = async () => {
  (scheme as any).ProductType.type!.nested!.values =
    await producttypeService.readAll();

  (scheme as any).WorkArea.type!.nested!.values =
    await workareaService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<Shift> = reactive({
  entityId: "ShiftId",

  Id: {
    russian: "ID",
    readonly: true,
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

  ProductTypeId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  ProductType: {
    type: {
      nested: {
        values: [],
        field: ["Name"],
      },
    },
  },

  ProductAmount: {
    russian: "Количество продукции",
    type: {
      primitive: "number",
    },
  },

  ShiftDate: {
    russian: "Дата смены",
    date: true,
    type: {
      primitive: "date",
    },
  },

  WorkAreaId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  WorkArea: {
    type: {
      nested: {
        values: [],
        field: ["Name"],
      },
    },
  },

  CreatedAt: {
    russian: "Дата создания",
    readonly: true,
    date: true,
    type: {
      primitive: "date",
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<Shift> = (entity) => {
  return {
    status: "success",
  };
};

const search = async (input: string) => {
  items.value = await service.search(input);
};

const sortOptions = ref(getDefaultSortOptions(scheme));
</script>

<template>
  <main class="w-screen h-screen">
    <Table
      :scheme
      :service
      :get-defaults
      :load
      :items
      :validate
      @on-search="search"
      v-model:sort-options="sortOptions"
    ></Table>
  </main>
</template>
