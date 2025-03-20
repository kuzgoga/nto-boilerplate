<script setup lang="ts">
import { Button, DatePicker } from 'primevue';
import { reactive, ref } from 'vue';
import { GetResult } from '../../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/reportservice';
import {toTimestamp} from "../../utils/date/converters.util.ts";
const from = ref(new Date)

const to = ref(new Date)

const data = ref<any>()

const load = async () => {

    data.value = await GetResult(toTimestamp(from.value), toTimestamp(to.value))
}
</script>

<template>
    <div class="flex items-center flex-col gap-2 mt-2">
        <DatePicker v-model="from" class="w-56" placeholder="Начало" />
        <DatePicker v-model="to" class="w-56" placeholder="Конец" />
        <Button severity="secondary" class="w-56" @click="load">Загрузить</Button>
    </div>
    <p v-if="data">
        Распиленное лесосырье
        {{ data.RawProductAmount?.join(', ') }}
    </p>
    <p v-if="data">
        Количество
        {{ data.RawProductCount }}
    </p>
    <p v-else>
        Нет данных
    </p>
</template>