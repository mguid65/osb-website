import React, { Component } from 'react';
import { Paper } from "@material-ui/core";
import classNames from 'classnames';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import CardContent from '@material-ui/core/CardContent';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import DownloadIcon from "@material-ui/icons/CloudDownload";

const styles = theme => ({
  root: {
    maxWidth: "90vw",
    ...theme.mixins.gutters(),
    paddingTop: theme.spacing.unit * 2,
    paddingBottom: theme.spacing.unit * 2,
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
  },
  title: {
    flex: "0 0 auto",
  },
});

class Download extends Component {
  render() {
    const { classes } = this.props;

    return (
      <Paper className={ classes.root } elevation={1}>
        <Typography gutterBottom variant="h6" component="h2">
          Downloads
        </Typography>
        <Typography gutterBottom variant="subtitle1" component="h2">
          Ubuntu 18.04
        </Typography>
          <List>
           <a href="https://opensystembench.com/release/release-1.0.0-Ubuntu-1804.tar"> 
            <ListItem>        
             <ListItemIcon>
              <DownloadIcon />
             </ListItemIcon>
             <ListItemText>
               v1.0.0 .tar
             </ListItemText>
            </ListItem>
           </a>
           <a href="https://opensystembench.com/release/release-1.0.0-Ubuntu-1804.tar.gz">
            <ListItem>
             <ListItemIcon>
              <DownloadIcon />
             </ListItemIcon>
             <ListItemText>
               v1.0.0 .tar.gz
             </ListItemText>
            </ListItem>
           </a>
           <a href="https://opensystembench.com/release/release-1.0.0-Ubuntu-1804.zip">
            <ListItem>
             <ListItemIcon>
              <DownloadIcon />
             </ListItemIcon>
             <ListItemText>
               v1.0.0 .zip
             </ListItemText>
            </ListItem>
           </a>
          </List>
      </Paper>
    );
  }
}

Download.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Download);
