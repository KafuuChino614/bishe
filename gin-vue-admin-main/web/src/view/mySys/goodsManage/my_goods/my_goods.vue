<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="商品名" prop="goodsName">
         <el-input v-model="searchInfo.goodsName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商品类型" prop="goodsType">
         <el-input v-model="searchInfo.goodsType" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="商品名" prop="goodsName" width="120" />
        <el-table-column align="left" label="商品类型" prop="goodsType" width="120" />
        <el-table-column align="left" label="商品单位" prop="goodsUnit" width="120" />
        <el-table-column align="left" label="商品进价" prop="goodsPrice" width="120" />
        <el-table-column align="left" label="商品库存" prop="goodsNum" width="120" />
        <el-table-column align="left" label="商品供货商" prop="goodsVender" width="120" />
        <el-table-column align="left" label="生产日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.manufactureData) }}</template>
         </el-table-column>
         <el-table-column align="left" label="过期日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.expirationDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateMy_goodsFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="商品名:"  prop="goodsName" >
              <el-input v-model="formData.goodsName" :clearable="true"  placeholder="请输入商品名" />
            </el-form-item>
            <el-form-item label="商品类型:"  prop="goodsType" >
              <!-- --------------------------------------------------------------------------------------------- -->
              <el-select v-model="formData.goodsType" placeholder="请选择商品类型" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in goodsTypeOptions" :key="key" :label="item" :value="item" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品单位:"  prop="goodsUnit" >
              <!-- <el-input v-model="formData.goodsUnit" :clearable="true"  placeholder="请输入商品单位" /> -->
              <el-select v-model="formData.goodsUnit" placeholder="请选择商品单位" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in goodsUnitOptions" :key="key" :label="item" :value="item" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品进价:"  prop="goodsPrice" >
              <el-input-number v-model="formData.goodsPrice"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="商品库存:"  prop="goodsNum" >
              <el-input v-model.number="formData.goodsNum" :clearable="true" placeholder="请输入商品库存" />
            </el-form-item>
            <el-form-item label="商品供货商:"  prop="goodsVender" >
              <!-- <el-input v-model="formData.goodsVender" :clearable="true"  placeholder="请输入商品供货商" /> -->
              <el-select v-model="formData.goodsVender" placeholder="请选择供货商" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in goodsVenderOptions" :key="key" :label="item" :value="item" />
              </el-select>
            </el-form-item>
            <el-form-item label="生产日期:"  prop="manufactureData" >
              <el-date-picker v-model="formData.manufactureData" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="过期日期:"  prop="expirationDate" >
              <el-date-picker v-model="formData.expirationDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" destroy-on-close>
          <template #header>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="商品名">
                        {{ formData.goodsName }}
                </el-descriptions-item>
                <el-descriptions-item label="商品类型">
                        {{ formData.goodsType }}
                </el-descriptions-item>
                <el-descriptions-item label="商品单位">
                        {{ formData.goodsUnit }}
                </el-descriptions-item>
                <el-descriptions-item label="商品价格">
                        {{ formData.goodsPrice }}
                </el-descriptions-item>
                <el-descriptions-item label="商品库存">
                        {{ formData.goodsNum }}
                </el-descriptions-item>
                <el-descriptions-item label="商品供货商">
                        {{ formData.goodsVender }}
                </el-descriptions-item>
                <el-descriptions-item label="生产日期">
                      {{ formatDate(formData.manufactureData) }}
                </el-descriptions-item>
                <el-descriptions-item label="过期日期">
                      {{ formatDate(formData.expirationDate) }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createMy_goods,
  deleteMy_goods,
  deleteMy_goodsByIds,
  updateMy_goods,
  findMy_goods,
  getMy_goodsList
} from '@/api/mySys/my_goods'

import {
  createMy_goodsType,
  deleteMy_goodsType,
  deleteMy_goodsTypeByIds,
  updateMy_goodsType,
  findMy_goodsType,
  getMy_goodsTypeList,
  getMy_goodsTypePublic
} from '@/api/mySys/my_goodsType'

import {
  createMy_goodsUnit,
  deleteMy_goodsUnit,
  deleteMy_goodsUnitByIds,
  updateMy_goodsUnit,
  findMy_goodsUnit,
  getMy_goodsUnitList
} from '@/api/mySys/my_goodsUnit'

import {
  createMy_vendor,
  deleteMy_vendor,
  deleteMy_vendorByIds,
  updateMy_vendor,
  findMy_vendor,
  getMy_vendorList
} from '@/api/mySys/my_vendor'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'My_goods'
})

// 自动化生成的字典（可能为空）以及字段


const formData = ref({
        goodsName: '',
        goodsType: '',
        goodsUnit: '',
        goodsPrice: 0,
        goodsNum: 0,
        goodsVender: '',
        manufactureData: new Date(),
        expirationDate: new Date(),
        })


// 验证规则
const rule = reactive({
               goodsName : [{
                   required: true,
                   message: '请输入商品名',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               goodsType : [{
                   required: true,
                   message: '请选择商品类型',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               goodsUnit : [{
                   required: true,
                   message: '请选择商品单位',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               goodsPrice : [{
                   required: true,
                   message: '请输入价格',
                   trigger: ['input','blur'],
               },
              ],
               goodsVender : [{
                   required: true,
                   message: '请选择供货商',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
              manufactureData : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               expirationDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMy_goodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============
var goodsTypeOptions = []
var goodsUnitOptions = []
var goodsVenderOptions=[]
// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
  var tmp=ref([])
  var tmp1=ref([])
  var tmp2=ref([])
  tmp.value = await (await getMy_goodsTypeList()).data.list
  tmp1.value=await (await getMy_goodsUnitList()).data.list
  tmp2.value=await (await getMy_vendorList()).data.list
  for(var i=0;i<tmp.value.length;i++){
    goodsTypeOptions.push(tmp.value[i].typeName)
  }
  for(var i=0;i<tmp1.value.length;i++){
    goodsUnitOptions.push(tmp1.value[i].unitName)
  }
  for(var i=0;i<tmp2.value.length;i++){
    goodsVenderOptions.push(tmp2.value[i].venderName)
  }
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMy_goodsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteMy_goodsByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMy_goodsFunc = async(row) => {
    const res = await findMy_goods({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remy_goods
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMy_goodsFunc = async (row) => {
    const res = await deleteMy_goods({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findMy_goods({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.remy_goods
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          goodsName: '',
          goodsType: '',
          goodsUnit: '',
          goodsPrice: 0,
          goodsNum: 0,
          goodsVender: '',
          manufactureData: new Date(),
          expirationDate: new Date(),
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        goodsName: '',
        goodsType: '',
        goodsUnit: '',
        goodsPrice: 0,
        goodsNum: 0,
        goodsVender: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
