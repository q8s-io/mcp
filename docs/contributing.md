# Contributing

Welcome to the Multi Cloud Platform contributing guide.  We are excited
about the prospect of you joining us!

## Before You Begin

- we strongly recommend you to use goland and unified configuration as follows:
![image](https://github.com/DoubleBabylol/mcp/blob/0.1.1/docs/images/config.png)

- We strongly recommend you to understand the main [develop.md](https://github.com/q8s-io/mcp/blob/master/docs/develop.md) and adhere to the contribution rules .

- Please be aware that all contributions to Multi Cloud Platform require time and commitment from project maintainers to direct and review work. This is done in additional to many other maintainer responsibilities, and direct engagement from maintainers is a finite resource.


## Begin

1. fork repo to your github
2. clone to the local branch and do some development
3. Initiate a pull request to the original github


## Implementation

Implementation PRs should
- mention the issue of the associated design proposal

Small features and flag changes require only unit/integration tests,
while larger changes require both unit/integration tests and e2e tests.


## Merge state meanings

- Merged:
  - Ready to be implemented.
- Unmerged:
  - Experience and design still being worked out.
  - Not a high priority issue but may implement in the future: revisit
    in 6 months.
  - Unintentionally dropped.
- Closed:
  - Not something we plan to implement in the proposed manner.
  - Not something we plan to revisit in the next 12 months.
