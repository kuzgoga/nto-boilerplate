import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/workareaservice";
import type { WorkArea } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class WorkAreaService implements IService<WorkArea> {
  async read(id: number) {
    return (await GetById(id)) as WorkArea;
  }

  async readAll() {
    return (await GetAll()) as WorkArea[];
  }

  async create(item: WorkArea) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: WorkArea) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as WorkArea[];
  }

  async sort(options: SortOptions<WorkArea>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as WorkArea[];
  }
}
