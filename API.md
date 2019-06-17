## automatabot API


### get a random challenge

`GET https://api.noopschallenge.com/automatabot/challenges/new`

`HTTP 200`

```
{
  "challengePath": "/automatabot/challenges/9IwtwpGWs2zjEel1BZ02",
  "challenge": {
    "rules": {
      "name": "mazectric",
      "survival": [ 1, 2, 3, 4 ],
      "birth": [ 3 ]
    },
    "cells": [
      [ 1, 1, 1, 0, 0 ],
      [ 1, 0, 1, 0, 0 ],
      [ 0, 1, 0, 1, 1 ],
      [ 0, 0, 1, 1, 0 ],
      [ 0, 0, 0, 1, 0 ]
    ],
    "generations": 29
  }
}
```


### get a challenge again

`GET https://api.noopschallenge.com/automatabot/challenges/keMwxJadhTgXxyVwJ610`

`HTTP 200`

```
{
  "challengePath": "/automatabot/challenges/undefined",
  "challenge": {
    "rules": {
      "name": "coral",
      "survival": [ 4, 5, 6, 7, 8 ],
      "birth": [ 3 ]
    },
    "cells": [
      [ 0, 1, 0, 0, 1 ],
      [ 0, 1, 0, 1, 0 ],
      [ 1, 0, 0, 0, 0 ],
      [ 1, 1, 1, 1, 1 ],
      [ 0, 1, 1, 1, 0 ]
    ],
    "generations": 27
  }
}
```


### solve a challenge

`POST https://api.noopschallenge.com/automatabot/challenges/EAKYU23GkZ9OSRHUsewo`


```
[
 [ 0, 0, 0, 0, 0 ],
 [ 0, 0, 0, 0, 0 ],
 [ 0, 1, 0, 0, 0 ],
 [ 0, 0, 0, 0, 0 ],
 [ 1, 1, 0, 0, 0 ]
]
```

`HTTP 200`

```
{ "result": "correct" }
```


### get all rules

`GET https://api.noopschallenge.com/automatabot/rules`

`HTTP 200`

```
[
  {
    "name": "2by",
    "survival": [ 1, 2, 5 ],
    "birth": [ 3, 6 ]
  },
  {
    "name": "34",
    "survival": [ 3, 4 ],
    "birth": [ 3, 4 ]
  },
...
]
```


### get a random automaton setup for a particular ruleset

`GET https://api.noopschallenge.com/automatabot/rules/conway/random`

`HTTP 200`

```
{
  "rules": {
    "name": "conway",
    "birth": [ 3 ],
    "survival": [ 2, 3 ]
  },
  "cells": [
    [ 0, 1, 0, 1, 0 ],
    [ 1, 0, 0, 0, 0 ],
    [ 0, 1, 1, 0, 0 ],
    [ 0, 1, 0, 0, 0 ],
    [ 1, 1, 0, 0, 0 ]
  ],
  "generations": 27
}
```

