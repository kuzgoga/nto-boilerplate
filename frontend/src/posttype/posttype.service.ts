import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
} from "../../bindings/app/internal/services/posttypeservice.ts";
import type { PostType } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type.ts";

export default class PostTypeService implements IService<PostType> {
  async read(id: number) {
    return (await GetById(id)) as PostType;
  }

  async readAll() {
    return (await GetAll()) as PostType[];
  }

  async create(item: PostType) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: PostType) {
    await Update(item);
  }

  async count() {
    return await Count();
  }
}
