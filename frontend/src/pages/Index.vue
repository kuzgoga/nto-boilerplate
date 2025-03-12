<script setup lang="ts">
import Calendar from '../components/calendar/Calendar.vue';
import { CalendarEvent } from '../components/calendar/event.class';

const events: CalendarEvent[] = []

const newEvent = new CalendarEvent(
  "Advanced algebra",
  "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Asperiores assumenda corporis doloremque et expedita molestias necessitatibus quam quas temporibus veritatis. Deserunt excepturi illum nobis perferendis praesentium repudiandae saepe sapiente voluptatem!",
  new Date("2025-03-11T12:05:00"),
  new Date("2025-03-12T13:35:00"),
  {
    color: 'pink',
  }
)

events.push(newEvent)
</script>

<template>
  <Calendar :events="events" />
</template>