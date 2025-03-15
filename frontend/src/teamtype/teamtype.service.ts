import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/teamtypeservice";
import type { TeamType } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class TeamTypeService implements IService<TeamType> {
  async read(id: number) {
    return (await GetById(id)) as TeamType;
  }

  async readAll() {
    return (await GetAll()) as TeamType[];
  }

  async create(item: TeamType) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: TeamType) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as TeamType[];
  }

  async sort(options: SortOptions<TeamType>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as TeamType[];
  }
}
