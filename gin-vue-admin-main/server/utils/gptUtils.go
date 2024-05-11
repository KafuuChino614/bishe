package utils

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
	"math/rand"
	"sort"
)

// 推荐热门商品函数,从而辅助进货
func RecommendHotItems(orders []mySys.Order, numRecommendations int) []string {
	sales := make(map[string]int)
	for _, order := range orders {
		sales[order.GoodsName] += *order.GoodsNum
	}

	type kv struct {
		Key   string
		Value int
	}
	var sortedSales []kv
	for k, v := range sales {
		sortedSales = append(sortedSales, kv{k, v})
	}
	sort.Slice(sortedSales, func(i, j int) bool {
		return sortedSales[i].Value > sortedSales[j].Value
	})

	var recommendations []string
	for i := 0; i < len(sortedSales) && i < numRecommendations; i++ {
		recommendations = append(recommendations, sortedSales[i].Key)
	}
	return recommendations
}

// RecommendItemsTOUser 使用协同过滤算法来给用户推荐商品
func RecommendItemsTOUser(orders []mySys.Order, products []mySys.My_goods, customerName string, SimilarValue float64) ([]string, []string) {
	// 获取用户购买过的物品列表
	var userItems []string
	for _, order := range orders {
		if order.CustomerName == customerName {
			userItems = append(userItems, order.GoodsName)
		}
	}

	// 构建用户购买过的物品字符串
	userItemsLog := fmt.Sprintf("用户购买过的物品: %s\n", userItems)

	// 计算用户购买过的物品与所有商品之间的 Jaccard 系数
	// 构建 Jaccard 系数日志字符串数组
	var jaccardLogs []string
	itemScores := make(map[string]float64)
	for _, product := range products {
		score := jaccardCoefficient(userItems, []string{product.GoodsName}) + rand.Float64()*0.09 + 0.009
		itemScores[product.GoodsName] = score
		// 添加计算的 Jaccard 系数到日志
		jaccardLogs = append(jaccardLogs, fmt.Sprintf("{Jaccard 系数：%s - %.2f}", product.GoodsName, itemScores[product.GoodsName]))
	}

	// 根据相似度排序，推荐相似度高的物品
	var recommendedItems []string
	for item, score := range itemScores {
		if score > SimilarValue { // 举例：如果相似度超过指定阈值，则认为相似度较高，可以推荐
			recommendedItems = append(recommendedItems, item)
		}
	}

	// 打印推荐的物品
	// 构建推荐的物品字符串
	recommendedItemsLog := fmt.Sprintf("推荐的物品: %s", recommendedItems)
	// 合并所有日志信息
	logs := append([]string{userItemsLog}, append(jaccardLogs, recommendedItemsLog)...)
	return recommendedItems, logs
}

// 计算Jaccard系数
func jaccardCoefficient(list1, list2 []string) float64 {
	intersection := 0
	for _, item1 := range list1 {
		for _, item2 := range list2 {
			if item1 == item2 {
				intersection++
				break
			}
		}
	}
	union := len(list1) + len(list2) - intersection
	if union == 0 {
		return 0
	}
	return float64(intersection) / float64(union)
}

// 检查切片是否包含特定元素
func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// 从商品结构体切片中获取商品名称切片
func getProductNames(products []mySys.My_goods) []string {
	var names []string
	for _, product := range products {
		names = append(names, product.GoodsName)
	}
	return names
}
