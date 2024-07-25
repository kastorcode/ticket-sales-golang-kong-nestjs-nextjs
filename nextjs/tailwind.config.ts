import type { Config } from "tailwindcss"

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        primary: "rgb(0,0,0)",
        secondary: "rgba(255,255,255,0.1)",
        bar: "rgb(0,0,0)",
        "btn-primary": "rgb(255,215,0)",
        "btn-primary-hover": "rgb(43, 36, 16)",
        available: "rgb(0,255,0)",
        occupied: "rgb(255,0,0)",
        chosen: "rgb(255,255,0)"
      },
      textColor: {
        default: "rgb(192,192,192)",
        "btn-primary": "rgb(43, 36, 16)",
        "btn-primary-hover": "rgb(255,215,0)",
        subtitle: "rgb(205,127,50)",
      },
      gridTemplateColumns: {
        "auto-fit-cards": "repeat(auto-fit, minmax(277px, 1fr))",
      },
    },
  },
  plugins: [],
}
export default config