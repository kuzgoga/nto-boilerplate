<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./task.service";
import type { Scheme } from "../types/scheme.type";
import { Task } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import ProducttypeService from "../producttype/producttype.service";
const producttypeService = new ProducttypeService();

import WorkshopService from "../workshop/workshop.service";
const workshopService = new WorkshopService();

import OrderService from "../order/order.service";
const orderService = new OrderService();

import PreptaskService from "../preptask/preptask.service";
const preptaskService = new PreptaskService();

const service = new Service();

const items = ref<Task[]>([]);

const load = async () => {
  (scheme as any).ProductType.type!.nested!.values =
    await producttypeService.readAll();

  (scheme as any).Workshops.type!.nested!.values =
    await workshopService.readAll();

  (scheme as any).Order.type!.nested!.values = await orderService.readAll();

  (scheme as any).PrepTasks.type!.nested!.values =
    await preptaskService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<Task> = reactive({
  entityId: "TaskId",

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
    russian: "Тип",
    type: {
      nested: {
        values: [],
        field: ["Name"],
      },
    },
  },

  Workshops: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
      },
    },
  },

  OrderId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  Order: {
    russian: "Заказ",
    type: {
      nested: {
        values: [],
        field: ["Description"],
      },
    },
  },

  PrepTasks: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
      },
    },
  },

  ProductionStart: {
    russian: "Дата начала производства",
    date: true,
    type: {
      primitive: "date",
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

  Amount: {
    russian: "Количество",
    type: {
      primitive: "number",
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<Task> = (entity) => {
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
