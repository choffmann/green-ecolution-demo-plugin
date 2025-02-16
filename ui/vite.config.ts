import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { federation } from "@module-federation/vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    federation({
      name: "demo_plugin",
      filename: "plugin.js",
      exposes: {
        "./app": "./src/App.tsx",
      },
      shared: ["react", "react-dom", "@green-ecolution/plugin-interface"],
    }),
  ],
  base: "",
  build: {
    target: "esnext",
  },
});
