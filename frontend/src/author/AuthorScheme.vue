<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./author.service.ts";
import type { Scheme } from "../types/scheme.type";
import { Author } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";

import PostService from "../post/post.service.ts";
const postService = new PostService();

import CommentService from "../comment/comment.service.ts";
const commentService = new CommentService();

const service = new Service();

const items = ref<Author[]>([]);

const load = async () => {
  (scheme as any).Posts.type!.nested!.values = await postService.readAll();

  (scheme as any).Comments.type!.nested!.values =
    await commentService.readAll();

  items.value = await service.readAll();
  return items.value;
};

onMounted(async () => {
  load();
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

  Comments: {
    russian: "Комментарии",
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

const validate: Validate<Author> = (entity) => {
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
