import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';
import logger from 'redux-logger';
import Redux from 'redux';

import generalReducer from '../features/general/generalSlice';
import customersReducer from '../features/customers/customersSlice';

const middlewares: Redux.Middleware<{}, any, Redux.Dispatch<Redux.AnyAction>>[] = [
  logger,
];

export const store = configureStore({
  reducer: {
    general: generalReducer,
    customers: customersReducer,
  },
  middleware: (defaultMiddleware) => defaultMiddleware().concat(middlewares)
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
