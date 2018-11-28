import React from "react";
import classNames from "classnames";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";

const MenuBar = props => {
  const { classes, state, onClick } = props;
  return (
      <AppBar
        position="absolute"
        className={classNames(
          classes.appBar,
          state.open && classes.appBarShift
        )}
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
            variant="h6"
            color="inherit"
            noWrap
            className={classes.title}
          >
            OpenSystemBench
          </Typography>
        </Toolbar>
      </AppBar>
  );
};

export default MenuBar;
