#!/bin/bash

set -e

python3 -m venv .ow-test
source .ow-test/bin/activate
pip3 --version