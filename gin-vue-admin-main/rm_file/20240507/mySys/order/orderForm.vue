<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单编号:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入订单编号" />
       </el-form-item>
        <el-form-item label="客户姓名:" prop="customer">
          <el-input v-model="formData.customer" :clearable="true"  placeholder="请输入客户姓名" />
       </el-form-item>
        <el-form-item label="商品名:" prop="goodsName">
          <el-input v-model="formData.goodsName" :clearable="true"  placeholder="请输入商品名" />
       </el-form-item>
        <el-form-item label="商品数量:" prop="goodsNum">
          <el-input v-model.number="formData.goodsNum" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品价格:" prop="goodsPrice">
          <el-input-number v-model="formData.goodsPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="折扣:" prop="discount">
          <el-input-number v-model="formData.discount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="总价:" prop="allPrice">
          <el-input-number v-model="formData.allPrice" :precision="2" :clearable="false"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createOrder,
  updateOrder,
  findOrder
} from '@/api/mySys/order'

defineOptions({
    name: 'OrderForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            uuid: '',
            customer: '',
            goodsName: '',
            goodsNum: 0,
            goodsPrice: 0,
            discount: 0,
            allPrice: 0,
        })
// 验证规则
const rule = reactive({
               customer : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               goodsName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               goodsNum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               goodsPrice : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reorder
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOrder(formData.value)
               break
             case 'update':
               res = await updateOrder(formData.value)
               break
             default:
               res = await createOrder(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
