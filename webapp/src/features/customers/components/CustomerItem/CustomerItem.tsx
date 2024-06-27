import React from "react";

import { useAppSelector } from "../../../../app/hooks";
import { Customer } from "../../../../models";

import "./CustomerItem.scss";

export interface CustomerItemProps {
  customer: Customer;
}

export function CustomerItem(props: CustomerItemProps) {
  const { customer } = props;

  const company = useAppSelector((state) =>
    state.general.companiesSlim.find((c) => c.id === customer.companyId)
  );

  return (
    <div className="customer-item-component">
      <div className="name">{`${customer.firstName} ${customer.lastName}`}</div>
      <div className="company">{company?.name}</div>
    </div>
  );
}
