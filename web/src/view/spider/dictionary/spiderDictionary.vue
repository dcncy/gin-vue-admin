<template>
  <div>
    <div class="dict-box flex gap-4">
      <div class="w-64 bg-white p-4">
        <div class="flex justify-between items-center">
          <span class="text font-bold">字典类型</span>
        </div>
        <el-scrollbar
            class="mt-4"
            style="height: calc(100vh - 300px)"
        >
          <div
              v-for="dictionary in dictionaryData"
              :key="dictionary.value"
              class="rounded flex justify-between items-center px-2 py-4 cursor-pointer mt-2 hover:bg-blue-50 hover:text-gray-800 group bg-gray-50"
              :class="selectID === dictionary.value && 'active'"
              @click="toDetail(dictionary)"
          >
            <span class="max-w-[160px] truncate">{{ dictionary.label }}</span>
          </div>
        </el-scrollbar>
      </div>
      <div class="flex-1 bg-white">
        <sysDictionaryDetail :sys-dictionary-i-d="selectID"/>
      </div>
    </div>
  </div>
</template>

<script setup>

import {
  getSysDictionaryDetailList,
} from '@/api/sysDictionaryDetail'

import { ref } from 'vue'
import sysDictionaryDetail from './spiderDictionaryDetail.vue'

defineOptions({
  name: 'SysDictionary',
})

const selectID = ref('CITY')

const dictionaryData = ref([])

const DEFAULT_SPIDER_DICT_CODE = 8
// 查询
const getTableData = async() => {
  const res = await getSysDictionaryDetailList({
    page: 1,
    pageSize: 100,
    sysDictionaryID: DEFAULT_SPIDER_DICT_CODE,
  })
  if (res.code === 0) {
    dictionaryData.value = res.data.list
  }
}

getTableData()

const toDetail = (row) => {
  selectID.value = row.value
}
</script>

<style>
.dict-box {
  height: calc(100vh - 240px);
}

.active {
  background-color: var(--el-color-primary) !important;
  color: #fff;
}
</style>
