package v2

import "github.com/ethereum/go-ethereum/common"

// ConditionalTriggerBucketCount is how many buckets a market's conditional orders fall into: the two
// order types crossed with the two trigger directions.
const ConditionalTriggerBucketCount = 4

// ConditionalTriggerBuckets holds one number per conditional order bucket of a market. It carries
// both the eligible counts a block's counting pass finds and the allowances the block's trigger
// budget grants for reading them.
type ConditionalTriggerBuckets struct {
	MarketHigher int
	MarketLower  int
	LimitHigher  int
	LimitLower   int
}

// UnboundedConditionalTriggerBuckets reads every eligible order of every bucket: the queries and
// governance driven paths that are not subject to the block's trigger budget.
func UnboundedConditionalTriggerBuckets() ConditionalTriggerBuckets {
	return ConditionalTriggerBuckets{MarketHigher: -1, MarketLower: -1, LimitHigher: -1, LimitLower: -1}
}

func (b ConditionalTriggerBuckets) Total() int {
	return b.MarketHigher + b.MarketLower + b.LimitHigher + b.LimitLower
}

// AsSlice lists the buckets in their canonical order, for budgeting them as interchangeable claims.
func (b ConditionalTriggerBuckets) AsSlice() []int {
	return []int{b.MarketHigher, b.MarketLower, b.LimitHigher, b.LimitLower}
}

// ConditionalTriggerBucketsFromSlice is the inverse of AsSlice.
func ConditionalTriggerBucketsFromSlice(values []int) ConditionalTriggerBuckets {
	return ConditionalTriggerBuckets{
		MarketHigher: values[0],
		MarketLower:  values[1],
		LimitHigher:  values[2],
		LimitLower:   values[3],
	}
}

// ConditionalBucket identifies one bucket of a market's conditional orders: an order type crossed
// with a trigger direction. Conditional orders are stored and walked one bucket at a time.
type ConditionalBucket struct {
	MarketOrders       bool
	TriggerPriceHigher bool
}

// ConditionalBucketsOf lists a market's four buckets in their canonical order.
func ConditionalBucketsOf() [ConditionalTriggerBucketCount]ConditionalBucket {
	return [ConditionalTriggerBucketCount]ConditionalBucket{
		{MarketOrders: true, TriggerPriceHigher: true},
		{MarketOrders: true, TriggerPriceHigher: false},
		{MarketOrders: false, TriggerPriceHigher: true},
		{MarketOrders: false, TriggerPriceHigher: false},
	}
}

// ConditionalBucketCursors carries the resume points of one conditional order type's two trigger
// directions. Each opaque payload carries its position, captured sweep endpoint and filtering state.
type ConditionalBucketCursors struct {
	Higher []byte
	Lower  []byte
}

// ConditionalTriggerCursors carries the resume points of all four buckets of one market.
type ConditionalTriggerCursors struct {
	MarketOrders ConditionalBucketCursors
	LimitOrders  ConditionalBucketCursors
}

// ConditionalBucketRead states how much of one order type's two buckets to read and from where.
// A negative limit reads every eligible order of that bucket and a zero limit reads none. Keep, when
// set, selects the orders worth reading: an order it rejects is neither returned nor counted against
// the limit, so orders the caller cannot act on do not consume its budget. ScanCeiling then bounds the
// walk however many are rejected; zero or less leaves it bounded only by the limits.
type ConditionalBucketRead struct {
	HigherLimit int
	LowerLimit  int
	ScanCeiling int
	Cursors     ConditionalBucketCursors
	Keep        func(subaccountID common.Hash) bool
}

// UnboundedConditionalBucketRead reads every eligible order of both buckets, from the start: the
// queries and governance driven paths, which are not subject to a block's trigger budget.
func UnboundedConditionalBucketRead() ConditionalBucketRead {
	return ConditionalBucketRead{HigherLimit: -1, LowerLimit: -1}
}

// WalkIsDone reports whether a bucket walk has read all it was asked for: the limit's worth of kept
// orders, or the ceiling's worth of entries whether kept or rejected.
func (r ConditionalBucketRead) WalkIsDone(kept, scanned, limit int) bool {
	if limit > 0 && kept >= limit {
		return true
	}

	return r.ScanCeiling > 0 && scanned >= r.ScanCeiling
}

// ConditionalInvalidationInlineBudget is how much of a block's conditional invalidation budget the
// callers that run inline may spend: everything but the reserve held for the deferred backlog.
func ConditionalInvalidationInlineBudget() int {
	return MaxConditionalOrderInvalidationCancelsPerBlock - ConditionalOrderInvalidationPendingReserve
}

