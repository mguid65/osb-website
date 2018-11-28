#!/bin/bash

sudo rm -r build
sudo rm -r server/build

npm run build

sudo cp -r build server/
