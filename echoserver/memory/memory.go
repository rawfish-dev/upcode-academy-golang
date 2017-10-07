package memory

import (
	"upcode-academy-golang/echoserver/common"
	"upcode-academy-golang/echoserver/interfaces"
)

/* An interface check is missing here */

type MemoryBuffer struct{}

func NewMemoryBuffer() *MemoryBuffer {
	return &MemoryBuffer{}
}

func (m *MemoryBuffer) BuildMessagePad() string {
	return common.BuildMessagePad("Memory Buffer")
}

func (m *MemoryBuffer) ReadAll() []string {
	/* A function call is missing here */
}

func dataSource() []string {
	return []string{
		"5 affordable private islands in Asia",
		"The Asteria condo sells for $27.1 million",
		"iPhone 8 Plus batteries are swelling out of their cases",
	}
}
