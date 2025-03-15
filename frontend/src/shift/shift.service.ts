import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/shiftservice";
import type { Shift } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class ShiftService implements IService<Shift> {
  async read(id: number) {
    return (await GetById(id)) as Shift;
  }

  async readAll() {
    return (await GetAll()) as Shift[];
  }

  async create(item: Shift) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: Shift) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as Shift[];
  }

  async sort(options: SortOptions<Shift>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as Shift[];
  }
}
