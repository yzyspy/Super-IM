package collection

func deduplicateUids(uids []uint64) []uint64 {
	uniqueUids := make(map[uint64]bool)
	for _, uid := range uids {
		uniqueUids[uid] = true
	}
	result := make([]uint64, 0, len(uniqueUids))
	for uid := range uniqueUids {
		result = append(result, uid)
	}
	return result
}
