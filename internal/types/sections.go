package types

type SectionInfo struct {
	LineCount int
	StartLine int
	Content   []string
}

type Sections struct {
	Preamble SectionInfo
	Data     SectionInfo
	Bss      SectionInfo
	Text     SectionInfo
}
