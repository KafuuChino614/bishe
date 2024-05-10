<template>
    <div class="gva-card-box">
      <!-- 大容器 -->
      <section class="container">
        <!-- 左容器 -->
        <section class="itemLeft">
          <ItemPage >
            <ItemOne/>
          </ItemPage>
           <ItemPage >
            <ItemTwo/>
          </ItemPage>
        </section>
        <!-- 中容器 -->
        <section class="itemCenter">
          <ItemPage >
            <ItemFive/>
          </ItemPage>
          <ItemPage >
            <MapPage/>
          </ItemPage>
        </section>
        <!-- 右容器 -->
        <section class="itemRight">
          <ItemPage >
            <ItemThree/>
          </ItemPage>
           <ItemPage >
            <ItemFour/>
          </ItemPage>
        </section>
      </section>
    </div>
  </template>



<script setup>
 import ItemPage from "@/view/mySys/dataAnalyseMange/components/itemPage.vue";
 import ItemOne from "@/view/mySys/dataAnalyseMange/components/itemOne.vue";
 import ItemTwo from "@/view/mySys/dataAnalyseMange/components/itemTwo.vue";
 import ItemThree from "@/view/mySys/dataAnalyseMange/components/itemThree.vue";
 import ItemFour from "@/view/mySys/dataAnalyseMange/components/itemFour.vue";
 import ItemFive from "@/view/mySys/dataAnalyseMange/components/itemFive.vue";
 import MapPage from "@/view/mySys/dataAnalyseMange/components/mapPage.vue";
import { provide } from "vue";
import * as echarts from "echarts";
provide("echarts", echarts);
// import"lib-flexible/flexible.js"
</script>

<style lang="less">

  // 大容器的样式
  .container {
    // background-image: url("/src/assets/bg.jpg");
    // 最大最小的宽度
    min-width: 1200px;
    max-width: 2048px;
    min-height: 650px;
    margin: 0 auto;
    padding: 10px 10px 0;
    display: flex;
    // 设置左右在页面的份数
    .itemLeft,
    .itemRight {
      flex: 3;
    }
    .itemCenter {
      flex: 3;

    }

  }
  </style>