# How to change the site and the docs

> **Translated by AI, waiting for a human review.**

Anyone can propose a change to the home page, the docs or the visuals. The path is the same one Prumo argues for in software: the idea is written down and approved before the code.

## Only the text of the docs

Fixed a mistake, made something clearer, added a source? The footer of each page says which file it comes from (for example, `docs/a-ideia.md`). Open that file [in the repository](https://github.com/vitorsantss/prumo/tree/main/docs), edit it on GitHub and open the PR directly, no issue needed.

Prefer to just report it? [Open an issue](https://github.com/vitorsantss/prumo/issues/new?template=site.yml) saying which page and what is wrong.

## How to write

- **Every citation has a link.** A claim about the world carries its source, preferably as a footnote (`text[^1]` and, at the end, `[^1]: Author, *Title* · URL`). A project, tool or challenge mentioned by name links to its page.
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
6. **Review and squash merge**, like every PR (ADR-0008).

## Test on your computer

```
cd site
npm install
npm run dev
```

Search only works after `npm run build && npm run preview`, because the index is generated at build time.
