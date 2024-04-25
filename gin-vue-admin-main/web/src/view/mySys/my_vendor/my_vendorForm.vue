<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="厂家名:" prop="venderName">
          <el-input v-model="formData.venderName" :clearable="true"  placeholder="请输入厂家名" />
       </el-form-item>
        <el-form-item label="厂家地址:" prop="venderAddr">
          <el-input v-model="formData.venderAddr" :clearable="true"  placeholder="请输入厂家地址" />
       </el-form-item>
        <el-form-item label="负责人:" prop="header">
          <el-input v-model="formData.header" :clearable="true"  placeholder="请输入负责人" />
       </el-form-item>
        <el-form-item label="厂家电话:" prop="venderPhone">
          <el-input v-model="formData.venderPhone" :clearable="true"  placeholder="请输入厂家电话" />
       </el-form-item>
        <el-form-item label="厂家类型:" prop="venderClass">
           <el-select v-model="formData.venderClass" placeholder="请选择厂家类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in venderClassOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
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
  createMy_vendor,
  updateMy_vendor,
  findMy_vendor
} from '@/api/mySys/my_vendor'

defineOptions({
    name: 'My_vendorForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const venderClassOptions = ref([])
const formData = ref({
            venderName: '',
            venderAddr: '',
            header: '',
            venderPhone: '',
            venderClass: '',
        })
// 验证规则
const rule = reactive({
               venderName : [{
                   required: true,
                   message: '请填写厂家名',
                   trigger: ['input','blur'],
               }],
               venderAddr : [{
                   required: true,
                   message: '请填写厂家地址',
                   trigger: ['input','blur'],
               }],
               header : [{
                   required: true,
                   message: '请填写负责人',
                   trigger: ['input','blur'],
               }],
               venderPhone : [{
                   required: true,
                   message: '请填写厂家电话',
                   trigger: ['input','blur'],
               }],
               venderClass : [{
                   required: true,
                   message: '请填写厂家类型',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMy_vendor({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remy_vendor
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    venderClassOptions.value = await getDictFunc('venderClass')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMy_vendor(formData.value)
               break
             case 'update':
               res = await updateMy_vendor(formData.value)
               break
             default:
               res = await createMy_vendor(formData.value)
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
