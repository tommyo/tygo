package emit

// emit directive on a string literal emits that value.
//
//tygo:emit
var _ = `export type OtherStructAsTuple=[
  a:number, 
  b:number, 
  c:string,
]
`

//tygo:emit This has no effect, only strings.
var _ = 12

// a non-string var is ignored. A var with no comment is ignored.

var foo = " "

// CustomMarshalled illustrates getting tygo to emit literal text
// This solves the problem of a struct field being marshalled into a tuple.
//
// emit directive on a struct emits the remainder of the directive line
//
//tygo:emit export type StructAsTuple=[a:number, b:number, c:string]
type CustomMarshalled struct {
	Content []StructAsTuple `json:"content"`
}
