import { useApi } from "../../api/useApi";
import type { PortfolioData } from "./types";

export function Home() {
  const {
    data: portfolio,
    loading,
    error,
  } = useApi<PortfolioData>("/api/portfolio");

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;
  if (!portfolio) return <div>No data available</div>;

  return (
    <main>
      <h1>{portfolio.title}</h1>
      <p>{portfolio.description}</p>
    </main>
  );
}
