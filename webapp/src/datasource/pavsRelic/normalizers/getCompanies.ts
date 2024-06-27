import { CompanyBase } from '../../../models';

export const getCompanySlim = (data: any): CompanyBase => {
  return {
    id: data.id,
    name: data.name,
  };
}

export const getCompaniesSlim = (data: any): CompanyBase[] => {
  return data.map(getCompanySlim);
}
