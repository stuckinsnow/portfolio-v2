import type { ComponentChildren } from "preact";

interface AppProps {
  children: ComponentChildren;
}

export function App({ children }: AppProps) {
  return <>{children}</>;
}
