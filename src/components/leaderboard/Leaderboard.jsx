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

// const ScoreMetaData = ({ classes, name, time, score }) => {
//   return (
//     <TableRow className={classes}>
//       <TableCell colSpan={2} />
//       <TableCell numeric>{name}</TableCell>
//       <TableCell numeric>{time}</TableCell>
//       <TableCell numeric>{score}</TableCell>
//     </TableRow>
//   );
// };

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
    rowsPerPage: 5
  };

  componentDidMount = async () => {
    const resultsRes = await fetch("http://localhost:8080/api/results");
    const results = await resultsRes.json();
    const usersRes = await fetch("http://localhost:8080/api/users");
    const users = await usersRes.json();
    const specsRes = await fetch("http://localhost:8080/api/specs");
    const specs = await specsRes.json();

    const d = results.map(result => {
      const total = result.scores.find(score => score.name === "total");
      return {
        id: result.ID,
        totalTime: total.time,
        totalScore: total.score,
        user: users.find(user => user.ID === result.UserID).Name,
        scores: result.scores,
        specs: specs.find(spec => spec.ResultID === result.ID)
      };
    });

    const sorted = d.sort((a, b) => {
      if (a.totalScore < b.totalScore) return 1;
      else if (b.totalScore < a.totalScore) return 1;
      else return 0;
    });

    const ranked = sorted.map((result, index) => {
      result.rank = index + 1;
      return result;
    });

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
    this.setState({ loading: true});
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
                      const isSelected = this.isSelected(data.ID);

                      return (
                        <React.Fragment key={data.ID}>
                          <TableRow
                            hover
                            onClick={event => this.handleClick(event, data.ID)}
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
