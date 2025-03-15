<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./workarea.service";
import type { Scheme } from "../types/scheme.type";
import { WorkArea } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import WorkshopService from "../workshop/workshop.service";
const workshopService = new WorkshopService();

import PreptaskService from "../preptask/preptask.service";
const preptaskService = new PreptaskService();

import ShiftService from "../shift/shift.service";
const shiftService = new ShiftService();

import TeamtaskService from "../teamtask/teamtask.service";
const teamtaskService = new TeamtaskService();

const service = new Service();

const items = ref<WorkArea[]>([]);

const load = async () => {
  (scheme as any).Workshop.type!.nested!.values =
    await workshopService.readAll();

  (scheme as any).PrepTasks.type!.nested!.values =
    await preptaskService.readAll();

  (scheme as any).Shifts.type!.nested!.values = await shiftService.readAll();

  (scheme as any).TeamTasks.type!.nested!.values =
    await teamtaskService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<WorkArea> = reactive({
  entityId: "WorkAreaId",

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

  Description: {
    russian: "Описание",
    type: {
      primitive: "string",
    },
  },

  Performance: {
    russian: "Производительность",
    type: {
      primitive: "number",
    },
  },

  WorkshopId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  Workshop: {
    russian: "Цех",
    type: {
      nested: {
        values: [],
        field: ["Name"],
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

  Shifts: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
      },
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
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<WorkArea> = (entity) => {
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
