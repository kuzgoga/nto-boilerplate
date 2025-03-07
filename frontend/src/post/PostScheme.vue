<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import S from './post.service.ts'
import type { Scheme } from '../types/scheme.type'
import { Post } from '../../bindings/app/internal/services/models.ts'

const service = new S

onMounted(async () => {
  	
})


const scheme: Scheme<Post> = reactive({
Id:{
  type: {
    primitive: "number",
  },
}, 
Text:{
  type: {
    primitive: "string",
  },
}, 
CreatedAt:{
  type: {
    primitive: "number",
  },
}, 

})

const getDefaults = ()  => getDefaultValues(scheme)

</script>

<template>
	<main class="w-screen h-screen">
		<Table :scheme :service :getDefaults></Table>
	</main>
</template>
