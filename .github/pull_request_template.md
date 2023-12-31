# PR Standards

## Opening a pull request should be able to meet the following requirements:

---

For Author:

- [ ]  Targeted PR against the correct branch
- [ ]  included the correct [type prefix](https://github.com/commitizen/conventional-commit-types/blob/v3.0.0/index.json) in the PR title
- [ ]  Linked to the GitHub issue with discussion and accepted design
- [ ]  Targets only one GitHub issue
- [ ]  Wrote unit and integration tests
- [ ]  Wrote relevant migration scripts if necessary
- [ ]  All CI checks have passed
- [ ]  Added relevant `godoc` [comments](https://blog.golang.org/godoc-documenting-go-code)

---

For Reviewer:

- [ ]  confirmed the correct [type prefix](https://github.com/commitizen/conventional-commit-types/blob/v3.0.0/index.json) in the PR title
- [ ]  Reviewers assigned
- [ ]  confirmed all author checklist items have been addressed

---;

After reviewer approval:

- [ ]  In case it targets the main branch, PR should be squashed and merged.
- [ ]  In case the PR targets a release branch, PR should be rebased.
