import React, { Component } from "react";
import {
  Paper,
  Typography,
  List,
  ListItem,
  ListItemIcon,
  ListItemText
} from "@material-ui/core";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';

const styles = theme => ({
  root: {
    maxWidth: "90vw",
    ...theme.mixins.gutters(),
    paddingTop: theme.spacing.unit * 2,
    paddingBottom: theme.spacing.unit * 2
  },
  container: {
    display: "flex",
    flexWrap: "wrap"
  },
  margin: {
    margin: theme.spacing.unit
  },
  title: {
    flex: "0 0 auto"
  },
  info: {
    marginLeft: theme.spacing.unit,
  },
});

class About extends Component {
  render() {
    const { classes } = this.props;

    return (
      <Paper className={classes.root} elevation={1}>
	<Grid container justify='center'>
	  <img src='https://opensystembench.com/favicon.ico' />
	</Grid>
        <Grid container justify='center'>
          <Typography gutterBottom variant="body2" className={classes.info}>
            OpenSystemBench(OSB) is a multi-platform CPU benchmarking application.
          </Typography>
        <Grid container justify='center'>
	</Grid>
	  <Typography gutterBotton variant="body2" className={classes.info}>
            Matthew Guidry, Nam Vu, Mason Walton, Adam Yu, Raphaela Mettig, Benjamin Geiss 2018
	  </Typography>
	</Grid>
      </Paper>
    );
  }
}

About.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(About);
