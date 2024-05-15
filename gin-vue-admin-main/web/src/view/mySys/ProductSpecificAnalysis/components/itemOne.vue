<template>
  <div>
    <h2>商品销售额/元</h2>
    <div class="chart" ref="myEchartsFive"></div>
  </div>
</template>

<script setup>
import { inject, onMounted, ref } from "vue";
import { getGoodsSalesPublic } from '@/api/mySys/order';

const $echarts = inject("echarts");
const myEchartsFive = ref(null);
const data = ref(null);

async function fetchData() {
  try {
    const response = await getGoodsSalesPublic();
    data.value = response.data.productCounts; // 只取 productCounts 部分的数据
    renderChart();
  } catch (error) {
    console.error("Error fetching product data:", error);
  }
}

onMounted(() => {
  fetchData();
});

const renderChart = () => {
  const chart = $echarts.init(myEchartsFive.value);

  const legendData = data.value.map(item => item.goodsName); // 商品类型作为图例数据
  const seriesData = data.value.map(item => ({ value: item.allPrice, name: item.goodsName })); // 商品类型及其数量作为饼状图数据

  chart.setOption({
    tooltip: {
      trigger: "item",
      formatter: "{a} <br/>{b}: {c} ({d}%)"
    },
    legend: {
      orient: "vertical",
      left: 10,
      data: legendData
    },
    series: [
      {
        name: "商品类型",
        type: "pie",
        radius: ["50%", "70%"],
        avoidLabelOverlap: false,
        label: {
          show: false,
          position: "center"
        },
        emphasis: {
          label: {
            show: true,
            fontSize: "30",
            fontWeight: "bold"
          }
        },
        labelLine: {
          show: false
        },
        data: seriesData
      }
    ]
  });
};
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
