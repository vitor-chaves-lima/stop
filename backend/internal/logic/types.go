package logic

// PaginationInfo holds metadata about the pagination state.
type PaginationInfo struct {
	TotalCount   int // Total number of items available
	PageCount    int // Total number of pages available
	CurrentPage  int // The current page number
	ItemsPerPage int // Number of items per page
}

// NewPaginationInfo creates and returns a new PaginationInfo structure.
// It calculates the PageCount based on the TotalCount and ItemsPerPage.
func NewPaginationInfo(totalCount, currentPage, itemsPerPage int) *PaginationInfo {
	pageCount := (totalCount + itemsPerPage - 1) / itemsPerPage

	return &PaginationInfo{
		TotalCount:   totalCount,
		PageCount:    pageCount,
		CurrentPage:  currentPage,
		ItemsPerPage: itemsPerPage,
	}
}
