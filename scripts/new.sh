#!/usr/bin/env bash
# Creates a new problem skeleton.
#   LeetCode: scripts/new.sh 1 two-sum      -> leetcode/0001-two-sum/
#   NeetCode: scripts/new.sh nc design-dynamic-array -> neetcode/design-dynamic-array/
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $0 <leetcode-number|nc> <slug>" >&2
  exit 1
fi

num=$1
slug=$(echo "$2" | tr '[:upper:]' '[:lower:]')
pkg=${slug//-/}

if [[ $num == nc ]]; then
  dir="neetcode/$slug"
  link="https://neetcode.io/problems/$slug"
else
  dir="leetcode/$(printf '%04d' "$((10#$num))")-$slug"
  link="https://leetcode.com/problems/$slug/"
fi

if [[ -e $dir ]]; then
  echo "$dir already exists" >&2
  exit 1
fi

mkdir -p "$dir"

cat > "$dir/solution.go" <<GO
// Package $pkg solves $link
package $pkg

func solve() {
}
GO

cat > "$dir/solution_test.go" <<GO
package $pkg

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "example 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve()
		})
	}
}
GO

echo "created $dir"
