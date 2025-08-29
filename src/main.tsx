import { render } from "preact";
import { Route, Router } from "preact-router";
import { App } from "./app.tsx";
import { Home } from "./pages/Home/Home.tsx";
import { Projects } from "./pages/Projects.tsx";
import "./styles/main.scss";

render(
  <App>
    <Router>
      <Route path="/" component={Home} />
      <Route path="/projects" component={Projects} />
    </Router>
  </App>,
  document.getElementById("app")!,
);
