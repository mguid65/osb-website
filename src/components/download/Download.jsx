import React, { Component } from 'react';
import { Paper } from "@material-ui/core";
import classNames from 'classnames';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

const styles = theme => ({
  root: {
    maxWidth: "90vw",
    margin: "auto",
    marginTop: theme.spacing.unit * 3
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  margin: {
    margin: theme.spacing.unit,
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200,
  }
});

class Download extends Component {
  render() {
    const { classes } = this.props;

    return (
      <Paper className={ classes.root } elevation={1}>

      </Paper>
    );
  }
}

Download.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Download);







