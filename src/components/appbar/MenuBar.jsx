import React from "react";
import classNames from "classnames";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import { deepOrange500 } from "@material-ui/core/colors";
import MuiThemeProvider from "@material-ui/core/styles/MuiThemeProvider";
import { createMuiTheme } from "@material-ui/core/styles";

const muiTheme = createMuiTheme({
  palette: {
    primary1color: deepOrange500
  }
});

const MenuBar = props => {
  const { classes, state, onClick } = props;
  return (
    <MuiThemeProvider theme={muiTheme}>
      <AppBar
        position="absolute"
        className={classNames(
          classes.appBar,
          state.open && classes.appBarShift
        )}
        color="primary"
      >
        <Toolbar disableGutters={!state.open} className={classes.toolbar}>
          <IconButton
            color="inherit"
            aria-label="Open drawer"
            onClick={onClick}
            className={classNames(
              classes.menuButton,
              state.open && classes.menuButtonHidden
            )}
          >
            <MenuIcon />
          </IconButton>
          <Typography
            variant="title"
            color="inherit"
            noWrap
            className={classes.title}
          >
            OpenSystemBench
          </Typography>
        </Toolbar>
      </AppBar>
    </MuiThemeProvider>
  );
};

export default MenuBar;
