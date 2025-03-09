import type { IEntity } from "./entity.type";

export type Validate<T extends IEntity> = (data: T, mode: "update" | "create") => { status: "error" | "success", message?: string }