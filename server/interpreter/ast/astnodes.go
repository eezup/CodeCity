// AUTO-GENERATED BY md2go.  DO NOT EDIT.

package ast

/********************************************************************/

// ESTree AST nodes are represented as Node objects, which may have
// any prototype inheritance but which implement the following
// interface:
//
// The type field is a string representing the AST variant type. Each
// subtype of Node is documented below with the specific string of
// its type field. You can use this field to determine which
// interface a node implements.
//
// The loc field represents the source location information of the
// node. If the node contains no information about the source
// location, the field is null; otherwise it is an object consisting
// of a start position (the position of the first character of the
// parsed source region) and an end position (the position of the
// first character after the parsed source region):
type nodeStuff struct {
	Type  string          `json:"type,omitempty"`
	Loc   *SourceLocation `json:"loc,omitempty"`
	Start int             `json:"start"`
	End   int             `json:"end"`
}

/********************************************************************/

// Each Position object consists of a line number (1-indexed) and a
// column number (0-indexed):
type SourceLocation struct {
	Source string    `json:"source,omitempty"`
	Start  *Position `json:"start,omitempty"`
	End    *Position `json:"end,omitempty"`
}

/********************************************************************/

type Position struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

/********************************************************************/

// Identifier is an identifier. Note that an identifier may be an
// expression or a destructuring pattern.
type Identifier struct {
	expressionStuff
	Name string `json:"name,omitempty"`
}

/********************************************************************/

// Literal is a literal token. Note that a literal can be an
// expression.
type Literal struct {
	expressionStuff
	Raw string `json:"raw,omitempty"`
}

/********************************************************************/

// Program is a complete program source tree.
type Program struct {
	nodeStuff
	Body *Statements `json:"body,omitempty"`
}

/********************************************************************/

// functionStuff is a function [declaration](#functiondeclaration) or
// [expression](#functionexpression).
type functionStuff struct {
	Id     *Identifier     `json:"id,omitempty"`
	Params []*Identifier   `json:"params,omitempty"`
	Body   *BlockStatement `json:"body,omitempty"`
}

/********************************************************************/

// Any statement.
type statementStuff struct {
	nodeStuff
}

/********************************************************************/

// ExpressionStatement is an expression statement, i.e., a statement
// consisting of a single expression.
type ExpressionStatement struct {
	statementStuff
	Expression *Expression `json:"expression,omitempty"`
}

/********************************************************************/

// BlockStatement is a block statement, i.e., a sequence of
// statements surrounded by braces.
type BlockStatement struct {
	statementStuff
	Body *Statements `json:"body,omitempty"`
}

/********************************************************************/

// EmptyStatement is an empty statement, i.e., a solitary semicolon.
type EmptyStatement struct {
	statementStuff
}

/********************************************************************/

// DebuggerStatement is a debugger statement.
type DebuggerStatement struct {
	statementStuff
}

/********************************************************************/

// WithStatement is a with statement.
type WithStatement struct {
	statementStuff
	Object *Expression `json:"object,omitempty"`
	Body   *Statement  `json:"body,omitempty"`
}

/********************************************************************/

// ReturnStatement is a return statement.
type ReturnStatement struct {
	statementStuff
	Argument *Expression `json:"argument,omitempty"`
}

/********************************************************************/

// LabeledStatement is a labeled statement, i.e., a statement
// prefixed by a break/continue label.
type LabeledStatement struct {
	statementStuff
	Label *Identifier `json:"label,omitempty"`
	Body  *Statement  `json:"body,omitempty"`
}

/********************************************************************/

// BreakStatement is a break statement.
type BreakStatement struct {
	statementStuff
	Label *Identifier `json:"label,omitempty"`
}

/********************************************************************/

// ContinueStatement is a continue statement.
type ContinueStatement struct {
	statementStuff
	Label *Identifier `json:"label,omitempty"`
}

/********************************************************************/

