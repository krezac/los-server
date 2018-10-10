#!/bin/sh
# Generate test coverage statistics for Go packages.
# based on https://github.com/mlafeldt/chef-runner/blob/v0.7.0/script/coverage
#
# Works around the fact that `go test -coverprofile` currently does not work
# with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909
#
# Usage: ./coverall.sh [--html|--coveralls]
#
#     --html      Additionally create HTML report and open it in browser
#

set -e

workdir=.cover
profile="$workdir/cover.out"
mode=count

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    go test -i "$@" # compile dependencies first before serializing go test invocations

    for pkg in "$@"; do
        f="$workdir/$(echo $pkg | tr / -).cover"
        go test -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
}

show_cover_report() {
    go tool cover -${1}="$profile" -o coverage.html
}

calculate_total_coverage() {
    go tool cover -func "$profile"  | tail -n 1 | sed -e 's/[^0-9]*\([0-9\.]*\)*/\1/g' > ${1}
}

generate_cover_data $(go list ./... | egrep -v '/vendor/')
show_cover_report html
calculate_total_coverage coverage_total.txt
echo "TOTAL: " `cat coverage_total.txt`"%"