import React, { Component } from "react";
import PropTypes from "prop-types";
import {
  TableHead,
  TableRow,
  TableCell,
  TableSortLabel,
  Tooltip
} from "@material-ui/core";
import ExpandMore from "@material-ui/icons/ExpandMore";

const headers = [
  {
    id: "username",
    numeric: false,
    disablePadding: true,
    label: "Username"
  },
  {
    id: "rank",
    numeric: true,
    disablePadding: false,
    label: "Rank"
  },
  {
    id: "totalTime",
    numeric: true,
    disablePadding: false,
    label: "Time (s)"
  },
  {
    id: "totalScore",
    numeric: true,
    disablePadding: false,
    label: "Score"
  }
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
          <TableCell padding="checkbox" width={50}>
            <Tooltip title="Expand All" position="bottom">
              <ExpandMore />
            </Tooltip>
          </TableCell>
          {headers.map(header => {
            return (
              <TableCell
                key={header.id}
                numeric={header.numeric}
                padding={header.disablePadding ? "none" : "default"}
                sortDirection={orderBy === header.id ? order : false}
              >
                <Tooltip
                  title="Sort"
                  placement={header.numeric ? "bottom-end" : "bottom-start"}
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
