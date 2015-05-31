#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

echo "=== install syncloud dependencies ==="
wget --no-check-certificate --progress=dot:mega -O get-pip.py https://bootstrap.pypa.io/get-pip.py 2>&1
python get-pip.py

pip2 install -U pytest
pip2 install -r /requirements.txt
python setup.py
pip2 freeze | grep syncloud

syncloud-platform-post-install

export TEAMCITY_VERSION=9
py.test -s verify.py