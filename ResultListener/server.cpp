#include <iostream>
#include "TCPServer.h"
#include <string>
#include <boost/algorithm/string/classification.hpp>
#include <boost/algorithm/string/split.hpp>
#include <vector>

TCPServer tcp;

void parseResultString(std::string str) {
    std::vector<std::string> splitResults;
    boost::split(splitResults, str, boost::is_any_of(";"));
    for(int i = 0; i < splitResults.size(); i++) {
      std::cout << splitResults[i] << "\n";
    }
}

void * loop(void * m) {
        pthread_detach(pthread_self());
        while(1) {
                srand(time(NULL));
                // char ch = 'a' + rand() % 26;
                // string s(1,ch);
                string str = tcp.getMessage();
                if( str != "" ) {
                        parseResultString(str);
                        tcp.Send("Success!");
                        tcp.clean();
                }
//              std::cout << "test";
//              std::cout.flush();
                usleep(1000);
        }
        tcp.detach();
}

int main()
{
        pthread_t msg;
        tcp.setup(47002);
        if( pthread_create(&msg, NULL, loop, (void *)0) == 0)
        {
                tcp.receive();
        }
        return 0;
}

