package v2

var (
	// ConditionalOrderGenerationPrefix || market || subaccount -> current admission generation.
	ConditionalOrderGenerationPrefix = []byte{0xaa}
	// ConditionalOrderBirthPrefix || market || subaccount || hash -> immutable admission generation.
	ConditionalOrderBirthPrefix = []byte{0xab}
	// ConditionalGenerationOrdersPrefix || market || subaccount || generation || bucket || hash -> true.
	// Only reduce-only orders enter this index. Generation zero uses the frozen legacy indexes.
	ConditionalGenerationOrdersPrefix = []byte{0xac}
	// ConditionalInvalidatedGenerationPrefix || market || subaccount -> last invalidated generation.
	ConditionalInvalidatedGenerationPrefix = []byte{0xad}
	// ConditionalGenerationROCountPrefix || market || subaccount [|| generation] -> count.
	// The shorter key counts all versioned reduce-only orders, including deferred cancellations.
	ConditionalGenerationROCountPrefix = []byte{0xae}
)
