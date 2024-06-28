// Search.test.tsx
import React from "react";
import { render, fireEvent, screen } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";
import { Search, SearchProps } from "./Search";

describe("Search component", () => {
  const mockOnInputChange = jest.fn();

  const defaultProps: SearchProps = {
    search: "",
    placeholder: "Search...",
    onInputChange: mockOnInputChange,
  };

  it("renders with placeholder and value correctly", () => {
    render(<Search {...defaultProps} />);

    const inputElement = screen.getByPlaceholderText("Search...");
    expect(inputElement).toBeInTheDocument();
    expect(inputElement).toHaveValue("");

    // Simulate typing in the input field
    fireEvent.change(inputElement, { target: { value: "test search" } });

    // Check if onChange callback was called
    expect(mockOnInputChange).toHaveBeenCalledTimes(1);
    expect(mockOnInputChange).toHaveBeenCalledWith(expect.any(Object));
  });

  it("renders with custom placeholder", () => {
    const { rerender } = render(
      <Search {...defaultProps} placeholder="Custom placeholder" />
    );

    const inputElement = screen.getByPlaceholderText("Custom placeholder");
    expect(inputElement).toBeInTheDocument();
  });
});