// IfStatement is an if statement.
type IfStatement struct {
	statementStuff
	Test       *Expression `json:"test,omitempty"`
	Consequent *Statement  `json:"consequent,omitempty"`
	Alternate  *Statement  `json:"alternate,omitempty"`
}

/********************************************************************/

// SwitchStatement is a switch statement.
type SwitchStatement struct {
	statementStuff
	Discriminant *Expression   `json:"discriminant,omitempty"`
	Cases        []*SwitchCase `json:"cases,omitempty"`
}

/********************************************************************/

// SwitchCase is a case (if test is an Expression) or default (if
// test === null) clause in the body of a switch statement.
type SwitchCase struct {
	nodeStuff
	Test       *Expression `json:"test,omitempty"`
	Consequent *Statements `json:"consequent,omitempty"`
}

/********************************************************************/

// ThrowStatement is a throw statement.
type ThrowStatement struct {
	statementStuff
	Argument *Expression `json:"argument,omitempty"`
}

/********************************************************************/

// TryStatement is a try statement. If handler is null then finalizer
// must be a BlockStatement.
type TryStatement struct {
	statementStuff
	Block     *BlockStatement `json:"block,omitempty"`
	Handler   *CatchClause    `json:"handler,omitempty"`
	Finalizer *BlockStatement `json:"finalizer,omitempty"`
}

/********************************************************************/

// CatchClause is a catch clause following a try block.
type CatchClause struct {
	nodeStuff
	Param *Identifier     `json:"param,omitempty"`
	Body  *BlockStatement `json:"body,omitempty"`
}

/********************************************************************/

// WhileStatement is a while statement.
type WhileStatement struct {
	statementStuff
	Test *Expression `json:"test,omitempty"`
	Body *Statement  `json:"body,omitempty"`
}

/********************************************************************/

// DoWhileStatement is a do/while statement.
type DoWhileStatement struct {
	statementStuff
	Body *Statement  `json:"body,omitempty"`
	Test *Expression `json:"test,omitempty"`
}

/********************************************************************/

// ForStatement is a for statement.
type ForStatement struct {
	statementStuff
	Init   *DeclOrID   `json:"init,omitempty"`
	Test   *Expression `json:"test,omitempty"`
	Update *Expression `json:"update,omitempty"`
	Body   *Statement  `json:"body,omitempty"`
}

/********************************************************************/

// ForInStatement is a for/in statement.
type ForInStatement struct {
	statementStuff
	Left  *DeclOrID   `json:"left,omitempty"`
	Right *Expression `json:"right,omitempty"`
	Body  *Statement  `json:"body,omitempty"`
}

/********************************************************************/

// FunctionDeclaration is a function declaration. Note that unlike in
// the parent interface Function, the id cannot be null.
type FunctionDeclaration struct {
	statementStuff
	functionStuff
	Id *Identifier `json:"id,omitempty"`
}

/********************************************************************/

// VariableDeclaration is a variable declaration.
type VariableDeclaration struct {
	statementStuff
	Declarations []*VariableDeclarator `json:"declarations,omitempty"`
	Kind         string                `json:"kind,omitempty"`
}

/********************************************************************/

// VariableDeclarator is a variable declarator.
type VariableDeclarator struct {
	nodeStuff
	Id   *Identifier `json:"id,omitempty"`
	Init *Expression `json:"init,omitempty"`
}

/********************************************************************/

// Any expression node. Since the left-hand side of an assignment may
// be any expression in general, an expression can also be a pattern.
type expressionStuff struct {
	nodeStuff
}

/********************************************************************/

// ThisExpression is a this expression.
type ThisExpression struct {
	expressionStuff
}

/********************************************************************/

// ArrayExpression is an array expression.
type ArrayExpression struct {
	expressionStuff
	Elements *Expressions `json:"elements,omitempty"`
}

/********************************************************************/

// ObjectExpression is an object expression.
type ObjectExpression struct {
	expressionStuff
	Properties []*Property `json:"properties,omitempty"`
}

