import React from "react";

import { useAppDispatch } from "../../../../app/hooks";
import { getCompaniesSlim } from "../../../general/generalSlice";
import { CustomersList } from "../../components";

import "./CustomersRoute.scss";

export function CustomersRoute() {
  const dispatch = useAppDispatch();

  React.useEffect(() => {
    dispatch(getCompaniesSlim());
  }, []);

  return (
    <div className="customers-route-component">
      <CustomersList />
    </div>
  );
}
