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

// @Summary 更新爬虫任务
// @Router /spider/taskStatus [get]
export const updateSpiderTaskStatus = (data) => {
  return service({
    url: '/spider/task',
    method: 'put',
    data: data
  })
}

// @Tags system
// @Summary 更新爬虫相关配置
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sysModel.System true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /spider/updateConfig [post]
export const updateSpiderConfig = (data) => {
  return service({
    url: '/spider/updateConfig',
    method: 'post',
    data
  })
}

// @Tags systrm
// @Summary 获取爬虫配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /spider/getConfig [post]
export const getSpiderConfig = () => {
  return service({
    url: '/spider/getConfig',
    method: 'post'
  })
}
