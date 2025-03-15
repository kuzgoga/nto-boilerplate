<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./teamtype.service";
import type { Scheme } from "../types/scheme.type";
import { TeamType } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

const service = new Service();

const items = ref<TeamType[]>([]);

const load = async () => {
  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<TeamType> = reactive({
  entityId: "TeamTypeId",

  Id: {
    russian: "ID",
    readonly: true,
    type: {
      primitive: "number",
    },
  },

  Name: {
    russian: "Название",
    type: {
      primitive: "string",
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<TeamType> = (entity) => {
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
