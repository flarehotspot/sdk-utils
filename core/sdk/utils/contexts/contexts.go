package contexts

type sysacctctx string
type clientctx string

var (
	ClientCtxKey  clientctx  = "clnt"
	SysAcctCtxKey sysacctctx = "adminacct"
)
