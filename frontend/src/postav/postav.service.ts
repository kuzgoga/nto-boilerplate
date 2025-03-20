import { GetAll, Create, Delete, GetById, Update, Count, SortedByOrder, SearchByAllTextFields } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/postavservice"
import type { Postav } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services"
import type { IService } from "../types/service.type"
import type { SortOptions } from "../types/sort-options.type";


export default class PostavService implements IService<Postav> {
	async read(id: number) {
		return await GetById(id) as Postav
	}

	async readAll() {
		return await GetAll() as Postav[]
	}

	async create(item: Postav) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}

	async update(item: Postav) {
		await Update(item)
	}

	async count() {
		return await Count()
	}

	async search(input: string) {
		return await SearchByAllTextFields(input) as Postav[]
	}

	async sort(options: SortOptions<Postav>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => {
        if (item[1] !== 'NONE') {
          return ({
            Name: item[0],
            Order: item[1],
          })
        }
      }).filter(item => !!item)
    )) as Postav[];
  }
}