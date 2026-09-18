# greenlight-fixtures

Protection-rule and dependency-update fixtures for
[greenlight](https://github.com/UiPath/greenlight). **This is not a real
project.** Nothing here ships anywhere and none of the dependencies are used.

Every branch under `fx/` is one protection scenario, with its own ruleset or
classic branch protection. Manifests are deliberately out of date so Dependabot
opens real pull requests against each one.

| Branch | Scenario |
|---|---|
| `fx/open` | No rules at all. Baseline: greenlight should merge straight away. |
| `fx/approval` | One required approval, no code owners. The app's own review satisfies the rule, so no bypass event should be emitted. |
| `fx/codeowners` | One required approval plus code owner review. An app cannot be a code owner, so bypass must fire here. |
| `fx/checks` | Requires the 'ci/build' status, which nothing produces automatically. Drive it with 'fixtures.sh status'. |
| `fx/convo` | Requires review thread resolution. Leave an unresolved comment to exercise it. |
| `fx/uptodate` | Strict required checks: the branch must be current with its base before merging. |
| `fx/queue` | Requires a merge queue. greenlight should enqueue rather than bypass. |
| `fx/classic` | Classic branch protection rather than a ruleset. Reviews are bypassable for an app here; status checks are not, by anyone. |

The `ci/build` status that some branches require is not produced by any
workflow. It is set by hand so that green, red and never-reported are all
reproducible:

```
scripts/fixtures.sh status <pr> ci/build failure
```
