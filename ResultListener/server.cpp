#include <iostream>
#include "TCPServer.h"
#include <string>
#include <boost/algorithm/string/classification.hpp>
#include <boost/algorithm/string/split.hpp>
#include <vector>

  TCPServer tcp;
/* Copyright 2008, 2010, Oracle and/or its affiliates. All rights reserved.

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; version 2 of the License.

There are special exceptions to the terms and conditions of the GPL
as it is applied to this software. View the full text of the
exception in file EXCEPTIONS-CONNECTOR-C++ in the directory of this
software distribution.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
*/

/* Standard C++ includes */

#include <stdlib.h>

  /*
  Include directly the different
  headers from cppconn/ and mysql_driver.h + mysql_util.h
  (and mysql_connection.h). This will reduce your build time!
  */
#include "mysql_connection.h"

#include <cppconn/driver.h>
#include <cppconn/exception.h>
#include <cppconn/resultset.h>
#include <cppconn/statement.h>
#include <cppconn/prepared_statement.h>

  using namespace std;

int MySQLUpload(std::vector<std::string> str) {
  cout << "Attempting to Insert Into Database...\n";
  cout.flush();

  try {
    sql::Driver * driver;
    sql::Connection * con;
//    sql::Statement * stmt;
//    sql::ResultSet * res;
    sql::PreparedStatement * pstmt;
    //cout<<"test1";
    //cout.flush();
    /* Create a connection */
    driver = get_driver_instance();
    con = driver->connect("tcp://127.0.0.1/leaderboard", "root", "");
    /* Connect to the MySQL test database */
    //cout<<"test2";
    //cout.flush();
    pstmt = con->prepareStatement("INSERT INTO scores(ResultAlias, NBodyTime, NBodyScore, PiDigitsTime, PiDigitsScore, MandelbrotTime, MandelbrotScore, SpectralNormTime, SpectralNormScore, BinaryTreesTime, BinaryTreesScore, TotalTime, TotalScore, Overclocked) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)");
    cout << str[0].size();
    cout.flush();
    
    if (str[0].empty() || string.size() > 19) { //Error handling for not filling out each entry
        return 1;
    }
    pstmt->setString(1, str[0]);
    //cout<<"test4";
    //cout.flush();
    pstmt->setDouble(2, stod(str[2]));
    pstmt->setDouble(3, stod(str[3]));
    pstmt->setDouble(4, stod(str[5]));
    pstmt->setDouble(5, stod(str[6]));
    pstmt->setDouble(6, stod(str[8]));
    pstmt->setDouble(7, stod(str[9]));
    pstmt->setDouble(8, stod(str[11]));
    pstmt->setDouble(9, stod(str[12]));
    pstmt->setDouble(10, stod(str[14]));
    pstmt->setDouble(11, stod(str[15]));
    pstmt->setDouble(12, stod(str[16]));
    pstmt->setDouble(13, stod(str[17]));
    pstmt->setInt(14, stoi(str[18]));
    //cout<<"test5";
    //cout.flush();

    pstmt->execute();
    //cout<<"test6";
    //cout.flush();
    delete pstmt;
    delete con;
  } catch (sql::SQLException & e) {
    //cout << "# ERR: SQLException in " << __FILE__;
    //cout << "(" << __FUNCTION__ << ") on line "» << __LINE__ << endl;
    //cout << "# ERR: " << e.what();
    cout << " (MySQL error code: " << e.getErrorCode();
    cout.flush();
    //cout << ", SQLState: " << e.getSQLState() << »
    //  " )" << endl;
    return 1;
  }

  cout << endl;

  return 0;
}

void * loop(void * m) {
  pthread_detach(pthread_self());
  std::vector < std::string > splitResults;
  while (1) {
    srand(time(NULL));
    // char ch = 'a' + rand() % 26;
    // string s(1,ch);
    string str = tcp.getMessage();
    if (str != "") {
      boost::split(splitResults, str, boost::is_any_of(";"));
        for (int i = 0; i < splitResults.size(); i++) {
        std::cout << splitResults[i] << "\n";
      }

     /// parseResultString(str, splitResults);
      int pass_flag = MySQLUpload(splitResults);
      if (pass_flag == 0) {
        tcp.Send("Upload Success!");
      } else {
        tcp.Send("Upload Failure!");
      }
      tcp.clean();
    }
    //              std::cout << "test";
    //              std::cout.flush();
    usleep(1000);
  }
  tcp.detach();
}

int main() {
  pthread_t msg;
  tcp.setup(47002);
  if (pthread_create( & msg, NULL, loop, (void * ) 0) == 0) {
    tcp.receive();
  }
  return 0;
}
