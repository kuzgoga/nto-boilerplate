import {
  GetAll,
  Create,
  Delete,
  GetById,
  Update,
  Count,
  SortedByOrder,
  SearchByAllTextFields,
} from "../../bindings/app/internal/services/producttypeservice";
import type { ProductType } from "../../bindings/app/internal/services";
import type { IService } from "../types/service.type";
import type { SortOptions } from "../types/sort-options.type";

export default class ProductTypeService implements IService<ProductType> {
  async read(id: number) {
    return (await GetById(id)) as ProductType;
  }

  async readAll() {
    return (await GetAll()) as ProductType[];
  }

  async create(item: ProductType) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: ProductType) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as ProductType[];
  }

  async sort(options: SortOptions<ProductType>) {
    return (await SortedByOrder(
      Object.entries(options).map((item) => ({
        Name: item[0],
        Order: item[1],
      })),
    )) as ProductType[];
  }
}
