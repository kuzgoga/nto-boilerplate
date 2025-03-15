import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/workshopservice";
import type { Workshop } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class WorkshopService implements IService<Workshop> {
  async read(id: number) {
    return (await GetById(id)) as Workshop;
  }

  async readAll() {
    return (await GetAll()) as Workshop[];
  }

  async create(item: Workshop) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: Workshop) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as Workshop[];
  }

  async sort(options: SortOptions<Workshop>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as Workshop[];
  }
}
