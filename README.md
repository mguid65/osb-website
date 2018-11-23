# osb-website

Repo for the [OpenSystemBench](https://github.com/mguid65/OpenSystemBench) website.

## Development

Run the following commands in separate terminals to run the website locally:

```
# create go project directory structure
mkdir -p go/src/github.com/<github-username>/

# go to folder
cd go/src/github.com/<github-username>/

# clone repo
git clone https://github.com/mguid65/osb-website.git

# go into project folder
cd osb-website

# build react packages
npm run build

# move build folder into server
sudo cp -r build server/

# run the server
go run server/main.go
```
