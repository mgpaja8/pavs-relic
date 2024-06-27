import React from "react";
import { useNavigate, useLocation } from "react-router-dom";

import { useAppDispatch, useAppSelector } from "../../../../app/hooks";
import { getCustomers } from "../../customersSlice";
import { CustomerItem } from "../CustomerItem/CustomerItem";
import { Search } from "../../../../components";
import { useDebounce, useIsFirstRender } from "../../../../lib/hooks";

import "./CustomersList.scss";

export function CustomersList() {
  const dispatch = useAppDispatch();
  const isFirstRender = useIsFirstRender();
  const navigate = useNavigate();
  const location = useLocation();

  const searchParams = new URLSearchParams(location.search);
  const searchQuery = searchParams.get("search") || "";

  const [search, setSearch] = React.useState(searchQuery);

  const customers = useAppSelector((state) => state.customers.customers);
  const getCustomersAPI = useAppSelector(
    (state) => state.customers.getCustomersAPI
  );

  React.useEffect(() => {
    fetchCustomers();
  }, []);

  React.useEffect(() => {
    // Initial request will be covered by initial render hook (hook above)
    if (!isFirstRender) {
      fetchCustomersDebounced();
      const searchParams = new URLSearchParams(location.search);
      searchParams.set("search", search);
      navigate({ search: searchParams.toString() }, { replace: true });
    }
  }, [search]);

  const fetchCustomers = () => {
    const payload = {
      search,
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

        <div>
          <Search
            search={search}
            placeholder="Search by first or last name"
            onInputChange={onSearchChange}
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
