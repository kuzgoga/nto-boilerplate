import type { IEntity } from "../../types/entity.type";
import type { Scheme } from "../../types/scheme.type";
import { getTomorrow } from "../date/getters";

export const getDefaultValues = <T extends IEntity>(scheme: Scheme<T>) => {
    const keys = Object.keys(scheme) as (keyof typeof scheme)[]
    let obj: any = {}

    for (let key of keys) {
        const primitive = scheme[key]?.type?.primitive
        if (primitive == 'string' || primitive == 'multiple') {
            obj[key] = ''
        } else if (primitive == 'date') {
            obj[key] = getTomorrow()
        } else if (primitive == 'boolean') {
            obj[key] = false
        } else if (primitive == 'number') {
            obj[key] = 1
        } else if (scheme[key].type?.many) {
            obj[key] = []
        } else if (scheme[key]?.type?.nested?.values) {
            obj[key] = scheme[key].type.nested.values[0]
        }
    }

    return obj as T
}