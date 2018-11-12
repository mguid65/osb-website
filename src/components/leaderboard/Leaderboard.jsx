import React, { Component } from "react";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import {
  Paper,
  Table,
  TableBody,
  TableRow,
  TableCell,
  TablePagination,
  CircularProgress,
  TableFooter
} from "@material-ui/core";
import { ExpandLess, ExpandMore } from "@material-ui/icons";
import LeaderboardHeader from "./LeaderboardHeader";
import LeaderboardToolbar from "./LeaderboardToolbar";
import LeaderboardPagination from "./LeaderboardPagination";
import { sleep, getData, stableSort, getSorting } from "./data";

const ScoreMetaData = ({ classes, name, time, score }) => {
  return (
    <TableRow className={classes}>
      <TableCell colSpan={2} />
      <TableCell numeric>{name}</TableCell>
      <TableCell numeric>{time}</TableCell>
      <TableCell numeric>{score}</TableCell>
    </TableRow>
  );
};

const leaderboardStyles = theme => ({
  root: {
    maxWidth: "90vw",
    margin: "auto",
    marginTop: theme.spacing.unit * 3
  },
  table: {
    minWidth: 550,
    tableLayout: "fixed"
  },
  tableWrapper: {
    overflowX: "auto"
  },
  row: {
    "&:nth-of-type(odd)": {
      backgroundColor: theme.palette.background.default
    }
  }
});

class Leaderboard extends Component {
  state = {
    order: "desc",
    orderBy: "TotalScore",
    selected: [],
    data: [],
    loading: true,
    page: 0,
    rowsPerPage: 5
  };

  componentDidMount = async () => {
    const data = await getData();
    await sleep(1000);
    this.setState({ data, loading: false });
  };

  handleRequestSort = (event, property) => {
    const orderBy = property;
    let order = "desc";
    if (this.state.orderBy === property && this.state.order === "desc") {
      order = "asc";
    }
    this.setState({ order, orderBy });
  };

  handleClick = (event, id) => {
    const { selected } = this.state;
    const selectedIndex = selected.indexOf(id);
    let newSelected = [];

    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selected, id);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selected.slice(1));
    } else if (selectedIndex === selected.length - 1) {
      newSelected = newSelected.concat(selected.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selected.slice(0, selectedIndex),
        selected.slice(selectedIndex + 1)
      );
    }

    this.setState({ selected: newSelected });
  };

  handleChangePage = (event, page) => {
    this.setState({ page });
  };

  handleChangeRowsPerPage = event => {
    this.setState({ rowsPerPage: event.target.value });
  };

  handleRefresh = () => {
    this.setState({ loading: true });
    sleep(1000).then(() =>
      getData().then(scores => this.setState({ data: scores, loading: false }))
    );
  };

  isSelected = id => this.state.selected.indexOf(id) !== -1;

  render() {
    const { classes } = this.props;
    const { data, order, orderBy, rowsPerPage, page, loading } = this.state;
    const emptyRows =
      rowsPerPage - Math.min(rowsPerPage, data.length - page * rowsPerPage);

    return (
      <Paper className={classes.root}>
        <LeaderboardToolbar onRefresh={this.handleRefresh} />
        <div className={classes.tableWrapper}>
          <Table className={classes.table} aria-labelledby="tableTitle">
            <LeaderboardHeader
              order={order}
              orderBy={orderBy}
              onRequestSort={this.handleRequestSort}
            />
            <TableBody>
              {loading ? (
                <TableRow style={{ height: 48 * rowsPerPage }}>
                  <TableCell colSpan={5} style={{ textAlign: "center" }}>
                    <CircularProgress size={50} color="primary" />
                  </TableCell>
                </TableRow>
              ) : (
                <React.Fragment>
                  {stableSort(data, getSorting(order, orderBy))
                    .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                    .map(score => {
                      const isSelected = this.isSelected(score.ID);
                      return (
                        <React.Fragment key={score.ID}>
                          <TableRow
                            hover
                            onClick={event => this.handleClick(event, score.ID)}
                            role="checkbox"
                            aria-checked={isSelected}
                            tabIndex={-1}
                            selected={isSelected}
                            className={classes.row}
                          >
                            <TableCell padding="checkbox">
                              {isSelected ? <ExpandLess /> : <ExpandMore />}
                            </TableCell>
                            <TableCell
                              component="th"
                              scope="row"
                              padding="none"
                            >
                              {score.ResultAlias}
                            </TableCell>
                            <TableCell numeric>{score.ID}</TableCell>
                            <TableCell numeric>{score.TotalTime}</TableCell>
                            <TableCell numeric>{score.TotalScore}</TableCell>
                          </TableRow>
                          {isSelected && (
                            <React.Fragment>
                              <ScoreMetaData
                                classes={classes.row}
                                name="N-Body"
                                time={score.NBodyTime}
                                score={score.NBodyScore}
                              />
                              <ScoreMetaData
                                classes={classes.row}
                                name="PI Digits"
                                time={score.PiDigitsTime}
                                score={score.PiDigitsScore}
                              />
                              <ScoreMetaData
                                classes={classes.row}
                                name="Mandelbrot"
                                time={score.MandelbrotTime}
                                score={score.MandelbrotScore}
                              />
                              <ScoreMetaData
                                classes={classes.row}
                                name="Spectral Norm"
                                time={score.SpectralNormTime}
                                score={score.SpectralNormScore}
                              />
                              <ScoreMetaData
                                classes={classes.row}
                                name="Binary Trees"
                                time={score.BinaryTreesTime}
                                score={score.BinaryTreesScore}
                              />
                            </React.Fragment>
                          )}
                        </React.Fragment>
                      );
                    })}
                  {emptyRows > 0 && (
                    <TableRow style={{ height: 48 * emptyRows }}>
                      <TableCell colSpan={5} />
                    </TableRow>
                  )}
                </React.Fragment>
              )}
            </TableBody>
            <TableFooter>
              <TableRow>
                <TablePagination
                  colSpan={5}
                  count={data.length}
                  rowsPerPage={rowsPerPage}
                  page={page}
                  onChangePage={this.handleChangePage}
                  onChangeRowsPerPage={this.handleChangeRowsPerPage}
                  ActionsComponent={LeaderboardPagination}
                />
              </TableRow>
            </TableFooter>
          </Table>
        </div>
      </Paper>
    );
  }
}

Leaderboard.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(leaderboardStyles)(Leaderboard);