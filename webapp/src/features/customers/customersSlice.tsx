import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

import { pavsRelicDatasource, GetCustomersParams } from "../../datasource";
import {
  APIResult,
  emptyAPIResult,
  loadingAPIResult,
  successAPIResult,
  errorAPIResult,
  ApiError,
  defaultError,
} from "../../lib/redux";
import { Customer } from "../../models";

export interface CustomersState {
  getCustomersAPI: APIResult<Customer[]>;

  customers: Customer[];
}

const initialState: CustomersState = {
  getCustomersAPI: emptyAPIResult(),

  customers: [],
};

export const getCustomers = createAsyncThunk<
  Customer[],
  GetCustomersParams,
  { rejectValue: ApiError }
>("customers/getCustomers", async (params, { rejectWithValue }) => {
  try {
    const data = await pavsRelicDatasource.getCustomers(params);

    return data;
  } catch (err: any) {
    const error = err.response.data?.errors?.[0] || defaultError;

    return rejectWithValue(error);
  }
});

export const customersSlice = createSlice({
  name: "customers",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getCustomers.pending, (state) => {
        state.getCustomersAPI = loadingAPIResult();
      })
      .addCase(getCustomers.fulfilled, (state, action) => {
        state.getCustomersAPI = successAPIResult(action.payload);
        state.customers = action.payload;
      })
      .addCase(getCustomers.rejected, (state, action) => {
        state.getCustomersAPI = errorAPIResult(action.payload!);
      });
  },
});

export default customersSlice.reducer;
