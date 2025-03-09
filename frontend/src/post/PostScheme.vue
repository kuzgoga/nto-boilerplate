<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./post.service.ts";
import type { Scheme } from "../types/scheme.type";
import { Post } from "../../bindings/app/internal/services";

import AuthorService from "../author/author.service.ts";
const authorService = new AuthorService();

const service = new Service();

onMounted(async () => {
  (scheme as any).Author.type!.nested!.values = await authorService.readAll();
});

const scheme: Scheme<Post> = reactive({
  entityId: "PostId",
  Id: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  Text: {
    russian: "Текст",
    type: {
      primitive: "string",
    },
  },

  Deadline: {
    russian: "Дедлайн",
    date: true,
    type: {
      primitive: "date",
    },
  },

  CreatedAt: {
    readonly: true,
    date: true,
    type: {
      primitive: "date",
    },
  },

  AuthorId: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },

  Author: {
    russian: "Автор",
    type: {
      nested: {
        values: [],
        field: ["Name"],
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
