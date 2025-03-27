import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  env: {
    CLERK_FRONTEND_API: process.env.CLERK_FRONTEND_API,
    CLERK_SECRET_KEY: process.env.CLERK_SECRET_KEY
  }
};

export default nextConfig;
