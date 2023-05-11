Simple test app using Bazel build system for the project:
[Dependent GitHub CI/CD gating on scale](https://github.com/polupanovaanna/go_git_actions)

It contains the ".github/workflows/action.yml` file, which is run as a check when a PR is opened.
This action collects data from pull requests, such as differences with master HEAD commits and Bazel 
targets affected by this change.

To build and run the project:

```bazel build //a_app:a_build //b_app:b_build //c_app:c_build```

```bazel run <target>```

Where target is one of the `//a_app:a_build, //b_app:b_build, //c_app:c_build`

There are 4 branches:

- `branch_a`: contains changes in a_app target 
- `branch_b`: contains changes in b_app target 
- `c_first`: change the variable name in c_app target
- `c_second`: add some new logic in c_app target

`c_first` and `c_second` are conflicting, so if two pull requests with both branches are opened,
the second couldn't be merged because of the failing checks.