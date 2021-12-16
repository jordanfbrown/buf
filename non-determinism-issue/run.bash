#!/usr/bin/env bash

set -euo pipefail

for i in $(seq 1 100); do
  rm -rf out
  mkdir -p out
  buf generate 2>/dev/null
  grep sources out/foo.py
done
