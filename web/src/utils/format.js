import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'
import { ref } from 'vue'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  } else {
    return ''
  }
}

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}

const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'
export const ReturnArrImg = (arr) => {
  const imgArr = []
  if (arr instanceof Array) { // 如果是数组类型
    for (const arrKey in arr) {
      if (arr[arrKey].slice(0, 4) !== 'http') {
        imgArr.push(path + arr[arrKey])
      } else {
        imgArr.push(arr[arrKey])
      }
    }
  } else { // 如果不是数组类型
    if (arr.slice(0, 4) !== 'http') {
      imgArr.push(path + arr)
    } else {
      imgArr.push(arr)
    }
  }
  return imgArr
}

export const onDownloadFile = (url) => {
  window.open(path + url)
}

export const spiderTaskStatus = ref([
  { name: '未开始', value: '0' },
  { name: '处理中', value: '1' },
  { name: '处理完成', value: '2' },
  { name: '处理失败', value: '3' },
  { name: '废弃', value: '4' },
])

export const spiderTaskStatusFormat = (status) => {
  const statusMap = {
    0: '未开始',
    1: '处理中',
    2: '处理完成',
    3: '处理失败',
    4: '废弃',
  }
  return statusMap[status]
}
export const spiderTaskStatusColorFormat = (status) => {
  const colorMap = {
    0: 'warning',
    1: 'primary',
    2: 'success',
    3: 'error',
    4: 'info',
  }
  return colorMap[status]
}
