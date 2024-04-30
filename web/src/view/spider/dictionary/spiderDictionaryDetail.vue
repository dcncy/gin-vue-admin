<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list justify-between">
        <span class="text font-bold">字典详细内容</span>
        <el-button
            type="primary"
            icon="plus"
            @click="openDialog"
        >新增字典项
        </el-button>
      </div>
      <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="id"
      >
        <el-table-column
            align="left"
            label="字典项"
            prop="name"
            min-width="200"
        />

        <el-table-column
            align="left"
            label="字典值"
            prop="code"
            min-width="200"
        />

        <el-table-column
            align="left"
            label="父结点"
            prop="parentCode"
            min-width="100"
        />

        <el-table-column
            align="center"
            label="操作"
            width="200"
        >
          <template #default="scope">
            <el-button
                type="primary"
                link
                icon="edit"
                @click="updateSysDictionaryDetailFunc(scope.row)"
            >变更
            </el-button>
            <el-button
                type="primary"
                link
                icon="delete"
                @click="deleteSysDictionaryDetailFunc(scope.row)"
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

    <el-dialog
        v-model="dialogFormVisible"
        :before-close="closeDialog"
        :title="type==='create'?'添加字典项':'修改字典项'"
    >
      <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="110px"
      >
        <el-form-item
            label="展示值"
            prop="label"
        >
          <el-input
              v-model="formData.name"
              placeholder="请输入字典项"
              clearable
              :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item
            label="字典值"
            prop="value"
        >
          <el-input
              v-model="formData.code"
              placeholder="请输入字典值"
              clearable
              :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item
            label="父结点"
            prop="extend"
        >
          <el-input
              v-model="formData.parentCode"
              placeholder="请输入父结点值"
              clearable
              :style="{width: '100%'}"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
              type="primary"
              @click="enterDialog"
          >确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import {
  findDictByType,
  createSpiderDictionary,
  updateSpiderDictionary,
  deleteSpiderDictionary, getSpiderDictionary,
} from '@/api/spider'

import { ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'SysDictionaryDetail',
})

const props = defineProps({
  sysDictionaryID: {
    type: Number,
    default: 0,
  },
})

const formData = ref({
  name: null,
  code: null,
  parentCode: null,
  type: null,
})
const rules = ref({
  name: [
    {
      required: true,
      message: '请输入字典项',
      trigger: 'blur',
    },
  ],
  code: [
    {
      required: true,
      message: '请输入字典值',
      trigger: 'blur',
    },
  ],
  parentCode: [
    {
      required: true,
      message: '请输入父结点值',
      trigger: 'blur',
    },
  ],
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
  const table = await findDictByType({
    page: page.value,
    pageSize: pageSize.value,
    type: String(props.sysDictionaryID),
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const type = ref('')
const dialogFormVisible = ref(false)

const updateSysDictionaryDetailFunc = async(row) => {
  const res = await getSpiderDictionary({ ID: row.id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: null,
    code: null,
    parentCode: null,
    type: String(props.sysDictionaryID),
  }
}
const deleteSysDictionaryDetailFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async() => {
    const res = await deleteSpiderDictionary({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功',
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

const dialogForm = ref(null)
const enterDialog = async() => {
  dialogForm.value.validate(async valid => {
    formData.value.type = String(props.sysDictionaryID)
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSpiderDictionary(formData.value)
        break
      case 'update':
        res = await updateSpiderDictionary(formData.value)
        break
      default:
        res = await createSpiderDictionary(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功',
      })
      closeDialog()
      getTableData()
    }
  })
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

watch(() => props.sysDictionaryID, () => {
  getTableData()
})

</script>

<style>
</style>
