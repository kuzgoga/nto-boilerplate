<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./posttype.service.ts";
import type { Scheme } from "../types/scheme.type";
import { PostType } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";

const service = new Service();

const items = ref<PostType[]>([]);

const load = async () => {
  items.value = await service.readAll();
  return items.value;
};

onMounted(async () => {
  load();
});

const scheme: Scheme<PostType> = reactive({
  entityId: "PostTypeId",

  Id: {
    hidden: true,
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

const validate: Validate<PostType> = (entity) => {
  return {
    status: "success",
  };
};
</script>

<template>
  <main class="w-screen h-screen">
    <Table :scheme :service :get-defaults :load :items :validate></Table>
  </main>
</template>
