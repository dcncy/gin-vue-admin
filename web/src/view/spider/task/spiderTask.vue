<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
            type="primary"
            icon="plus"
            @click="addSpiderTask"
        >新增爬虫任务
        </el-button>
      </div>
      <el-table
          :data="tableData"
          row-key="ID"
      >
        <el-table-column
            align="left"
            label="ID"
            min-width="50"
            prop="id"
        />
        <el-table-column
            align="left"
            label="任务名称"
            min-width="180"
            prop="taskName"
        />
        <el-table-column
            align="left"
            label="任务链接前缀"
            min-width="350"
            prop="taskUrlPrefix"
        />
        <el-table-column
            align="left"
            label="任务链接后缀"
            min-width="200"
            prop="taskUrlSuffix"
        />
        <el-table-column
            align="center"
            label="总页数"
            min-width="80"
            prop="pageNum"
        >
        </el-table-column>
        <el-table-column
            align="center"
            label="状态"
            min-width="100"
            prop="status"
        >
          <template #default="scope">
            <div>
              <el-tag :type="spiderTaskStatusColorFormat(scope.row.status)"
                      class="clickable-tag"
                      @click="tagClicked(scope.row)"
              >
                {{ spiderTaskStatusFormat(scope.row.status) }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column
            align="center"
            label="开始时间"
            min-width="200"
            prop="startTime"
        >
          <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
        </el-table-column>
        <el-table-column
            align="center"
            label="完成时间"
            min-width="200"
            prop="endTime"
        >
          <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
        </el-table-column>
        <el-table-column
            align="center"
            label="操作"
            min-width="150"
            fixed="right"
        >
          <template #default="scope">
            <el-button
                type="primary"
                link
                icon="edit"
                @click="openEdit(scope.row)"
            >编辑
            </el-button>
            <el-button
                type="primary"
                link
                icon="delete"
                @click="deleteSpiderTaskFunc(scope.row)"
            >删除
            </el-button>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer
        v-model="addUserDialog"
        size="50%"
        :show-close="false"
        :close-on-press-escape="false"
        :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">爬虫任务</span>
          <div>
            <el-button @click="closeAddUserDialog">取 消</el-button>
            <el-button
                type="primary"
                @click="enterAddSpiderTaskDialog"
            >确 定
            </el-button>
          </div>
        </div>
      </template>

      <el-form
          ref="userForm"
          :rules="rules"
          :model="spiderTask"
          label-width="150px"
      >
        <el-form-item
            label="任务名称"
            prop="taskName"
        >
          <el-input v-model="spiderTask.taskName"/>
        </el-form-item>
        <el-form-item
            label="任务链接前缀"
            prop="taskUrlPrefix"
        >
          <el-input v-model="spiderTask.taskUrlPrefix"/>
        </el-form-item>
        <el-form-item
            label="任务链接后缀"
            prop="taskUrlSuffix"
        >
          <el-input v-model="spiderTask.taskUrlSuffix"/>
        </el-form-item>
        <el-form-item
            label="总页数"
            prop="pageNum"
        >
          <el-input v-model="spiderTask.pageNum"/>
        </el-form-item>
        <el-form-item
            label="状态"
            prop="status"
        >
          <el-select
              v-model="spiderTask.status"
              clearable
              placeholder="请选择"
          >
            <el-option
                v-for="item in spiderTaskStatus"
                :key="item.value"
                :label="`${item.name}【${item.value}】`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>

import {
  createSpiderTask,
  getSpiderTaskList,
  deleteSpiderTask,
  updateSpiderTaskStatus, startSpiderTask,
} from '@/api/spider'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate, spiderTaskStatusFormat, spiderTaskStatusColorFormat, spiderTaskStatus } from '@/utils/format'

defineOptions({
  name: 'User',
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getSpiderTaskList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const deleteSpiderTaskFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async() => {
    const res = await deleteSpiderTask({ id: row.id })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      await getTableData()
    }
  })
}

// 弹窗相关
const spiderTask = ref({
  taskName: '',
  taskUrlPrefix: '',
  taskUrlSuffix: '',
  pageNum: '',
  status: '0',
})

const rules = ref({
  taskName: [
    { required: true, message: '请输入爬虫名称', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' },
    { max: 255, message: '最大255位字符', trigger: 'blur' },
  ],
  taskUrlPrefix: [
    { required: true, message: '请输入任务链接前缀', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' },
    { max: 255, message: '最大255位字符', trigger: 'blur' },
  ],
  taskUrlSuffix: [
    { required: true, message: '请输入任务链接后缀', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' },
    { max: 255, message: '最大255位字符', trigger: 'blur' },
  ],
  pageNum: [
    { required: true, message: '请输入总页数', trigger: 'blur' },
    {
      pattern: /^([1-9][0-9]{0,5}|1000000)$/,
      message: '请输入非0正整数，且小于100万',
      trigger: 'blur',
    },
  ],
  status: [
    { required: true, message: '请选择状态值', trigger: 'blur' },
  ],
})

const userForm = ref(null)
// 确认新增爬虫任务
const enterAddSpiderTaskDialog = async() => {
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...spiderTask.value,
      }
      if (dialogFlag.value === 'add') {
        const res = await createSpiderTask(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await updateSpiderTaskStatus(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '更新成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  spiderTask.value = {
    taskName: '',
    taskUrlPrefix: '',
    taskUrlSuffix: '',
    pageNum: '',
    status: '0',
  }
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addSpiderTask = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  spiderTask.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const tagClicked = (row) => {
  const item = JSON.parse(JSON.stringify(row))
  if (item.status === '0') {
    ElMessageBox.confirm('确定要开始执行这个任务吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }).then(async() => {
      const res = await startSpiderTask({ id: row.id })
      if (res.code === 0) {
        ElMessage.success('启动任务成功')
        await getTableData()
      }
    })
  }
}
</script>

<style lang="scss">
.header-img-box {
  @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}

.clickable-tag {
  cursor: pointer;
}
</style>
