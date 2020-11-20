// Code generated from Z:/higo-express/test\FnucExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package fnucExpr // FnucExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FnucExprListener is a complete listener for a parse tree produced by FnucExprParser.
type FnucExprListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterFuncCall is called when entering the FuncCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncArgs is called when entering the FuncArgs production.
	EnterFuncArgs(c *FuncArgsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitFuncCall is called when exiting the FuncCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncArgs is called when exiting the FuncArgs production.
	ExitFuncArgs(c *FuncArgsContext)
}
