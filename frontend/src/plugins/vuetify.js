import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

import { createVuetify } from 'vuetify'

const podcatsDarkTheme = {
  dark: true,
  colors: {
    'background': '#1F1F23',
    'surface': '#3F3E41',
    'primary': '#F8EF12',
    'primary-darken-1': '#3700B3',
    'secondary': '#FFFFFF',
    'secondary-darken-1': '#A5AC8C',
    'error': '#ff1744',
    'info': '#4dd0e1',
    'success': '#76ff03',
    'warning': '#ffd600',
  },
}

export default createVuetify({
  theme: {
    defaultTheme: 'podcatsDarkTheme',
    variations: {
      colors: ['surface'],
      lighten: 1,
    },
    themes: {
      podcatsDarkTheme,
    },
  },
})
