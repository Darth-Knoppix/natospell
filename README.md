# NATO Spell

Transform a phrase into phonetic spelling for use over phone/radio.

e.g. `./natospell hello there`

```plain
Original: h     e    l    l    o     t     h     e    r     e
Transformed: Hotel Echo Lima Lima Oscar Tango Hotel Echo Romeo Echo
```

## Installation

Check [releases](https://github.com/Darth-Knoppix/natospell/releases) for the latest version. There are binaries for a a few common systems to choose from.


## Dev: Releasing

1. Add release token: `export GITHUB_TOKEN="YOUR_GH_TOKEN"`

2. Tag the release:
```sh
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
```

3. Release: `goreleaser release`
