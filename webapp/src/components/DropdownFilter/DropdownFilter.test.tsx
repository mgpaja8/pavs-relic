// DropdownFilter.test.tsx
import React from "react";
import { render, fireEvent, screen } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";
import { DropdownFilter, DropdownOption } from "./DropdownFilter";

describe("DropdownFilter component", () => {
  const mockOnOptionChange = jest.fn();

  const options: DropdownOption[] = [
    { id: "1", label: "Option 1" },
    { id: "2", label: "Option 2" },
    { id: "3", label: "Option 3" },
  ];

  const defaultProps = {
    options: options,
    selectedOption: "",
    onOptionChange: mockOnOptionChange,
  };

  it("renders with options and default select message", () => {
    render(<DropdownFilter {...defaultProps} />);

    const selectElement = screen.getByRole("combobox");
    expect(selectElement).toBeInTheDocument();
    expect(selectElement).toHaveValue("");

    const defaultOption = screen.getByText("Select an option");
    expect(defaultOption).toBeInTheDocument();

    options.forEach((option) => {
      const renderedOption = screen.getByText(option.label);
      expect(renderedOption).toBeInTheDocument();
    });
  });

  it("calls onOptionChange when an option is selected", () => {
    render(<DropdownFilter {...defaultProps} />);

    const selectElement = screen.getByRole("combobox");
    fireEvent.change(selectElement, { target: { value: "2" } });

    expect(mockOnOptionChange).toHaveBeenCalledTimes(1);
    expect(mockOnOptionChange).toHaveBeenCalledWith("2");
  });

  it("renders with pre-selected option", () => {
    render(<DropdownFilter {...defaultProps} selectedOption="2" />);

    const selectElement = screen.getByRole("combobox");
    expect(selectElement).toHaveValue("2");
  });
});
