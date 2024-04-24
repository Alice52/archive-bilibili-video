package c

const ( // 0: 未同步, 1: 同步中, 2: 同步完成
	SyncStatusTodo  = 0
	SyncStatusDoing = 1
	SyncStatusDone  = 2
)

const ( // video archived type
	ArchivedTypeFav  = 0
	ArchivedTypeCoin = 1
	ArchivedTypeLike = 2
	ArchivedTypeView = 3
)