import { GetAll, Create, Delete, GetById, Update, Count, SortedByOrder, SearchByAllTextFields } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/drymodeservice"
import type { DryMode } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services"
import type { IService } from "../types/service.type"
import type { SortOptions } from "../types/sort-options.type";


export default class DryModeService implements IService<DryMode> {
	async read(id: number) {
		return await GetById(id) as DryMode
	}

	async readAll() {
		return await GetAll() as DryMode[]
	}

	async create(item: DryMode) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}

	async update(item: DryMode) {
		await Update(item)
	}

	async count() {
		return await Count()
	}

	async search(input: string) {
		return await SearchByAllTextFields(input) as DryMode[]
	}

	async sort(options: SortOptions<DryMode>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => {
        if (item[1] !== 'NONE') {
          return ({
            Name: item[0],
            Order: item[1],
          })
        }
      }).filter(item => !!item)
    )) as DryMode[];
  }
}