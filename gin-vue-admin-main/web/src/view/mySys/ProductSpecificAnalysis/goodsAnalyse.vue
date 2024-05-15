<template>
    <div class="gva-card-box">
      <!-- 大容器 -->
      <section class="container">
        <!-- 左容器 -->
        <section class="itemLeft">
          <ItemPage >
            <ItemOne></ItemOne>
          </ItemPage>
        </section>
        <!-- 右容器 -->
        <section class="itemRight">
          <ItemPage >
            <ItemTwo></ItemTwo>
          </ItemPage>
        </section>
      </section>
    </div>
  </template>



<script setup>
 import ItemPage from "@/view/mySys/ProductSpecificAnalysis/components/itemPage.vue";
 import ItemOne from "@/view/mySys/ProductSpecificAnalysis/components/itemOne.vue";
 import ItemTwo from "@/view/mySys/ProductSpecificAnalysis/components/itemTwo.vue";

import { provide } from "vue";
import * as echarts from "echarts";
provide("echarts", echarts);
</script>

<style lang="less">

  // 大容器的样式
  .container {
    // background-image: url("/src/assets/bg.jpg");
    // 最大最小的宽度
    min-width: 1200px;
    max-width: 2048px;
    min-height: 750px;
    margin: 0 auto;
    padding: 10px 10px 0;
    display: flex;
    // 设置左右在页面的份数
    .itemLeft,
    .itemRight {
      flex: 1;
    }

  }
  </style>