<script setup lang="ts">
import Table from "../table/Table.vue";
import {onMounted, reactive, watch} from "vue";
import { getDefaultValues } from "../utils/structs/defaults.util";
import Service from "./post.service.ts";
import type { Scheme } from "../types/scheme.type";
import { Author, Post } from "../../bindings/app/internal/services";
import { ref } from "vue";
import type { Validate } from "../types/validate.type";

import AuthorService from "../author/author.service.ts";
const authorService = new AuthorService();

import PosttypeService from "../posttype/posttype.service.ts";
const posttypeService = new PosttypeService();

import CommentService from "../comment/comment.service.ts";
import { SortedByOrder } from "../../bindings/app/internal/services/postservice.ts";
import {getDefaultSortOptions} from "../utils/structs/default-sort-options.util.ts";

const commentService = new CommentService();

const service = new Service();

const items = ref<Post[]>([]);

const load = async () => {
    (scheme as any).Author.type!.nested!.values = await authorService.readAll();

    (scheme as any).PostType.type!.nested!.values =
        await posttypeService.readAll();

    (scheme as any).Comments.type!.nested!.values =
        await commentService.readAll();

    items.value = await service.sort(sortOptions.value) ;
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
        }
    },
});

const getDefaults = () => {
    return ({ ...getDefaultValues(scheme), AuthorId: 1, PostTypeId: 1 });
};

const validate: Validate<Post> = (entity) => {
    return {
        status: "success",
    };
};

const search = async (input: string) => {
    items.value = await service.search(input)
}

const sortOptions = ref(getDefaultSortOptions(scheme))


</script>

<template>
    <Table :scheme :service :get-defaults :load :items :validate @on-search="search" v-model:sort-options="sortOptions"></Table>
</template>
