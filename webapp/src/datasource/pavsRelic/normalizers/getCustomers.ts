import { Customer } from '../../../models';

export const getCustomer = (data: any): Customer => {
  return {
    id: data.id,
    companyId: data.company_id,
    firstName: data.first_name,
    lastName: data.last_name,
  };
}

export const getCustomers = (data: any): Customer[] => {
  return data.map(getCustomer);
}
