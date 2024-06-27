import React from "react";

import { useAppDispatch, useAppSelector } from "../../../../app/hooks";
import { getCustomers } from "../../customersSlice";
import { CustomerItem } from "../CustomerItem/CustomerItem";
import { Search, DropdownFilter } from "../../../../components";
import {
  useDebounce,
  useIsFirstRender,
  useUrlSearchParam,
} from "../../../../lib/hooks";
import { companiesToDropdownOptions } from "../../../../models";

import "./CustomersList.scss";

export function CustomersList() {
  const dispatch = useAppDispatch();
  const isFirstRender = useIsFirstRender();

  const [search, setSearch] = useUrlSearchParam("search");
  const [companyId, setCompanyId] = useUrlSearchParam("company_id");

  const customers = useAppSelector((state) => state.customers.customers);
  const getCustomersAPI = useAppSelector(
    (state) => state.customers.getCustomersAPI
  );
  const companiesSlim = useAppSelector((state) => state.general.companiesSlim);

  React.useEffect(() => {
    fetchCustomers();
  }, []);

  React.useEffect(() => {
    if (!isFirstRender) {
      fetchCustomersDebounced();
    }
  }, [search]);

  React.useEffect(() => {
    if (!isFirstRender) {
      fetchCustomers();
    }
  }, [companyId]);

  const fetchCustomers = () => {
    const payload = {
      search,
      companyId,
    };

    dispatch(getCustomers(payload));
  };

  const fetchCustomersDebounced = useDebounce(fetchCustomers, 500);

  const onSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    setSearch(e.target.value);
  };

  return (
    <div className="customers-list-component">
      <div className="header">
        <div className="title">Customers</div>

        <div className="filters-container">
          <Search
            search={search}
            placeholder="Search by first or last name"
            onInputChange={onSearchChange}
          />

          <DropdownFilter
            selectedOption={companyId}
            onOptionChange={setCompanyId}
            options={companiesToDropdownOptions(companiesSlim)}
          />
        </div>
      </div>

      <div className="customers-container">
        {getCustomersAPI.loading ? (
          <div>Loading...</div>
        ) : customers.length > 0 ? (
          customers.map((c) => <CustomerItem customer={c} key={c.id} />)
        ) : (
          <div className="no-results">No results for your filters.</div>
        )}
      </div>
    </div>
  );
}
