// Code generated from Z:/higo-express/test\FnucExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package fnucExpr // FnucExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFnucExprListener is a complete listener for a parse tree produced by FnucExprParser.
type BaseFnucExprListener struct{}

var _ FnucExprListener = &BaseFnucExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFnucExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFnucExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFnucExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFnucExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseFnucExprListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseFnucExprListener) ExitStart(ctx *StartContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseFnucExprListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseFnucExprListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterFuncCall is called when production FuncCall is entered.
func (s *BaseFnucExprListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production FuncCall is exited.
func (s *BaseFnucExprListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncArgs is called when production FuncArgs is entered.
func (s *BaseFnucExprListener) EnterFuncArgs(ctx *FuncArgsContext) {}

// ExitFuncArgs is called when production FuncArgs is exited.
func (s *BaseFnucExprListener) ExitFuncArgs(ctx *FuncArgsContext) {}
