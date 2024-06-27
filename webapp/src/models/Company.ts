import { DropdownOption } from '../components'

export interface CompanyBase {
  id: string;
  name: string;
}

export const companiesToDropdownOptions = (companies: CompanyBase[]): DropdownOption[] => {
  return companies.map(company => ({
    id: company.id,
    label: company.name,
  }));
}
