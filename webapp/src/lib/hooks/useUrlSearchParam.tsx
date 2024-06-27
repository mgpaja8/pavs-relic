import { useState, useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";

export function useUrlSearchParam(param: string) {
  const navigate = useNavigate();
  const location = useLocation();

  const [value, setValue] = useState(() => {
    const searchParams = new URLSearchParams(location.search);
    return searchParams.get(param) || "";
  });

  useEffect(() => {
    const searchParams = new URLSearchParams(location.search);
    if (value) {
      searchParams.set(param, value);
    } else {
      searchParams.delete(param);
    }
    navigate({ search: searchParams.toString() }, { replace: true });
  }, [value, param, navigate, location.search]);

  return [value, setValue] as const;
}
