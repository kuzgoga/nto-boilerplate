import { GetAll, Create, Delete, GetById, Update, Count, SortedByOrder, SearchByAllTextFields } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/workresultservice"
import type { WorkResult } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services"
import type { IService } from "../types/service.type"
import type { SortOptions } from "../types/sort-options.type";


export default class WorkResultService implements IService<WorkResult> {
	async read(id: number) {
		return await GetById(id) as WorkResult
	}

	async readAll() {
		return await GetAll() as WorkResult[]
	}

	async create(item: WorkResult) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}

	async update(item: WorkResult) {
		await Update(item)
	}

	async count() {
		return await Count()
	}

	async search(input: string) {
		return await SearchByAllTextFields(input) as WorkResult[]
	}

	async sort(options: SortOptions<WorkResult>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => {
        if (item[1] !== 'NONE') {
          return ({
            Name: item[0],
            Order: item[1],
          })
        }
      }).filter(item => !!item)
    )) as WorkResult[];
  }
}