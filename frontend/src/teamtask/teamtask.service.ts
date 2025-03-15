import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/teamtaskservice";
import type { TeamTask } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class TeamTaskService implements IService<TeamTask> {
  async read(id: number) {
    return (await GetById(id)) as TeamTask;
  }

  async readAll() {
    return (await GetAll()) as TeamTask[];
  }

  async create(item: TeamTask) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: TeamTask) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as TeamTask[];
  }

  async sort(options: SortOptions<TeamTask>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as TeamTask[];
  }
}
