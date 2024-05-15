<template>
  <div>
    <h2 class="chart-title">商品类型利润分析</h2>
    <div class="chart" ref="profitChart"></div>
  </div>
</template>

<script setup>
import { inject, onMounted, ref } from "vue";
import { getAllOrderProfitPublic } from '@/api/mySys/order'

const $echarts = inject("echarts");
const profitChart = ref(null);
const data = ref(null);

async function getState() {
  try {
    data.value = (await getAllOrderProfitPublic()).data;
    renderChart(data.value);
  } catch (error) {
    console.error("Error fetching profit data:", error);
  }
}
onMounted(async () => {
  await getState();
});

const renderChart = (profitData) => {
  const chart = $echarts.init(profitChart.value);

  const xAxisData = []; // 存放商品类型的名称
  const profitDataList = []; // 存放利润数据

  // 遍历传入的利润数据，生成 x 轴和利润数据
  for (const key in profitData.profitData) {
    xAxisData.push(key); // 商品类型的名称
    profitDataList.push(profitData.profitData[key]); // 利润数据
  }

  chart.setOption({
    title: {
      left: 'center',
      textStyle: {
        color: '#333', // 设置标题颜色
        fontSize: 18, // 设置标题字体大小
        fontWeight: 'bold' // 设置标题字体粗细
      }
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      }
    },
    xAxis: {
      type: 'category',
      data: xAxisData,
      axisLabel: {
        rotate: 45, // 适当旋转标签，防止重叠
        textStyle: {
          color: '#333' // 设置标签字体颜色
        }
      },
      axisLine: {
        lineStyle: {
          color: '#ccc' // 调整轴线颜色
        }
      }
    },
    yAxis: {
      type: 'value',
      name: '（单位：元）',
      axisLabel: {
        textStyle: {
          color: ' #ff0' // 设置标签字体颜色
        }
      },
      axisLine: {
        lineStyle: {
          color: '#ccc' // 调整轴线颜色
        }
      }
    },
    series: [
      {
        type: 'bar',
        data: profitDataList,
        itemStyle: {
          color: ' #FFCC99' // 调整柱状图颜色
        }
      }
    ]
  });
};
</script>

<style>
.chart-title {
  height: 48px;
  color:  #fff; /* 修改标题颜色 */
  line-height: 48px;
  font-size: 20px; /* 修改标题字体大小 */
  text-align: center;
}
.chart {
  height: 360px;
}
</style>
