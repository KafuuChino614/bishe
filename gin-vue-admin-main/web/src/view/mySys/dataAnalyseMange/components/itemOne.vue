<template>
<div>
    <h2>商品销售图</h2>
    <div class="chart" id="oneChart">图表的容器</div>
  </div>
</template>

<script setup>

import { inject, onMounted, reactive ,ref} from "vue";
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList
} from '@/api/mySys/order'
let $echarts = inject("echarts");


let data = reactive({});
let xdata = reactive([]);
let ydata = reactive([]);

function setData() {
  xdata = data.map((v) => v.goodsName);
  ydata = data.map((v) => v.goodsNum);
}

async function getState() {
  var orderList=ref([])
  orderList.value=await (await getOrderList()).data.list
  let newList = []
  //处理数据
  for (var i = 0; i < orderList.value.length; i++) {
    let goodsName = orderList.value[i].goodsName;
    let goodsNum = orderList.value[i].goodsNum;
    // 检查 newList 中是否已存在相同 goodsName 的订单信息
    let existingOrder = newList.find(order => order.goodsName === goodsName);
    // 如果存在相同的 goodsName，则将其对应的 goodsNum 相加
    if (existingOrder) {
      existingOrder.goodsNum += goodsNum;
    } else {
      // 否则，将当前订单的 goodsName 和 goodsNum 添加到 processedOrderList 中
      newList.push({ goodsName, goodsNum });
    }
  }
  data=newList  


}

onMounted(() => {
  let myChart = $echarts.init(document.getElementById("oneChart"));
  // 调用请求
  getState().then(() => {
    setData();
    myChart.setOption({
      grid: {
        top: "3%",
        left: "1%",
        right: "6%",
        bottom: "3%",
        containLabel: true
      },
      xAxis: {
        type: "value",
        axisLine: {
          lineStyle: {
            color: "#fff"
          }
        }
      },
      yAxis: {
        type: "category",
        data: xdata,
        axisLine: {
          lineStyle: {
            color: "#fff"
          }
        }
      },
      series: [
        {
          data: ydata,
          type: "bar",
          itemStyle: {
            normal: {
              barBorderRadius: [0, 20, 20, 0],
              color: new $echarts.graphic.LinearGradient(0, 0, 1, 0, [
                {
                  offset: 0,
                  color: "#005eaa"
                },
                {
                  offset: 0.5,
                  color: "#339ca8"
                },
                {
                  offset: 1,
                  color: "#cda819"
                }
              ])
            }
          }
        }
      ]
    });
  });
});

</script>


<style scoped>

.chart {
  height: 360px;
}
h2 {
  height: 48px;
  color: #fff;
  line-height: 48px;
  font-size: 20px;
  text-align: center;
}
</style>