package fasthash

import (
	"testing"

	"github.com/matrixorigin/matrixone/pkg/common/mpool"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/stretchr/testify/require"
)

// #cgo CFLAGS: -DPNG_DEBUG=1 -I./cgo
// #cgo LDFLAGS: -L/Users/tal/go_project/fasthash/cgo -lmo -lm
func TestIntHashMap_Iterator(t *testing.T) {
	{
		m := mpool.MustNewZero()
		mp, err := NewIntHashMap(false, 0, 0, m)
		require.NoError(t, err)
		rowCount := 10
		vecs := []*vector.Vector{
			newVector(rowCount, types.T_int32.ToType(), m, false, []int32{
				-1, -1, -1, 2, 2, 2, 3, 3, 3, 4,
			}),
			newVector(rowCount, types.T_uint32.ToType(), m, false, []uint32{
				1, 1, 1, 2, 2, 2, 3, 3, 3, 4,
			}),
		}
		itr := mp.NewIterator()
		vs, _, err := itr.Insert(0, rowCount, vecs)
		require.NoError(t, err)
		require.Equal(t, []uint64{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}, vs)
		vs, _ = itr.Find(0, rowCount, vecs, nil)
		require.Equal(t, []uint64{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}, vs)
		for _, vec := range vecs {
			vec.Free(m)
		}
		mp.Free()
		require.Equal(t, int64(0), m.Stats().NumCurrBytes.Load())
	}
	{
		m := mpool.MustNewZero()
		mp, err := NewIntHashMap(true, 0, 0, m)
		require.NoError(t, err)
		ts := []types.Type{
			types.New(types.T_int8, 0, 0, 0),
			types.New(types.T_int16, 0, 0, 0),
		}
		vecs := newVectorsWithNull(ts, false, Rows, m)
		itr := mp.NewIterator()
		vs, _, err := itr.Insert(0, Rows, vecs)
		require.NoError(t, err)
		require.Equal(t, []uint64{1, 2, 1, 3, 1, 4, 1, 5, 1, 6}, vs[:Rows])
		vs, _ = itr.Find(0, Rows, vecs, nil)
		require.Equal(t, []uint64{1, 2, 1, 3, 1, 4, 1, 5, 1, 6}, vs[:Rows])
		for _, vec := range vecs {
			vec.Free(m)
		}
		mp.Free()
		require.Equal(t, int64(0), m.Stats().NumCurrBytes.Load())
	}
	{
		m := mpool.MustNewZero()
		mp, err := NewIntHashMap(true, 0, 0, m)
		require.NoError(t, err)
		ts := []types.Type{
			types.New(types.T_int64, 0, 0, 0),
		}
		vecs := newVectorsWithNull(ts, false, Rows, m)
		itr := mp.NewIterator()
		vs, _, err := itr.Insert(0, Rows, vecs)
		require.NoError(t, err)
		require.Equal(t, []uint64{1, 2, 1, 3, 1, 4, 1, 5, 1, 6}, vs[:Rows])
		vs, _ = itr.Find(0, Rows, vecs, nil)
		require.Equal(t, []uint64{1, 2, 1, 3, 1, 4, 1, 5, 1, 6}, vs[:Rows])
		for _, vec := range vecs {
			vec.Free(m)
		}
		mp.Free()
		require.Equal(t, int64(0), m.Stats().NumCurrBytes.Load())
	}
	{
		m := mpool.MustNewZero()
		mp, err := NewIntHashMap(true, 0, 0, m)
		require.NoError(t, err)
		ts := []types.Type{
			types.New(types.T_char, 1, 0, 0),
		}
		vecs := newVectorsWithNull(ts, false, Rows, m)
		itr := mp.NewIterator()
		vs, _, err := itr.Insert(0, Rows, vecs)
		require.NoError(t, err)
		require.Equal(t, []uint64{1, 2, 1, 3, 1, 4, 1, 5, 1, 6}, vs[:Rows])
		vs, _ = itr.Find(0, Rows, vecs, nil)
		require.Equal(t, []uint64{1, 2, 1, 3, 1, 4, 1, 5, 1, 6}, vs[:Rows])
		for _, vec := range vecs {
			vec.Free(m)
		}
		mp.Free()
		require.Equal(t, int64(0), m.Stats().NumCurrBytes.Load())
	}
	{
		m := mpool.MustNewZero()
		mp, err := NewIntHashMap(true, 0, 0, m)
		require.NoError(t, err)
		ts := []types.Type{
			types.New(types.T_char, 1, 0, 0),
		}
		vecs := newVectors(ts, false, Rows, m)
		itr := mp.NewIterator()
		vs, _, err := itr.Insert(0, Rows, vecs)
		require.NoError(t, err)
		require.Equal(t, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, vs[:Rows])
		vs, _ = itr.Find(0, Rows, vecs, nil)
		require.Equal(t, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, vs[:Rows])
		for _, vec := range vecs {
			vec.Free(m)
		}
		mp.Free()
		require.Equal(t, int64(0), m.Stats().NumCurrBytes.Load())
	}
	{
		m := mpool.MustNewZero()
		mp, err := NewIntHashMap(false, 0, 0, m)
		require.NoError(t, err)
		ts := []types.Type{
			types.New(types.T_char, 1, 0, 0),
		}
		vecs := newVectorsWithNull(ts, false, Rows, m)
		itr := mp.NewIterator()
		vs, _, err := itr.Insert(0, Rows, vecs)
		require.NoError(t, err)
		require.Equal(t, []uint64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}, vs[:Rows])
		vs, _ = itr.Find(0, Rows, vecs, nil)
		require.Equal(t, []uint64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}, vs[:Rows])
		for _, vec := range vecs {
			vec.Free(m)
		}
		mp.Free()
		require.Equal(t, int64(0), m.Stats().NumCurrBytes.Load())
	}
}
