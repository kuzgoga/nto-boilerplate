<script setup lang="ts">
import { Button } from 'primevue';
import NavCard from '../components/cards/NavCard.vue';

</script>

<template>
  <div class="flex w-screen h-screen items-center gap-5 justify-center">
    <NavCard :title="'Пользователь'" :to="'/user'" :content="'Может кидать зигу и не может пользоваться туалетной бумагой'" />
  </div>
  <footer class="fixed w-full bottom-10 flex items-center justify-center gap-2">
    <Button severity="secondary">Импортировать данные</Button>
    <Button severity="secondary">Экспортировать данные</Button>
  </footer>
</template>