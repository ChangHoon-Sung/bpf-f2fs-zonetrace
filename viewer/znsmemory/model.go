package znsmemory

const SegmentSize = 512

type SegmentType int

const (
	NotChanged SegmentType = iota - 2 // no color

	UnknownSegment // gray

	// HotDataSegment is type for frequently accessed data segments
	HotDataSegment // red

	// WarmDataSegment is type for commonly accessed data segments
	WarmDataSegment // yellow

	// ColdDataSegment is type for infrequently accessed data segments
	ColdDataSegment // blue

	// HotNodeSegment is type for frequently accessed node segments
	HotNodeSegment // pink

	// WarmNodeSegment is type for commonly accessed node segments
	WarmNodeSegment // orange

	// ColdNodeSegment is type for infrequently accessed node segments
	ColdNodeSegment // sky blue

	EmptySegment // black
)

func (t SegmentType) IsValid() bool {
	return HotDataSegment <= t && t <= ColdNodeSegment
}

type ValidMap []byte

type Segment struct {
	ValidMap    ValidMap
	SegmentType SegmentType
	DirtyCount  uint64
}

type SegmentUpdateEvent struct {
	ZoneNo         int
	SegmentNo      int
	SegmentType    SegmentType
	ZoneDirtyCount uint64
}

type SitEntryUpdate struct {
	SegmentFullNo int
	ZoneNo        int
	SegmentType   SegmentType
	ValidMap      ValidMap
}

type ZoneInfo struct {
	MountPath               string
	TotalZone               int
	BlockPerSegment         int
	TotalBlockPerZone       int
	AvailableBlockPerZone   int
	TotalSegmentPerZone     int
	AvailableSegmentPerZone int
}

type FileInfo struct {
	FileSegments []FileSegment
}

type FileSegment struct {
	ZoneIndex    int
	SegmentIndex int
	ValidMap     ValidMap // 해당하는것을 1로 바꾼 512bit
}

type Zone struct {
	ZoneDirtyCount   uint64
	ZoneNo           int
	Segments         []Segment
	LastSegmentType  SegmentType
	SegmentTypeCount map[SegmentType]int
}

func (z *Zone) FrequentSegmentType() SegmentType {
	if z.ZoneDirtyCount == 0 {
		return EmptySegment
	}
	var maxType SegmentType
	maxCount := 0
	for t, count := range z.SegmentTypeCount {
		if count > maxCount {
			maxCount = count
			maxType = t
		}
	}
	return maxType
}

type ZonedStorage struct {
	ZoneInfo
	Zones []*Zone
}
