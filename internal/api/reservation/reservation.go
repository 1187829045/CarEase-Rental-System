package api

import (
	"net/http"
	"strconv"
	"time"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/dao/model"
	_struct "car.rental/struct/reservation"
	"github.com/gin-gonic/gin"
)

func CreateReservation(c *gin.Context) {
	var form _struct.ReservationCreateForm
	if err := c.ShouldBindJSON(&form); err != nil || form.CarID == 0 || form.UserID == 0 || form.StoreID == 0 || form.StartTime.IsZero() || form.EndTime.IsZero() || !form.StartTime.Before(form.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	reservation := &model.Reservation{
		CarID:     form.CarID,
		UserID:    form.UserID,
		StoreID:   form.StoreID,
		StartTime: form.StartTime,
		EndTime:   form.EndTime,
		Status:    0,
	}
	if err := dao.CreateReservation(reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	order := &model.Order{
		ReservationID: reservation.ReservationID,
		CarID:         reservation.CarID,
		UserID:        reservation.UserID,
		StoreID:       reservation.StoreID,
		StartTime:     reservation.StartTime,
		EndTime:       reservation.EndTime,
		OrderType:     0,
		Status:        0,
	}
	if err := dao.CreateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toReservationInfo(reservation),
	})
}

func ListReservations(c *gin.Context) {
	var q _struct.ReservationQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	reservations, err := dao.ListReservations(q.CarID, q.UserID, q.StoreID, q.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	items := make([]_struct.ReservationInfo, 0, len(reservations))
	for _, reservation := range reservations {
		items = append(items, toReservationInfo(reservation))
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": items,
	})
}

func GetReservationDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	reservation, err := dao.GetReservationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toReservationInfo(reservation),
	})
}

func UpdateReservation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	var form _struct.ReservationUpdateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	reservation, err := dao.GetReservationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if form.CarID != nil {
		reservation.CarID = *form.CarID
	}
	if form.UserID != nil {
		reservation.UserID = *form.UserID
	}
	if form.StoreID != nil {
		reservation.StoreID = *form.StoreID
	}
	if form.StartTime != nil {
		reservation.StartTime = *form.StartTime
	}
	if form.EndTime != nil {
		reservation.EndTime = *form.EndTime
	}
	if form.Status != nil {
		reservation.Status = *form.Status
	}
	if reservation.StartTime.IsZero() || reservation.EndTime.IsZero() || !reservation.StartTime.Before(reservation.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	if err := dao.UpdateReservation(reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toReservationInfo(reservation),
	})
}

func CancelReservation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	if err := dao.UpdateReservationStatus(uint(id), 1); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	reservation, err := dao.GetReservationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toReservationInfo(reservation),
	})
}

