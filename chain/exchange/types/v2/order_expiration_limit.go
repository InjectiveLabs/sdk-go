package v2

// MaxOrdersPerExpirationBlock bounds new GTB admissions across all markets,
// sides and addresses. Existing and imported orders above the limit are retained.
// Changing this protocol limit requires a chain upgrade.
const MaxOrdersPerExpirationBlock uint64 = 1_000
