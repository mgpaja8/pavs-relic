import React from "react";

export interface DropdownOption {
  id: string;
  label: string;
}

interface DropdownFilterProps {
  options: DropdownOption[];
  selectedOption?: string;
  onOptionChange: (option: string) => void;
}

export const DropdownFilter: React.FC<DropdownFilterProps> = ({
  options,
  selectedOption,
  onOptionChange,
}) => {
  return (
    <div className="dropdown-filter">
      <select
        value={selectedOption || ""}
        onChange={(e) => onOptionChange(e.target.value)}
      >
        <option value="">Select an option</option>

        {options.map((option) => (
          <option key={option.id} value={option.id}>
            {option.label}
          </option>
        ))}
      </select>
    </div>
  );
};
