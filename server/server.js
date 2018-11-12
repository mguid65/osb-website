const express = require("express");
const cors = require("cors");
const app = express();
const mysql = require("mysql");

const pool = mysql.createPool({
  host: process.env.HOST,
  user: process.env.USER,
  password: process.env.PASS,
  database: process.env.DB,
  supportBigNumbers: true,
  waitForConnections: true,
  connectionLimit: 10
});

app.get("/scores", cors(), (req, res) => {
  pool.getConnection((err, conn) => {
    if (err) throw err;

    conn.query("SELECT * FROM scores", (err, rows) => {
      res.json(rows);
      conn.release();
      if (err) throw err;
    });
  });
});

const PORT = process.env.PORT || 8080;
app.listen(PORT, () =>
  console.log(`Server listening at http://localhost:${PORT}/`)
);
