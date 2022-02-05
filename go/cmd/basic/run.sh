#!/usr/bin/env bash
set -eou pipefail
killall sb 2>/dev/null||true

clear >/dev/null 2>&1
go build -o sb .||{ ./fix.sh && clear && go build -o sb .; }


cu(){
  reset
}
trap cu EXIT

eval ./sb ${@:-}
