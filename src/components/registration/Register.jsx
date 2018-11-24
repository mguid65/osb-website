import React, { Component } from "react";
import { Paper } from "@material-ui/core";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import InputAdornment from "@material-ui/core/InputAdornment";
import TextField from "@material-ui/core/TextField";
import Visibility from "@material-ui/icons/Visibility";
import VisibilityOff from "@material-ui/icons/VisibilityOff";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";

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
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200
  },
  title: {
    flex: "0 0 auto"
  }
});

class Register extends Component {
  state = {
    showPassword: false
  };

  handleClickShowPassword = () => {
    this.setState(state => ({ showPassword: !state.showPassword }));
  };

  render() {
    const { classes } = this.props;

    return (
      <Paper className={classes.root} elevation={1}>
        <Typography gutterBottom variant="h6" component="h2">
          Registration
        </Typography>
        <form
          className={classes.container}
          action="/api/users/register"
          method="post"
        >
          <TextField
            name="email"
            id="filled-email-input"
            type="email"
            variant="filled"
            label="Email"
            className={classes.textField}
            margin="normal"
          />
          <TextField
            name="username"
            id="filled-name"
            variant="filled"
            label="Username"
            className={classes.textField}
            margin="normal"
          />
          <TextField
            name="password"
            id="filled-adornment-password"
            type={this.state.showPassword ? "text" : "password"}
            variant="filled"
            className={classes.textField}
            margin="normal"
            label="Password"
            InputProps={{
              endAdornment: (
                <InputAdornment variant="filled" position="end">
                  <IconButton
                    aria-label="Toggle password visibility"
                    onClick={this.handleClickShowPassword}
                  >
                    {this.state.showPassword ? (
                      <VisibilityOff />
                    ) : (
                      <Visibility />
                    )}
                  </IconButton>
                </InputAdornment>
              )
            }}
          />
          <Button size="small" type="submit" className={classes.margin}>
            Register
          </Button>
        </form>
        <Typography component="p">
          Register an account with OpenSystemBench for score submission from our
          desktop clients.
        </Typography>
      </Paper>
    );
  }
}

Register.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(Register);
