<template>
  <div class="system">
    <el-form
        ref="form"
        :model="config"
        label-width="240px"
    >
      <!--  System start  -->
      <el-collapse v-model="spiderConfigModel">
        <el-collapse-item
            title="爬虫配置"
            name="config-model"
        >
          <el-form-item label="封面图存储地址">
            <el-input v-model.number="config.spider.picPath.cover"/>
          </el-form-item>
          <el-form-item label="户型图存储地址">
            <el-input v-model.number="config.spider.picPath.layout"/>
          </el-form-item>
          <el-form-item label="请求头Cookie">
            <el-input v-model.number="config.spider.header.cookie"/>
          </el-form-item>
          <el-form-item label="请求头Agent">
            <el-input v-model.number="config.spider.header.agent"/>
          </el-form-item>
          <el-form-item label="爬虫休眠时长">
            <el-input v-model.number="config.spider.sleepSecond"/>
          </el-form-item>
        </el-collapse-item>
      </el-collapse>
    </el-form>
    <div class="mt-4">
      <el-button
          type="primary"
          @click="update"
      >立即更新
      </el-button>
      <el-button
          type="primary"
          @click="reload"
      >重启服务（开发中）
      </el-button>
    </div>
  </div>
</template>

<script setup>
import { updateSpiderConfig, getSpiderConfig } from '@/api/spider'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Config',
})

const spiderConfigModel = reactive(['config-model'])
const config = ref({
  spider: {
    picPath: {
      cover: '',
      layout: '',
    },
    header: {
      cookie: '',
      agent: '',
    },
    sleepSecond: '3',
  },
})

const initForm = async() => {
  const res = await getSpiderConfig()
  if (res.code === 0) {
    config.value = res.data
  }
}
initForm()
const reload = () => {
}
const update = async() => {
  const res = await updateSpiderConfig(config.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置文件设置成功',
    })
    await initForm()
  }
}
</script>

<style lang="scss">
.system {
  @apply bg-white p-9 rounded;
  h2 {
    @apply p-2.5 my-2.5 text-lg shadow;
  }
}
</style>
