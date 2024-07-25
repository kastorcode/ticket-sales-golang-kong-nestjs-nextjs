/** @type {import('next').NextConfig} */
const nextConfig = {
  basePath: "/nextjs",
  experimental: {
    serverActions: {
      allowedOrigins: ["localhost:8000", "host.docker.internal:8000"]
    }
  },
  images: {
    remotePatterns: [
      { hostname: "*.epicgames.com" },
      { hostname: "*.googleusercontent.com" },
      { hostname: "*.hypb.st" },
      { hostname: "meups.com.br" }
    ]
  }
}

export default nextConfig