/********************************************************************/

// Property is a literal property in an object expression can have
// either a string or number as its value. Ordinary property
// initializers have a kind value "init"; getters and setters have
// the kind values "get" and "set", respectively.
type Property struct {
	nodeStuff
	Key   *LitOrID    `json:"key,omitempty"`
	Value *Expression `json:"value,omitempty"`
}

/********************************************************************/

// FunctionExpression is a function expression.
type FunctionExpression struct {
	functionStuff
	expressionStuff
}

/********************************************************************/

// UnaryExpression is a unary operator expression.
type UnaryExpression struct {
	expressionStuff
	Operator *string     `json:"operator,omitempty"`
	Prefix   bool        `json:"prefix"`
	Argument *Expression `json:"argument,omitempty"`
}

/********************************************************************/

//  is a unary operator token.
//
// UpdateExpression is an update (increment or decrement) operator
// expression.
type UpdateExpression struct {
	expressionStuff
	Operator *string     `json:"operator,omitempty"`
	Argument *Expression `json:"argument,omitempty"`
	Prefix   bool        `json:"prefix"`
}

/********************************************************************/

//  is an update (increment or decrement) operator token.
//
// BinaryExpression is a binary operator expression.
type BinaryExpression struct {
	expressionStuff
	Operator *string     `json:"operator,omitempty"`
	Left     *Expression `json:"left,omitempty"`
	Right    *Expression `json:"right,omitempty"`
}

/********************************************************************/

//  is a binary operator token.
//
// AssignmentExpression is an assignment operator expression.
type AssignmentExpression struct {
	expressionStuff
	Operator *string     `json:"operator,omitempty"`
	Left     *Expression `json:"left,omitempty"`
	Right    *Expression `json:"right,omitempty"`
}

/********************************************************************/

//  is an assignment operator token.
//
// LogicalExpression is a logical operator expression.
type LogicalExpression struct {
	expressionStuff
	Operator *string     `json:"operator,omitempty"`
	Left     *Expression `json:"left,omitempty"`
	Right    *Expression `json:"right,omitempty"`
}

/********************************************************************/

//  is a logical operator token.
//
// MemberExpression is a member expression. If computed is true, the
// node corresponds to a computed (a[b]) member expression and
// property is an Expression. If computed is false, the node
// corresponds to a static (a.b) member expression and property is an
// Identifier.
type MemberExpression struct {
	expressionStuff
	Object   *Expression `json:"object,omitempty"`
	Property *Expression `json:"property,omitempty"`
	Computed bool        `json:"computed"`
}

/********************************************************************/

// ConditionalExpression is a conditional expression, i.e., a ternary
// ?/: expression.
type ConditionalExpression struct {
	expressionStuff
	Test       *Expression `json:"test,omitempty"`
	Alternate  *Expression `json:"alternate,omitempty"`
	Consequent *Expression `json:"consequent,omitempty"`
}

/********************************************************************/

// CallExpression is a function or method call expression.
type CallExpression struct {
	expressionStuff
	Callee    *Expression  `json:"callee,omitempty"`
	Arguments *Expressions `json:"arguments,omitempty"`
}

/********************************************************************/

// NewExpression is a new expression.
type NewExpression struct {
	expressionStuff
	Callee    *Expression  `json:"callee,omitempty"`
	Arguments *Expressions `json:"arguments,omitempty"`
}

/********************************************************************/

// SequenceExpression is a sequence expression, i.e., a
// comma-separated sequence of expressions.
type SequenceExpression struct {
	expressionStuff
	Expressions *Expressions `json:"expressions,omitempty"`
}

/********************************************************************/

var nodeTypes = map[string]node{
	"Program":            (*Program)(nil),
	"SwitchCase":         (*SwitchCase)(nil),
	"CatchClause":        (*CatchClause)(nil),
	"VariableDeclarator": (*VariableDeclarator)(nil),
	"Property":           (*Property)(nil),
}

