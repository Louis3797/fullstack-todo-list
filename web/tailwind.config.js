const colors = require("tailwindcss/colors");

module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      transparent: "transparent",
      primary: {
        DEFAULT: "#DBE2EF",
        600: "#c5cbd7",
      },
      secondary: "#3F72AF",
      accent: "#112D4E",
      white: "#ffffff",
      red: colors.red,
    },
    extend: {},
  },
  plugins: [],
};
