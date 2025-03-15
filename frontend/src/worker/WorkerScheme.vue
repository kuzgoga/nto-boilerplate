<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./worker.service";
import type { Scheme } from "../types/scheme.type";
import { Worker } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import TeamtaskService from "../teamtask/teamtask.service";
const teamtaskService = new TeamtaskService();

import WorkshopService from "../workshop/workshop.service";
const workshopService = new WorkshopService();

const service = new Service();

const items = ref<Worker[]>([]);

const load = async () => {
  (scheme as any).TeamTasks.type!.nested!.values =
    await teamtaskService.readAll();

  (scheme as any).Workshop.type!.nested!.values =
    await workshopService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<Worker> = reactive({
  entityId: "WorkerId",

  Id: {
    russian: "ID",
    readonly: true,
    type: {
      primitive: "number",
    },
  },

  Name: {
    russian: "Имя",
    type: {
      primitive: "string",
    },
  },

  TeamTasks: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
      },
    },
  },

  Workshop: {
    type: {
      nested: {
        values: [],
        field: ["Name"],
      },
    },
  },

  WorkshopId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<Worker> = (entity) => {
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
