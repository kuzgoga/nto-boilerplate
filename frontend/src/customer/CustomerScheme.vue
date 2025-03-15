<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./customer.service";
import type { Scheme } from "../types/scheme.type";
import { Customer } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import OrderService from "../order/order.service";
const orderService = new OrderService();

const service = new Service();

const items = ref<Customer[]>([]);

const load = async () => {
  (scheme as any).Orders.type!.nested!.values = await orderService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<Customer> = reactive({
  entityId: "CustomerId",

  Id: {
    russian: "ID",
    readonly: true,
    type: {
      primitive: "number",
    },
  },

  Title: {
    russian: "Название",
    type: {
      primitive: "string",
    },
  },

  Contact: {
    russian: "Контакт",
    type: {
      primitive: "string",
    },
  },

  Orders: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
      },
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<Customer> = (entity) => {
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
