// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/sqlc-dev/sqlc/internal/sql/ast"
	"github.com/sqlc-dev/sqlc/internal/sql/catalog"
)

var funcsIntarray = []*catalog.Function{
	{
		Name: "_int_contained",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_int_contains",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_int_different",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_int_inter",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "_int_overlap",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_int_same",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_int_union",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "_intbig_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "intbig_gkey"},
	},
	{
		Name: "_intbig_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "intbig_gkey"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "boolop",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "query_int"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "bqarr_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "query_int"},
	},
	{
		Name: "bqarr_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "query_int"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "icount",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "idx",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "intarray_del_elem",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "intarray_push_array",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "intarray_push_elem",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "intset",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "intset_subtract",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "intset_union_elem",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "querytree",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "query_int"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "rboolop",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "query_int"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "sort",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "sort",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "sort_asc",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "sort_desc",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "subarray",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "subarray",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "uniq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
}

func Intarray() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsIntarray
	return s
}
