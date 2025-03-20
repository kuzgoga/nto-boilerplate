import { GetAll, Create, Delete, GetById, Update, Count, SortedByOrder, SearchByAllTextFields } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/exporterservice"
import type { Exporter } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services"
import type { IService } from "../types/service.type"
import type { SortOptions } from "../types/sort-options.type";


export default class ExporterService implements IService<Exporter> {
	async read(id: number) {
		return await GetById(id) as Exporter
	}

	async readAll() {
		return await GetAll() as Exporter[]
	}

	async create(item: Exporter) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}

	async update(item: Exporter) {
		await Update(item)
	}

	async count() {
		return await Count()
	}

	async search(input: string) {
		return await SearchByAllTextFields(input) as Exporter[]
	}

	async sort(options: SortOptions<Exporter>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => {
        if (item[1] !== 'NONE') {
          return ({
            Name: item[0],
            Order: item[1],
          })
        }
      }).filter(item => !!item)
    )) as Exporter[];
  }
}