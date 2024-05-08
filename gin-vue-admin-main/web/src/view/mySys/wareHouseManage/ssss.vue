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
          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
  
      <!-- ------------------------------------------------------------------------------------------------------------------------ -->
      <div class="gva-table-box">
          <div class="gva-btn-list">
              <el-button type="primary" icon="plus" @click="openDialog">商品入库</el-button>
              <!-- <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">商品出库</el-button> -->
          </div>
          <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
          >
  
          <el-table-column align="left" label="仓库名" prop="wareHouseName" width="350" />
          <el-table-column align="left" label="商品名" prop="goodsName" width="350" />
          <el-table-column align="left" label="商品数量" prop="num" width="350" />
          <el-table-column align="left" label="商品单价（预计售价）" prop="price" width="350" />
          <el-table-column align="left" label="商品总价" prop="priceAll" width="350" />
          <el-table-column align="left" label="操作" fixed="right" min-width="240">
              <template #default="scope">
              <el-button type="primary" link icon="edit" class="table-button" @click="updateWareHouseInfoFunc(scope.row)">出库</el-button>
              <el-button type="primary" link icon="edit" class="table-button" @click="updateWareHouseInfoPriceFunc(scope.row)">修改价格</el-button>
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
      <!-- ------------------------------------------修改价格----------------------------------------------- -->
      <el-drawer size="800" v-model="dialogFormVisible_price" :show-close="false" :before-close="closeDialog">
         <template #header>
                <div class="flex justify-between items-center">
                  <span class="text-lg">修改价格</span>
                  <div>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                    <el-button @click="closeDialog">取 消</el-button>
                  </div>
                </div>
              </template>
  
            <el-form :model="formData_price" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
              <el-form-item label="商品单价:"  prop="price" >
                <el-input-number v-model="formData_price.price"  style="width:100%" :precision="2" :clearable="true" />
              </el-form-item>
  
            </el-form>
      </el-drawer>
      <!-- ------------------------------------------修改价格END!!!!----------------------------------------------- -->
  
  
      <!-- ------------------------------------------新增入库单----------------------------------------------- -->
      <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
         <template #header>
                <div class="flex justify-between items-center">
                  <span class="text-lg">{{type==='create'?'入库商品':'修改'}}</span>
                  <div>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                    <el-button @click="closeDialog">取 消</el-button>
                  </div>
                </div>
              </template>
  
            <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
              <el-form-item label="仓库名:"  prop="wareHouseID" >
                <el-select v-model="formData.wareHouseID" placeholder="请选择仓库" style="width:100%" :clearable="true" @change="changeWareHouse">
                  <el-option v-for="(item,key) in wareHouseNameOptions" :value="item.ID" :key="key" :label="item.wareHouseName" />
                </el-select>
              </el-form-item>
              <el-form-item label="商品名:"  prop="goodsID" >
                <el-select v-model="formData.goodsID" placeholder="请选择商品" style="width:100%" :clearable="true" @change="changeGoods">
                  <el-option v-for="(item,key) in goodsNameOptions" :key="key" :label="item.goodsName" :value="item.ID" />
                </el-select>
              </el-form-item>
              <el-form-item label="最大可入库数量:">
                <el-input v-model.number="formData.numAlow" disabled="true"/>
              </el-form-item>
              <el-form-item label="入库数量:"  prop="num" >
                <el-input v-model.number="formData.num" :clearable="true" />
              </el-form-item>
            </el-form>
      </el-drawer>
        <!-- ------------------------------------------新增入库单 END!----------------------------------------------- -->
  
  
  
  
  
  
  
  <!-- ------------------------------------------新增出库单----------------------------------------------- -->
  <el-drawer size="800" v-model="dialogFormVisible_out" :show-close="false" :before-close="closeDialog">
         <template #header>
                <div class="flex justify-between items-center">
                  <span class="text-lg">出库商品</span>
                  <div>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                    <el-button @click="closeDialog">取 消</el-button>
                  </div>
                </div>
              </template>
  
            <el-form :model="formData_out" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
              <el-form-item label="仓库名:"  prop="wareHouseName" >
                <el-input v-model.number="formData_out.wareHouseName" :clearable="false" disabled="true"/>
              </el-form-item>
              <el-form-item label="商品名:"  prop="goodsName" >
                <el-input v-model.number="formData_out.goodsName" :clearable="false" disabled="true"/>
              </el-form-item>
              <el-form-item label="最大可出库数量:">
                <el-input v-model.number="formData_out.numAlow" disabled="true"/>
              </el-form-item>
              <el-form-item label="出库数量:"  prop="num" >
                <el-input v-model.number="formData_out.num" :clearable="true"  @input="goodsNumInput"/>
              </el-form-item>
              <el-form-item label="商品单价:"  prop="price" >
                <el-input-number v-model="formData_out.price"  style="width:100%" :precision="2" :clearable="true"  disabled="true" />
              </el-form-item>
            </el-form>
            <el-form :model="formData_order" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <!-- TODO -->
              <!-- 客户姓名 从后台传过来的数据选择 -->
              <el-form-item label="客户姓名:"  prop="customer" >
                <el-select v-model="formData_order.customer" placeholder="请选择客户" style="width:100%" :clearable="true">
                  <el-option v-for="(item,key) in customerNameOptions" :value="item.name" :key="key" :label="item.name" />
                </el-select>
              </el-form-item>
              <!-- 折扣 -->
              <el-form-item label="折扣:"  prop="discount" >
                <el-select v-model="formData_order.discount" placeholder="请选择折扣" style="width:100%" :clearable="true"  @change="changeDiscount">
                  <el-option v-for="(item,key) in DiscountOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
              <!-- 出库商品总价 -->
              <el-form-item label="出库商品总价:"  prop="allPrice" >
                <el-input-number v-model="formData_order.allPrice"  style="width:100%" :precision="2" :clearable="false" disabled="true" />
              </el-form-item>
            </el-form>
      </el-drawer>
  <!-- ------------------------------------------新增出库单------END!!!!!!!----------------------------------------- -->
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
  
  import {
    createMy_goods,
    deleteMy_goods,
    deleteMy_goodsByIds,
    updateMy_goods,
    findMy_goods,
    getMy_goodsList
  } from '@/api/mySys/my_goods'
  
  import {
    createWareHouse,
    deleteWareHouse,
    deleteWareHouseByIds,
    updateWareHouse,
    findWareHouse,
    getWareHouseList
  } from '@/api/mySys/wareHouse'
  import {
    createCustomer,
    deleteCustomer,
    deleteCustomerByIds,
    updateCustomer,
    findCustomer,
    getCustomerList
  } from '@/api/mySys/customer'
  import {
    createOrder,
    deleteOrder,
    deleteOrderByIds,
    updateOrder,
    findOrder,
    getOrderList
  } from '@/api/mySys/order'
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
          numAlow:'',//可入库数量
          })
  const formData_out = ref({
          wareHouseID: '',
          goodsID: '',
          num: 0,
          price: 0,
          wareHouseName: '',
          goodsName: '',
          numAlow:'',//最大可出库数量
          })
  const formData_price = ref({
          wareHouseID: '',
          goodsID: '',
          num: 0,
          price: 0,
          wareHouseName: '',
          goodsName: '',
          numAlow:'',//可入库数量
          })
  
  const formData_order = ref({
          uuid: '',
          customer: '',
          wareHouseName: '',
          goodsName: '',
          goodsType: '',
          goodsUnit:'',
          goodsNum: 0,
          goodsPrice:0,
          discount:'1',
          allPrice:0,
          })
  // 验证规则
  const rule = reactive({
                 wareHouseID : [{
                     required: true,
                     message: '',
                     trigger: ['input','blur'],
                 },
                ],
                 goodsID : [{
                     required: true,
                     message: '',
                     trigger: ['input','blur'],
                 },
                ],
                 num : [
                  {
                     required: true,
                     message: '数量范围不合法',
                     trigger: ['input','blur'],
                 },
                ],
                 price : [{
                     required: true,
                     message: '',
                     trigger: ['input','blur'],
                 },
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
  
                 customer : [{
                     required: true,
                     message: '请选择顾客',
                     trigger: ['input','blur'],
                 },
                ],
                discount : [{
                     required: true,
                     message: '请选择折扣',
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
    const table = await getWareHouseInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      //遍历tableData.value设置priceAll的值
      for(var i=0;i<tableData.value.length;i++){
        tableData.value[i].priceAll=tableData.value[i].price*tableData.value[i].num
      }
  
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  
  getTableData()
  
  // ============== 表格控制部分结束 ===============
  var goodsNameOptions=[]
  var wareHouseNameOptions=[]
  var customerNameOptions=[]  //客户名选项
  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () =>{
    var tmp0=ref([])
    var tmp=ref([])
    var tmp1=ref([])
    tmp0.value=await (await getCustomerList()).data.list
    tmp.value = await (await getWareHouseList()).data.list
    tmp1.value=await(await getMy_goodsList()).data.list
     //获取所有客户
     for(var i=0;i<tmp0.value.length;i++){
      customerNameOptions.push(tmp0.value[i])
    }
    for(var i=0;i<tmp.value.length;i++){
      wareHouseNameOptions.push(tmp.value[i])
    }
    for(var i=0;i<tmp1.value.length;i++){
      goodsNameOptions.push(tmp1.value[i])
    }
    DiscountOptions.value = await getDictFunc('Discount')
  }
  // 获取需要的字典 可能为空 按需保留
  setOptions()
  // 根据选中wareHose的ID改变表单数值
  const changeWareHouse = async () =>{
    for(var i=0;i<wareHouseNameOptions.length;i++){
      if (wareHouseNameOptions[i].ID==formData.value.wareHouseID)formData.value.wareHouseName=wareHouseNameOptions[i].wareHouseName
    }
  }
  // 根据选中goods的ID改变表单数值
  const changeGoods = async () =>{
    for(var i=0;i<goodsNameOptions.length;i++){
      if(goodsNameOptions[i].ID==formData.value.goodsID){
        formData.value.goodsName=goodsNameOptions[i].goodsName
        formData.value.numAlow=goodsNameOptions[i].goodsNum
        formData.value.price=goodsNameOptions[i].goodsPrice
      }
    }
  
  }
  //改变数量 修改总价格
  const goodsNumInput = async () =>{
    formData_order.value.allPrice=formData_out.value.price*formData_out.value.num*Number(formData_order.value.discount)
  
  }
  //改变折扣 修改总价格
  const changeDiscount = async () =>{
    formData_order.value.allPrice=formData.value.goodsPrice*formData.value.goodsNum*Number(formData_order.value.discount)
  
  }
  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')
  
  
  // 更新行  (出库操作)
  var num_out  //用来记录查询到的商品数量
  const updateWareHouseInfoFunc = async(row) => {
      const res = await findWareHouseInfo({ ID: row.ID })
      type.value = 'update'
      if (res.code === 0) {
          formData_out.value = res.data.rewareHouseInfo
          formData_out.value.numAlow= res.data.rewareHouseInfo.num
          num_out=formData_out.value.numAlow
          formData_out.value.num= 0
          dialogFormVisible_out.value = true
      }
  }
  //更新价格
  // 更新行
  const updateWareHouseInfoPriceFunc = async(row) => {
      const res = await findWareHouseInfo({ ID: row.ID })
      type.value = 'updatePrice'
      if (res.code === 0) {
          formData_price.value = res.data.rewareHouseInfo
          dialogFormVisible_price.value = true
      }
  }
  
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
          }
  
  
          formData_out.value = {
          wareHouseID: '',
          goodsID: '',
          num: 0,
          price: 0,
          wareHouseName: '',
          goodsName: '',
        }
  
        formData_price.value = {
          wareHouseID: '',
          goodsID: '',
          num: 0,
          price: 0,
          wareHouseName: '',
          goodsName: '',
          }
          formData_order.value={
          uuid: '',
          customer: '',
          wareHouseName: '',
          goodsName: '',
          goodsType: '',
          goodsUnit: '',
          goodsNum: 0,
          goodsPrice:0,
          discount:'1',
          allPrice:0,
          }
  
  }
  // 弹窗确定
  const enterDialog = async () => {
       elFormRef.value?.validate( async (valid) => {
              //数量范围校验
              if(formData.value.num>formData.value.numAlow||formData.value.num<0){
                ElMessage({
                    type:'error',
                    message: '数量输入范围不合法，入库失败'
                  })
                  closeDialog()
                  return
              }
  
              if(type.value=='update'){
                if(formData_out.value.num>formData_out.value.numAlow||formData_out.value.num<0){
                ElMessage({
                    type:'error',
                    message: '数量输入范围不合法，出库失败'
                  })
                  closeDialog()
                  return
              }
            }
  
               if (!valid) return
                let res
                switch (type.value) {
                  case 'create':
                    res = await createWareHouseInfo(formData.value)
                    break
                  case 'update':
                    //出库商品  减少库存中商品的数量 
                    //封装formData_order数据 自动生成一个订单
                    var goodsID=formData_out.value.goodsID
                    var tmp=ref([])
                    tmp.value = await (await getMy_goodsList()).data.list
                    //根据商品id查询商品类型 和计量单位
                    for(var i=0;i<tmp.value.length;i++){
                       if(tmp.value[i].ID==goodsID){
                        formData_order.value.goodsType=tmp.value[i].goodsType
                        formData_order.value.goodsUnit=tmp.value[i].goodsUnit
                        break
                      }
                    }
                    formData_order.value.uuid=generateRandomString()
                    formData_order.value.wareHouseName=formData_out.value.wareHouseName
                    formData_order.value.goodsName=formData_out.value.goodsName
                    formData_order.value.goodsNum=formData_out.value.num
                    formData_order.value.goodsPrice=formData_out.value.price
                    formData_out.value.num=formData_out.value.numAlow-formData_out.value.num
                    res = await updateWareHouseInfo(formData_out.value)
                    await createOrder(formData_order.value)
                    break
                  case 'updatePrice':
                    res = await updateWareHouseInfo(formData_price.value)
                    break
                  default:
                    res = await createWareHouseInfo(formData.value)
                    break
                }
                if(res.code === 0&&type.value=='updatePrice'){
                  ElMessage({
                    type: 'success',
                    message: '修改成功'
                  })
                  closeDialog()
                  setOptions()
                  getTableData()
                  return
                }
  
                if(res.code === 0&&type.value=='update'){
                  ElMessage({
                    type: 'success',
                    message: '出库成功'
                  })
                  closeDialog()
                  setOptions()
                  getTableData()
                  return
                }
  
                if (res.code === 0) {
                  ElMessage({
                    type: 'success',
                    message: '入库成功'
                  })
                  closeDialog()
                  setOptions()
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
  
  <style>
  
  </style>
  