import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  async rewrites() {
    return [
      {
        source: '/api/',
        destination: 'http://localhost:8080/',
      },
    ];
  },
};

export default nextConfig;
