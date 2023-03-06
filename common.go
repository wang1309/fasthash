package hashtable

import "unsafe"

const (
	kInitialCellCntBits = 10
	kInitialCellCnt     = 1 << kInitialCellCntBits

	kLoadFactorNumerator   = 1
	kLoadFactorDenominator = 2

	//kTwoLevelBucketCntBits = 8
	//kTwoLevelBucketCnt     = 1 << kTwoLevelBucketCntBits
	//kMaxTwoLevelBucketCnt  = kTwoLevelBucketCnt - 1
)

type Aggregator interface {
	StateSize() uint8
	ResultSize() uint8
	NeedsInit() bool

	Init(state unsafe.Pointer)
	AddBatch(states []unsafe.Pointer, values unsafe.Pointer)
	MergeBatch(lstates, rstates []unsafe.Pointer)

	Finalize(states, results []unsafe.Pointer)
}
