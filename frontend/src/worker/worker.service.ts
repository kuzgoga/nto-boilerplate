import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/workerservice";
import type { Worker } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class WorkerService implements IService<Worker> {
  async read(id: number) {
    return (await GetById(id)) as Worker;
  }

  async readAll() {
    return (await GetAll()) as Worker[];
  }

  async create(item: Worker) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: Worker) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as Worker[];
  }

  async sort(options: SortOptions<Worker>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as Worker[];
  }
}
