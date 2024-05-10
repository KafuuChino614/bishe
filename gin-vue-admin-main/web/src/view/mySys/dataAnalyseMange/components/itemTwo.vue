<template>
  <div>
    <h2>周销图</h2>
    <div class="chart" id="myEchartsTwo">图表的容器</div>
  </div>
</template>

<script setup>
import { inject, onMounted, reactive } from "vue";
import {
  getetOrderPublic
} from '@/api/mySys/order'
let $echarts = inject("echarts");

let data = reactive({});

async function getState() {
  data=(await getetOrderPublic()).data

}

onMounted(() => {
  getState().then(() => {

    let legendData = []; // 存放图例数据
    let xAxisData = data.Day; // X 轴数据
    let seriesData = []; // 存放系列数据

    // 处理系列数据
    for (let type in data.weeklySales) {
      legendData.push(type); // 添加图例数据
      let seriesItem = {
        name: type,
        type: "line",
        data: [], // 存放每个类型对应的销售额数据
      };
      for (let i = 0; i < xAxisData.length; i++) {
        let day = xAxisData[i];
        let sales = data.weeklySales[type][day] || 0; // 获取销售额，如果没有则为0
        seriesItem.data.push(sales); // 添加销售额数据
      }
      seriesData.push(seriesItem); // 添加到系列数据中
    }

    let myChart = $echarts.init(document.getElementById("myEchartsTwo"));
    myChart.setOption({
      tooltip: {
        trigger: "axis",
        axisPointer: {
          type: "cross",
          label: {
            backgroundColor: "#e6b600",
          },
        },
      },
      legend: {
        data: legendData, // 使用处理后的图例数据
      },
      xAxis: {
        type: "category",
        boundaryGap: false,
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
      series: seriesData, // 使用处理后的系列数据
    });
  });
});
// 生成渐变颜色
function generateColor() {
  // 生成随机的RGB颜色
  const r = Math.floor(Math.random() * 256);
  const g = Math.floor(Math.random() * 256);
  const b = Math.floor(Math.random() * 256);
  return new $echarts.graphic.LinearGradient(0, 0, 0, 1, [
    { offset: 0, color: `rgb(${r}, ${g}, ${b})` },
    { offset: 1, color: "rgb(0, 0, 0)" }, // 黑色作为渐变的结束色
  ]);
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