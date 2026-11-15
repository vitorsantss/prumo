# First contribution, step by step

> **Translated by AI, waiting for a human review.**

First, read [Start here](comece-aqui.md). Then:

1. **Pick a task** among the ones [marked for newcomers](https://github.com/vitorsantss/prumo/issues?q=is%3Aopen+label%3A%22good+first+issue%22). Read the whole task: it says what to do and when it is done.
2. **Say you are taking it** with a comment on the task, so two people do not do the same thing.
3. **Fork** the repository on GitHub and clone your fork:
   ```sh
   git clone https://github.com/<your-user>/prumo
   cd prumo
   ```
4. **Create a branch** with a name that says what it is:
   ```sh
   git checkout -b docs/review-start-here-translation
   ```
5. **Make the change.** If you touch code, run the tests (requires Go 1.24 or newer):
   ```sh
   go test ./...
   ```
6. **Commit** with a message saying why. The project writes commits in Portuguese, but English is fine for a first contribution. If you used AI, you can credit it with a `Co-Authored-By` line.
7. **Open the PR** to the original repository. Write `Closes #<task number>` in the description and fill in the template, including "how it was made".
8. **Follow the review.** A maintainer comments on the PR; adjust it on the same branch. Once approved, it is squash-merged: the PR title becomes the commit (ADR-0008).

> **Note.** The repository's automated tests are currently paused because of an account issue. Until they are back, reviewers run the tests by hand.

Questions at any step: ask on the task itself or in the [repository discussions](https://github.com/vitorsantss/prumo/discussions).
