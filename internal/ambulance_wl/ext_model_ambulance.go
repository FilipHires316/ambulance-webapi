package ambulance_wl

import (
	"time"

	"slices"
)

func (a *Ambulance) reconcileWaitingList() {
	slices.SortFunc(a.WaitingList, func(left, right WaitingListEntry) int {
		if left.WaitingSince.Before(right.WaitingSince) {
			return -1
		} else if left.WaitingSince.After(right.WaitingSince) {
			return 1
		} else {
			return 0
		}
	})

	if len(a.WaitingList) == 0 {
		return
	}


	if a.WaitingList[0].EstimatedStart.Before(a.WaitingList[0].WaitingSince) {
		a.WaitingList[0].EstimatedStart = a.WaitingList[0].WaitingSince
	}

	if a.WaitingList[0].EstimatedStart.Before(time.Now()) {
		a.WaitingList[0].EstimatedStart = time.Now()
	}

	nextEntryStart :=
		a.WaitingList[0].EstimatedStart.
			Add(time.Duration(a.WaitingList[0].EstimatedDurationMinutes) * time.Minute)
	for i := range a.WaitingList[1:] {
		idx := i + 1
		if a.WaitingList[idx].EstimatedStart.Before(nextEntryStart) {
			a.WaitingList[idx].EstimatedStart = nextEntryStart
		}
		if a.WaitingList[idx].EstimatedStart.Before(a.WaitingList[idx].WaitingSince) {
			a.WaitingList[idx].EstimatedStart = a.WaitingList[idx].WaitingSince
		}

		nextEntryStart =
			a.WaitingList[idx].EstimatedStart.
				Add(time.Duration(a.WaitingList[idx].EstimatedDurationMinutes) * time.Minute)
	}
}
