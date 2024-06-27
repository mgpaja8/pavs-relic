import React from "react";

import "./Search.scss";

export interface SearchProps {
  search: string;
  placeholder?: string;

  onInputChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

export function Search(props: SearchProps) {
  const { search, placeholder, onInputChange } = props;

  return (
    <div className="search-component">
      <input
        value={search}
        placeholder={placeholder}
        onChange={onInputChange}
      />
    </div>
  );
}
