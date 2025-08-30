import { useApi } from "../../api/useApi";
import type { PortfolioData, TodoData } from "./types";

export function Home() {
  const {
    data: portfolio,
    loading,
    error,
  } = useApi<PortfolioData>("/api/portfolio");

  const {
    data: todoData,
    loading: todosLoading,
    error: todosError,
  } = useApi<TodoData>("/api/todos");

  if (loading || todosLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;
  if (todosError) return <div>Error loading todos: {todosError}</div>;
  if (!portfolio) return <div>No data available</div>;

  return (
    <main>
      <h1>{portfolio.title}</h1>
      <p>{portfolio.description}</p>

      <h2>Todo Items</h2>
      <table>
        <thead>
          <tr>
            <th>Item</th>
          </tr>
        </thead>
        <tbody>
          {todoData?.todos?.map((todo, index) => (
            <tr key={index}>
              <td>{todo}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </main>
  );
}
