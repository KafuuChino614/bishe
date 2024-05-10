<template>
  <div>
    <h2>库存商品数量分析</h2>
    <div class="chart" id="myEchartsFour">图表的容器</div>
  </div>
</template>

<script setup>
import { inject, onMounted, reactive } from "vue";
import {
  getWareHouseInfoPublic
} from '@/api/mySys/wareHouseInfo'
const $echarts = inject("echarts");

let data = reactive({});

const getState = async () => {
  data.value = (await getWareHouseInfoPublic()).data;
  renderChart();
};

onMounted(() => {
  getState();
});
function renderChart() {
  const myChart = $echarts.init(document.getElementById("myEchartsFour"));
  const { wareHouseProductCounts } = data.value;

  const wareHouseNames = Object.keys(wareHouseProductCounts).sort(); // 对仓库名称进行排序

  const xAxisData = [];
  const seriesData = [];

  // 遍历每个仓库
  for (const wareHouseName of wareHouseNames) {
    xAxisData.push(wareHouseName); // 将仓库名添加到 X 轴数据中

    const productTypes = wareHouseProductCounts[wareHouseName];
    const productTypeData = {};

    // 遍历每个产品类型
    for (const productType in productTypes) {
      const products = productTypes[productType];

      // 遍历每个商品
      for (const productName in products) {
        const productCount = products[productName];
        
        // 将商品名及其数量添加到该仓库的数据中
        if (!productTypeData[productName]) {
          productTypeData[productName] = Array(xAxisData.length - 1).fill(0); // 初始化数据，长度减一是因为 xAxisData 已经添加了当前仓库的名称
        }
        productTypeData[productName].push(productCount);
      }
    }

    // 将该仓库的数据添加到 seriesData 中
    for (const productName in productTypeData) {
      seriesData.push({
        name: productName,
        type: "bar",
        stack: wareHouseName,
        label: { show: true },
        data: productTypeData[productName],
      });
    }
  }
  myChart.setOption({
    grid: {
      left: "3%",
      right: "4%",
      bottom: "3%",
      containLabel: true,
    },
    xAxis: {
      type: "category",
      data: xAxisData,
      axisLine: {
        lineStyle: {
          color: "#fff",
        },
      },
    },
    yAxis: {
      type: "value",
      axisLine: {
        lineStyle: {
          color: "#fff",
        },
      },
    },
    legend: {},
    tooltip: {
      trigger: "axis",
      axisPointer: {
        type: "shadow",
      },
    },
    series: seriesData,
  });
}

</script>


<style>
h2 {
  height: 48px;
  color: #fff;
  line-height: 48px;
  font-size: 20px;
  text-align: center;
}
.chart {
  height: 360px;
}
</style>