package libplanet

import "strconv"

type RuneSlotInfo struct {
	SlotIndex int64 `abi:"slotIndex"`
	RuneId    int64 `abi:"runeId"`
}

func extractRuneSlotInfo(value []any) RuneSlotInfo {
	slotIndexValue, _ := strconv.Atoi(value[0].(string))
	runeIdValue, _ := strconv.Atoi(value[1].(string))
	return RuneSlotInfo{
		SlotIndex: int64(slotIndexValue),
		RuneId:    int64(runeIdValue),
	}
}
