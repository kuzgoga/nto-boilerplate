import { GetAll, Create, Delete, GetById, Update, Count, SortedByOrder, SearchByAllTextFields } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/woodspecservice"
import type { WoodSpec } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services"
import type { IService } from "../types/service.type"
import type { SortOptions } from "../types/sort-options.type";


export default class WoodSpecService implements IService<WoodSpec> {
	async read(id: number) {
		return await GetById(id) as WoodSpec
	}

	async readAll() {
		return await GetAll() as WoodSpec[]
	}

	async create(item: WoodSpec) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}

	async update(item: WoodSpec) {
		await Update(item)
	}

	async count() {
		return await Count()
	}

	async search(input: string) {
		return await SearchByAllTextFields(input) as WoodSpec[]
	}

	async sort(options: SortOptions<WoodSpec>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => {
        if (item[1] !== 'NONE') {
          return ({
            Name: item[0],
            Order: item[1],
          })
        }
      }).filter(item => !!item)
    )) as WoodSpec[];
  }
}