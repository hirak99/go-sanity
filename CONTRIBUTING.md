# How to Contribute

We'd love to accept your patches and contributions to this project. There are
just a few small guidelines you need to follow.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Code Reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows [Google's Open Source Community
Guidelines](https://opensource.google/conduct/).

## Developing

If you want to develop as part of your code, it may be convenient to use a local copy. To do that, add the followign to the go.mod of your project -

```
...
require github.com/hirak99/go-sanity v0.1.4
replace github.com/hirak99/go-sanity v0.1.4 => ./sanity
...
```

Then copy the .go files from this project, to the subdirectory ./sanity of your project.

## Releasing

For most people, a maintainer will release a new version after pulling in your changes.

If you have access to the git repo, simply follow this process to release a version.

```bash
# Add a tag locally incrementing version.
git tag v1.0.0

# Update the tag on to the repo.
git push origin v1.0.0
```