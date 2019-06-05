package sqlbuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestFuncAVG(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertClauseSerialize(t, AVGf(table1ColFloat), "AVG(table1.colFloat)")
	})

	t.Run("integer", func(t *testing.T) {
		assertClauseSerialize(t, AVGi(table1ColInt), "AVG(table1.colInt)")
	})
}

func TestFuncBIT_AND(t *testing.T) {
	assertClauseSerialize(t, BIT_AND(table1ColInt), "BIT_AND(table1.colInt)")
}

func TestFuncBIT_OR(t *testing.T) {
	assertClauseSerialize(t, BIT_OR(table1ColInt), "BIT_OR(table1.colInt)")
}

func TestFuncBOOL_AND(t *testing.T) {
	assertClauseSerialize(t, BOOL_AND(table1ColBool), "BOOL_AND(table1.colBool)")
}

func TestFuncBOOL_OR(t *testing.T) {
	assertClauseSerialize(t, BOOL_OR(table1ColBool), "BOOL_OR(table1.colBool)")
}

func TestFuncEVERY(t *testing.T) {
	assertClauseSerialize(t, EVERY(table1ColBool), "EVERY(table1.colBool)")
}

func TestFuncMIN(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertClauseSerialize(t, MINf(table1ColFloat), "MIN(table1.colFloat)")
	})

	t.Run("integer", func(t *testing.T) {
		assertClauseSerialize(t, MINi(table1ColInt), "MIN(table1.colInt)")
	})
}

func TestFuncMAX(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertClauseSerialize(t, MAXf(table1ColFloat), "MAX(table1.colFloat)")
		assertClauseSerialize(t, MAXf(Float(11.2222)), "MAX($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertClauseSerialize(t, MAXi(table1ColInt), "MAX(table1.colInt)")
		assertClauseSerialize(t, MAXi(Int(11)), "MAX($1)", int64(11))
	})
}

func TestFuncSUM(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertClauseSerialize(t, SUMf(table1ColFloat), "SUM(table1.colFloat)")
		assertClauseSerialize(t, SUMf(Float(11.2222)), "SUM($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertClauseSerialize(t, SUMi(table1ColInt), "SUM(table1.colInt)")
		assertClauseSerialize(t, SUMi(Int(11)), "SUM($1)", int64(11))
	})
}

func TestFuncCOUNT(t *testing.T) {
	assertClauseSerialize(t, COUNT(STAR), "COUNT(*)")
	assertClauseSerialize(t, COUNT(table1ColFloat), "COUNT(table1.colFloat)")
	assertClauseSerialize(t, COUNT(Float(11.2222)), "COUNT($1)", float64(11.2222))
}

func TestFuncABS(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertClauseSerialize(t, ABSf(table1ColFloat), "ABS(table1.colFloat)")
		assertClauseSerialize(t, ABSf(Float(11.2222)), "ABS($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertClauseSerialize(t, ABSi(table1ColInt), "ABS(table1.colInt)")
		assertClauseSerialize(t, ABSi(Int(11)), "ABS($1)", int64(11))
	})
}

func TestFuncSQRT(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertClauseSerialize(t, SQRTf(table1ColFloat), "SQRT(table1.colFloat)")
		assertClauseSerialize(t, SQRTf(Float(11.2222)), "SQRT($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertClauseSerialize(t, SQRTi(table1ColInt), "SQRT(table1.colInt)")
		assertClauseSerialize(t, SQRTi(Int(11)), "SQRT($1)", int64(11))
	})
}

func TestFuncCBRT(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertClauseSerialize(t, CBRTf(table1ColFloat), "CBRT(table1.colFloat)")
		assertClauseSerialize(t, CBRTf(Float(11.2222)), "CBRT($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertClauseSerialize(t, CBRTi(table1ColInt), "CBRT(table1.colInt)")
		assertClauseSerialize(t, CBRTi(Int(11)), "CBRT($1)", int64(11))
	})
}

