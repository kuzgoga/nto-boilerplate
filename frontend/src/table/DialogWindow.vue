<script setup lang="ts" generic="T extends IEntity">
import { Button, DatePicker, Dialog, InputNumber, InputText, Select, Textarea, ToggleSwitch } from 'primevue';
import type { IEntity } from '../types/entity.type';
import type { Scheme } from '../types/scheme.type';
import type { IService } from '../types/service.type';
import { manyStructsView } from '../utils/structs/structs-view.util';
import { type UnwrapRef } from 'vue';
import { toDate, toTimestamp } from '../utils/date/converters.util';
import MultiSelect from '../components/selects/MultiSelect.vue';

const showCreate = defineModel<boolean>('show')
const createItem = defineModel<T>('item')
const items = defineModel<UnwrapRef<T[]>>('items')

const props = defineProps<{
    scheme: Scheme<T>,
    getDefaults: () => Partial<T>,
    service: IService<T>,
    updateMode?: boolean
}>()

type Key = keyof T
const keys = Object.keys(props.scheme) as Key[]

const emits = defineEmits<{
    (e: 'onSave', data: T): void
    (e: 'onSaveUpdate', data: T): void
    (e: 'onSaveCreate', data: T): void
}>()
</script>

<template>
    <Dialog v-model:visible="showCreate">
        <div class="flex flex-col justify-center gap-5">
            <div v-for="key in keys" v-show="!props.scheme[key].hidden && !props.scheme[key].readonly"
                class="flex items-center gap-5">
                <h1 class="w-[200px]">{{ props.scheme[key].russian }}</h1>
                <div>
                    <div v-if="props.scheme[key]?.customWindow?.[props.updateMode ? 'update' : 'create']">
                        <slot :name="<string>key + (props.updateMode ? 'Update' : 'Create')"></slot>
                    </div>
                    <div v-else-if="props.scheme[key]?.customWindow?.common">
                        <slot :name="<string>key"></slot>
                    </div>
                    <InputNumber class="w-[300px]" v-model:model-value="<number>createItem![key]"
                        v-else-if="props.scheme[key]?.type?.primitive === 'number'" />
                    <InputText class="w-[300px]" v-model:model-value="<string>createItem![key]"
                        v-else-if="props.scheme[key].type?.primitive === 'string'" />
                    <DatePicker class="w-[300px]" :default-value="toDate(createItem![key] as number)" @value-change="v => {
                        createItem![key] = toTimestamp(v as Date) as any
                    }" show-time
                        v-else-if="props.scheme[key].type?.primitive === 'date'" />
                    <Textarea class="w-[300px]" v-model="<string>createItem![key]"
                        v-else-if="props.scheme[key].type?.primitive === 'multiple'" />
                    <ToggleSwitch class="w-[300px]" v-model:model-value="<boolean>createItem![key]"
                        v-else-if="props.scheme[key].type?.primitive === 'boolean'" />
                    <Select v-else-if="props.scheme[key].type?.nested?.values && !props.scheme[key]?.many"
                        v-model:model-value="createItem![key]" :options="props.scheme[key].type.nested.values"
                        :placeholder="`Выберите ${props.scheme[key].russian}`" class="w-[300px]">
                        <template #option="{ option }">
                            {{ manyStructsView(option, props.scheme[key].type.nested.field) }}
                        </template>
                        <template #value="{ value }">
                            {{ manyStructsView(value, props.scheme[key].type.nested.field) }}
                        </template>
                    </Select>
                    <MultiSelect v-else-if="props.scheme[key].type?.nested?.values && props.scheme[key]?.many" class="w-[300px]"
                        v-model="<T[]>createItem![key]" :options="props.scheme[key].type.nested.values"
                        :path="props.scheme[key].type.nested.field" :entity-id="props.scheme.entityId" />
                </div>
            </div>
        </div>
        <template #footer>
            <Button severity="success" @click="async () => {
                if (props.updateMode) {
                    await props.service.update(createItem as T)
                    await emits('onSaveUpdate', createItem as T)
                    await emits('onSave', createItem as T)
                } else {
                    await props.service.create(createItem as T)
                    await emits('onSaveCreate', createItem as T)
                    await emits('onSave', createItem as T)
                }
                items = await props.service.readAll() as UnwrapRef<T[]>
                showCreate = false
            }">Сохранить</Button>
        </template>
    </Dialog>
</template>