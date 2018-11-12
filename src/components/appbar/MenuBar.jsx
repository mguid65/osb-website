import React from "react";
import classNames from "classnames";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import Badge from "@material-ui/core/Badge";
import MenuIcon from "@material-ui/icons/Menu";
import NotificationsIcon from "@material-ui/icons/Notifications";

const MenuBar = props => {
  const { classes, state, onClick } = props;
  return (
    <AppBar
      position="absolute"
      className={classNames(classes.appBar, state.open && classes.appBarShift)}
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
          Open System Bench
        </Typography>
        <IconButton color="inherit">
          <Badge badgeContent={4} color="secondary">
            <NotificationsIcon />
          </Badge>
        </IconButton>
      </Toolbar>
    </AppBar>
  );
};

export default MenuBar;
