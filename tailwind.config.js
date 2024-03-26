/** @type {import('tailwindcss').Config} */
export default {
  content: ["./**/*.templ,", "./**/*.html", "./**/*.go"],
  theme: {
    extend: {
      backgroundImage: {
        "topography-pattern": "url('/img/topography.svg')",
      },
      colors: {
        "blue-black-gradient": "linear-gradient(to right, blue, black)",
      },
    },
  },
  plugins: [],
};
