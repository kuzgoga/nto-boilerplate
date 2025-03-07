<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import S from './author.service.ts'
import type { Scheme } from '../types/scheme.type'
import { Author } from '../../bindings/app/internal/services/models.ts'

const service = new S

onMounted(async () => {
  	
})


const scheme: Scheme<Author> = reactive({
Id:{
  type: {
    primitive: "number",
  },
}, 
Name:{
  type: {
    primitive: "string",
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
