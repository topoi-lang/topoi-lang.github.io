## Identifier and Keywords

```
identifier           ::= letter (letter | "_" | "'" )* | "_" identifier
qualified_identifier ::= (identifier ".")+
```

When referencing something inside a module or a record, or a tuple, we use a qualified name (`qualified_identifier`).

### Comment and Documentation

```
prefix                 ::= letter (letter | "_" | "'" )* | "_" prefix

comment                ::= "--" (char)* newline
doc_comment            ::= "--|" (char)* newline
prefixed_doc_comment   ::= "--|" prefix "|" (char)* newline
block_comment          ::= "--[" (char)* "]--"
prefixed_block_comment ::= "--[" prefix "|" (char)* "]--" 
```

Content enclosed in documentation comment (`doc_comment`) and prefixed block comment (`prefixed_block_comment`) are conversed in AST and will be output or generated as web content.

The prefix is just a name that will let compiler pass to know what generator used to process the content inside the doc comment or block comment. eg.: LaTex, KaTeX, markdown...

## Primitive value and types
