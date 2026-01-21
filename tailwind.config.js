/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.templ"],
  theme: {
    extend: {
      colors: {
        vapula: {
          neon: '#10b981', // Verde esmeralda brillante
          dark: '#052e16',  // Verde muy oscuro para fondos
          copper: '#d97706', // Color cobre/ámbar metálico
        },
        darkbg: {
          950: '#0a0a0a', // Negro casi puro
          900: '#121212', // Gris muy oscuro para tarjetas
        }
      },
      fontFamily: {
        sans: ['Inter', 'sans-serif'],
      },
    },
  },
  plugins: [],
};
