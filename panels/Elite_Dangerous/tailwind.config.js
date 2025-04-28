/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html", // Ensure this path is correct
  ],
  theme: {
    extend: {},
  },
  safelist: [
    {
      pattern: /^(border|bg)-(blue|red|sky|orange|lime)-(200|400|300\/40)$/,
    },
  ],
  plugins: [],
  preflight: false,
  mode: "jit",
};
