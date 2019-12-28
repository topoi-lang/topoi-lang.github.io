# Foreword

This document is intended as a reference manual for the Topoi language.
It lists the language constructs, and gives their precise syntax and informal semantics.
It is by no means a tutorial introduction to the language.

> In future we will make a proper landing page for the language.
But right now it serves as our internal team's draft.

## Notations

The syntax of the languages is given in BNF-like notation.

Where square brackets `[...]` denote optional components, curly brackets `{...}` denotes zero, one or several repetitions of the encolesd components.

Curly brackets with a trailing asterisk sign `{...}*` denote one or several repetitions of the encolosed components.

Parentheses `(...)` denote grouping of components.

|Notation|Repetition|
|-|-|
|`{...}?`|0 or 1|
|`{...}*`|0 or more|
|`{...}+`|1 or more|
