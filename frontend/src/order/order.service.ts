import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/orderservice";
import type { Order } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class OrderService implements IService<Order> {
  async read(id: number) {
    return (await GetById(id)) as Order;
  }

  async readAll() {
    return (await GetAll()) as Order[];
  }

  async create(item: Order) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: Order) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as Order[];
  }

  async sort(options: SortOptions<Order>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as Order[];
  }
}
