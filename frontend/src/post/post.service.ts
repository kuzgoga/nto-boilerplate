import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count, SortedByOrder,
} from "../../bindings/app/internal/services/postservice.ts";
import type { Post } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type.ts";
import type {SortOptions} from "../types/sort-options.type.ts";
import {SortField} from "../../bindings/app/internal/utils";
import { SearchByAllTextFields } from "../../bindings/app/internal/services/postservice.ts";

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

  async sort(options: SortOptions<Post>) {
    return await SortedByOrder(Object.entries(options).map(item => ({Name: item[0], Order: item[1]}))) as Post[]
  }

  async search(input: string): Promise<Post[]> {
    return await SearchByAllTextFields(input) as Post[]
  }
}
