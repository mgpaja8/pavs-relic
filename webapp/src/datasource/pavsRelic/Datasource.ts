import axios, { AxiosInstance } from 'axios';

import * as Interface from './Interface';
import * as normalizers from './normalizers';
import { Customer, CompanyBase } from '../../models';

export interface DatasourceConfig {
  baseURL: string;
}

export class Datasource implements Interface.Interface {
  private client: AxiosInstance;

  constructor(config: DatasourceConfig) {
    const { baseURL } = config;

    this.client = axios.create({
      baseURL,
    });
  }

  async getCompaniesSlim(): Promise<CompanyBase[]> {
    const { data } = await this.client.get('/companies/slim');

    return normalizers.getCompaniesSlim(data.companies);
  }

  async getCustomers(params: Interface.GetCustomersParams): Promise<Customer[]> {
    const urlParams = new URLSearchParams({
      search: params.search ?? '',
      company_id: params.companyId ?? '',
    })

    const { data } = await this.client.get(`/customers?${urlParams.toString()}`);

    return normalizers.getCustomers(data.customers || []);
  }
}
