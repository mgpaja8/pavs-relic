import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

import { pavsRelicDatasource } from "../../datasource";
import {
  APIResult,
  emptyAPIResult,
  loadingAPIResult,
  successAPIResult,
  errorAPIResult,
  ApiError,
  defaultError,
} from "../../lib/redux";
import { CompanyBase } from "../../models";

export interface GeneralState {
  getCompaniesSlimAPI: APIResult<CompanyBase[]>;

  companiesSlim: CompanyBase[];
}

const initialState: GeneralState = {
  getCompaniesSlimAPI: emptyAPIResult(),

  companiesSlim: [],
};

export const getCompaniesSlim = createAsyncThunk<
  CompanyBase[],
  void,
  { rejectValue: ApiError }
>("general/getCompaniesSlim", async (_, { rejectWithValue }) => {
  try {
    const data = await pavsRelicDatasource.getCompaniesSlim();

    return data;
  } catch (err: any) {
    const error = err.response.data?.errors?.[0] || defaultError;

    return rejectWithValue(error);
  }
});

export const generalSlice = createSlice({
  name: "general",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getCompaniesSlim.pending, (state) => {
        state.getCompaniesSlimAPI = loadingAPIResult();
      })
      .addCase(getCompaniesSlim.fulfilled, (state, action) => {
        state.getCompaniesSlimAPI = successAPIResult(action.payload);
        state.companiesSlim = action.payload;
      })
      .addCase(getCompaniesSlim.rejected, (state, action) => {
        state.getCompaniesSlimAPI = errorAPIResult(action.payload!);
      });
  },
});

export default generalSlice.reducer;
