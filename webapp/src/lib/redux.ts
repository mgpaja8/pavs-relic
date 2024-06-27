export interface ApiError {
  code: number;
  detail: string;
  title: string;
}

export const defaultError: ApiError = {
  code: -1,
  detail: '',
  title: ''
};

export interface APIResult<T> {
  loading: boolean;
  error?: ApiError;
  value?: T;
}

export function emptyAPIResult<T>(): APIResult<T> {
  return {
    loading: false,
    error: undefined,
    value: undefined
  };
}

export function loadingAPIResult<T>(): APIResult<T> {
  return {
    loading: true,
    error: undefined,
    value: undefined
  };
}

export function errorAPIResult<T>(error: ApiError): APIResult<T> {
  return {
    loading: false,
    error,
    value: undefined
  };
}

export function successAPIResult<T>(value: T): APIResult<T> {
  return {
    loading: false,
    error: undefined,
    value
  };
}
