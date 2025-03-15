<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./teamtask.service";
import type { Scheme } from "../types/scheme.type";
import { TeamTask } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";
import { getDefaultSortOptions } from "../utils/structs/default-sort-options.util";

import TeamtypeService from "../teamtype/teamtype.service";
const teamtypeService = new TeamtypeService();

import WorkerService from "../worker/worker.service";
const workerService = new WorkerService();

import WorkareaService from "../workarea/workarea.service";
const workareaService = new WorkareaService();

const service = new Service();

const items = ref<TeamTask[]>([]);

const load = async () => {
  (scheme as any).TeamType.type!.nested!.values =
    await teamtypeService.readAll();

  (scheme as any).TeamLeader.type!.nested!.values =
    await workerService.readAll();

  (scheme as any).WorkArea.type!.nested!.values =
    await workareaService.readAll();

  items.value = await service.sort(sortOptions.value);
  return items.value;
};

onMounted(() => {
  load();
});

const scheme: Scheme<TeamTask> = reactive({
  entityId: "TeamTaskId",

  Id: {
    russian: "ID",
    readonly: true,
    type: {
      primitive: "number",
    },
  },

  TeamTypeId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  TeamType: {
    type: {
      nested: {
        values: [],
        field: ["Name"],
      },
    },
  },

  TeamLeaderId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  TeamLeader: {
    type: {
      nested: {
        values: [],
        field: ["Name"],
      },
    },
  },

  TeamMembers: {
    hidden: true,
    many: true,
    type: {
      nested: {
        values: [],
        field: [""],
      },
    },
  },

  WorkStartDate: {
    russian: "Дата начала работ",
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

  ShiftDuties: {
    russian: "Обязанности смены",
    type: {
      primitive: "string",
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<TeamTask> = (entity) => {
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