func TestFuncCEIL(t *testing.T) {
	assertClauseSerialize(t, CEIL(table1ColFloat), "CEIL(table1.colFloat)")
	assertClauseSerialize(t, CEIL(Float(11.2222)), "CEIL($1)", float64(11.2222))
}

func TestFuncFLOOR(t *testing.T) {
	assertClauseSerialize(t, FLOOR(table1ColFloat), "FLOOR(table1.colFloat)")
	assertClauseSerialize(t, FLOOR(Float(11.2222)), "FLOOR($1)", float64(11.2222))
}

func TestFuncROUND(t *testing.T) {
	assertClauseSerialize(t, ROUND(table1ColFloat), "ROUND(table1.colFloat)")
	assertClauseSerialize(t, ROUND(Float(11.2222)), "ROUND($1)", float64(11.2222))

	assertClauseSerialize(t, ROUND(table1ColFloat, Int(2)), "ROUND(table1.colFloat, $1)", int64(2))
	assertClauseSerialize(t, ROUND(Float(11.2222), Int(1)), "ROUND($1, $2)", float64(11.2222), int64(1))
}

func TestFuncSIGN(t *testing.T) {
	assertClauseSerialize(t, SIGN(table1ColFloat), "SIGN(table1.colFloat)")
	assertClauseSerialize(t, SIGN(Float(11.2222)), "SIGN($1)", float64(11.2222))
}

func TestFuncTRUNC(t *testing.T) {
	assertClauseSerialize(t, TRUNC(table1ColFloat), "TRUNC(table1.colFloat)")
	assertClauseSerialize(t, TRUNC(Float(11.2222)), "TRUNC($1)", float64(11.2222))

	assertClauseSerialize(t, TRUNC(table1ColFloat, Int(2)), "TRUNC(table1.colFloat, $1)", int64(2))
	assertClauseSerialize(t, TRUNC(Float(11.2222), Int(1)), "TRUNC($1, $2)", float64(11.2222), int64(1))
}

func TestFuncLN(t *testing.T) {
	assertClauseSerialize(t, LN(table1ColFloat), "LN(table1.colFloat)")
	assertClauseSerialize(t, LN(Float(11.2222)), "LN($1)", float64(11.2222))
}

func TestFuncLOG(t *testing.T) {
	assertClauseSerialize(t, LOG(table1ColFloat), "LOG(table1.colFloat)")
	assertClauseSerialize(t, LOG(Float(11.2222)), "LOG($1)", float64(11.2222))
}

func TestFuncCOALESCE(t *testing.T) {
	assertClauseSerialize(t, COALESCE(table1ColFloat), "COALESCE(table1.colFloat)")
	assertClauseSerialize(t, COALESCE(Float(11.2222), NULL, String("str")), "COALESCE($1, NULL, $2)", float64(11.2222), "str")
}

func TestFuncNULLIF(t *testing.T) {
	assertClauseSerialize(t, NULLIF(table1ColFloat, table2ColInt), "NULLIF(table1.colFloat, table2.colInt)")
	assertClauseSerialize(t, NULLIF(Float(11.2222), NULL), "NULLIF($1, NULL)", float64(11.2222))
}

func TestFuncGREATEST(t *testing.T) {
	assertClauseSerialize(t, GREATEST(table1ColFloat), "GREATEST(table1.colFloat)")
	assertClauseSerialize(t, GREATEST(Float(11.2222), NULL, String("str")), "GREATEST($1, NULL, $2)", float64(11.2222), "str")
}

func TestFuncLEAST(t *testing.T) {
	assertClauseSerialize(t, LEAST(table1ColFloat), "LEAST(table1.colFloat)")
	assertClauseSerialize(t, LEAST(Float(11.2222), NULL, String("str")), "LEAST($1, NULL, $2)", float64(11.2222), "str")
}

func TestInterval(t *testing.T) {
	query := INTERVAL(`6 years 5 months 4 days 3 hours 2 minutes 1 second`)

	queryData := &queryData{}

	err := query.serialize(select_statement, queryData)

	assert.NilError(t, err)
	assert.Equal(t, queryData.buff.String(), `INTERVAL $1`)
}
