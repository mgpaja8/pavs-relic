import { useRef, useEffect, useCallback } from "react";

type Timer = ReturnType<typeof setTimeout>;
type SomeFunction = (...args: any[]) => void;

export function useDebounce<Func extends SomeFunction>(
  func: Func,
  delay = 1000
) {
  const timer = useRef<Timer | null>(null);

  useEffect(() => {
    return () => {
      if (timer.current) {
        clearTimeout(timer.current);
      }
    };
  }, []);

  const debouncedFunction = useCallback(
    (...args: Parameters<Func>) => {
      if (timer.current) {
        clearTimeout(timer.current);
      }
      timer.current = setTimeout(() => {
        func(...args);
      }, delay);
    },
    [func, delay]
  );

  return debouncedFunction;
}
