import React, { Component } from 'react';
import { Paper } from "@material-ui/core";
import classNames from 'classnames';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import IconButton from '@material-ui/core/IconButton';
import InputAdornment from '@material-ui/core/InputAdornment';
import TextField from '@material-ui/core/TextField';
import MenuItem from '@material-ui/core/MenuItem';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';
import Button from '@material-ui/core/Button';

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

class Register extends Component {
  constructor() {
    super();
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleSubmit(event){
    event.preventDefault();
    const data = new FormData(event.target);
    fetch('https://opensystembench.com/api/users/register', {
      method: 'POST',
      body: data,
    });
  }

  render() {
    const { classes } = this.props;

    return (
      <Paper className={ classes.root }>
      <form className={ classes.container } onSubmit={this.handleSubmit}>
        <TextField
	  id="filled-email-input"
	  label="Email"
	  className={ classes.textField }
          margin="normal"
        />
        <TextField
          id="filled-name"
          label="Username"
          className={ classes.textField }
          margin="normal"
        />
        <TextField
          id="filled-password-input"
          className= { classes.textField }
	  margin="normal"
          label="Password"
        />
        <Button size="small" type="submit" className={classes.margin}>Register</Button>
      </form>
      </Paper>
    );
  }
}

Register.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Register);
