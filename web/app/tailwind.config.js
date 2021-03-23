const colors = require('tailwindcss/colors')

module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class', // or 'media' or 'class'
  theme: {
    extend: {},
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      primary: {
        50: '#efeded',
        100: '#eae6e7',
        200: '#6a5354',
        300: '#4a2e2f',
        400: '#3a1b1c',
        500: '#351516',
        600: '#2a090a', // Base
        700: '#260809',
        800: '#1c0607',
        900: '#130405',
      },
      footer: '#321f28',
      gray: colors.trueGray,
      white: colors.white,
      black: colors.black,
    }
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
