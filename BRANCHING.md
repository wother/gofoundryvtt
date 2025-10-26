# Branch Strategy

This document outlines the branching strategy for the gofoundryvtt project.

## Branch Structure

### `main`
- **Protected**: Yes
- **Purpose**: Production-ready code
- **Merge Strategy**: Squash merge from `develop` or hotfix branches
- **Status**: All checks must pass (CI, tests, coverage >= 80%)
- **Version**: Released versions only (v0.1.0, v1.0.0, etc.)

### `develop`
- **Protected**: Yes
- **Purpose**: Integration branch for features
- **Merge Strategy**: Squash merge from feature branches
- **Status**: All checks must pass
- **Version**: Pre-release versions (v0.1.0-rc.1, etc.)

### Feature Branches
- **Naming**: `feature/<description>` or `feat/<description>`
- **Purpose**: New features and enhancements
- **Base**: `develop`
- **Merge Target**: `develop`
- **Lifecycle**: Delete after merge

### Bugfix Branches
- **Naming**: `bugfix/<description>` or `fix/<description>`
- **Purpose**: Bug fixes
- **Base**: `develop`
- **Merge Target**: `develop`
- **Lifecycle**: Delete after merge

### Hotfix Branches
- **Naming**: `hotfix/<version>`
- **Purpose**: Critical production fixes
- **Base**: `main`
- **Merge Target**: `main` AND `develop`
- **Lifecycle**: Delete after merge

### Release Branches
- **Naming**: `release/<version>`
- **Purpose**: Prepare release (version bump, changelog, final testing)
- **Base**: `develop`
- **Merge Target**: `main` AND `develop`
- **Lifecycle**: Delete after merge

## Branch Protection Rules

### For `main`:
- Require pull request reviews before merging (1 approval minimum)
- Require status checks to pass before merging:
  - CI tests (all platforms)
  - Lint checks
  - Format checks
  - Coverage >= 80%
- Require branches to be up to date before merging
- Require signed commits (optional, recommended)
- Include administrators in restrictions
- Restrict force pushes
- Restrict deletions

### For `develop`:
- Require pull request reviews before merging (1 approval minimum)
- Require status checks to pass before merging:
  - CI tests
  - Lint checks
  - Coverage >= 80%
- Require branches to be up to date before merging
- Restrict force pushes (allow for rebase workflows)
- Restrict deletions

## Workflow

### Adding a New Feature
```bash
git checkout develop
git pull origin develop
git checkout -b feature/my-feature
# ... make changes ...
git add .
git commit -m "feat(scope): description"
git push origin feature/my-feature
# Create PR to develop
```

### Fixing a Bug
```bash
git checkout develop
git pull origin develop
git checkout -b bugfix/fix-issue
# ... make changes ...
git add .
git commit -m "fix(scope): description"
git push origin bugfix/fix-issue
# Create PR to develop
```

### Creating a Release
```bash
git checkout develop
git pull origin develop
git checkout -b release/v1.0.0
# Update VERSION file, CHANGELOG.md
git add .
git commit -m "chore: prepare release v1.0.0"
git push origin release/v1.0.0
# Create PR to main
# After merge to main, tag the release
git checkout main
git pull origin main
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
# Merge back to develop
git checkout develop
git merge main
git push origin develop
```

### Hotfix for Production
```bash
git checkout main
git pull origin main
git checkout -b hotfix/v1.0.1
# ... fix critical issue ...
git add .
git commit -m "fix: critical security issue"
git push origin hotfix/v1.0.1
# Create PR to main
# After merge, tag the hotfix
git checkout main
git pull origin main
git tag -a v1.0.1 -m "Hotfix v1.0.1"
git push origin v1.0.1
# Merge back to develop
git checkout develop
git merge main
git push origin develop
```

## Commit Message Format

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Code style changes (formatting, no code change)
- `refactor`: Code refactoring (no feature or bug fix)
- `perf`: Performance improvements
- `test`: Adding or updating tests
- `chore`: Maintenance tasks (dependencies, build, etc.)
- `ci`: CI/CD changes

### Examples
```
feat(actors): implement GetActor and ListActors
fix(client): handle connection timeout correctly
docs(readme): update installation instructions
test(scenes): add unit tests for scene operations
chore(deps): update dependencies
```

## Version Numbering

Follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version: Incompatible API changes
- **MINOR** version: New functionality (backwards compatible)
- **PATCH** version: Bug fixes (backwards compatible)

### Pre-release Versions
- `v0.0.0-pre`: Initial development
- `v0.1.0-alpha.1`: Alpha releases
- `v0.1.0-beta.1`: Beta releases
- `v0.1.0-rc.1`: Release candidates

### Current Version
- `v0.0.0-pre`: Pre-release development phase

## GitHub Actions

All workflows are currently set to **manual trigger only** (`workflow_dispatch`). 

To enable automatic runs:
1. Edit workflow files in `.github/workflows/`
2. Uncomment the appropriate trigger sections
3. Commit and push changes

Available workflows:
- **CI** (`ci.yml`): Run tests, lint, format checks
- **Release** (`release.yml`): Build and publish releases
- **Security** (`security.yml`): Security scanning
