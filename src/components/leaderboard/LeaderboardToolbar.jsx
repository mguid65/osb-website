import React from "react";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import { lighten } from "@material-ui/core/styles/colorManipulator";
import { Toolbar, Typography, Tooltip, IconButton } from "@material-ui/core";
import RefreshIcon from "@material-ui/icons/Refresh";

const toolbarStyles = theme => ({
  root: {
    paddingRight: theme.spacing.unit
  },
  highlight:
    theme.palette.type === "light"
      ? {
          color: theme.palette.secondary.main,
          backgroundColor: lighten(theme.palette.secondary.light, 0.85)
        }
      : {
          color: theme.palette.text.primary,
          backgroundColor: theme.palette.secondary.dark
        },
  spacer: {
    flex: "1 1 100%"
  },
  actions: {
    color: theme.palette.text.secondary
  },
  title: {
    flex: "0 0 auto"
  }
});

const LeaderboardToolbar = ({ classes, onRefresh }) => {
  return (
    <Toolbar>
      <div className={classes.title}>
        <Typography variant="title" id="tableTitle">
          Leaderboard
        </Typography>
      </div>
      <div className={classes.spacer} />
      <div className={classes.actions}>
        <Tooltip title="Refresh list">
          <IconButton aria-label="Refresh list" onClick={onRefresh}>
            <RefreshIcon />
          </IconButton>
        </Tooltip>
      </div>
    </Toolbar>
  );
};

LeaderboardToolbar.propTypes = {
  classes: PropTypes.object.isRequired,
  onRefresh: PropTypes.func.isRequired
};

export default withStyles(toolbarStyles)(LeaderboardToolbar);
