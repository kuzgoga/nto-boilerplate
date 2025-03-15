<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./workshop.service";
import type { Scheme } from "../types/scheme.type";
import { Workshop } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import WorkareaService from "../workarea/workarea.service";
const workareaService = new WorkareaService();

import TaskService from "../task/task.service";
const taskService = new TaskService();

import WorkerService from "../worker/worker.service";
const workerService = new WorkerService();

const service = new Service();

const items = ref<Workshop[]>([]);

const load = async () => {
  (scheme as any).WorkAreas.type!.nested!.values =
    await workareaService.readAll();

  (scheme as any).Tasks.type!.nested!.values = await taskService.readAll();

  (scheme as any).Workers.type!.nested!.values = await workerService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<Workshop> = reactive({
  entityId: "WorkshopId",

  Id: {
    russian: "ID",
    readonly: true,
    type: {
      primitive: "number",
    },
  },

  Name: {
    russian: "Наименование",
    type: {
      primitive: "string",
    },
  },

  WorkAreas: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
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

  Workers: {
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

const validate: Validate<Workshop> = (entity) => {
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
