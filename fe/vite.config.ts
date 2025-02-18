import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import path from "path";

// https://vite.dev/config/
export default defineConfig({
   plugins: [react()],
   server: {
      proxy: {
         "lh3.googleusercontent.com": {
            target: "https://lh3.googleusercontent.com",
            changeOrigin: true,
         },
      },
   },
   resolve: {
      alias: {
         "@": path.resolve(__dirname, "./src"),
      },
   },
});
