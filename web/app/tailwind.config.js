const colors = require('tailwindcss/colors')
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class', // or 'media' or 'class'
  theme: {
    extend: {
      maxWidth: {
        '4.5xl': '60rem',
        '8xl': '88rem',
      },
      spacing: {
        '4.5': '1.125rem',
      },
    },
    screens: {
      'xs': '425px',
      'xs-max': { 'max': '425px' },
      ...defaultTheme.screens,
      '3xl': '1904px',
    },
    fontFamily: {
      ...defaultTheme.fontFamily,
      'awesome': '"Font Awesome 5 Free"',
      'awesome-brands': '"Font Awesome 5 Brands"'
    },
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
      secondary: {
        50: '#f9f3f5',
        100: '#e3cad2',
        200: '#c99baa',
        300: '#ad667c',
        400: '#9a425e',
        500: '#8e2b4a',
        600: '#811336', // Base
        700: '#640f2a',
        800: '#470a1e',
        900: '#2a0612',
      },
      footer: '#321f28',
      black: { lighter: "#050505", DEFAULT: colors.black },
      gray: colors.trueGray,
      green: colors.green,
      orange: colors.orange,
      red: colors.red,
      rose: colors.rose,
      'warm-gray': colors.warmGray,
      white: colors.white,
      yellow: colors.yellow,
    }
  },
  variants: {
    extend: {
      opacity: ['disabled'],
    },
  },
  plugins: [],
  corePlugins: {
    container: false, // use custom container
  }
}
