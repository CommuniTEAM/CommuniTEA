import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 3000,
    // Look into whether watch option (hot reload) should be removed for production
    watch: {
      usePolling: true,
    },
  },
});
