import React, { Component } from "react";
import PropTypes from "prop-types";
import {
  TableHead,
  TableRow,
  TableCell,
  TableSortLabel,
  Tooltip
} from "@material-ui/core";

const headers = [
  { id: "ResultAlias", numeric: false, label: "User" },
  { id: "ID", numeric: true, label: "Rank" },
  { id: "TotalTime", numeric: true, label: "Time (s)" },
  { id: "TotalScore", numeric: true, label: "Score" }
];

class LeaderboardHeader extends Component {
  createSortHandler = property => event => {
    console.log(property);
    this.props.onRequestSort(event, property);
  };

  render() {
    const { order, orderBy } = this.props;

    return (
      <TableHead>
        <TableRow>
          {headers.map(header => {
            return (
              <TableCell
                key={header.id}
                numeric={header.numeric}
                sortDirection={orderBy === header.id ? order : false}
              >
                <Tooltip
                  title="Sort"
                  placement={header.numeric ? "bottom-end" : "bottom-start"}
                  enterDelay={300}
                >
                  <TableSortLabel
                    active={orderBy === header.id}
                    direction={order}
                    onClick={this.createSortHandler(header.id)}
                  >
                    {header.label}
                  </TableSortLabel>
                </Tooltip>
              </TableCell>
            );
          }, this)}
        </TableRow>
      </TableHead>
    );
  }
}

LeaderboardHeader.propTypes = {
  order: PropTypes.string.isRequired,
  orderBy: PropTypes.string.isRequired,
  onRequestSort: PropTypes.func.isRequired
};

export default LeaderboardHeader;
