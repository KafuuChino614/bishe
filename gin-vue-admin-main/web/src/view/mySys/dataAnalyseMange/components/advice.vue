<template>
    <div class="container">
      <!-- 进货建议模块 -->
      <div class="module">
        <h2>进货建议</h2>
        <div class="form">
          <label for="recommendationInput">推荐商品数量：</label>
          <input type="number" id="recommendationInput" v-model="recommendationCount" />
          <button @click="getPurchaseAdvice" data-action="purchase">获取建议</button>
        </div>
        <div class="description">
          <h3>进货建议：</h3>
          <textarea :value="purchaseAdvice" readonly></textarea>
        </div>
      </div>
  
      <!-- 用户喜好分析模块 -->
      <div class="module">
        <h2>用户喜好分析</h2>
        <div class="form">
          <label for="customerSelect">选择客户：</label>
          <el-select id="customerSelect" v-model="selectedCustomer" placeholder="请选择客户" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in customerNameOptions" :value="item.name" :key="key" :label="item.name" />
          </el-select>
          <label for="jaccardInput">Jaccard系数阈值：</label>
          <input type="number" id="jaccardInput" v-model="jaccardThreshold" />
          <button @click="getUserRecommendations" data-action="recommend">获取推荐</button>
        </div>
        <div class="description">
          <h3>用户喜好商品推荐：</h3>
          <textarea :value="userRecommendations" readonly></textarea>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
import {
  getOrderGPTPublic
} from '@/api/mySys/order'
import {
  getCustomerList
} from '@/api/mySys/customer'
  import { ref } from 'vue';


var customerNameOptions=ref([])
  // 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
   var tmp0=ref([])
   var tmp=[]
    tmp0.value=await (await getCustomerList()).data.list
    //获取所有客户
    for(var i=0;i<tmp0.value.length;i++){
      tmp.push(tmp0.value[i])
    }
    customerNameOptions.value=tmp
    console.log(customerNameOptions.value);
}
setOptions()
  // 前端数据
  const recommendationCount = ref();
  const selectedCustomer = ref();
  const jaccardThreshold = ref();
  const purchaseAdvice = ref('');
  const userRecommendations = ref('');
  
// 向后端发送请求获取进货建议
const getPurchaseAdvice = async () => {
  try {
    const response = await getOrderGPTPublic({
      numRecommendations: recommendationCount.value,
      customerName: selectedCustomer.value,
      action: 'purchase' // 添加按钮值为 "purchase"
    });
    purchaseAdvice.value = response.data.AdviceJinHuo;
  } catch (error) {
    console.error('Failed to get purchase advice:', error);
  }
};

// 向后端发送请求获取用户喜好商品推荐
const getUserRecommendations = async () => {
  try {
    const response = await getOrderGPTPublic({
      customerName: selectedCustomer.value,
      similarVlue: jaccardThreshold.value,
      action: 'recommend' // 添加按钮值为 "recommend"
    });
    userRecommendations.value = response.data.strUserAlike;
  } catch (error) {
    console.error('Failed to get user recommendations:', error);
  }
};
</script>
  
  <style scoped>
  .container {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  
  .module {
    border: 1px solid #2f7ed8; /* 深蓝色边框 */
    background-color: #f0f6ff; /* 淡蓝色背景 */
    padding: 20px;
    height: 450px
  }
  
  .module h2 {
    margin-top: 0;
    color: #2f7ed8; /* 深蓝色标题 */
  }
  
  .form {
    display: flex;
    align-items: center;
    gap: 10px;
  }
  
  .form label {
    font-weight: bold;
    color: #2f7ed8; /* 深蓝色标签 */
  }
  
  .description {
    margin-top: 10px;
  }
  
  textarea {
    width: 100%;
    height: 200px;
    resize: vertical;
  }
  </style>
  