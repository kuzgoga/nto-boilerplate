<script setup lang="ts">
import Table from "../table/Table.vue";
import { onMounted, reactive } from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./post.service.ts";
import type { Scheme } from "../types/scheme.type";
import { Post } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";

import AuthorService from "../author/author.service.ts";
const authorService = new AuthorService();

import PosttypeService from "../posttype/posttype.service.ts";
const posttypeService = new PosttypeService();

import CommentService from "../comment/comment.service.ts";
const commentService = new CommentService();

const service = new Service();

const items = ref<Post[]>([]);

const load = async () => {
    (scheme as any).Author.type!.nested!.values = await authorService.readAll();

    (scheme as any).PostType.type!.nested!.values =
        await posttypeService.readAll();

    (scheme as any).Comments.type!.nested!.values =
        await commentService.readAll();

    items.value = await service.readAll();
    return items.value;
};

onMounted(async () => {
    await load();
});

const scheme: Scheme<Post> = reactive({
    entityId: "PostId",

    Id: {
        hidden: true,
        russian: 'Id',
        type: {
            primitive: "number",
        }
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
        russian: "Дата создания",
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

    PostTypeId: {
        hidden: true,
        type: {
            primitive: "number",
        },
    },

    PostType: {
        russian: "Тип",
        type: {
            nested: {
                values: [],
                field: ["Name"],
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
        customWindow: {
            create: true,
        }
    },
});

const getDefaults = () => getDefaultValues(scheme);

const validate: Validate<Post> = (entity) => {
    return {
        status: "success",
    };
};
</script>

<template>
    <main class="w-screen h-screen">
        <Table :scheme :service :get-defaults :load :items :validate>
            <template #CommentsCreate="{ data }">
                <p>{{ data.Comments }}</p>
            </template>
        </Table>
    </main>
</template>
