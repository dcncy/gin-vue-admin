import service from '@/utils/request'

// @Summary 新增爬虫任务
// @Router /spider/task [post]
export const createSpiderTask = (data) => {
  return service({
    url: '/spider/task',
    method: 'post',
    data: data
  })
}

// @Summary 分页获取爬虫任务列表
// @Router /spider/taskList [get]
export const getSpiderTaskList = (data) => {
  return service({
    url: '/spider/taskList',
    method: 'post',
    data: data
  })
}

// @Summary 删除爬虫任务
// @Router /spider/task [delete]
export const deleteSpiderTask = (data) => {
  return service({
    url: '/spider/task',
    method: 'delete',
    data: data
  })
}

// @Summary 获取爬虫任务信息
// @Router /spider/taskById [get]
export const getSpiderTaskInfo = () => {
  return service({
    url: '/spider/taskById',
    method: 'post'
  })
}

// @Summary 更新爬虫任务状态
// @Router /spider/taskStatus [get]
export const updateSpiderTaskStatus = (data) => {
  return service({
    url: '/spider/taskStatus',
    method: 'put',
    data: data
  })
}
