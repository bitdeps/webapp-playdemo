# webapp-playdemo

This project is a playful demonstration of a Golang application, showcasing various development approaches:

- **Simplified GitFlow:** Implements a basic GitFlow with separate development (dev) and production (prod) releases. Note that it doesn't fully automate Continuous Integration (CI) for release branches.
- **Automatic Version Bumping:** Utilizes the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for automatic version incrementation.
- **GoLang CI Linting:** Integrates [GoLang CI lint](https://golangci-lint.run/) to support multiple linting tools available for Go applications.
- **Container Image Releases:** Generates container images for both pre-release (development environments) and production environments.
- **Go Application Release:** Utilizes [GoReleaser](https://goreleaser.com/) to create a GitHub release, marking a version milestone.
- **Configuration Code Separation:** Centralizes configuration in [a separate repository](https://github.com/dennybaa/cue-values/tree/main), facilitating GitOps CD tools. Updated values enable automated deployment through GitOps Continuous Delivery (CD) processes.

## TL;DR;

The workflow is as follows:

1. **Pull Request (PR):** Executes configured lints during a PR.
2. **PR Merge into Dev Branch:** Creates a pre-release [semver](https://semver.org/) tag (e.g., `0.1.5-dev.2`). Pushes a container image with the corresponding tag to the [Harbor container registry](https://c8n.io/). Updates version values in the values repository for specified applications.
3. **Development PR (against a development branch):** Labeled as `release: pending`. Offers a convenient helper for creating a release PR into the production branch (main).
4. **Release PR:** Runs all lints again. Once approved and merged, GoReleaser creates a GitHub release. Subsequently, the release container image is pushed into the container registry, and prod version values are updated.
