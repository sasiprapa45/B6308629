import * as React from "react";

import AppBar from "@mui/material/AppBar";

import Box from "@mui/material/Box";

import Toolbar from "@mui/material/Toolbar";

import Typography from "@mui/material/Typography";

import IconButton from "@mui/material/IconButton";

import MenuIcon from "@mui/icons-material/Menu";

import { createTheme, ThemeProvider } from '@mui/material/styles';
import { green } from '@mui/material/colors';


const theme = createTheme({

  palette: {
    primary: {
      main: green[400],
    },
    secondary: {
      main: '#e8f5e9',
    },
  },
});


function Navbar() {

 return (
<ThemeProvider theme={theme}>
   <Box sx={{ flexGrow: 1 }}>

     <AppBar position="static">

       <Toolbar>

         <IconButton

           size="large"

           edge="start"

           color="secondary"

           aria-label="menu"

           sx={{ mr: 5 }}

         >

           <MenuIcon />

         </IconButton>

         <Typography variant="h6" component="div"  color="secondary" sx={{ flexGrow: 1 }}>

           ระบบผู้ป่วยของแพทย์

         </Typography>

       </Toolbar>

     </AppBar>

    </Box>
   </ThemeProvider>

 );

}

export default Navbar;