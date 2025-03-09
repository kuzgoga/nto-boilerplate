<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./comment.service.ts";
import type { Scheme } from "../types/scheme.type";
import { Comment } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";

import AuthorService from "../author/author.service.ts";
const authorService = new AuthorService();

import PostService from "../post/post.service.ts";
const postService = new PostService();

const service = new Service();

const items = ref<Comment[]>([]);

const load = async () => {
  (scheme as any).Author.type!.nested!.values = await authorService.readAll();

  (scheme as any).Posts.type!.nested!.values = await postService.readAll();

  items.value = await service.readAll();
  return items.value;
};

onMounted(async () => {
  load();
});

const scheme: Scheme<Comment> = reactive({
  entityId: "CommentId",

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

const validate: Validate<Comment> = (entity) => {
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
