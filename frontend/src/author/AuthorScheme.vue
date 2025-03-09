<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive, ref } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./author.service.ts";
import type { Scheme } from "../types/scheme.type";
import { Author } from "../../bindings/app/internal/services";

import PostService from "../post/post.service.ts";
import type { Validate } from "../types/validate.type.ts";
const postService = new PostService();

const service = new Service();

const items = ref<Author[]>([])

async function load() {
  (scheme as any).Posts.type!.nested!.values = await postService.readAll();
  items.value = await service.readAll();
  return items.value;
}

onMounted(async () => {
  load()
});

const scheme: Scheme<Author> = reactive({
  entityId: "AuthorId",
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

const validate: Validate<Author> = (data, mode): ReturnType<Validate<Author>> => {
  return {
    status: 'success',
  }
}
</script>

<template>
  <main class="w-screen h-screen">
    <Table :scheme :service :get-defaults :validate :items :load></Table>
  </main>
</template>
