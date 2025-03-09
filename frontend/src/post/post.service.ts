import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
} from "../../bindings/app/internal/services/postservice.ts";
import type { Post } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type.ts";

export default class PostService implements IService<Post> {
  async read(id: number) {
    return (await GetById(id)) as Post;
  }

  async readAll() {
    return (await GetAll()) as Post[];
  }

  async create(item: Post) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: Post) {
    await Update(item);
  }

  async count() {
    return await Count();
  }
}
