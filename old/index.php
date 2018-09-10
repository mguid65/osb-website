<!DOCTYPE html>
<html>
 <head class="header">
  <link href="main.css" type="text/css" rel="stylesheet"/>
  <title>OpenSystemBench</title>
  <meta content="width=device-width, initial-scale=1"/>
 </head>
 <body class="back">
  <div class="topbar">
   <div class="topcontainer">
    <h2 class="topbartext">OpenSystemBench</h2>
   </div>
  </div>
  <br>
  <div class="normalbodymargin">
   <div>
    <div class="wide-card card-2">
     <p class="content cardheadtext">RELEASE: Ubuntu-17.10-GUI</p>
     <p class="content regulartext">
      This is a BETA. It has 5 different algorithms for testing CPU Single/Multi Core and IO.
      You can get a score and a time and submit it to the server. You can also save results to a CSV file.
      This has not been tested in other versions of Linux.
     </p>
     <a class="card-1 btn regulartext" href="/OpenSystemBench" download="OpenSystemBench" style="text-decoration: none"><b>Download</b></a>
    </div>
    <div class="wide-card card-2">
     <p class="content cardheadtext">RELEASE: Ubuntu-17.10-CLI</p>
     <p class="content regulartext">
      This is a BETA commandline version. It only allows standard tests. It also allows for submissions.<br> 
     </p>
     <a class="card-1 btn regulartext" href="/OpenSystemBenchCLI" download="OpenSystemBenchCLI" style="text-decoration: none"><b>Download</b></a>
    </div>
    <div class="wide-card card-2">
     <p class="content cardheadtext">Source</p>
     <p class="content regulartext">
      The source code is available at:
      <a href="https://github.com/mguid65/OpenSystemBench" style="text-decoration: none"><b>OpenSystemBench</b></a>
     </p>
    </div>
   </div>
   <div>
   <table class="table table-striped">
    <thead>
     <tr>
      <th>Rank</th>
      <th>Name</th>
      <th>Score</th>
     </tr>
    </thead>
    <tbody>
     <?php 
     $servername = "localhost";
     $username = "reader";
     $password = "2j6d2ve";
     $dbname = "leaderboard";

     $conn = new mysqli($servername, $username, $password, $dbname);

     if($conn->connect_error) {
       die("Connection failed: ". $conn->connect_error);
     }
     if (isset($_GET["page"])) { $page = $_GET["page"]; } else { $page=1;};
     $start_from = ($page-1) * 10;
     $sql = "SELECT ResultAlias, TotalScore FROM scores ORDER by TotalScore desc LIMIT " . $start_from . ", 10";
     $result = $conn->query($sql);
     $counter = $start_from + 1;
     if ($result->num_rows > 0) {
       while($row = $result->fetch_assoc()) {
         echo "<tr><td>" . $counter. "</td><td>" . $row[ResultAlias]. "</td><td>" . $row[TotalScore]. "</td></tr>";
         $counter++; 
       }
     }
     else {
       echo "0 Results";
     }
     ?>
    </tbody>
   </table>
    <?php 
     $sql = "SELECT COUNT(ID) AS total FROM scores";
     $res = $conn->query($sql);
     $row = $res->fetch_assoc();
     $total_pages = ceil($row["total"] / 10); // calculate total pages with results
     echo "<span class='pageNum'>";
     for ($i=1; $i<=$total_pages; $i++) {  // print links for all pages
            echo "<a href='index.php?page=".$i."'";
            if ($i==$page)  echo " class='page'";
            echo ">".$i."</a> "; 
     };
     echo "</span>";

     $conn->close(); 
    ?>
  </div>
  <footer class="footnote absolute">Matthew Guidry, Scott Wilder 2018</footer>
 </body>
</html>
