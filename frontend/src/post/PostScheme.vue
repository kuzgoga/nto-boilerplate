<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util.ts'
import S from './post.service.ts'
import type { Scheme } from '../types/scheme.type'
import { Post } from '../../bindings/app/internal/services/models.ts'

const service = new S

onMounted(async () => {

})


const scheme: Scheme<Post> = reactive({
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
  CreatedAt: {
    hidden: true,
    type: {
      primitive: "number",
    },
  },
})

const getDefaults = () => getDefaultValues(scheme)

</script>

<template>
  <main class="w-screen h-screen">
    <Table :scheme :service :getDefaults></Table>
  </main>
</template>
