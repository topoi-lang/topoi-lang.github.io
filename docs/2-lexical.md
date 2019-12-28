## Identifier and Keywords

```
identifier           
  = letter (letter | "_" | "'" )* 
  | "_" identifier

qualified_identifier 
  = (identifier ".")+
```

When referencing something inside a module or a record, or a tuple, we use a qualified name (`qualified_identifier`).

### Comment and Documentation

```
prefix                 
  = letter (letter | "_" | "'" )* 
  | "_" prefix

comment
  = "--" (char)* newline

doc_comment 
  = "--|" (char)* newline

prefixed_doc_comment
  = "--|" prefix "|" (char)* newline

block_comment 
  = "--[" (char)* "]--"

prefixed_block_comment 
  = "--[" prefix "|" (char)* "]--" 
```

Content enclosed in documentation comment (`doc_comment`) and prefixed block comment (`prefixed_block_comment`) are conversed in AST and will be output or generated as web content.

The prefix is just a name that will let compiler pass to know what generator used to process the content inside the doc comment or block comment. eg.: LaTex, KaTeX, markdown...

## Primitive value and types

## Full Grammar

```
program 
  = {definition}+

definition 
  = constant_definition
  | type_definition

constant_definition
  = assignable [":" type] ":=" expression

assignable
  = identifier
  | destructurable

destructurable
  = "(" 
    {identifier [":" type] ["as" assignable]","}* 
    ["..." identifier] 
  ")"

expression
  = function_abstraction
  | function_application
  | identifier
  | record
  | adt_construction

function_abstraction
  = "{" {assignable "->" expression ","}+ "}"

function_application
  = expression "." expression

record
  = "(" {record_body ","}*")"

record_body
  = identifier [":" type] [":=" expression]
  | spread_expression

spread_expression
  = "..." expression

type
  = identifier
  | function_type
  | record_type
  | type_function_application

type_function_application
  = function_application

function_type
  = record_type "->" (record_type|function_type)

record_type
  = "(" {identifier [":" type]}* ")"

type_definition
  =  "type" identifier [":" type ":=" type_body]

type_body
  = {"|" constructor [record_type] ":" type}+

constructor
  = identifier
```