const colors = require("tailwindcss/colors");
const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  mode: "jit",
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: "class", // or 'media' or 'class'
  theme: {
    extend: {
      height: {
        footer: "var(--footer-height)",
      },
      lineHeight: {
        nav: "var(--nav-height)",
      },
      maxWidth: {
        "4.5xl": "60rem",
        "8xl": "88rem",
      },
      spacing: {
        nav: "var(--nav-height)",
        4.5: "1.125rem",
        5.5: "1.375rem",
      },
      transitionDelay: {
        25: "25ms",
      },
      keyframes: {
        slide: {
          "0%": { "margin-left": "-25%" },
          "100%": { "margin-left": "125%" },
        },
      },
      animation: {
        slide: "slide 1.7s cubic-bezier(0, 0, 0.26, 0.74) infinite",
      },
    },
    screens: {
      xs: "430px",
      "xs-max": { max: "430px" },
      ...defaultTheme.screens,
      "3xl": "1904px",
    },
    fontFamily: {
      ...defaultTheme.fontFamily,
    },
    colors: {
      transparent: "transparent",
      current: "currentColor",
      primary: {
        lighter: "#4f051e",
        DEFAULT: "#26020e", // Base
        darker: "#160108",
      },
      secondary: {
        50: "#f9f3f5",
        100: "#e3cad2",
        200: "#c99baa",
        300: "#ad667c",
        400: "#9a425e",
        500: "#8e2b4a",
        600: "#811336", // Base
        700: "#640f2a",
        800: "#470a1e",
        900: "#2a0612",
      },
      footer: "#371b28",
      black: { lighter: "#050505", DEFAULT: colors.black },
      blue: colors.blue,
      gray: colors.neutral,
      green: colors.green,
      orange: colors.orange,
      red: colors.red,
      rose: colors.rose,
      "warm-gray": colors.stone,
      white: colors.white,
      yellow: colors.yellow,
    },
  },
  variants: {
    extend: {
      display: ["group-hover"],
      opacity: ["disabled"],
    },
  },
  plugins: [],
  corePlugins: {
    container: false, // use custom container
  },
};
