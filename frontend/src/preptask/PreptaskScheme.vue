<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./preptask.service";
import type { Scheme } from "../types/scheme.type";
import { PrepTask } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import TaskService from "../task/task.service";
const taskService = new TaskService();

import WorkareaService from "../workarea/workarea.service";
const workareaService = new WorkareaService();

const service = new Service();

const items = ref<PrepTask[]>([]);

const load = async () => {
  (scheme as any).Task.type!.nested!.values = await taskService.readAll();

  (scheme as any).WorkArea.type!.nested!.values =
    await workareaService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<PrepTask> = reactive({
  entityId: "PrepTaskId",

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

  TaskId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  Task: {
    russian: "Задача",
    type: {
      nested: {
        values: [],
        field: ["Description"],
      },
    },
  },

  WorkAreaId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  WorkArea: {
    russian: "Рабочая зона",
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

  Deadline: {
    russian: "Крайний срок",
    date: true,
    type: {
      primitive: "date",
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<PrepTask> = (entity) => {
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
