import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
} from "../../bindings/app/internal/services/commentservice.ts";
import type { Comment } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type.ts";

export default class CommentService implements IService<Comment> {
  async read(id: number) {
    return (await GetById(id)) as Comment;
  }

  async readAll() {
    return (await GetAll()) as Comment[];
  }

  async create(item: Comment) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: Comment) {
    await Update(item);
  }

  async count() {
    return await Count();
  }
}
