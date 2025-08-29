import { useState, useEffect } from "preact/hooks";

interface UseApiReturn<T> {
  data: T | null;
  loading: boolean;
  error: string | null;
  refetch: () => void;
}

export function useApi<T = any>(
  endpoint: string,
  options?: RequestInit
): UseApiReturn<T> {
  const [state, setState] = useState<UseApiReturn<T>>({
    data: null,
    loading: true,
    error: null,
    refetch: () => {},
  });

  const fetchData = async () => {
    setState(prev => ({ ...prev, loading: true, error: null }));
    
    try {
      const response = await fetch(endpoint, options);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setState(prev => ({ ...prev, data, loading: false }));
    } catch (err) {
      setState(prev => ({
        ...prev,
        error: err instanceof Error ? err.message : 'An error occurred',
        loading: false,
      }));
    }
  };

  useEffect(() => {
    fetchData();
  }, [endpoint]);

  return {
    ...state,
    refetch: fetchData,
  };
}