<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./order.service";
import type { Scheme } from "../types/scheme.type";
import { Order } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import ProducttypeService from "../producttype/producttype.service";
const producttypeService = new ProducttypeService();

import CustomerService from "../customer/customer.service";
const customerService = new CustomerService();

import TaskService from "../task/task.service";
const taskService = new TaskService();

const service = new Service();

const items = ref<Order[]>([]);

const load = async () => {
  (scheme as any).ProductType.type!.nested!.values =
    await producttypeService.readAll();

  (scheme as any).Customer.type!.nested!.values =
    await customerService.readAll();

  (scheme as any).Tasks.type!.nested!.values = await taskService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<Order> = reactive({
  entityId: "OrderId",

  Id: {
    russian: "ID",
    readonly: true,
    type: {
      primitive: "number",
    },
  },

  Status: {
    russian: "Статус",
    type: {
      primitive: "string",
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

  ProductAmount: {
    russian: "Количество продукции",
    type: {
      primitive: "number",
    },
  },

  CustomerId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  Customer: {
    russian: "Клиент",
    type: {
      nested: {
        values: [],
        field: ["Title"],
      },
    },
  },

  Tasks: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
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

  DeadlineDate: {
    russian: "Крайний срок",
    date: true,
    type: {
      primitive: "date",
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<Order> = (entity) => {
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
