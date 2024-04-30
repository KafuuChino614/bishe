<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="仓库ID:" prop="wareHouseID">
          <el-input v-model.number="formData.wareHouseID" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品ID:" prop="goodsID">
          <el-input v-model.number="formData.goodsID" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品数量:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品单价:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="仓库名字:" prop="wareHouseName">
          <el-input v-model="formData.wareHouseName" :clearable="false"  placeholder="请输入仓库名字" />
       </el-form-item>
        <el-form-item label="商品名字:" prop="goodsName">
          <el-input v-model="formData.goodsName" :clearable="true"  placeholder="请输入商品名字" />
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
  createWareHouseInfo,
  updateWareHouseInfo,
  findWareHouseInfo
} from '@/api/mySys/wareHouseInfo'

defineOptions({
    name: 'WareHouseInfoForm'
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
            wareHouseID: 0,
            goodsID: 0,
            num: 0,
            price: 0,
            wareHouseName: '',
            goodsName: '',
        })
// 验证规则
const rule = reactive({
               wareHouseID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               goodsID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               num : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               price : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               wareHouseName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               goodsName : [{
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
      const res = await findWareHouseInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewareHouseInfo
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
               res = await createWareHouseInfo(formData.value)
               break
             case 'update':
               res = await updateWareHouseInfo(formData.value)
               break
             default:
               res = await createWareHouseInfo(formData.value)
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
