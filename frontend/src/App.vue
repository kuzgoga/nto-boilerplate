<script setup lang="ts">
import { ref } from 'vue';
import type { IEntity } from './types/entity.type';
import type { IService } from './types/service.type';
import type { Scheme } from './types/scheme.type';
import Table from './table/Table.vue';
import { getDefaultValues } from './utils/structs/defaults.util';

class Entity implements IEntity {
    constructor(public Id: number, public Name: string, public Region: string) { }
}

class Service implements IService<Entity> {
    private readonly entities = ref<Entity[]>([])
    private maxId = 0

    async create(data: Entity): Promise<void | never> {
        this.maxId++
        this.entities.value.push({ ...data, Id: this.maxId })
    }
    async delete(id: number): Promise<void | never> {
        this.entities.value = this.entities.value.filter(el => el.Id != id)
    }
    async read(id: number): Promise<Entity | undefined> {
        return this.entities.value.find(el => el.Id == id)
    }
    async readAll(): Promise<Entity[]> {
        return this.entities.value
    }
    async update(data: Entity): Promise<void | never> {
        this.entities.value = this.entities.value.map(el => el.Id == data.Id ? data : el)
    }
}

const service = new Service

const scheme: Scheme<Entity> = {
    Id: {
        hidden: true,
        russian: 'Id'
    },
    Name: {
        type: {
            primitive: 'string'
        },
        russian: 'Имя'
    },
    Region: {
        type: {
            nested: {
                field: [],
                values: ['kemerovo', 'kuzbass', 'berlin']
            }
        },
        russian: "Место жительства"
    },
}
</script>

<template>
    <Table :scheme :service :get-defaults="() => getDefaultValues(scheme)"></Table>
</template>
