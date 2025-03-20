import { GetAll, Create, Delete, GetById, Update, Count, SortedByOrder, SearchByAllTextFields } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services/receiverservice"
import type { Receiver } from "../../bindings/github.com/kuzgoga/nto-boilerplate/internal/services"
import type { IService } from "../types/service.type"
import type { SortOptions } from "../types/sort-options.type";


export default class ReceiverService implements IService<Receiver> {
	async read(id: number) {
		return await GetById(id) as Receiver
	}

	async readAll() {
		return await GetAll() as Receiver[]
	}

	async create(item: Receiver) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}

	async update(item: Receiver) {
		await Update(item)
	}

	async count() {
		return await Count()
	}

	async search(input: string) {
		return await SearchByAllTextFields(input) as Receiver[]
	}

	async sort(options: SortOptions<Receiver>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => {
        if (item[1] !== 'NONE') {
          return ({
            Name: item[0],
            Order: item[1],
          })
        }
      }).filter(item => !!item)
    )) as Receiver[];
  }
}