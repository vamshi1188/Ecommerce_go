import { red, blue, green, yellow, purple, orange, pink, cyan, lime } from '@mui/material/colors';
import { createTheme } from '@mui/material/styles';

// A custom theme for this app
const theme = createTheme({
  palette: {
    primary: {
      main: '#556cd6', // Primary color
      light: '#778beb',
      dark: '#335bb6',
    },
    secondary: {
      main: '#19857b', // Secondary color
      light: '#45a29e',
      dark: '#116d62',
    },
    error: {
      main: red.A400, // Error color
      light: red.A200,
      dark: red.A700,
    },
    success: {
      main: green[500], // Success color
      light: green[300],
      dark: green[700],
    },
    warning: {
      main: yellow[800], // Warning color
      light: yellow[600],
      dark: yellow[900],
    },
    info: {
      main: blue[500], // Info color
      light: blue[300],
      dark: blue[700],
    },
    pink: {
      main: pink[500], // Pink color
      light: pink[300],
      dark: pink[700],
    },
    orange: {
      main: orange[500], // Orange color
      light: orange[300],
      dark: orange[700],
    },
    cyan: {
      main: cyan[500], // Cyan color
      light: cyan[300],
      dark: cyan[700],
    },
    lime: {
      main: lime[500], // Lime color
      light: lime[300],
      dark: lime[700],
    },
    contrastThreshold: 3, // Contrast threshold for text
    tonalOffset: 0.2, // Shift the color to darker or lighter
  },
  typography: {
    fontFamily: [
      'Roboto',
      'Helvetica Neue',
      'Arial',
      'sans-serif',
    ].join(','),
  },
});

export default theme;
