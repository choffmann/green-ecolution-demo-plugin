import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { federation } from "@module-federation/vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    federation({
      name: "remote",
      filename: "plugin.js",
      exposes: {
        "./demo_plugin": "./src/App.tsx",
      },
      shared: {
        // ...pkg.dependencies,
        react: { singleton: true },
        "react-dom": { singleton: true },
        "@green-ecolution/plugin-interface": { singleton: true },
      },
    }),
  ],
  base: "",
  build: {
    target: "esnext",
  },
});
