import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/postservice";
import type { Post } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

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

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as Post[];
  }

  async sort(options: SortOptions<Post>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as Post[];
  }
}
