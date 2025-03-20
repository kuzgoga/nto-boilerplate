
 <script setup lang="ts">
import { Button } from 'primevue';
import NavCard from '../components/cards/NavCard.vue';
import { useErrorStore } from '../stores/error.store'; 
import { useNavModalStore } from '../stores/nav-modal.store';
import { routes } from '../router/router';
import NavModal from '../components/modals/NavModal.vue';

const errorStore = useErrorStore()
const nm = useNavModalStore()

// const importFromExcel = async () => {
//   try {
//     await ImportAllEntities()
//   } catch(e) {
//     errorStore.summon((e as Error).message)
//   }
// }

// const exportFromExcel = async () => {
//   try {
//     await ExportAllEntities()
//   } catch(e) {
//     errorStore.summon((e as Error).message)
//   }
// }
</script>

<template>
    
  <div class="flex w-screen h-screen items-center gap-5 justify-center">
    <NavCard :title="'Приемка сырья'" :to="'/ps'" />
    <NavCard :title="'Производство'" :to="'/dev'" />
    <NavCard :title="'Отгрузка продукции'" :to="'/del'"  />
  </div>
</template>