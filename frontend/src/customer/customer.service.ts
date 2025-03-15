import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/customerservice";
import type { Customer } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class CustomerService implements IService<Customer> {
  async read(id: number) {
    return (await GetById(id)) as Customer;
  }

  async readAll() {
    return (await GetAll()) as Customer[];
  }

  async create(item: Customer) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: Customer) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as Customer[];
  }

  async sort(options: SortOptions<Customer>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as Customer[];
  }
}
