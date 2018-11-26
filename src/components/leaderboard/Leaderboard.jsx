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
import { getData, stableSort, getSorting } from "./data";

const Metadata = ({ classes, specs }) => {
  return (
    <React.Fragment>
      <TableRow classNames={classes}>
        <TableCell colSpan={3} numeric>
          Vendor: {specs.vendor}
        </TableCell>
        <TableCell numeric>Model: {specs.model}</TableCell>
        <TableCell numeric>Clock Speed: {specs.speed}</TableCell>
      </TableRow>
      <TableRow>
        <TableCell colSpan={3} numeric>
          Threads: {specs.threads}
        </TableCell>
        <TableCell numeric>Overclocked: {String(specs.overclocked)}</TableCell>
        <TableCell numeric>Byte Order: {specs.speed}</TableCell>
      </TableRow>
      <TableRow>
        <TableCell colSpan={3} numeric>
          Physical Memory: {specs.physical}
        </TableCell>
        <TableCell numeric>Virtual Memory: {specs.virtual}</TableCell>
        <TableCell numeric>Swap Memory: {specs.swap}</TableCell>
      </TableRow>
    </React.Fragment>
  );
};

const leaderboardStyles = theme => ({
  root: {
    maxWidth: "90vw",
    margin: "auto"
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
    orderBy: "score",
    selected: [],
    data: [],
    loading: true,
    page: 0,
    rowsPerPage: 10
  };

  componentDidMount = async () => {
    const ranked = await getData();
    this.setState({ data: ranked, loading: false });
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

  handleRefresh = async () => {
    this.setState({ loading: true });
    const ranked = await getData();
    this.setState({ data: ranked, loading: false });
  };

  isSelected = id => this.state.selected.indexOf(id) !== -1;

  render() {
    const { classes } = this.props;
    const { data, order, orderBy, rowsPerPage, page, loading } = this.state;
    const emptyRows =
      rowsPerPage - Math.min(rowsPerPage, data.length - page * rowsPerPage);

    return (
      <Paper className={classes.root} elevation={1}>
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
                    .map(data => {
                      const isSelected = this.isSelected(data.id);

                      return (
                        <React.Fragment key={data.id}>
                          <TableRow
                            hover
                            onClick={event => this.handleClick(event, data.id)}
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
                              {data.user}
                            </TableCell>
                            <TableCell numeric>{data.rank}</TableCell>
                            <TableCell numeric>{data.totalTime}</TableCell>
                            <TableCell numeric>{data.totalScore}</TableCell>
                          </TableRow>
                          {isSelected &&
                            data.scores != null &&
                            data.scores.map(score => {
                              return (
                                <TableRow className={classes}>
                                  <TableCell colSpan={3} numeric>
                                    {score.name}
                                  </TableCell>
                                  <TableCell numeric>{score.time}</TableCell>
                                  <TableCell numeric>{score.score}</TableCell>
                                </TableRow>
                              );
                            })}
                          {isSelected &&
                            data.specs != null &&
                            data.specs.specs != null && (
                              <Metadata specs={data.specs.specs} />
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
