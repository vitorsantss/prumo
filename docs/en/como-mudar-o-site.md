# How to change the site and the docs

> **Translated by AI, waiting for a human review.**

Anyone can propose a change to the home page, the docs or the visuals. The path is the same one Prumo argues for in software: the idea is written down and approved before the code.

## Only the text of the docs

Fixed a mistake, made something clearer, added a source? The footer of each page says which file it comes from (for example, `docs/a-ideia.md`). Open that file [in the repository](https://github.com/vitorsantss/prumo/tree/main/docs), edit it on GitHub and open the PR directly, no issue needed.

Prefer to just report it? [Open an issue](https://github.com/vitorsantss/prumo/issues/new?template=site.yml) saying which page and what is wrong.

## How to write

- **Every citation has a link.** A claim about the world carries its source, preferably as a footnote (`text[^1]` and, at the end, `[^1]: Author, *Title* · URL`). A project, tool or challenge mentioned by name links to its page.
- **Which link to use.** The official page when there is one (project, specification, docs). Paper or standard: the DOI (`https://doi.org/...`). Book: its [Open Library](https://openlibrary.org) record, which shows the ISBN and where to find it. Prefer the primary source over a text that comments on it.
- **A source without a link: how to tell if it is reliable.**
  1. Look for an identifier: DOI, ISBN, arXiv or standard number. Found one, use it as the link.
  2. Check who published it: an author with recognized work in the field, a publisher, a peer-reviewed journal, an official body or the project owner.
  3. Check it still holds: a page marked "no longer maintained", an inactive standard or an old version changes what we can claim.
  4. Read the passage: does the source really say what we claim? A third-party summary does not count.
  5. Did not pass? Do not cite it. If it is our own experience, write it as our experience, without attributing it to anyone.
- **Internal citations too.** Write `ADR-0007` or `#50` as usual: the site turns them into links to the ADR and the issue. If the ADR is not merged yet, the build warns and the link appears by itself after the merge.

## Home page content

Hypotheses, the timeline and the home page texts live in `site/src/data/landing.json`, in Portuguese and English. If the project state changed (a hypothesis was tested, a milestone was reached), update that file in the same PR. Checking this automatically is issue #50.

## Visuals, components or navigation

1. **Open an issue** with the ["Mudança no site ou na documentação"](https://github.com/vitorsantss/prumo/issues/new?template=site.yml) template: the problem you saw, the proposal and who is affected.
2. **Show how it looks**, in light and dark themes. Screenshots, an Excalidraw sketch or a prototype all work; what matters is comparing before and after.
3. **Check the rules** of the [site layer of the design system](https://github.com/vitorsantss/prumo/blob/main/docs/design/site.md). If the proposal breaks a rule, say which one and why: sometimes the rule is what needs to change.
4. **Wait for approval on the issue** before coding. That is the human gate: it avoids a large PR being rejected because of the design.
5. **Open the PR** with:
   - the code;
   - `docs/design/site.md` updated, if a new rule or component came up;
   - an ADR, if a new dependency or pattern came in;
   - screenshots in light and dark themes, on phone, laptop and wide screen;
   - `cd site && npm run build` passing.
6. **Review and squash merge**, like every PR ([ADR-0008](../adr/0008-merge-por-squash.md)).

## Test on your computer

```
cd site
npm install
npm run dev
```

Search only works after `npm run build && npm run preview`, because the index is generated at build time.
