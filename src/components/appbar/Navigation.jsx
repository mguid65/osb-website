import React, { Component } from "react";
import PropTypes from "prop-types";
import classNames from "classnames";
import { withStyles } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import Drawer from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import { secondaryListItems } from "./listItems";
import Leaderboard from "../leaderboard/Leaderboard";
import Register from "../registration/Register";
import Download from "../download/Download";
import MenuBar from "./MenuBar";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import InsertChartIcon from "@material-ui/icons/InsertChart";
import DownloadIcon from "@material-ui/icons/CloudDownload";
import PersonAdd from "@material-ui/icons/PersonAdd";

import { Route, NavLink, HashRouter } from "react-router-dom";

const drawerWidth = 240;

const styles = theme => ({
  root: {
    display: "flex"
  },
  toolbar: {
    paddingRight: 24
  },
  toolbarIcon: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-end",
    padding: "0 8px",
    ...theme.mixins.toolbar
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen
    })
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen
    })
  },
  menuButton: {
    marginLeft: 12,
    marginRight: 36
  },
  menuButtonHidden: {
    display: "none"
  },
  title: {
    flexGrow: 1
  },
  drawerPaper: {
    position: "relative",
    whiteSpace: "nowrap",
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen
    })
  },
  drawerPaperClose: {
    overflowX: "hidden",
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen
    }),
    width: theme.spacing.unit * 7,
    [theme.breakpoints.up("sm")]: {
      width: theme.spacing.unit * 9
    }
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    padding: theme.spacing.unit * 3,
    height: "100vh",
    overflow: "auto"
  },
  chartContainer: {
    marginLeft: -22
  },
  tableContainer: {
    height: 320
  },
  menuText: {
    textDecoration: "none"
  }
});

class Dashboard extends Component {
  state = {
    open: false
  };

  handleDrawerOpen = () => {
    this.setState({ open: true });
  };

  handleDrawerClose = () => {
    this.setState({ open: false });
  };

  render() {
    const { classes } = this.props;

    return (
      <HashRouter>
        <React.Fragment>
          <CssBaseline />
          <div className={classes.root}>
            <MenuBar
              classes={classes}
              state={this.state}
              onClick={this.handleDrawerOpen}
            />
            <Drawer
              variant="permanent"
              classes={{
                paper: classNames(
                  classes.drawerPaper,
                  !this.state.open && classes.drawerPaperClose
                )
              }}
              open={this.state.open}
            >
              <div className={classes.toolbarIcon}>
                <IconButton onClick={this.handleDrawerClose}>
                  <ChevronLeftIcon />
                </IconButton>
              </div>
              <Divider />
              <List>
                <React.Fragment>
                  <NavLink to="/">
                    <ListItem button>
                      <ListItemIcon>
                        <InsertChartIcon />
                      </ListItemIcon>
                      <ListItemText
                        className={classes.menuText}
                        primary="Leaderboard"
                      />
                    </ListItem>
                  </NavLink>
                  <NavLink to="/downloads">
                    <ListItem button>
                      <ListItemIcon>
                        <DownloadIcon />
                      </ListItemIcon>
                      <ListItemText
                        className={classes.menuText}
                        primary="Download"
                      />
                    </ListItem>
                  </NavLink>
                  <NavLink to="/register">
                    <ListItem button>
                      <ListItemIcon>
                        <PersonAdd />
                      </ListItemIcon>
                      <ListItemText
                        className={classes.menuText}
                        primary="Register"
                      />
                    </ListItem>
                  </NavLink>
                </React.Fragment>
              </List>
              <Divider />
              <List>{secondaryListItems}</List>
            </Drawer>
            <main className={classes.content}>
              <div className={classes.appBarSpacer} />
              <div className={classes.tableContainer}>
                <Route exact path="/" component={Leaderboard} />
                <Route path="/register" component={Register} />
                <Route path="/downloads" component={Download} />
              </div>
            </main>
          </div>
        </React.Fragment>
      </HashRouter>
    );
  }
}

Dashboard.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(Dashboard);
