import { useApi } from "../../api/useApi";
import { PortfolioData } from "./types";

export function Home() {
  const { data: portfolio, loading, error } = useApi<PortfolioData>("/api/portfolio");

  switch (true) {
    case loading: return <div>Loading...</div>;
    case !!error: return <div>Error: {error}</div>;
    case !portfolio: return <div>No data available</div>;
  }

  return (
    <main>
      <h1>Welcome</h1>
      <h2>Skills</h2>
      <div id="skills">
        {portfolio.skills.map((skill, index) => (
          <span key={index} class="skill">
            {skill}
          </span>
        ))}
      </div>
      <h2>Projects</h2>
      <div id="projects">
        {portfolio.projects.map((project, index) => (
          <div key={index} class="project">
            <h3>{project.name}</h3>
            <p>{project.description}</p>
          </div>
        ))}
      </div>
    </main>
  );
}