func ConfirmReservation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	reservation, err := dao.GetReservationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if reservation.Status == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	if reservation.Status != 2 {
		reservation.Status = 2
		if err := dao.UpdateReservation(reservation); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
	}
	order, err := dao.GetOrderByReservationID(reservation.ReservationID)
	if err != nil {
		if err.Error() == "订单信息不存在" {
			order = &model.Order{
				ReservationID: reservation.ReservationID,
				CarID:         reservation.CarID,
				UserID:        reservation.UserID,
				StoreID:       reservation.StoreID,
				StartTime:     reservation.StartTime,
				EndTime:       reservation.EndTime,
				OrderType:     1,
				Status:        0,
			}
			if err := dao.CreateOrder(order); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
	} else {
		order.OrderType = 1
		if order.Status == 0 {
			order.Status = 0
		}
		if err := dao.UpdateOrder(order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
	}
	order.StartTime = reservation.StartTime
	order.EndTime = reservation.EndTime
	if err := dao.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func ListOrders(c *gin.Context) {
	var q _struct.OrderQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	orders, err := dao.ListOrders(q.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	items := make([]_struct.OrderInfo, 0, len(orders))
	for _, order := range orders {
		items = append(items, toOrderInfo(order))
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": items,
	})
}

func CreateOrder(c *gin.Context) {
	var form _struct.OrderCreateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	var order *model.Order
	if form.ReservationID != nil && *form.ReservationID != 0 {
		reservation, err := dao.GetReservationByID(*form.ReservationID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
		if reservation.Status == 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  consts.ErrInvalidParameter,
			})
			return
		}
		if reservation.Status != 2 {
			reservation.Status = 2
			if err := dao.UpdateReservation(reservation); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				})
				return
			}
		}
		order, err = dao.GetOrderByReservationID(reservation.ReservationID)
		if err != nil {
			if err.Error() == "订单信息不存在" {
				order = &model.Order{
					ReservationID: reservation.ReservationID,
					CarID:         reservation.CarID,
					UserID:        reservation.UserID,
					StoreID:       reservation.StoreID,
					StartTime:     reservation.StartTime,
					EndTime:       reservation.EndTime,
					OrderType:     1,
					Status:        0,
				}
				if err := dao.CreateOrder(order); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": http.StatusInternalServerError,
						"msg":  err.Error(),
					})
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				})
				return
			}
		} else {
			order.OrderType = 1
			order.StartTime = reservation.StartTime
			order.EndTime = reservation.EndTime
			if err := dao.UpdateOrder(order); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				})
				return
			}
		}
	} else {
		if form.CarID == 0 || form.UserID == 0 || form.StoreID == 0 || form.StartTime.IsZero() || form.EndTime.IsZero() || !form.StartTime.Before(form.EndTime) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  consts.ErrInvalidParameter,
			})
			return
		}
		order = &model.Order{
			ReservationID: 0,
			CarID:         form.CarID,
			UserID:        form.UserID,
			StoreID:       form.StoreID,
			StartTime:     form.StartTime,
			EndTime:       form.EndTime,
			OrderType:     1,
			Status:        0,
		}
		if err := dao.CreateOrder(order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func GetOrderDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	order, err := dao.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func PickupOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	var form _struct.OrderPickupForm
	if err := c.ShouldBindJSON(&form); err != nil || form.Mileage < 0 || form.Fuel < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	order, err := dao.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	order.Status = 1
	order.PickupInspection = form.Inspection
	order.PickupMileage = form.Mileage
	order.PickupFuel = form.Fuel
	order.PickupPhotos = form.Photos
	if err := dao.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func ReturnOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	var form _struct.OrderReturnForm
	if err := c.ShouldBindJSON(&form); err != nil || form.Mileage < 0 || form.Fuel < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	order, err := dao.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	order.ReturnInspection = form.Inspection
	order.ReturnMileage = form.Mileage
	order.ReturnFuel = form.Fuel
	order.ReturnPhotos = form.Photos
	if time.Now().After(order.EndTime) {
		order.Status = 3
	} else {
		order.Status = 2
	}
	if err := dao.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func ExtendOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	var form _struct.OrderExtendForm
	if err := c.ShouldBindJSON(&form); err != nil || form.EndTime.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	order, err := dao.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if !form.EndTime.After(order.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	order.EndTime = form.EndTime
	if err := dao.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func DamageReportOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	var form _struct.OrderDamageReportForm
	if err := c.ShouldBindJSON(&form); err != nil || form.Report == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	order, err := dao.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	order.DamageReport = form.Report
	if err := dao.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func CancelOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	if err := dao.UpdateOrderStatus(uint(id), 4); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	order, err := dao.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": toOrderInfo(order),
	})
}

func toReservationInfo(reservation *model.Reservation) _struct.ReservationInfo {
	return _struct.ReservationInfo{
		ID:        reservation.ReservationID,
		CarID:     reservation.CarID,
		UserID:    reservation.UserID,
		StoreID:   reservation.StoreID,
		StartTime: reservation.StartTime,
		EndTime:   reservation.EndTime,
		Status:    reservation.Status,
	}
}

func toOrderInfo(order *model.Order) _struct.OrderInfo {
	return _struct.OrderInfo{
		ID:               order.OrderID,
		ReservationID:    order.ReservationID,
		CarID:            order.CarID,
		UserID:           order.UserID,
		StoreID:          order.StoreID,
		StartTime:        order.StartTime,
		EndTime:          order.EndTime,
		OrderType:        order.OrderType,
		Status:           order.Status,
		PickupInspection: order.PickupInspection,
		ReturnInspection: order.ReturnInspection,
		PickupMileage:    order.PickupMileage,
		ReturnMileage:    order.ReturnMileage,
		PickupFuel:       order.PickupFuel,
		ReturnFuel:       order.ReturnFuel,
		PickupPhotos:     order.PickupPhotos,
		ReturnPhotos:     order.ReturnPhotos,
		DamageReport:     order.DamageReport,
	}
}