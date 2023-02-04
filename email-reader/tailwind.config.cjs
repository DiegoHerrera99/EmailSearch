/** @type {import('tailwindcss').Config} */

const colors = require('tailwindcss/colors')

module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    colors: {
      white: colors.white,
      red: colors.red,
      primario: '#233044',
      secundario: '#1b2635',
      gray: '#9197a1',
      mostaza: '#E68A00',
      blueButtons: '#146F8E'
    },
  },
  plugins: [],
}
