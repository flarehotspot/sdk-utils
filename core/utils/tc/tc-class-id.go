package tc

import (
	"fmt"
	"sort"

	jobque "github.com/flarehotspot/flarehotspot/core/utils/job-que"
)

const (
	TcClassIdRoot TcClassId = iota
	TcClassIdDefault
	TcClassIdUser
	startId
)

var (
	usedClassIds []TcClassId
	tmpClassIds  []TcClassId
	q            *jobque.JobQues = jobque.NewJobQues()
)

type TcClassId uint

func (self TcClassId) String() string {
	return fmt.Sprintf("1:%x", int(self))
}

func (self TcClassId) Uint() uint {
	return uint(self)
}

func (self TcClassId) Cancel() {
	q.Exec(func() (interface{}, error) {
		removeClassId(tmpClassIds, self)
		return nil, nil
	})
}

func (self TcClassId) Commit() {
	q.Exec(func() (interface{}, error) {
		removeClassId(tmpClassIds, self)
		usedClassIds = append(usedClassIds, self)
		return nil, nil
	})
}

func (self TcClassId) Return() {
	q.Exec(func() (interface{}, error) {
		removeClassId(usedClassIds, self)
		return nil, nil
	})
}

func GetAvailableId() TcClassId {
	sym, _ := q.Exec(func() (interface{}, error) {
		classids := orderedIds()

		for i := 0; i < len(classids); i++ {
			expected := (i * 2) + int(startId)
			if classids[i] != TcClassId(expected) {
				return TcClassId(expected), nil
			}
		}

		return TcClassId((len(classids) * 2) + int(startId)), nil
	})

	classid := sym.(TcClassId)
	tmpClassIds = append(tmpClassIds, classid)

	return classid
}

func removeClassId(classids []TcClassId, id TcClassId) {
	for i, curr := range classids {
		if id == curr {
			classids = append(classids[:i], classids[i+1:]...)
			break
		}
	}
}

func orderedIds() []TcClassId {
	classids := []TcClassId{}

	for _, id := range tmpClassIds {
		classids = append(classids, id)
	}

	for _, id := range usedClassIds {
		classids = append(classids, id)
	}

	sort.Slice(classids, func(i, j int) bool {
		return classids[i] < classids[j]
	})

	return classids
}
