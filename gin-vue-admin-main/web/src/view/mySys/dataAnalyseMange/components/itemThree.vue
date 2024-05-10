<template>
  <div>
      <h2>库存商品类型数量分析</h2>
      <div class="chart" id="myEcharts">
          图表的容器
      </div>
  </div>
</template>

<script setup>
import {
  createWareHouseInfo,
  deleteWareHouseInfo,
  deleteWareHouseInfoByIds,
  updateWareHouseInfo,
  findWareHouseInfo,
  getWareHouseInfoList
} from '@/api/mySys/wareHouseInfo'
import { inject, onMounted, reactive ,ref} from "vue";
let $echarts = inject("echarts");
let $http = inject("axios");

let data = reactive({});

async function getState() {
  var wareHouseInfoList=ref([])
  wareHouseInfoList.value=await (await getWareHouseInfoList()).data.list
  let newList = []
  //处理数据
  for (var i = 0; i < wareHouseInfoList.value.length; i++) {
    let goodsType = wareHouseInfoList.value[i].goodsType;
    let num = wareHouseInfoList.value[i].num;
    // 检查 newList 中是否已存在相同 goodsType 
    let existingOrder = newList.find(item => item.name === goodsType);
    // 如果存在相同的 goodsType，则将其对应的 goodsNum 相加
    if (existingOrder) {
      existingOrder.value += num;
    } else {
      // 否则，将当前订单的 goodsName 和 goodsNum 添加到 processedOrderList 中
      newList.push({ name: goodsType, value: num });
    }
    
  }
  data=newList

}

onMounted(() => {
  getState().then(() => {


    let myChart = $echarts.init(document.getElementById("myEcharts"));
    myChart.setOption({
      legend: {
        top: "bottom",
      },
      tooltip: {
        show: true,
      },
      series: [
        {
          type: "pie",
          data: data,
          radius: [10, 100],
          center: ["50%", "45%"],
          roseType: "area",
          itemStyle: {
            borderRadius: 10,
          },
        },
      ],
    });
  });
});

</script>

<style>
h2 {
  height: 48px;
  color: #fff;
  line-height: 48px;
  font-size: 20px;
  text-align: center;
}
.chart{
  height: 360px;
}
</style>