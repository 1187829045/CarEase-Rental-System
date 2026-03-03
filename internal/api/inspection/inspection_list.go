package api

import (
	"sync"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	"car.rental/struct/inspection"
	"github.com/gin-gonic/gin"
)

// GetInspectionList 获取检测报告列表
func GetInspectionList(c *gin.Context) {
	var q inspection.InspectionListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 构建响应
	resp := inspection.InspectionListResp{}
	resp.Counts = make(map[int8]int64)

	// 并发查询各种状态的数量
	wg := sync.WaitGroup{}
	wg.Add(3)

	// 待审核数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountInspectionReportsByStatus(0)
		if err == nil {
			resp.Counts[0] = count
		}
	}()

	// 已通过数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountInspectionReportsByStatus(1)
		if err == nil {
			resp.Counts[1] = count
		}
	}()

	// 已拒绝数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountInspectionReportsByStatus(2)
		if err == nil {
			resp.Counts[2] = count
		}
	}()
	wg.Wait()
	reports, err := dao.ListInspectionReports(q.CarID, q.InspectorID, q.Type, q.Status)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// 转换为响应格式
	resp.Items = make([]*inspection.InspectionReportResponse, len(reports))
	for i, report := range reports {
		resp.Items[i] = &inspection.InspectionReportResponse{
			ReportID:       report.ReportID,
			CarID:          report.CarID,
			InspectorID:    report.InspectorID,
			Type:           report.Type,
			Mileage:        report.Mileage,
			Exterior:       report.Exterior,
			Interior:       report.Interior,
			Notes:          report.Notes,
			Photos:         report.Photos,
			InspectorName:  report.InspectorName,
			InspectionTime: report.InspectionTime.Format("2006-01-02 15:04:05"),
			Status:         report.Status,
		}
	}

	response.Success(c, resp)
}
