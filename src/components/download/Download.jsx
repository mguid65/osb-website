import React, { Component } from "react";
import {
  Paper,
  Typography,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from "@material-ui/core";
import DownloadIcon from "@material-ui/icons/CloudDownload";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";

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
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200
  },
  title: {
    flex: "0 0 auto"
  }
});

class Download extends Component {
  render() {
    const { classes } = this.props;

    return (
      <Paper className={classes.root} elevation={1}>
        <Typography gutterBottom variant="h6" component="h2">
          Downloads
        </Typography>
        <Typography gutterBottom variant="subtitle1" component="h2">
          Ubuntu 18.04
        </Typography>
        <List>
          <a href="https://opensystembench.com/downloads/release-1.0.0-Ubuntu-1804.tar" download>
            <ListItem>
              <ListItemIcon>
                <DownloadIcon />
              </ListItemIcon>
              <ListItemText>v1.0.0 .tar</ListItemText>
            </ListItem>
          </a>
          <a href="https://opensystembench.com/downloads/release-1.0.0-Ubuntu-1804.tar.gz" download>
            <ListItem>
              <ListItemIcon>
                <DownloadIcon />
              </ListItemIcon>
              <ListItemText>v1.0.0 .tar.gz</ListItemText>
            </ListItem>
          </a>
          <a href="https://opensystembench.com/downloads/release-1.0.0-Ubuntu-1804.zip" download>
            <ListItem>
              <ListItemIcon>
                <DownloadIcon />
              </ListItemIcon>
              <ListItemText>v1.0.0 .zip</ListItemText>
            </ListItem>
          </a>
        </List>
        <Typography gutterBottom variant="subtitle1" component="h2">
          Windows 32-bit
        </Typography>
        <List>
          <a href="https://opensystembench.com/downloads/OpenSystemBenchWIN32.zip" download>
            <ListItem>
              <ListItemIcon>
                <DownloadIcon />
              </ListItemIcon>
              <ListItemText>v1.0.0 .zip</ListItemText>
            </ListItem>
          </a>
        </List>
      </Paper>
    );
  }
}

Download.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(Download);
