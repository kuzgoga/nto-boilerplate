import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/preptaskservice";
import type { PrepTask } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class PrepTaskService implements IService<PrepTask> {
  async read(id: number) {
    return (await GetById(id)) as PrepTask;
  }

  async readAll() {
    return (await GetAll()) as PrepTask[];
  }

  async create(item: PrepTask) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: PrepTask) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as PrepTask[];
  }

  async sort(options: SortOptions<PrepTask>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as PrepTask[];
  }
}
