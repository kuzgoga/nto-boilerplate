<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./author.service.ts";
import type { Scheme } from "../types/scheme.type";
import { Author } from "../../bindings/app/internal/services";

import PostService from "../post/post.service.ts";
const postService = new PostService();

const service = new Service();

onMounted(async () => {
  (scheme as any).Posts.type!.nested!.values = await postService.readAll();
});

const scheme: Scheme<Author> = reactive({
  Id: {
    hidden: true,
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

  Posts: {
    russian: "Посты",
    many: true,
    type: {
      nested: {
        values: [],
        field: ["Text"],
      },
    },
  },
});

const getDefaults = () => getDefaultValues(scheme);
</script>

<template>
  <main class="w-screen h-screen">
    <Table :scheme :service :getDefaults></Table>
  </main>
</template>
