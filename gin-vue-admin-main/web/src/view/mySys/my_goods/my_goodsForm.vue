<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品名:" prop="goodsName">
          <el-input v-model="formData.goodsName" :clearable="true"  placeholder="请输入商品名" />
       </el-form-item>
        <el-form-item label="商品类型:" prop="goodsType">
          <el-input v-model="formData.goodsType" :clearable="true"  placeholder="请输入商品类型" />
          <!-- TODO  -->
          <!-- <el-select v-model="formData.goodsType" placeholder="请选择商品类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in goodsTypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select> -->
       </el-form-item>
        <el-form-item label="商品单位:" prop="goodsUnit">
          <el-input v-model="formData.goodsUnit" :clearable="true"  placeholder="请选择商品单位" />
       </el-form-item>
        <el-form-item label="商品价格:" prop="goodsPrice">
          <el-input-number v-model="formData.goodsPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="商品库存:" prop="goodsNum">
          <el-input v-model.number="formData.goodsNum" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品供货商:" prop="goodsVender">
          <el-input v-model="formData.goodsVender" :clearable="true"  placeholder="请输入商品供货商" />
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
  createMy_goods,
  updateMy_goods,
  findMy_goods
} from '@/api/mySys/my_goods'




defineOptions({
    name: 'My_goodsForm'
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
            goodsName: '',
            goodsType: '',
            goodsUnit: '',
            goodsPrice: 0,
            goodsNum: 0,
            goodsVender: '',
        })
// 验证规则
const rule = reactive({
               goodsName : [{
                   required: true,
                   message: '请输入商品名',
                   trigger: ['input','blur'],
               }],
               goodsType : [{
                   required: true,
                   message: '请选择商品类型',
                   trigger: ['input','blur'],
               }],
               goodsUnit : [{
                   required: true,
                   message: '请选择商品单位',
                   trigger: ['input','blur'],
               }],
               goodsPrice : [{
                   required: true,
                   message: '请输入价格',
                   trigger: ['input','blur'],
               }],
               goodsVender : [{
                   required: true,
                   message: '请选择供货商',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMy_goods({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remy_goods
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
               res = await createMy_goods(formData.value)
               break
             case 'update':
               res = await updateMy_goods(formData.value)
               break
             default:
               res = await createMy_goods(formData.value)
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
