package mySys

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
	mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.MySysServiceGroup.OrderService

// CreateOrder 创建订单
// @Tags Order
// @Summary 创建订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.Order true "创建订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /order/createOrder [post]
func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var order mySys.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.CreatedBy = utils.GetUserID(c)

	if err := orderService.CreateOrder(&order); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除订单
// @Tags Order
// @Summary 删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.Order true "删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := orderService.DeleteOrder(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除订单
// @Tags Order
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /order/deleteOrderByIds [delete]
func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := orderService.DeleteOrderByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新订单
// @Tags Order
// @Summary 更新订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.Order true "更新订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var order mySys.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.UpdatedBy = utils.GetUserID(c)

	if err := orderService.UpdateOrder(order); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询订单
// @Tags Order
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.Order true "用id查询订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	ID := c.Query("ID")
	if reorder, err := orderService.GetOrder(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// GetOrderList 分页获取订单列表
// @Tags Order
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.OrderSearch true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo mySysReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetOrderPublic 返回周销售表
// @Tags Order
// @Summary 不需要鉴权的订单接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.OrderSearch true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderPublic(c *gin.Context) {
	// 获取当前时间
	now := time.Now()

	// 计算本周的起始时间（本周一的零点）
	thisMonday := now.AddDate(0, 0, -int(now.Weekday())+1)
	thisMonday = time.Date(thisMonday.Year(), thisMonday.Month(), thisMonday.Day(), 0, 0, 0, 0, thisMonday.Location())

	// 获取本周的结束时间（本周日的23:59:59）
	thisSunday := thisMonday.AddDate(0, 0, 6).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	fmt.Println(thisMonday)
	fmt.Println(thisSunday)
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	day := [7]string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}
	var pageInfo mySysReq.OrderSearch
	pageInfo.StartCreatedAt = &thisMonday
	pageInfo.EndCreatedAt = &thisSunday
	list, _, _ := orderService.GetOrderInfoList(pageInfo)
	//处理数据
	// 初始化周销表数据结构
	weeklySales := make(map[string]map[string]float64)

	for _, order := range list {
		// 确保订单创建时间非空
		if !order.CreatedAt.IsZero() {
			// 调试输出订单创建时间

			// 获取订单创建时间所在的星期
			dayIndex := int(order.CreatedAt.Weekday()) - 1 // 减去1以匹配数组索引
			if dayIndex < 0 {
				dayIndex += 7 // 修正为负索引
			}
			dayName := day[dayIndex]

			// 根据商品类型统计销售量
			typeName := order.GoodsType
			if typeName != "" {
				if weeklySales[typeName] == nil {
					weeklySales[typeName] = make(map[string]float64)
				}
				// 解引用指针以获取销售额的值，然后再加到weeklySales中
				if order.AllPrice != nil {
					if _, ok := weeklySales[typeName][dayName]; !ok {
						weeklySales[typeName][dayName] = *order.AllPrice
					} else {
						weeklySales[typeName][dayName] += *order.AllPrice
					}
				}
			}
		}
	}
	response.OkWithDetailed(gin.H{
		"Day":         day,
		"weeklySales": weeklySales,
	}, "获取成功", c)
}

// GetOrderPublic 返回周销售表
// @Tags Order
// @Summary 不需要鉴权的订单接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.OrderSearch true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetAllOrderProfitPublic(c *gin.Context) {
	//此函数，查询所有订单，然后根据订单信息中的商品名查询商品进价，商品类型，然后计算出利润，统计每个商品类型的利润和，然后返回给前端
	var orders []mySys.Order
	var products []mySys.My_goods
	// 查询所有订单
	if err := global.GVA_DB.Find(&orders).Error; err != nil {
		response.FailWithMessage("查询订单失败", c)
		return
	}
	// 查询所有商品
	if err := global.GVA_DB.Find(&products).Error; err != nil {
		response.FailWithMessage("查询商品失败", c)
		return
	}

	// 创建以商品类型为键的利润映射
	profitMap := make(map[string]float64)

	// 遍历订单列表
	for _, order := range orders {
		// 遍历商品列表
		for _, product := range products {
			// 检查商品名是否匹配
			if product.GoodsName == order.GoodsName {
				// 累加到对应商品类型的利润中
				profitMap[order.GoodsType] += (*order.AllPrice - (float64(*order.GoodsNum) * *product.GoodsPrice))
			}
		}
	}

	response.OkWithDetailed(gin.H{
		"profitData": profitMap,
	}, "获取成功，利润数据接口", c)

}

