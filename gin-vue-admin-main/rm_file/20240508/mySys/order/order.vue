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
      
        <el-form-item label="订单编号" prop="uuid">
         <el-input v-model="searchInfo.uuid" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="客户姓名" prop="customer">
         <el-input v-model="searchInfo.customer" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="仓库名" prop="wareHouseName">
         <el-input v-model="searchInfo.wareHouseName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商品名" prop="goodsName">
         <el-input v-model="searchInfo.goodsName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商品类型" prop="goodsType">
         <el-input v-model="searchInfo.goodsType" placeholder="搜索条件" />

        </el-form-item>
        <!-- <el-form-item label="商品价格" prop="goodsPrice">
            
             <el-input v-model.number="searchInfo.goodsPrice" placeholder="搜索条件" />

        </el-form-item> -->
           <el-form-item label="折扣" prop="discount">
            <el-select v-model="searchInfo.discount" clearable placeholder="请选择" @clear="()=>{searchInfo.discount=undefined}">
              <el-option v-for="(item,key) in DiscountOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
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
        
        <el-table-column align="left" label="订单编号" prop="uuid" width="180" />
        <el-table-column align="left" label="客户姓名" prop="customer" width="120" />
        <el-table-column align="left" label="仓库名" prop="wareHouseName" width="120" />
        <el-table-column align="left" label="商品名" prop="goodsName" width="120" />
        <el-table-column align="left" label="商品类型" prop="goodsType" width="120" />
        <el-table-column align="left" label="商品数量" prop="goodsNum" width="120" />
        <el-table-column align="left" label="商品价格" prop="goodsPrice" width="120" />
        <el-table-column align="left" label="折扣" prop="discount" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.discount,DiscountOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="总价" prop="allPrice" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <!-- <el-button type="primary" link icon="edit" class="table-button" @click="updateOrderFunc(scope.row)">变更</el-button> -->
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
            <!--订单编号 创建时自动生成 -->
            <!-- <el-form-item label="订单编号:"  prop="uuid" >
              <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入订单编号" />
            </el-form-item> -->
            <!-- 客户姓名 从后台传过来的数据选择 -->
            <el-form-item label="客户姓名:"  prop="customer" >
              <el-select v-model="formData.customer" placeholder="请选择客户" style="width:100%" :clearable="true">
                <el-option v-for="(item,key) in customerNameOptions" :value="item.name" :key="key" :label="item.name" />
              </el-select>
            </el-form-item>
            <!-- 从warehouse里查看仓库  然后根据仓库名 查warehouseinfo表 选出里边的商品 商品类型 商品售价 进行预加载   --> 
            <el-form-item label="仓库名:"  prop="wareHouseName" >
              <el-select v-model="formData.wareHouseName" placeholder="请选择仓库" style="width:100%" :clearable="true" @change="changeWareHouse">
                <el-option v-for="(item,key) in wareHouseNameOptions" :value="item.wareHouseName" :key="key" :label="item.wareHouseName" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品名:"  prop="goodsName" >
              <el-select v-model="formData.goodsName" placeholder="请选择商品" style="width:100%" :clearable="true" @change="changeGoodsName" >
                <el-option v-for="(item,key) in goodsList" :value="item.goodsName" :key="key" :label="item.goodsName" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品类型:"  prop="goodsType" >
              <el-input v-model="formData.goodsType" :clearable="true"  placeholder="请输入商品类型" disabled="true"/>
            </el-form-item>
            <el-form-item label="最大库存数量:">
              <el-input v-model.number="formData.numAlow" disabled="true" />
            </el-form-item>
            <el-form-item label="商品数量:"  prop="goodsNum" >
              <el-input v-model.number="formData.goodsNum" :clearable="true" placeholder="请输入商品数量" @input="goodsNumInput"/>
            </el-form-item>
            <el-form-item label="商品价格:"  prop="goodsPrice" >
              <el-input-number v-model="formData.goodsPrice"  style="width:100%" :precision="2" :clearable="true"  disabled="true"/>
            </el-form-item>
            <el-form-item label="折扣:"  prop="discount" >
              <el-select v-model="formData.discount" placeholder="请选择折扣" style="width:100%" :clearable="true"  @change="changeDiscount">
                <el-option v-for="(item,key) in DiscountOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="总价:"  prop="allPrice" >
              <el-input-number v-model="formData.allPrice"  style="width:100%" :precision="2" :clearable="false" disabled="true" />
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
                <el-descriptions-item label="订单编号">
                        {{ formData.uuid }}
                </el-descriptions-item>
                <el-descriptions-item label="客户姓名">
                        {{ formData.customer }}
                </el-descriptions-item>
                <el-descriptions-item label="仓库名">
                        {{ formData.wareHouseName }}
                </el-descriptions-item>
                <el-descriptions-item label="商品名">
                        {{ formData.goodsName }}
                </el-descriptions-item>
                <el-descriptions-item label="商品类型">
                        {{ formData.goodsType }}
                </el-descriptions-item>
                <el-descriptions-item label="商品数量">
                        {{ formData.goodsNum }}
                </el-descriptions-item>
                <el-descriptions-item label="商品价格">
                        {{ formData.goodsPrice }}
                </el-descriptions-item>
                <el-descriptions-item label="折扣">
                        {{ filterDict(formData.discount,DiscountOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="总价">
                        {{ formData.allPrice }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList
} from '@/api/mySys/order'
import {
  createMy_goods,
  deleteMy_goods,
  deleteMy_goodsByIds,
  updateMy_goods,
  findMy_goods,
  getMy_goodsList
} from '@/api/mySys/my_goods'
import {
  createWareHouseInfo,
  deleteWareHouseInfo,
  deleteWareHouseInfoByIds,
  updateWareHouseInfo,
  findWareHouseInfo,
  getWareHouseInfoList
} from '@/api/mySys/wareHouseInfo'
import {
  createCustomer,
  deleteCustomer,
  deleteCustomerByIds,
  updateCustomer,
  findCustomer,
  getCustomerList
} from '@/api/mySys/customer'
import {
  createWareHouse,
  deleteWareHouse,
  deleteWareHouseByIds,
  updateWareHouse,
  findWareHouse,
  getWareHouseList
} from '@/api/mySys/wareHouse'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'


defineOptions({
    name: 'Order'
})

// 自动化生成的字典（可能为空）以及字段
const DiscountOptions = ref([])
const formData = ref({
        uuid: '',
        customer: '',
        wareHouseName: '',
        goodsName: '',
        goodsType: '',
        goodsNum: 0,
        goodsPrice: 0,
        discount: '',
        allPrice: 0,
        numAlow:'',
        })
const formData_updateNum = ref({
        ID:'',
        wareHouseID: '',
        goodsID: '',
        num: 0,
        price: 0,
        wareHouseName: '',
        goodsName: '',
      })


// 验证规则
const rule = reactive({
               customer : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               wareHouseName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               goodsName : [{
                   required: true,
                   message: '',
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
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               goodsNum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               goodsPrice : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               discount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
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
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
  
// ============== 表格控制部分结束 ===============
var customerNameOptions=[]  //客户名选项
var wareHouseNameOptions=[]  //仓库选项
var goodsNameOptions=[]  //商品名选项
// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
  var tmp0=ref([])
  var tmp=ref([])
  tmp0.value=await (await getCustomerList()).data.list
  tmp.value = await (await getWareHouseList()).data.list
  //获取所有客户
  for(var i=0;i<tmp0.value.length;i++){
    customerNameOptions.push(tmp0.value[i])
  }
  //获取所有仓库名
  for(var i=0;i<tmp.value.length;i++){
    wareHouseNameOptions.push(tmp.value[i])
  }

  DiscountOptions.value = await getDictFunc('Discount')
}


const goodsList=ref([])//选中仓库中商品list
// 根据选中的仓库做出一些操作:根据warehouseinfo查找出该库存里的商品信息
const changeWareHouse = async () =>{
  var tmp=[]
  var warehouseInfoList=ref([])// wareHouseInfoList 仓库表信息
  warehouseInfoList.value=await (await getWareHouseInfoList()).data.list
  var warehouseName=formData.value.wareHouseName

  //根据仓库名比对出对应的商品
  for(var i=0;i<warehouseInfoList.value.length;i++){
    if(warehouseName==warehouseInfoList.value[i].wareHouseName){
      tmp.push(warehouseInfoList.value[i])
    }
  }
  goodsList.value=tmp
  formData.value.goodsName=''
}

// 根据选中的商品名做出一些操作:
const changeGoodsName = async () =>{
  var goodsName=formData.value.goodsName
  //遍历goodsList 根据名字比对
  for(var i=0;i<goodsList.value.length;i++){
    if(goodsName==goodsList.value[i].goodsName){
      formData.value.numAlow=goodsList.value[i].num
      formData.value.goodsPrice=goodsList.value[i].price
      //封装formData_updataNum
      console.log("end!!!!!!!!!!!!!!!!!!!!!!!!!!")
      console.log(goodsList.value[i])
      formData_updateNum.value.ID=goodsList.value[i].ID
      formData_updateNum.value.wareHouseID=goodsList.value[i].wareHouseID
      formData_updateNum.value.goodsID=goodsList.value[i].goodsID
      formData_updateNum.value.num=goodsList.value[i].num
      formData_updateNum.value.price=goodsList.value[i].price
      formData_updateNum.value.wareHouseName=goodsList.value[i].wareHouseName
      formData_updateNum.value.goodsName=goodsList.value[i].goodsName
      break
    }
  }

  //根据商品名 查找其对应的商品类型
  var tmplist=ref([])
  tmplist.value=await (await getMy_goodsList()).data.list

  for(var i=0;i<tmplist.value.length;i++){
    if(goodsName==tmplist.value[i].goodsName){
      formData.value.goodsType=tmplist.value[i].goodsType
      break
    }
  }
}
//改变折扣 修改总价格
const changeDiscount = async () =>{
  formData.value.allPrice=formData.value.goodsPrice*formData.value.goodsNum*Number(formData.value.discount)

}

//改变数量 修改总价格
const goodsNumInput = async () =>{
  formData.value.allPrice=formData.value.goodsPrice*formData.value.goodsNum*Number(formData.value.discount)

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
            deleteOrderFunc(row)
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
      const res = await deleteOrderByIds({ IDs })
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
const updateOrderFunc = async(row) => {
    const res = await findOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reorder
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOrderFunc = async (row) => {
    const res = await deleteOrder({ ID: row.ID })
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
  const res = await findOrder({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reorder
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          uuid: '',
          customer: '',
          wareHouseName: '',
          goodsName: '',
          goodsType: '',
          goodsNum: 0,
          goodsPrice: 0,
          discount: '',
          allPrice: 0,
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
        uuid: '',
        customer: '',
        wareHouseName: '',
        goodsName: '',
        goodsType: '',
        goodsNum: 0,
        goodsPrice: 0,
        discount: '',
        allPrice: 0,
        }
      formData_updateNum.value = {
        ID:'',
        wareHouseID: '',
        goodsID: '',
        num: 0,
        price: 0,
        wareHouseName: '',
        goodsName: '',
      }
}
// 弹窗确定
const enterDialog = async () => {
   //数量范围校验
   if(formData.value.goodsNum>formData.value.numAlow||formData.value.goodsNum<0){
              ElMessage({
                  type:'error',
                  message: '数量范围不合法，订单生成失败'
                })
                closeDialog()
                return
    }
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  //生成18位字符串作为订单编号
                  formData.value.uuid=generateRandomString()
                  //TODO 修改下库中商品的数量
                  formData_updateNum.value.num=formData_updateNum.value.num-formData.value.goodsNum
                  res = await createOrder(formData.value)
                  await updateWareHouseInfo(formData_updateNum.value)
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
                closeDialog()
                getTableData()
              }
      })
}
//生成18位随机数
function generateRandomString() {
    let result = '';
    const characters = '0123456789';
    const charactersLength = characters.length;
    for (let i = 0; i < 18; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}
</script>


