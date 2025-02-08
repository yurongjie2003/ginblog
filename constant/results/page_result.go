package results

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const (
	MaxPageSize = 50 // 单页最大数量
	MinPageSize = 5  // 单页最小数量
)

type PageResult struct {
	Total   int64       `json:"total"`   // 总记录数
	Records interface{} `json:"records"` // 当前页数据集合
}

type PageParams struct {
	page     int
	Offset   int
	PageSize int
}

func normalizePageSize(size int) int {
	if size > MaxPageSize {
		return MaxPageSize
	}
	if size < MinPageSize {
		return MinPageSize
	}
	return size
}

func NewPageParams(page, pageSize int) *PageParams {
	return &PageParams{
		page:     page,
		PageSize: normalizePageSize(pageSize),
		Offset:   (page - 1) * pageSize,
	}
}

// 默认分页：第一页，size默认最小值
var defaultPageParams = NewPageParams(1, MinPageSize)

func GetPageParams(c *gin.Context) *PageParams {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Printf("Invalid page parameter: %s", pageStr)
		return defaultPageParams
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		log.Printf("Invalid pageSize parameter: %s", pageSizeStr)
		return defaultPageParams
	}

	if page <= 0 || pageSize <= 0 {
		log.Printf("Invalid pageParams parameter: page-%d page-size-%d", page, pageSize)
		return defaultPageParams
	}

	return NewPageParams(page, pageSize)
}
