import React from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";
import Users from "./components/Users";
import PatientsCreate from "./components/PatientsCreate";
import VerticalTabs from "./components/Tabs";
import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';

import { createTheme, ThemeProvider } from '@mui/material/styles';
import { green } from '@mui/material/colors';

import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';

import Icon from '@mui/material/Icon';
import AccountBoxIcon from '@mui/icons-material/AccountBox';

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
  

export default function App() {
return (
    <Router>
        <div>
        <Navbar /> 
        <Box sx={{ flexGrow: 1 }}>
            <Grid container spacing={2}>
              <Grid item xs={3}>
                  <Grid item xs={12}>
                  
                  </Grid>
                  <Grid item xs={12}>
                  <VerticalTabs />
                  </Grid>
              </Grid>
              <Grid item xs={9}>
                <Routes>
                  <Route path="/" element={<PatientsCreate />} />
                  <Route path="/patient" element={<Users />} />
           
                  {/* <Route path="/create" element={<PatientsCreate />} /> */}
 
                </Routes>
              </Grid> 
            </Grid> 
          </Box>
      
                
        </div>
   
    </Router>
  );
}
