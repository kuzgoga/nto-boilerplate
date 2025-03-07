import type { IEntity } from "./entity.type";
import type { Scheme } from "./scheme.type";
import type { IService } from "./service.type";

export interface TableProps<T extends IEntity> {
    service: IService<T>
    scheme: Scheme<T>
    name?: string
    getDefaults: () => Partial<T>
    validate?: (data: T) => never | void
}