import { Customer, CompanyBase } from "../../models";

export interface GetCustomersParams {
  search?: string;
  companyId?: string;
}

export interface Interface {
  getCompaniesSlim: () => Promise<CompanyBase[]>;

  getCustomers: (params: GetCustomersParams) => Promise<Customer[]>;
}
