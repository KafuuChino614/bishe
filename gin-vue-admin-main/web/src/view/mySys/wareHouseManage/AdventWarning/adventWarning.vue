<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
  
        <el-form-item label="仓库名字" prop="wareHouseName">
         <el-input v-model="searchInfo.wareHouseName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商品名字" prop="goodsName">
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

    <!-- ------------------------------------------------------------------------------------------------------------------------ -->
    <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column align="left" label="仓库名" prop="wareHouseName" width="180" />
        <el-table-column align="left" label="商品名" prop="goodsName" width="180" />
        <el-table-column align="left" label="商品类型" prop="goodsType" width="180" />
        <el-table-column align="left" label="计量单位" prop="goodsUnit" width="180" />
        <el-table-column align="left" label="商品数量" prop="num" width="180" />
        <el-table-column align="left" label="商品单价（预计售价）" prop="price" width="240" />
        <el-table-column align="left" label="生产日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.manufactureData) }}</template>
         </el-table-column>
         <el-table-column align="left" label="过期日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.expirationDate) }}</template>
         </el-table-column>
         <el-table-column align="left" label="备注" width="180">
          <span style="color: red;">(临期产品)保质期低于或等于30天</span>
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
    </div>
</template>

<script setup>
import {
  createWareHouseInfo,
  deleteWareHouseInfo,
  deleteWareHouseInfoByIds,
  updateWareHouseInfo,
  findWareHouseInfo,
  getWareHouseInfoList
} from '@/api/mySys/wareHouseInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'


defineOptions({
    name: 'WareHouseInfo'
})

// 自动化生成的字典（可能为空）以及字段
const DiscountOptions = ref([])
const formData = ref({
        wareHouseID: '',
        goodsID: '',
        num: 0,
        price: 0,
        wareHouseName: '',
        goodsName: '',
        goodsType:'',
        goodsUnit:'',
        numAlow:'',//可入库数量
        manufactureData: new Date(),
        expirationDate: new Date(),
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
  const table = await getWareHouseInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    // 获取当前系统时间，与expirationDate做对比，筛选出保质期
    const currentDate = new Date(); // 获取当前系统时间
    const tmp = table.data.list.filter(item => {
      const expirationDate = new Date(item.expirationDate); // 将expirationDate转换为日期对象
      const timeDiff = expirationDate.getTime() - currentDate.getTime(); // 计算保质期剩余时间
      const daysDiff = timeDiff / (1000 * 3600 * 24); // 转换为天数
      return daysDiff <= 30; // 返回保质期小于等于30天的数据
    });

    tableData.value = tmp; // 将保质期小于等于30天的数据赋值给tableData.value
    total.value = tableData.length
    page.value = 1
    pageSize.value = tableData.length
  }
}

getTableData()

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
  DiscountOptions.value = await getDictFunc('Discount')
}
// 获取需要的字典 可能为空 按需保留
setOptions()

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')


// 弹窗控制标记
const dialogFormVisible = ref(false)
const dialogFormVisible_out = ref(false)
const dialogFormVisible_price = ref(false)
// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWareHouseInfo({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewareHouseInfo
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          wareHouseID: '',
          goodsID: 0,
          num: 0,
          price: 0,
          wareHouseName: '',
          goodsName: '',
          numAlow:'',//可入库数量
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
    dialogFormVisible_out.value = false
    dialogFormVisible_price.value=false

    formData.value = {
        wareHouseID: '',
        goodsID: '',
        num: 0,
        price: 0,
        wareHouseName: '',
        goodsName: '',
        goodsType:'',
        goodsUnit:'',
        manufactureData: new Date(),
        expirationDate: new Date(),
        }

}


</script>

<style>

</style>