var statementTypes = map[string]statement{
	"ExpressionStatement": (*ExpressionStatement)(nil),
	"BlockStatement":      (*BlockStatement)(nil),
	"EmptyStatement":      (*EmptyStatement)(nil),
	"DebuggerStatement":   (*DebuggerStatement)(nil),
	"WithStatement":       (*WithStatement)(nil),
	"ReturnStatement":     (*ReturnStatement)(nil),
	"LabeledStatement":    (*LabeledStatement)(nil),
	"BreakStatement":      (*BreakStatement)(nil),
	"ContinueStatement":   (*ContinueStatement)(nil),
	"IfStatement":         (*IfStatement)(nil),
	"SwitchStatement":     (*SwitchStatement)(nil),
	"ThrowStatement":      (*ThrowStatement)(nil),
	"TryStatement":        (*TryStatement)(nil),
	"WhileStatement":      (*WhileStatement)(nil),
	"DoWhileStatement":    (*DoWhileStatement)(nil),
	"ForStatement":        (*ForStatement)(nil),
	"ForInStatement":      (*ForInStatement)(nil),
	"FunctionDeclaration": (*FunctionDeclaration)(nil),
	"VariableDeclaration": (*VariableDeclaration)(nil),
}

var expressionTypes = map[string]expression{
	"Identifier":            (*Identifier)(nil),
	"Literal":               (*Literal)(nil),
	"ThisExpression":        (*ThisExpression)(nil),
	"ArrayExpression":       (*ArrayExpression)(nil),
	"ObjectExpression":      (*ObjectExpression)(nil),
	"FunctionExpression":    (*FunctionExpression)(nil),
	"UnaryExpression":       (*UnaryExpression)(nil),
	"UpdateExpression":      (*UpdateExpression)(nil),
	"BinaryExpression":      (*BinaryExpression)(nil),
	"AssignmentExpression":  (*AssignmentExpression)(nil),
	"LogicalExpression":     (*LogicalExpression)(nil),
	"MemberExpression":      (*MemberExpression)(nil),
	"ConditionalExpression": (*ConditionalExpression)(nil),
	"CallExpression":        (*CallExpression)(nil),
	"NewExpression":         (*NewExpression)(nil),
	"SequenceExpression":    (*SequenceExpression)(nil),
}

// $VAR1 = {
//           'Expression' => [
//                             'Identifier',
//                             'Literal',
//                             'ThisExpression',
//                             'ArrayExpression',
//                             'ObjectExpression',
//                             'FunctionExpression',
//                             'UnaryExpression',
//                             'UpdateExpression',
//                             'BinaryExpression',
//                             'AssignmentExpression',
//                             'LogicalExpression',
//                             'MemberExpression',
//                             'ConditionalExpression',
//                             'CallExpression',
//                             'NewExpression',
//                             'SequenceExpression'
//                           ],
//           'Function' => [
//                           'FunctionDeclaration',
//                           'FunctionExpression'
//                         ],
//           'Literal' => [
//                          'RegExpLiteral'
//                        ],
//           'Node' => [
//                       'Program',
//                       'Statement',
//                       'SwitchCase',
//                       'CatchClause',
//                       'VariableDeclarator',
//                       'Expression',
//                       'Property',
//                       'Pattern'
//                     ],
//           'Statement' => [
//                            'ExpressionStatement',
//                            'BlockStatement',
//                            'EmptyStatement',
//                            'DebuggerStatement',
//                            'WithStatement',
//                            'ReturnStatement',
//                            'LabeledStatement',
//                            'BreakStatement',
//                            'ContinueStatement',
//                            'IfStatement',
//                            'SwitchStatement',
//                            'ThrowStatement',
//                            'TryStatement',
//                            'WhileStatement',
//                            'DoWhileStatement',
//                            'ForStatement',
//                            'ForInStatement',
//                            'Declaration',
//                            'FunctionDeclaration',
//                            'VariableDeclaration'
//                          ]
//         };