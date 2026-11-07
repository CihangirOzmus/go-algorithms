.PHONY: new test vet

# make new N=1 SLUG=two-sum        -> leetcode/0001-two-sum
# make new N=nc SLUG=some-problem  -> neetcode/some-problem
new:
	@scripts/new.sh $(N) $(SLUG)

# make test            -> all problems
# make test P=0020     -> only problem 0020
test:
ifdef P
	go test ./leetcode/$(P)-*/...
else
	go test ./...
endif

vet:
	go vet ./...
