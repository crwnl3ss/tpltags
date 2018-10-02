#### tpltags

*tpltags* - set of usefull template filters and tags.

### Iterator
Iterator generates slice of uint from 0 to n-1.
Usage:
```
{{ range Iterator 5 }} {{ . }} {{ end }}
```
Result:
```
01234
```