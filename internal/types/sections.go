package types

type SectionInfo struct {
	LineCount int
	StartLine int
	Content   []string
}

type Sections struct {
	Preamble SectionInfo
	Data     []string
	Bss      []string
	Text     []string
}