// GetOrderPublic 分析订单,给用户进行进货建议，，给特定的用户推荐商品
// @Tags Order
// @Summary 不需要鉴权的订单接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.OrderSearch true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderGPTPublic(c *gin.Context) {
	// 从前端获取数据
	numRecommendations, _ := strconv.Atoi(c.Query("numRecommendations"))
	customerName := c.Query("customerName")
	similarVlue, _ := strconv.ParseFloat(c.Query("similarVlue"), 64)
	fmt.Println("前端传来的数据：")
	fmt.Println("商品推荐数量:", numRecommendations)
	fmt.Println("顾客姓名:", customerName)
	fmt.Println("Jaccard系数阈值:", similarVlue)
	// 根据按钮的值来决定执行的功能
	action := c.Query("action")
	fmt.Println("按钮的值:", action)
	var products []mySys.My_goods
	var orders []mySys.Order
	//获取所有订单
	if err := global.GVA_DB.Find(&orders).Error; err != nil {
		response.FailWithMessage("查询订单失败", c)
		return
	}
	//获取所有商品
	// 查询所有商品
	if err := global.GVA_DB.Find(&products).Error; err != nil {
		response.FailWithMessage("查询商品失败", c)
		return
	}
	switch action {
	case "purchase":
		fmt.Println("执行进货建议功能：")
		// 进货建议
		strPurchase := utils.RecommendHotItems(orders, numRecommendations)
		response.OkWithDetailed(gin.H{
			"AdviceJinHuo": strPurchase,
		}, "获取进货建议成功", c)
	case "recommend":
		fmt.Println("执行用户喜好商品推荐功能：")
		// 结合用户喜好进行商品推荐
		strUserAlike, str := utils.RecommendItemsTOUser(orders, products, customerName, similarVlue)
		response.OkWithDetailed(gin.H{
			"strUserAlike": str,
			"str":          strUserAlike,
		}, "获取用户喜好商品推荐成功", c)
	default:
		fmt.Println("未知操作：", action)
		response.FailWithMessage("未知操作", c)
	}
}

// GetOrderPublic 获取商品销售金额
// @Tags Order
// @Summary 不需要鉴权的订单接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.OrderSearch true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetGoodsSalesPublic(c *gin.Context) {
	//获取所有订单信息
	var orders []mySys.Order
	//获取所有订单
	if err := global.GVA_DB.Find(&orders).Error; err != nil {
		response.FailWithMessage("查询订单失败", c)
		return
	}
	//根据订单统计出所有商品的销售额，并作封装处理，返回给前端用来做一个饼状展示图
	// 统计商品销售额
	productCounts := make(map[string]float64)
	for _, order := range orders {
		productCounts[order.GoodsName] += *order.AllPrice // 商品总销售额累计
	}

	// 封装返回数据
	var result []gin.H
	for goodsName, allPrice := range productCounts {
		result = append(result, gin.H{
			"goodsName": goodsName,
			"allPrice":  allPrice,
		})
	}
	response.OkWithDetailed(gin.H{
		"productCounts": result,
	}, "获取成功，获取商品销售金额据接口", c)
}

// GetOrderPublic 获取商品利润
// @Tags Order
// @Summary 不需要鉴权的订单接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.OrderSearch true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetGoodsProfitPublic(c *gin.Context) {
	var products []mySys.My_goods
	var orders []mySys.Order
	//获取所有订单信息
	if err := global.GVA_DB.Find(&orders).Error; err != nil {
		response.FailWithMessage("查询订单失败", c)
		return
	}
	//获取所有商品商品信息
	if err := global.GVA_DB.Find(&products).Error; err != nil {
		response.FailWithMessage("查询商品失败", c)
		return
	}
	// 统计每个商品的总利润，并作封装处理，返回给前端用来做一个饼状展示图
	// 统计每个商品的总利润
	productProfits := make(map[string]float64)
	for _, order := range orders {
		for _, product := range products {
			if order.GoodsName == product.GoodsName {
				productProfits[order.GoodsName] += *order.AllPrice - *product.GoodsPrice*float64(*order.GoodsNum) // 商品总利润累计
			}
		}
	}
	// 封装返回数据
	var result []gin.H
	for goodsName, allProfit := range productProfits {
		result = append(result, gin.H{
			"goodsName": goodsName,
			"allProfit": allProfit,
		})
	}
	response.OkWithDetailed(gin.H{
		"productCounts": result,
	}, "获取成功，商品利润数据接口", c)
}
