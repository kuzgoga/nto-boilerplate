<script setup lang="ts">
import { Qalendar } from "qalendar";

const config = {
  month: {
    showTrailingAndLeadingDates: false,
  },
  locale: 'ru-RU',
  defaultMode: 'month',
  isSilent: true,
};
</script>

<template>
  <Qalendar :config></Qalendar>
</template>

<style>
 @import "qalendar/dist/style.css";
</style>