/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['static/**/*{html,tmpl,gohtml}', 'node_modules/preline/dist/*.js'],
  theme: {
    extend: {},
  },
  plugins: [
    require('preline/plugin'),
  ],
}

