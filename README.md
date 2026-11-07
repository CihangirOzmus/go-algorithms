# go-algorithms

My LeetCode / NeetCode solutions in Go.

## Layout

```
leetcode/NNNN-slug/     one package per LeetCode problem (4-digit number + URL slug)
  solution.go           main solution; alternatives go in solution_<approach>.go
  solution_test.go
neetcode/slug/          NeetCode-only problems (not on LeetCode)
internal/ds/            shared ListNode / TreeNode + test builders (NewList, NewTree)
scripts/new.sh          problem scaffold
```

- Package name = slug without dashes (`0020-valid-parentheses` → `package validparentheses`).
- Every problem is its own package, so `Constructor`, `ListNode` and helper
  names never collide.
- Topics and difficulty are tracked in the table below, not in the folder tree.

## Usage

```sh
make new N=1 SLUG=two-sum          # -> leetcode/0001-two-sum/
make new N=nc SLUG=some-problem    # -> neetcode/some-problem/
make test                          # all problems
make test P=0020                   # one problem
```

## Problems

| #    | Problem | Difficulty | Topics | NeetCode |
|------|---------|------------|--------|----------|
| 0020 | [Valid Parentheses](leetcode/0020-valid-parentheses) | Easy | Stack, String | ✅ |
| 0027 | [Remove Element](leetcode/0027-remove-element) | Easy | Array, Two Pointers | ✅ |
| 0155 | [Min Stack](leetcode/0155-min-stack) | Medium | Stack, Design | ✅ |
| 0206 | [Reverse Linked List](leetcode/0206-reverse-linked-list) | Easy | Linked List | ✅ |
| 0485 | [Max Consecutive Ones](leetcode/0485-max-consecutive-ones) | Easy | Array | |
| 0682 | [Baseball Game](leetcode/0682-baseball-game) | Easy | Stack | ✅ |
| 0707 | [Design Linked List](leetcode/0707-design-linked-list) | Medium | Linked List, Design | ✅ |
| 1299 | [Replace Elements with Greatest Element on Right Side](leetcode/1299-replace-elements-with-greatest-element-on-right-side) | Easy | Array | ✅ |
| 1385 | [Find the Distance Value Between Two Arrays](leetcode/1385-find-the-distance-value-between-two-arrays) | Easy | Array, Binary Search | |
| 1472 | [Design Browser History](leetcode/1472-design-browser-history) | Medium | Linked List, Design | ✅ |
| 1929 | [Concatenation of Array](leetcode/1929-concatenation-of-array) | Easy | Array | ✅ |
| 2089 | [Find Target Indices After Sorting Array](leetcode/2089-find-target-indices-after-sorting-array) | Easy | Array, Sorting | |
| 2409 | [Count Days Spent Together](leetcode/2409-count-days-spent-together) | Easy | Math, String | |
