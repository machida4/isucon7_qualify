#!/bin/bash -xe
cd /home/isucon/isubata/webapp/go
git pull
make
sudo systemctl restart isubata.golang.service
sudo systemctl restart nginx.service