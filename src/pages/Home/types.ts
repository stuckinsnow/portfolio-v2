export interface Project {
  name: string;
  description: string;
}

export interface PortfolioData {
  projects: Project[];
  skills: string[];
}