const (
	// MaxConditionalOrderTriggersPerMarketDirection bounds how far a block walks one market's
	// conditional orders of a single order type and trigger direction. It caps the eligible count a
	// market can claim when the block budget is shared out, and with it the scan a saturated market
	// costs; the counting pass decodes nothing it walks.
	MaxConditionalOrderTriggersPerMarketDirection = 1000

	// MaxConditionalOrderDiscoveryScansPerBlock bounds that walking summed over every market and
	// bucket. The number of markets is not bounded by the protocol, so a per-market cap alone leaves
	// the block's discovery cost growing with the market count; each market's share of this is what
	// it may actually walk.
	MaxConditionalOrderDiscoveryScansPerBlock = 20000

	// MaxConditionalOrderTriggersPerBlock bounds the conditional orders selected for triggering
	// across every market. Filtering paused accounts may decode additional orders, subject to the
	// separate discovery scan ceiling.
	// Orders past the bound are left untouched in state and are picked up by a later block for as
	// long as their trigger condition still holds.
	MaxConditionalOrderTriggersPerBlock = 5000

	// MaxConditionalOrderInvalidationCancelsPerBlock bounds how many conditional orders a block
	// cancels while invalidating subaccounts left without locked margin, across both the matching
	// path and the EndBlocker stage. Failed attempts also consume this budget; unfinished
	// subaccounts retain their queue position for the next block.
	MaxConditionalOrderInvalidationCancelsPerBlock = 2000

	// MaxConditionalOrderInvalidationVisitsPerBlock bounds how many durably deferred invalidations a
	// block looks at. Visiting costs a read and a revalidation whether or not anything is cancelled,
	// so the cancellation budget alone would leave a large backlog of entries that cancel nothing
	// setting the block's cost.
	MaxConditionalOrderInvalidationVisitsPerBlock = 500

	// ConditionalOrderInvalidationPendingReserve is the part of the cancellation budget only the
	// drain of durably deferred invalidations may spend. Matching runs before the drain, so without a
	// reserve a block busy closing positions spends the whole budget before the oldest deferred
	// subaccount is ever looked at, and the backlog never clears.
	ConditionalOrderInvalidationPendingReserve = 500
)

var (
	// PendingConditionalOrderInvalidationPrefix || sequence(8B) -> marketID || subaccountID || direction || generation(8B)
	// carries the conditional-order invalidations a block could not finish within its budget. The
	// working flags live in the transient store and are gone at commit, so a deferred invalidation
	// has to be durable or it is simply lost. Keying by an increasing sequence makes the drain
	// first-in-first-out: a deferred subaccount cannot be held behind later ones forever, which
	// keying by market and subaccount would allow.
	PendingConditionalOrderInvalidationPrefix = []byte{0xa4}
	// PendingConditionalOrderInvalidationSeqKey -> uint64, the sequence the next deferred
	// invalidation is filed under.
	PendingConditionalOrderInvalidationSeqKey = []byte{0xa5}
	// ConditionalTriggerCursorPrefix || marketID || isMarketOrders || direction -> opaque sweep state.
	// Records how far a bounded block walked one conditional order bucket. Orders the block reads
	// but cannot act on stay in state at the crossed end of the range, where the next block would
	// meet them again, so the walk resumes past them and wraps when it runs out.
	ConditionalTriggerCursorPrefix = []byte{0xa6}
	// TransientConditionalInvalidationBudgetKey -> uint64, the conditional cancellations already attempted
	// this block while invalidating subaccounts left without locked margin. Both the matching path
	// and the EndBlocker stage spend from it, so the block's total is bounded whichever runs first.
	TransientConditionalInvalidationBudgetKey = []byte{0xa7}
)

// ConditionalInvalidationMode identifies which part of the shared budget a caller may use.
type ConditionalInvalidationMode uint8

const (
	ConditionalInvalidationInline ConditionalInvalidationMode = iota
	ConditionalInvalidationPending
)

// ConditionalInvalidationStatus distinguishes finished work from work that must remain queued.
// Failed is the zero value so panic recovery cannot accidentally retire unfinished work.
type ConditionalInvalidationStatus uint8

const (
	ConditionalInvalidationFailed ConditionalInvalidationStatus = iota
	ConditionalInvalidationCompleted
	ConditionalInvalidationDeferred
)
