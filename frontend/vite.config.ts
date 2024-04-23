import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), "");

  return {
    define: {
      "process.env": env,
    },
    plugins: [react(), tsconfigPaths()],
    server: {
      host: true,
      port: 3000,
      strictPort: true,
      watch: {
        usePolling: true,
      },
    },
  };
});
